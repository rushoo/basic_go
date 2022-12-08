package main

import (
	"fmt"
	"sync"
)

/*
一、冒泡排序bubbleSort：
从第一个数开始，依次对每对相邻元素比较大小，若左大于边则交换两数。这样每轮结束时可以把当前最大数放到（相对）最后。
1、n个数只需要排序n-1轮 2、每轮都把最大数置最后，故每轮排序次数递减 3、最优情况是已排序的顺序数列，最差是逆序数列
复杂度分析：
空间：整个冒泡排序过程中，只用了一个可以重复利用的临时变量，用于交换需要调换位置的元素，所以空间复杂度为O(1)。
时间：外层n-1次，内层循环最多的时候执行n-1次，最少的时候执行1次，平均n/2次，时间复杂度是O(n^2)

二、插入排序：
类似于整理手上的扑克牌，开始时以最左边第一张牌为有序数列，用第二张牌与之比较，如果较小则插入左边，较大就插入右边。
每次整牌后左边的有序数列就向右扩充一位。为了代码方便，这里是从有序数列从右到左逐次比较的。
复杂度分析：
空间：整个插入排序过程中，只用了一个可以重复利用的临时变量，用于交换需要调换位置的元素，所以空间复杂度为O(1)。
时间：在最坏情况下，数组完全逆序，第i个元素插入时要比较i-1次，因此，最坏情况下的比较次数是 1 + 2 + 3 + ... + (N - 1)，时间复杂度是O(n^2)

三、归并排序
先不断将slice向下折半拆分,然后将最终拆分得到的每一对slice向上顺序合并，这样每次合并结束数列的元素都是顺序的，最终全部元素都是顺序的。
要点是设计拆分和合并。
复杂度分析：
归并排序将数组不断递归分割最终到1时，应该做了logn次，而合并时应该有对每一个元素都扫描了一次共计n次，所以时间复杂度为O(n*logn)
每次归并都临时使用了一个数组，结束当层递归后释放，递归栈深度为logn，期间临时最大数组长度为n，所以空间复杂度为O(n)
归并排序不是原地排序，每次归并都需要创建新数组（相对而言这是很耗时的），再将两个小数组的值排序给到新数组。

四、快速排序
先对全部数组处理得到左小列、中、右大列，再对左小列、右大列如此处理。当每个子数组的粒度为1的时候，数组就全部排好序了。
复杂度分析：
时间: 最好情况每次递归都平分数组，一共需要递归logn次，每次需要n时间，复杂度为O(n*logn)，最坏情况每次都把数组分成1和n-1，一共需要递归n次，每次需要n时间，总体复杂度为O(n^2)。平均总体时间复杂度为O(nlogn)。
空间: 和时间复杂度相关，每次递归需要的空间是固定的，总体空间复杂度即为递归层数，因此平均/最好空间复杂度为O(logn)，最坏空间复杂度为O(n)

五、堆排序（略）
*/
type dataType interface {
	//约定排序函数支持的类型
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func merge[T dataType](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0 //通过移动k来给新slice赋值
	//将两个顺序数列合并成为一个顺序数列
	//递归拆分最后被insertSort的子数列必然是顺序数列
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	//当一边遍历结束另一边还有
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

func MergeSort[T dataType](data []T) []T {
	if len(data) == 0 || len(data) == 1 {
		return data
	}
	//先后对左右半归并排序，再合并一起
	left := MergeSort(data[:len(data)/2])
	right := MergeSort(data[len(data)/2:])
	return merge(left, right)
}
func InsertSort[T dataType](data []T) {
	end := len(data)
	for i := 1; i < end; i++ {
		tmp := data[i] //临时记录此时参与排序的元素
		j := i - 1     //从有序数列的从右到左开始比较，也相当于每次都向右扩充了一位
		for j >= 0 && tmp < data[j] {
			data[j+1] = data[j] //逆序,则右移一位
			j--
		}
		data[j+1] = tmp
	}
}
func BubbleSort[T dataType](data []T) {
	end := len(data)
	for i := 0; i < end-1; i++ {
		for j := 0; j < end-i-1; j++ {
			if data[j] >= data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func partition[T dataType](data []T, low, high int) int {
	//选取了最左边的第一个元素作为pivot，同时左右两端双指针
	var pivot, i, j = data[low], low, high
	for i < j {
		//左起，遇到大于pivot的数停下来
		for data[i] <= pivot && i < high {
			i++
		}
		//右起，遇到小于等于pivot的数停下来
		for data[j] > pivot && j > low {
			j--
		}
		//交换这两个数，再继续扫描
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	//考虑停止时，data[i]一定是大于pivot的，而j==i时，满足j左移的条件
	//最终j==i-1时停止，此时的data[j]==data[i-1]是一个比pivot小的数
	//  data[low], data[j] = data[j], data[low]
	//  不用这种方式是因为pivot可以不一定选最左边元素
	data[low] = data[j]
	data[j] = pivot
	return j
}
func quickSort[T dataType](data []T, low, high int) {
	// 当low >= high时不处理，相当于空数组或者仅一个元素（视为已排序）
	if low < high {
		pivot := partition(data, low, high)
		quickSort(data, low, pivot-1)
		quickSort(data, pivot+1, high)
	}
}
func QuickSort[T dataType](data []T) {
	length := len(data)
	quickSort(data, 0, length-1)
}

func partition2[T dataType](data []T) int {
	//单向双指针版，i指针不停右移直到逆序数时，将mid指针右移，再交换两者指向的值
	//此时mid左边的数都比小，直到i扫描结束，所有较小的数都在mid左边了
	pivot := data[0]
	mid := 0
	i := 1
	for i < len(data) {
		if data[i] < pivot {
			mid += 1
			data[i], data[mid] = data[mid], data[i]
		}
		i += 1
	}
	data[0], data[mid] = data[mid], data[0]
	return mid
}

// 并发快排+插入排序
const MAX = 10000
const SIZE = 100_000_000

func concurrentQuicksort[T dataType](data []T, wg *sync.WaitGroup) {
	for len(data) >= 30 {
		mid := partition2(data)
		var portion []T
		if mid < len(data)/2 {
			portion = data[:mid] //较少数列
			data = data[mid+1:]  //较多数列，继续拆分
		} else {
			portion = data[mid+1:]
			data = data[:mid]
		}
		//优先排序较小数列，大数列（>10000）并发处理，小数列直接快排
		if len(portion) > MAX {
			wg.Add(1)
			go func(data []T) {
				defer wg.Done()
				concurrentQuicksort(data, wg)
			}(portion)
		} else {
			concurrentQuicksort(portion, wg)
		}
	}
	//当数据量较小时，使用插入排序对最后一个个小数列排序
	InsertSort(data)
}

/*
堆排序是一种利用树结构的排序，将给定排序数列的元素构造大顶堆或者小顶堆，这样就实现了升序/降序排列了。例如：
slice2 := []int{10, 35, 100, 80, 30, 75, 40, 50, 60}
heap2 := NewHeap[int](slice2)
heap2.Insert(90)
fmt.Println(heap2.Items)

	func NewHeap[T Ordered](input []T) *Heap[T] {
		heap := &Heap[T]{}
		for i := 0; i < len(input); i++ {
			heap.Insert(input[i])
		}
		return heap
	}

	func (heap *Heap[T]) Insert(value T) {
		//先将元素插入，再调整堆结构
		heap.Items = append(heap.Items, value)
		heap.buildHeap(len(heap.Items) - 1)
	}

	func (heap *Heap[T]) buildHeap(index int) {
		var parent int
		if index > 0 {
			//当新插入的元素大于其父节点元素时，堆结构被破坏，须重新调整结构
			parent = (index - 1) / 2
			if heap.Items[index] > heap.Items[parent] {
				heap.swap(index, parent)
			}
			heap.buildHeap(parent)
		}
	}
*/
func main() {
	//numbers := []float64{3.5, -2.4, 12.8, 9.1}
	numbers := []float64{-1, -4, 3, 5, -3, 0, 1, 7}
	QuickSort[float64](numbers)
	fmt.Println(numbers)
}
