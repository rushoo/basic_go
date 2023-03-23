package sort

import "sync"

// 冒泡排序 与 选择排序
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
func selectSort(s []int) {
	if len(s) < 2 {
		return
	}
	for end := len(s) - 1; end > 0; end-- {
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

// 插入排序 与 希尔排序
func insertSort(s []int) {
	//假象以下从左起第二张牌开始
	for i := 1; i < len(s); i++ {
		//暂记当前待抽出的牌
		tmp := s[i]
		//如果它左边的牌的牌比它大，那就应该放在左边牌的左边
		j := i - 1
		for s[j] > tmp {
			//将左边较大的牌向右挪以下，方便这张牌的插入
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
func shellSort(s []int) {
	n := len(s)
	//划定组员间隔为step，折半缩减至1
	for step := n / 2; step > 0; step = step / 2 {
		//从左往右，对任一元素，将其在所属组内做插入排序
		for i := 0; i < n; i++ {
			//记录当前处理到的元素索引和值
			j, tmp := i, s[i]
			//并对其所在组做插入排序过程，j >= step避免向左越界
			for j >= step && tmp < s[j-step] {
				s[j] = s[j-step]
				j = j - step
			}
			s[j] = tmp
		}
	}
}

// 计数排序
func countingSort(s []int, maxValue int) {
	//将元素出现次数作为频率数组的索引，这里考虑了0，所以+1
	frequency := make([]int, maxValue+1)
	//-1值填充，以便统计0值，若统计的最小数是负数值，也需要做类似处理
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

// 快速排序与优化
func partition(s []int, low, high int) int {
	// 直接判断退出可节省部分计算量，千万数据测试时可以看到这部分差别
	if high-low < 2 {
		return low
	}
	//通过交换，选定中间元素为基准值
	s[low], s[(high+low)/2] = s[(high+low)/2], s[low]
	var pivot = s[low]
	//双指针分别从数组最左边与最右边出发，指针停止条件：指向数与基准数逆序。
	var l, r = low, high
	for l < r {
		//先移动右指针，所以最后两指针则会合于较小的数
		for r > l && s[r] > pivot {
			r--
		}
		for l < r && s[l] <= pivot {
			l++
		}
		if l < r {
			s[l], s[r] = s[r], s[l]
		}
	}
	//考虑终止时的情况：
	//	1、l没动，r一路畅通与l重合  2、r和l合于中间某一处  3、r没动，l畅通至与r重合
	// 这三种情况下重合位置都有分界属性，除了第一种实质原地交换，后两种情况都指向了较小值。
	// 所以直接直接交换选定基准值和终止元素值即可，终止索引为期望分界标。
	s[low] = s[r]
	s[r] = pivot
	return r
}
func quickSort(s []int, low, high int) {
	if low < high {
		pivot := partition(s, low, high)
		quickSort(s, low, pivot-1)
		quickSort(s, pivot+1, high)
	}
}
func quickSort2(s []int, low, high int) {
	//当元素较少是时，直接做插入排序
	if high-low < 50 {
		//slice左闭右开
		insertSort(s[low : high+1])
		return
	}
	if low < high {
		p := partition(s, low, high)
		if p-low > high-p {
			quickSort2(s, low, p-1)
			quickSort2(s, p+1, high)
		} else {
			quickSort2(s, p+1, high)
			quickSort2(s, low, p-1)
		}
	}
}
func quickSort3(s []int, low, high int, wg *sync.WaitGroup) {
	//当元素较少时，直接做插入排序
	if high-low < 50 {
		insertSort(s)
		return
	}
	if low < high {
		p := partition(s, low, high)
		//如果前半段较长就先处理前半段
		if p-low > high-p {
			//先处理前半段
			if p-low > 10000 {
				//大数组并发处理
				wg.Add(1)
				go func(data []int) {
					defer wg.Done()
					quickSort3(data, low, p-1, wg)
				}(s)
			} else {
				//小数组直接快排
				quickSort2(s, low, p-1)
			}
			//再处理后半段
			if high-p > 10000 {
				wg.Add(1)
				go func(data []int) {
					defer wg.Done()
					quickSort3(data, p+1, high, wg)
				}(s)
			} else {
				quickSort2(s, p+1, high)
			}
		} else {
			//先处理后半段
			if high-p > 10000 {
				wg.Add(1)
				go func(data []int) {
					defer wg.Done()
					quickSort3(data, p+1, high, wg)
				}(s)
			} else {
				quickSort2(s, p+1, high)
			}
			//再处理前半段
			if p-low > 10000 {
				wg.Add(1)
				go func(data []int) {
					defer wg.Done()
					quickSort3(data, low, p-1, wg)
				}(s)
			} else {
				quickSort2(s, low, p-1)
			}
		}
	}
}
func QSort(s []int) {
	// 再套一层函数外壳方便传入waitGroup
	var wg sync.WaitGroup
	quickSort3(s, 0, len(s)-1, &wg)
	wg.Wait()
}
func partition3(s []int, begin, end int) (int, int) {
	if end-begin < 2 {
		return begin, begin
	}
	//通过交换，选定中间元素为基准值
	s[begin], s[(begin+end)/2] = s[(begin+end)/2], s[begin]
	var pivot = s[begin]
	//l,r分别是两段-开-边界：小部分的右边界和大部分的左边界,也是中间实际开始和最后终止的地方
	//mid是中间组的右边界，实际指向是待处理的数或最终的 大-1st
	l, r, mid := begin, end, begin+1
	for mid <= r {
		if s[mid] > pivot {
			//将大于pivot的数换到末尾，同时尾指针左移一位，紧接着继续处理换过来的这个数
			s[mid], s[r] = s[r], s[mid]
			r--
		} else if s[mid] < pivot {
			//将最左端的pivot数换到此时中间组的最右端
			s[mid], s[l] = s[l], s[mid]
			mid++
			l++
		} else {
			// 相等数直接扩充
			mid++
		}
	}
	return l, r
}
func quickSort32(s []int, begin, end int) {
	//当元素较少是时，直接做插入排序
	if end-begin < 30 {
		//slice左闭右开
		insertSort(s[begin : end+1])
		return
	}
	if begin < end {
		l, r := partition3(s, begin, end)
		quickSort32(s, begin, l-1)
		quickSort32(s, r+1, end)
	}
}

// 归并排序
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
func mergeSort1(s []int) []int {
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

// 桶排序

// 基数排序
