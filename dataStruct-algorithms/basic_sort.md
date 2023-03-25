### 基础排序算法

##### #冒泡排序
冒泡排序是简单直观的排序算法，类似水中的气泡上升过程中因为压力变化越来越大的过程。每趟冒泡排序的处理过程都会将数列中的最大数
放到相对最靠后的位置。说相对是因为 m 趟处理过程后，后 m 个元素已经是排好序的了，不必重复处理。所以冒泡排序也是一种从后往前的
交换排序。
```
// 冒泡排序
func bubbleSort(s []int) {
    if len(s) < 2 {
    	return
    }
    m := len(s)
    for m > 1 {
        //从第一个元素开始遍历，直到最后一个元素
    	for i := 0; i < m-1; i++ {
        	//每次依次将第一个元素与后一个相比较，逆序交换
    		if s[i] > s[i+1] {
    			s[i], s[i+1] = s[i+1], s[i]
    		}
    	}
    	//更新”相对最后“位置索引
    	m--
    }
}
```

##### #选择排序
冒泡排序中，每一趟中间过程的比较与交换是不必要的，可以暂时记住这个最大数，直到全部遍历完成后直接拿去和相对最后的元素交换。   
这本质上也是一种从后往前的选择排序。
```
//选择排序
func selectSort(s []int) {
    if len(s) < 2 {
    	return
    }
    for end := len(s) - 1; end > 0 ;end--{
    	j := 0
    	tmp := s[0]
    	for i := 1; i < end+1; i++ {
    		//使用tmp记住较大的数，而非直接交换,并记下该较大数的位置
    		if s[i] > tmp {
    			tmp = s[i]
    			j = i
    		}
    	}
    	//将该较大数与end交换,这样一趟排序仅一次交换，也实现了“冒泡”的效果
    	s[j], s[end] = s[end], s[j]
    }
}
```

##### #插入排序
前面两个是从后往前做排序，先依次将后若干个元素排好序，直到将整个数列排序好。另一种思路是从前往后排序，将前面的元素先排好序，   
直到将整个数列排序完成。当然就此而言，无论是从后往前还是从前往后做排序，并没有本质上的区别。
```
//插入排序
func insertSort(s []int) {
    //设想一下从左起第二张牌开始
    for i := 1; i < len(s); i++ {
    	//暂记当前待抽出的牌，如果是当前最大的也就跳过了本趟
    	tmp := s[i]
    	//如果它左边的牌的牌比它大，那就应该放在左边牌的左边
    	j := i - 1
    	for s[j] > tmp {
    	    //将左边较大的牌向右挪一下，方便这张牌的插入
    		s[j+1] = s[j]
    		//然后继续比较更左边一位的牌
    		j--
    		if j < 0 {
    			break
    		}
    	}
    	//此时j=-1，或者s[j]<tmp，应该插入s[j+1]的位置
    	s[j+1] = tmp
    }
}
```
基于50_000个随机数数列，比较上述这几种简单排序的性能，结果如下：
```
$ go test -bench='Sort$' -benchtime=5s  .
goos: windows
goarch: amd64
pkg: tmp
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkBubbleSort-16          1000000000               2.505 ns/op
BenchmarkSelectSort-16          1000000000               0.9302 ns/op
BenchmarkInsertSort-16          1000000000               0.3142 ns/op
PASS
ok      tmp     81.096s
```
可见，就执行性能而言，插入排序 > 选择排序 > 冒泡排序，而插入排序之所以在这里是最优的，是因为它每一趟不必完全遍历，可能顺序数时       
仅一次比较就结束本趟了，节省了很多计算量。     
     
##### #计数排序
计数排序是一种空间换时间的排序，对于高度聚合的数据，可以构造一个数组，每项的索引和值与原来数列中元素和重复出现次数对应。
```
func countingSort(s []int, maxValue int) {
    frequency := make([]int, maxValue+1)
    //-1值填充，以便统计0值
    for i, _ := range frequency {
    	frequency[i] = -1
    }
    for _, v := range s {
    	//统计待排序数组的元素频率
    	frequency[v]++
    }
    //	依据频率数组将元素回填
    j := 0 //元素回填时的索引
    for i := 0; i < maxValue; i++ {
    	counter := frequency[i]
    	for counter > -1 {
    		s[j] = i
    		counter--
    		j++
    	}
    }
}
```
比如样本范围0~100、样本容量为10_000_000的这种高聚合数列，计数排序几乎是瞬间完成的，结果如下：
```
$ go test -bench='Sort$' -benchtime=4s -count=2 .
goos: windows
goarch: amd64
pkg: tmp
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkCountingSort-16        1000000000               0.01050 ns/op
BenchmarkCountingSort-16        1000000000               0.01005 ns/op

// 对比结果为三分法快速排序
Benchmark_32_QuickSort-16       1000000000               0.1956 ns/op
Benchmark_32_QuickSort-16       1000000000               0.1980 ns/op
PASS
ok      tmp     5.363s
```

##### #希尔排序
希尔排序是插入排序的改进，插入排序在数据有序时效率是很高的。希尔排序通过将数组分组排序，逐渐增加数据的顺序程度，优化插入排序的效率。
```
func shellSort(s []int) {
    n := len(s)
    //划定组员间隔为step，折半缩减至1
    for step := n / 2; step > 0; step = step / 2 {
    	//从左往右，对任一元素，将其在所属组内做插入排序
    	for i := 0; i < n; i++ {
    		//记录当前处理到的元素索引和值
    		j, tmp := i, s[i]
    		//并对其所在组做插入排序，j >= step避免向左越界
    		for j >= step && tmp < s[j-step] {
    			s[j] = s[j-step]
    			j = j - step
    		}
    		s[j] = tmp
    	}
    }
}
```
比如样本范围0~10000、样本容量为100_000，希尔排序对比普通插入排序，优化是非常明显的，对比结果如下：
```
$ go test -bench='Sort$' -benchtime=4s -count=2 .
goos: windows
goarch: amd64
pkg: tmp
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkShellSort-16           1000000000               0.005516 ns/op
BenchmarkShellSort-16           1000000000               0.006002 ns/op

BenchmarkInsertSort-16          1000000000               1.194 ns/op
BenchmarkInsertSort-16          1000000000               1.195 ns/op
PASS
ok      tmp     41.909s
```
另外样本范围0~100_000、样本容量为10_000_000时，希尔排序和快速排序的对比，可能希尔排序的应用场景很狭窄：
```
$ go test -bench='Sort$' -benchtime=4s -count=2 .
goos: windows
goarch: amd64
pkg: tmp
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_2_QuickSort-16        1000000000               0.8157 ns/op
Benchmark_2_QuickSort-16        1000000000               0.8254 ns/op

Benchmark_32_QuickSort-16       1000000000               0.4986 ns/op
Benchmark_32_QuickSort-16       1000000000               0.4968 ns/op

BenchmarkShellSort-16           1000000000               1.568 ns/op
BenchmarkShellSort-16           1000000000               1.465 ns/op
PASS
ok      tmp     99.254s
```

###### 以上所有结果的测试用例和详细过程在[这里](test/sort/main_test.go)
