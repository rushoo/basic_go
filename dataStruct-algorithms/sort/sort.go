package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
一、冒泡排序bubbleSort：
从第一个数开始，依次对每对相邻元素比较大小，若左大于边则交换两数。这样每轮结束时可以把当前最大数放到最后。
重复操作
要点：
1、n个数只需要排序n-1轮
2、每轮都把最大数置最后，故每轮排序次数递减
3、最优情况是已排序的顺序数列，最差是逆序数列

二、插入排序：
类似于整理手上的扑克牌，以data[i],i:=1开始，直到i=len(data)-1时结束
每次用data[i]与左边的牌从右到左依次比较，若遇到逆序，就抽出这张牌插入到逆序牌的右边
这样每次结束左边的牌都是顺序的，而下一次就是把右边的牌按顺序插入到左边

三、快速排序quickSort
本质上就是以某值为基准不断拆分为大值和小值数列。
1、从数列中挑出一个元素称为基准（pivot），一般用第一个元素
2、取两个下标low、high同时分别自左向右和自右向左扫描
3、在下标相遇前，若data[low]>pivot而data[high]<pivot,则交换两值
4、最后pivot和data[j]的值互换，并记下j的值
此时data[low,j]所有元素不大于pivot，data[j+1, high]中所有元素大于pivot，对这两个子序列重复以上操作。
要点：
1、合适的pivot于性能影响很大，否则可能带来频繁的数据交换
2、较大数据量时递归深度可能会很恐怖
3、并发改造应以无状态的数列拆分做切入点

四、归并排序
1、不断将slice向下折半拆分
2、将拆分的每一对slice向上顺序合并，这样每次合并得到的结果都是顺序数列
*/
type dataType interface {
	~int | ~float64 | ~string
}

const MAX = 10000
const SIZE = 100_000_000

func IsSorted[T dataType](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

// golang中字符串比较大小是逐字节比较Utf8编码

// 冒泡排序
func bubbleSort[T dataType](data []T) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// 插入排序
func insertSort[T dataType](data []T) {
	for i := 1; i < len(data); i++ {
		p := data[i]
		j := i - 1
		//j >= 0 && p < data[j] 表示先执行第一个判断，再执行第二个，顺序写反会越界panic
		for j >= 0 && p < data[j] {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = p
	}
}

// 快速排序
func partition[T dataType](data []T, low, high int) int {
	var pivot, i, j = data[low], low, high
	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}
		for data[j] > pivot && j > low {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[low] = data[j]
	data[j] = pivot
	return j
}
func quickSort[T dataType](data []T, low, high int) {
	if low < high {
		pivot := partition(data, low, high)
		quickSort(data, low, pivot-1)
		quickSort(data, pivot+1, high)
	}
}

// 并发快排+插入排序
func concurrentQuicksort[T dataType](data []T, wg *sync.WaitGroup) {
	for len(data) >= 30 {
		mid := Partition(data)
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
	insertSort(data)
}
func Partition[T dataType](data []T) int {
	//单向，简化版
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
func QSort[T dataType](data []T) {
	var wg sync.WaitGroup
	concurrentQuicksort(data, &wg)
	wg.Wait()
}

// 合并两个数列
func merge[T dataType](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i, j, k := 0, 0, 0
	// 例如{1，2，3，4}、{5，7，11，14}
	for i < len(left) && j < len(right) {
		//将两个顺序数列合并成为一个顺序数列
		//递归拆分最后被insertSort的子数列必然是顺序数列
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
		//	result{1，2，3，4}
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
		//	result{5，7，11，14}
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
		//	result{1，2，3，4，5，7，11，14}
	}
	return result
}
func mergeSort[T dataType](data []T) []T {
	if len(data) > 100 {
		//先拆分成两个数列
		middle := len(data) / 2
		left := data[:middle]
		right := data[middle:]
		//通过递归调用，先逐层拆分，再逐层归并
		data = merge(mergeSort(right), mergeSort(left))
	} else {
		// 当数列长度小于100时，返回插入排序的数列
		insertSort(data)
	}
	return data
}
func concurrentMergeSort[T dataType](data []T) []T {
	if len(data) > 1 {
		if len(data) <= MAX {
			return mergeSort(data)
		} else { // Concurrent
			middle := len(data) / 2
			left := data[:middle]
			right := data[middle:]

			var wg sync.WaitGroup
			wg.Add(2)
			var data1, data2 []T
			go func() {
				defer wg.Done()
				data1 = concurrentMergeSort(left)
			}()
			go func() {
				defer wg.Done()
				data2 = concurrentMergeSort(right)
			}()
			wg.Wait()
			return merge(data1, data2)
		}
	}
	return nil
}

func main() {
	data := make([]float64, SIZE)
	for i := 0; i < SIZE; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	/*	 data2 := make([]float64, SIZE)
		 copy(data2, data)
	*/
	start := time.Now()
	result := mergeSort[float64](data)
	fmt.Println("Elapsed time for mergeSort = ", time.Since(start))
	fmt.Println("Is sorted: ", IsSorted(result))

	data = make([]float64, SIZE)
	for i := 0; i < SIZE; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	start = time.Now()
	result1 := concurrentMergeSort(data)
	fmt.Println("Elapsed time for concurrent mergesort = ", time.Since(start))
	fmt.Println("Sorted: ", IsSorted(result1))

	data = make([]float64, SIZE)
	for i := 0; i < SIZE; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	data2 := make([]float64, SIZE)
	copy(data2, data)

	start = time.Now()
	quickSort(data, 0, len(data)-1)
	fmt.Println("Elapsed time for quicksort = ", time.Since(start))
	fmt.Println("Is sorted: ", IsSorted(data))

	start = time.Now()
	QSort[float64](data2)
	fmt.Println("Elapsed time for concurrent quicksort = ", time.Since(start))
	fmt.Println("Is sorted: ", IsSorted(data2))

	//	结果对比
	//Elapsed time for mergeSort =  13.7588855s
	//Is sorted:  true
	//Elapsed time for concurrent mergesort =  3.7632447s
	//Sorted:  true
	//Elapsed time for quicksort =  10.7013013s
	//Is sorted:  true
	//Elapsed time for concurrent quicksort =  1.9793596s
	//Is sorted:  true
}

//func main() {
//	numbers := []float64{3.5, -2.4, 12.8, 9.1}
//	names := []string{"Zachary", "John", "Moe", "Alex", "Robert"}
//	bubbleSort[float64](numbers)
//	bubbleSort[string](names)
//	fmt.Println(numbers)
//	fmt.Println(names)
//
//	numbers = []float64{3.5, -2.4, 12.8, 9.1}
//	names = []string{"Zachary", "John", "Moe", "Alex", "Robert"}
//	insertSort[float64](numbers)
//	insertSort[string](names)
//	fmt.Println("insertSort: ", numbers)
//	fmt.Println("insertSort: ", names)
//
//	numbers = []float64{3.5, -2.4, 12.8, 9.1}
//	names = []string{"Zachary", "John", "Moe", "Alex", "Robert"}
//	quickSort[float64](numbers, 0, len(numbers)-1)
//	quickSort[string](names, 0, len(names)-1)
//	fmt.Println(numbers)
//	fmt.Println(names)
//}
