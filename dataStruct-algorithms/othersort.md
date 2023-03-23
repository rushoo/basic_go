### 快速排序

#### #归并排序
归并排序和快速排序类似，也是一种分治思想的应用，先将数组不断折半拆散，再借用临时数组合并分散的数组元素，在合并的过程中完成排序。    
因为递归调用和额外开辟的临时空间，所以常常结合插入排序一起使用。另外数组的拆分有合并是独立的，所以可以使用并发。
##### #基本归并排序
```
func mergeSort(s []int) []int {
	//先拆分成两个数列
	if len(s) > 1 {
		middle := len(s) / 2
		left := s[:middle]
		right := s[middle:]
		//通过递归调用，先逐层拆分，再逐层归并
		s = merge(mergeSort(right), mergeSort(left))
	}
	return s
}
// 合并两个-有序-数列
func merge(left, right []int) []int {
	//额外开辟一块空间用于存放两数组元素
	result := make([]int, len(left)+len(right))
	i, j := 0, 0
	r := 0
	//将left的元素依次和right元素比较，将较小的数直接放入result，
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[r] = left[i]
			i++
		} else {
			result[r] = right[j]
			j++
		}
		r++
	}
	//可能存在一个数组元素已经完全复制，另一个数组还有若干元素未复制到result,则将剩余的全部复制过去
	for i < len(left) {
		result[r] = left[i]
		i++
		r++
	}
	for j < len(right) {
		result[r] = right[j]
		j++
		r++
	}
	return result
}
```
##### #归并排序改进
1、数列较小时直接使用插入排序，可以避免相当程度的递归和临时空间。   2、并发处理
```
func mergeSort(s []int) []int {
	//先拆分成两个数列
	if len(s) > 100 {
		middle := len(s) / 2
		left := s[:middle]
		right := s[middle:]
		//通过递归调用，先逐层拆分，再逐层归并
		s = merge(mergeSort(right), mergeSort(left))
	} else {
		insertSort(s)
	}
	return s
}
func concurrentMergeSort(data []int) []int {
	//较小数列直接归并了
	if len(data) <= 10000 {
		return mergeSort(data)
	} else {
		// Concurrent，left-slice、right-slice
		mid := len(data) / 2
		ls := data[:mid]
		rs := data[mid:]

		var wg sync.WaitGroup
		wg.Add(2)
		var data1, data2 []int
		go func() {
			defer wg.Done()
			data1 = concurrentMergeSort(ls)
		}()
		go func() {
			defer wg.Done()
			data2 = concurrentMergeSort(rs)
		}()
		wg.Wait()
		return merge(data1, data2)
	}
}
```
以下为数组长度10_000_000，不同并发粒度下，改进前后运行情况对比，并发本身带来的消耗也是不容忽视的。
```
//1000
BenchmarkConcurrentMergeSort-16         1000000000               0.3314 ns/op
BenchmarkMergeSort-16                   1000000000               0.9974 ns/op

//2000
BenchmarkConcurrentMergeSort-16         1000000000               0.2976 ns/op
BenchmarkMergeSort-16                   1000000000               0.8249 ns/op

// 5000
BenchmarkConcurrentMergeSort-16         1000000000               0.2707 ns/op
BenchmarkMergeSort-16                   1000000000               0.7969 ns/op

//10000
BenchmarkConcurrentMergeSort-16         1000000000               0.2792 ns/op
BenchmarkMergeSort-16                   1000000000               0.8098 ns/op
```
