package main

/*
参考：https://zhuanlan.zhihu.com/p/273666774
bmap-reference:https://github.com/golang/go/blob/0bb6115dd6246c047335a75ce4b01a07c291befd/src/cmd/compile/internal/gc/reflect.go#L83
-map数据结构
map的结构体为hmap
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
mapextra的结构体
// mapextra holds fields that are not present on all maps.
type mapextra struct {
    // 如果 key 和 value 都不包含指针，并且可以被 inline(<=128 字节)
    // 就使用 hmap的extra字段 来存储 overflow buckets，这样可以避免 GC 扫描整个 map
    // 然而 bmap.overflow 也是个指针。这时候我们只能把这些 overflow 的指针
    // 都放在 hmap.extra.overflow 和 hmap.extra.oldoverflow 中了
    // overflow 包含的是 hmap.buckets 的 overflow 的 buckets
    // oldoverflow 包含扩容时的 hmap.oldbuckets 的 overflow 的 bucket
    overflow    *[]*bmap
    oldoverflow *[]*bmap

    // 指向空闲的 overflow bucket 的指针
    nextOverflow *bmap
}
bmap结构体
// A bucket for a Go map.
type bmap struct {
    // tophash包含此桶中每个键的哈希值最高字节（高8位）信息（也就是前面所述的high-order bits）。
    // 如果tophash[0] < minTopHash，tophash[0]则代表桶的搬迁（evacuation）状态。
    tophash [bucketCnt]uint8
}
bmap结构体其实不止包含 tophash 字段，还有些字段是在编译期间通过bmap()函数确定的，具体如下：
field = append(field, makefield("topbits", arr))
field = append(field, makefield("keys", arr))
field = append(field, makefield("elems", arr))
field = append(field, makefield("overflow", otyp))
所以准确的结构应该如下：
type bmap struct {
    topbits  [8]uint8
    keys     [8]key-type
    values   [8]value-type
    overflow uint-ptr
}
在8个键值对数据后面有一个overflow指针，因为桶中最多只能装8个键值对，如果有多余的键值对落到了当前桶，那么就需要再构建一个桶（称为溢出桶），通过overflow指针链接起来。
makeBucket为map创建用于保存buckets的数组。当桶的数量大于等于16个时，正常情况下就会额外创建2^(b-4)个溢出桶，所以正常桶和溢出桶在内存中的存储空间是连续的，只是被 hmap 中的不同字段引用而已。

map操作
-插入key
假定key经过哈希计算后得到64bit位的哈希值。如果B=5，buckets数组的长度，即桶的数量是32（2的5次方）。
例如，现要置一key于map中，该key经过哈希后，得到的哈希值如下：
10010111 	00001111011011001...0001111001010100010		00110
哈希值低位（low-order bits）用于选择桶，哈希值高位（high-order bits）用于在一个独立的桶中区别出键。
当B等于5时，那么我们选择的哈希值低位也是5位，即00110，对应于6号桶；再用哈希值的高8位，找到此key在桶中的位置。
高8位10010111，十进制为151.最开始桶中还没有key，那么新加入的key和value就会被放入第一个key空位和value空位。
当两个不同的key落在了同一个桶中，这时就发生了哈希冲突。go的解决方式是链地址法：在桶中按照顺序寻到第一个空位，若有位置，则将其置于其中；
否则，判断是否存在溢出桶，若有溢出桶，则去该桶的溢出桶中寻找空位，如果没有溢出桶，则添加溢出桶，并将其置溢出桶的第一个空位（非扩容的情况）。
通过对mapassign的代码分析之后，发现该函数并没有将插入key对应的value写入对应的内存，而是返回了value应该插入的内存地址。
赋值的最后一步实际上是编译器额外生成的汇编指令来完成的，可见靠 runtime 有些工作是没有做完的。所以，在go中，编译器和 runtime 配合，才能完成一些复杂的工作。
-查找key
如果是查找key，那么我们会根据高位哈希值去桶中的每个cell中找，若在桶中没找到，并且overflow不为nil，那么继续去溢出桶中寻找，直至找到，
如果所有的cell都找过了，还未找到，则返回key类型的默认值（例如key是int类型，则返回0）。
以下是查找的核心逻辑
双重循环遍历：外层循环是从桶到溢出桶遍历；内层是桶中的cell遍历
跳出循环的条件有三种：
第一种是已经找到key值；第二种是当前桶再无溢出桶；第三种是当前桶中有cell位的tophash值是emptyRest，这个值在前面解释过，它代表此时的桶后面的cell还未利用，所以无需再继续遍历。
// 当h.flags对应的值为hashWriting（代表有其他goroutine正在往map中写key）时，那么位计算的结果不为0，因此抛出以下错误。throw("concurrent map read and map write")
// 这也表明，go的map是非并发安全的
-删除key
b.tophash[i] = emptyRest
h.count--
删除一个key实际上是修改了当前 key 的标记，而不是直接删除了内存里面的数据。
所以如果用 map 做缓存，而每次更新只是部分更新，更新的 key 如果偏差比较大，有可能会有内存逐渐增长而不释放的问题。

-遍历map
结论：迭代 map 的结果是无序的
map在遍历时，并不是从固定的0号bucket开始遍历的，每次遍历，都会从一个随机值序号的bucket，再从其中随机的cell开始遍历。然后再按照桶序遍历下去，直到回到起始桶结束。

-map扩容 --渐进式扩容
装载因子是决定哈希表是否进行扩容的关键指标。在go的map扩容中，除了装载因子会决定是否需要扩容，溢出桶的数量也是扩容的另一关键指标。
为了保证访问效率，当map将要添加、修改或删除key时，都会检查是否需要扩容，扩容实际上是以空间换时间的手段。
map扩容条件，主要是两点:
1、判断已经达到装载因子的临界点，即元素个数 >= 桶（bucket）总数 * 6.5，这时候说明大部分的桶可能都快满了（即平均每个桶存储的键值对达到6.5个），如果插入新元素，有大概率需要挂在溢出桶（overflow bucket）上。
2、判断溢出桶是否太多，粗糙的判断方法是比较溢出桶的数量和2^B。
对于第2点，其实算是对第 1 点的补充。因为在装载因子比较小的情况下，有可能 map 的查找和插入效率也很低，
而第 1 点识别不出来这种情况。表面现象就是计算装载因子的分子比较小，即 map 里元素总数少，但是桶数量多（真实分配的桶数量多，包括大量的溢出桶）。
在某些场景下，比如不断的增删，这样会造成overflow的bucket数量增多，但负载因子又不高，未达不到第 1 点的临界值，就不能触发扩容来缓解这种情况。
这样会造成桶的使用率不高，值存储得比较稀疏，查找插入效率会变得非常低，复杂度可能会接近O（n），因此有了第 2 点判断指标。
两种情况官方采用了不同的解决方案：
-针对 1，将 B + 1，新建一个buckets数组，新的buckets大小是原来的2倍，然后旧buckets数据搬迁到新的buckets。该方法我们称之为增量扩容。
-针对 2，并不扩大容量，buckets数量维持不变，重新做一遍类似增量扩容的搬迁动作，把松散的键值对重新排列一次，以使bucket的使用率更高，进而保证更快的存取。该方法我们称之为等量扩容。
对于 2 的解决方案，其实存在一个极端的情况：如果插入 map 的 key 哈希都一样，那么它们就会落到同一个 bucket 里，这时整个哈希表已经退化成了一个链表，操作效率变成了 O(n)。移动元素其实解决不了问题。
为什么每次至多搬迁2个bucket？这其实是一种性能考量，如果map存储了数以亿计的key-value，一次性搬迁将会造成比较大的延时，因此才采用逐步搬迁策略。
知识点1：bucket序号的变化
对于等量扩容而言，由于buckets的数量不变，因此可以按照序号来搬迁。增量扩容时，假设B从5变成6。那么决定key值落入哪个bucket的低位哈希值就会发生变化（从取5位变为取6位），取新的低位hash值得过程称为rehash。
知识点2：确定搬迁区间
增量扩容到原来的 2 倍，桶的数量是原来的 2 倍，前一半桶被称为bucket x，后一半桶被称为bucket y。一个 bucket 中的 key 可能会分裂到两个桶中去，分别位于bucket x的桶，或bucket y中的桶。
而对于同一个桶而言，搬迁到bucket x和bucket y桶序号的差别是老的buckets大小，即2^old_B。确定了要搬迁到的目标 bucket 后，搬迁操作就比较好进行了。将源 key/value 值 copy 到目的地相应的位置。

-使用建议
从map设计可以知道，它并不是一个并发安全的数据结构。同时对map进行读写时，程序很容易出错。因此，要想在并发情况下使用map，建议并发安全的map——sync.Map。
遍历map的结果是无序的，在使用中，应该注意到该点。
通过map的结构体可以知道，它其实是通过指针指向底层buckets数组。所以和slice一样，尽管go函数都是值传递，但是，当map作为参数被函数调用时，在函数内部对map的操作同样会影响到外部的map。
有个特殊的key值math.NaN，它每次生成的哈希值是不一样的，这会造成m[math.NaN]是拿不到值的，而且多次对它赋值，会让map中存在多个math.NaN的key。
对于等量扩容的极端情况，比如反复插入删除同一个key值的数据会将map实际上变成链表结构，所以使用中要避免这类情况。


另外：
JDK 1.8 HashMap 采用位桶 + 链表 + 红黑树实现。（当链表长度超过阈值 “8” 时，将链表转换为红黑树）
简单说明
HashMap 的每一个元素，都是链表的一个节点（entry）。新增一个元素时，会先计算 key 的 hash 值，找到存入数组的位置。如果该位置已经有节点（链表头），则存入该节点的最后一个位置（链表尾）。
所以 HashMap 就是一个数组（bucket），数组上每一个元素都是一个节点（节点和所有下一个节点组成一个链表）或者为空，显然同一个链表上的节点 hash 值都一样。

C++中unordered_map的底层是用哈希表来实现的，通过key的哈希路由到每一个桶（即数组）用来存放内容。通过key来获取value的时间复杂度就是O（1）。因为key的哈希容易碰撞，所以需要对碰撞做处理。
unordered_map里的每一个数组（桶）里面存的其实是一个链表，key的哈希冲突以后会加到链表的尾部，这是再通过key获取value的时间复杂度就变成O(n），当碰撞很多的时候查询就会变慢。
为了优化这个时间复杂度，map的底层就把这个链表转换成了红黑树，这样虽然插入增加了复杂度，但提高了频繁哈希碰撞时的查询效率，使查询效率变成O(log n)。
*/

/*
----------sync.Map揭秘-------------
为了保证访问效率，当map将要添加、修改或删除key时，都会检查是否需要扩容，扩容实际上是以空间换时间的手段。
*/
import "fmt"

//func main() {
//	myTable := NewTable()
//	var words []string
//	mapCollection := make(map[string]string)
//	for i := 0; i < 50_000_000; i++ {
//		word := strconv.Itoa(i)
//		words = append(words, word)
//		myTable.Insert(word)
//		mapCollection[word] = word
//	}
//	fmt.Println("Benchmark test begins to test words: ", length)
//	start := time.Now()
//	for i := 0; i < length; i++ {
//		if myTable.IsPresent(words[i]) == false {
//			fmt.Println("Word not found in table: ", words[i])
//		}
//	}
//	elapsed := time.Since(start)
//	fmt.Println("Time to test all words in myTable: ", elapsed)
//
//	start = time.Now()
//	for i := 0; i < len(mapCollection); i++ {
//		// value,ok := mapCollection[key]
//		_, present := mapCollection[words[i]]
//		if !present {
//			fmt.Println("Word not found in mapCollection: ", words[i])
//		}
//	}
//	elapsed = time.Since(start)
//	fmt.Println("Time to test words in mapCollection: ", elapsed)
//}

//	func main() {
//		text := "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
//		pattern := "816406286208998628034825342"
//		start := time.Now()
//		_, _ = bruteSearch(text, pattern)
//		elapsed := time.Since(start)
//		fmt.Println("Computation time using BruteForceSearch: ", elapsed)
//		start = time.Now()
//		_, _ = Search(text, pattern)
//		elapsed = time.Since(start)
//		fmt.Println("Computation time using Search: ", elapsed)
//		fmt.Println(bruteSearch(text, pattern))
//		fmt.Println(Search(text, pattern))
//	}
//func main() {
//	set1 := Set[int]{}
//	set1.Insert(3)
//	set1.Insert(5)
//	set1.Insert(7)
//	set1.Insert(9)
//	set2 := Set[int]{}
//	set2.Insert(3)
//	set2.Insert(6)
//	set2.Insert(8)
//	set2.Insert(9)
//	set2.Insert(11)
//	set2.Delete(11)
//	fmt.Println("Items in set1: ", set1.Items())
//	fmt.Println("Items in set2: ", set2.Items())
//	fmt.Println("5 in set1: ", set1.In(5))
//	fmt.Println("5 in set2: ", set2.In(5))
//	fmt.Println("Union of set1 and set2: ", set1.Union(set2).Items())
//	fmt.Println("Intersection of set1 and set2: ", set1.Intersection(set2).Items())
//	fmt.Println("Difference of set2 with respect to set1: ", set2.Difference(set1).Items())
//	fmt.Println("Size of this difference: ", set1.Intersection(set2).Size())
//}

func main() {
	map1 := make(map[string]string, 0)
	map1["xxxxxxxx"] = "oooooooo"
	map2 := make(map[int]int, 9)
	map2[11111111] = 8888888888
	fmt.Println(map1, map2)
}
