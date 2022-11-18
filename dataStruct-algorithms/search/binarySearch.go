package main

import (
	"fmt"
	"runtime"
	"time"
)

type dataType interface {
	~int | ~float64 | ~string
}

const size = 100_000_000

// linear-search
func linearSearch[T dataType](data []T, target T) bool {
	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return true
		}
	}
	return false
}

// 在片段内查找目标数，根据区间参数调整片段值，若存在就在channel内发送通知查找成功
func searchSegment[T dataType](data []T, target T, left, right int, ch chan<- bool) {
	for i := left; i < right; i++ {
		if data[i] == target {
			ch <- true
		}
	}
	ch <- false
}
func concurrentLinearSearch[T dataType](data []T, target T) bool {
	//将大的slice按可用cpu数量切成n个小片，再起n个goroutine分别处理每个片段
	n := runtime.NumCPU()
	ch := make(chan bool)
	segSize := int(float64(len(data)) / float64(n))
	for index := 0; index < n; index++ {
		go searchSegment(data, target, segSize*index, segSize*index+segSize, ch)
	}
	num := 0
	//	考虑取整舍掉的部分
	total := n * segSize
	if n*total < len(data) {
		go searchSegment(data, target, total, len(data), ch)
		num-- //容量加一
	}
	//	统计收到的goroutine信号
	for {
		select {
		case value := <-ch:
			if value == true {
				return true
			}
			num++
			if num == n {
				return false
			}
		}
	}
	return false
}

// binary search：二分查找/折半查找,查找的数列应该是有序的
func binarySearch[T dataType](data []T, target T) bool {
	// 设置左、右区间，
	low, high := 0, len(data)-1
	for low < high {
		//每次查询之后以此时区间的median为轴向左/右收敛一位，最后low==high
		median := (low + high) / 2
		if data[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	//最后data[low]==data[high],判断最终这个数是否是target即可
	//但要考虑空data[],避免data[low]引起的panic
	if low == len(data) || data[low] != target {
		return false
	}
	return true
}
func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = float64(i) // is sorted
	}

	start := time.Now()
	data = append(data, -10.0)
	result := binarySearch[float64](data, -10.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using binarySearch = ", elapsed)
	fmt.Println("Result of search(-10.0) is ", result)

	start = time.Now()
	result = binarySearch[float64](data, float64(size/2))
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using binarySearch = ", elapsed)
	fmt.Println("Result of search is ", result)
}
