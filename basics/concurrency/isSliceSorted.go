package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"time"
)

/*
给定一个[]float64，判定其元素是否已经从小到大排序
1、将slice与排序后的slice比较，不一致则未排序
2、依次比较slice内相邻两个数，如有顺序不一致则未排序
3、
*/

const size = 1_000_000_000

func isSorted1(data []float64) bool {
	var data1 = make([]float64, len(data))
	copy(data1, data)
	sort.Float64s(data1)
	for i := 0; i < len(data); i++ {
		if data[i] != data1[i] {
			return false
		}
	}
	return true
}
func isSorted2(data []float64) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func isSegmentSorted(data []float64, a, b int, ch chan<- bool) {
	// 左区间a+1，右区间b-1
	for i := a + 1; i <= b-1; i++ {
		if data[i] < data[i-1] {
			ch <- false
		}
	}
	ch <- true
}

func isSorted3(data []float64) bool {
	ch := make(chan bool)
	//将分段数量设置为当前进程可用cpu数量
	segNum := runtime.NumCPU()
	//那么每段slice的元素个数应该是：[总数/分段数]
	segSize := int(float64(len(data)) / float64(segNum))
	var leftInterval, rightInterval int
	for index := 0; index < segNum; index++ {
		leftInterval = index * segSize
		rightInterval = index*segSize + segSize
		go isSegmentSorted(data, leftInterval, rightInterval, ch)
	}
	num := 0
	//考虑边缘误差，如55/8取整后应该还有七个数未参与排序
	if rightInterval < len(data) {
		num -= 1 //计数器容量应 +1
		//仅余一个数，就和前一个数验证顺序
		if rightInterval == len(data)-1 {
			go func() {
				//借用匿名函数再起一个goroutine往channel发送数据
				ch <- data[rightInterval-1] < data[len(data)-1]
			}()
		} else {
			// 两个数及以上就可直接排序
			go isSegmentSorted(data, rightInterval, len(data), ch)
		}
	}
	// 获取goroutines对片段的检查结果，若存在结果为false，那就是未排序
	// 设置计数器来统计执行结果为true的goroutines，当计数器的值达到segNum时表面所有goroutines检查成功
	for {
		select {
		case value := <-ch:
			if value == false {
				return false
			}
			num += 1
			if num == segNum {
				return true
			}
		}
	}
	return true
}

func main() {
	data1 := make([]float64, size)
	for i := 0; i < size; i++ {
		data1[i] = 100.0 * rand.Float64()
	}
	data2 := make([]float64, size)
	for i := 0; i < size; i++ {
		data2[i] = float64(2 * i)
	}

	start := time.Now()
	result := isSorted2(data1)
	fmt.Println("\n data1 Sorted ? : ", result)
	fmt.Println("elapsed using sorted2", time.Since(start))

	start = time.Now()
	result = isSorted2(data2)
	fmt.Println("data2 Sorted? : ", result)
	fmt.Println("elapsed using sorted2:", time.Since(start))

	start = time.Now()
	result = isSorted3(data1)
	fmt.Println("\ndata1 Sorted? : ", result)
	fmt.Println("elapsed using concurrent sorted3", time.Since(start))

	start = time.Now()
	result = isSorted3(data2)
	fmt.Println("data2 Sorted? : ", result)
	fmt.Println("elapsed using concurrent sorted3:", time.Since(start))
}
