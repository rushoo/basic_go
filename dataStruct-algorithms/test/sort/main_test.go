package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 辅助函数
func gens(n int) []int {
	//重置随机数，否则每次生成的序列是一样的
	rand.Seed(time.Now().UnixNano())
	//给定cap以优化空间
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Int()
	}
	return s
}
func isSorted(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}
	return true
}

func Test_bubbleSort(t *testing.T) {
	var s = gens(10000)
	bubbleSort(s)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func Test_selectSort(t *testing.T) {
	var s = gens(10000)
	selectSort(s)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func Test_insertSort(t *testing.T) {
	var s = gens(10000)
	insertSort(s)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func Test_shellSort(t *testing.T) {
	var s = gens(SIZE)
	shellSort(s)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func Test_countingSort(t *testing.T) {
	var s = gens(SIZE)
	countingSort(s, SIZE)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func BenchmarkBubbleSort(b *testing.B) {
	//数据准备前停止计时器
	b.StopTimer()
	var s = gens(50_000)
	b.StartTimer()
	bubbleSort(s)
	b.StopTimer()
}
func BenchmarkSelectSort(b *testing.B) {
	//数据准备前停止计时器
	b.StopTimer()
	var s = gens(50_000)
	b.StartTimer()
	selectSort(s)
	b.StopTimer()
}
func BenchmarkInsertSort(b *testing.B) {
	//数据准备前停止计时器
	b.StopTimer()
	var s = gens(50_000)
	b.StartTimer()
	insertSort(s)
	b.StopTimer()
}
func BenchmarkShellSort(b *testing.B) {
	b.StopTimer()
	var s = gens(SIZE)
	b.StartTimer()
	shellSort(s)
	b.StopTimer()
}
func BenchmarkCountingSort(b *testing.B) {
	b.StopTimer()
	var s = gens(SIZE)
	b.StartTimer()
	countingSort(s, SIZE)
	b.StopTimer()
}

const SIZE = 50_000
const QuickSIZE = 1_000_000

func Test_partition(t *testing.T) {
	var s = gens(SIZE)
	p := partition(s, 0, len(s)-1)
	var fail bool
	for _, v := range s[:p] {
		if v > s[p] {
			fail = true
		}
	}
	for _, v := range s[p+1:] {
		if v < s[p] {
			fail = true
		}
	}
	if fail {
		t.Errorf("Not right!")
	}
}
func Test_partition3(t *testing.T) {
	var s = gens(0)
	l, r := partition3(s, 0, len(s)-1)
	var fail bool
	pivot := s[l]
	for _, v := range s[:l] {
		//左边不应 >=pivot
		if v >= pivot {
			fail = true
		}
	}
	for _, v := range s[r+1:] {
		//右边不应 <=pivot
		if v <= pivot {
			fail = true
		}
	}
	for _, v := range s[l : r+1] {
		//中间不应 !=pivot
		if v != pivot {
			fail = true
		}
	}
	if fail {
		t.Errorf("Not right!")
	}
	fmt.Println("l:", l, "  r:", r)
}
func Test_quickSort(t *testing.T) {
	var s = gens(SIZE)
	quickSort(s, 0, len(s)-1)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func Test_quickSort2(t *testing.T) {
	var s = gens(SIZE)
	quickSort2(s, 0, len(s)-1)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func Test_quickSort3(t *testing.T) {
	var s = gens(SIZE)
	QSort(s)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func Test_quickSort32(t *testing.T) {
	var s = gens(SIZE)
	quickSort32(s, 0, len(s)-1)
	if !isSorted(s) {
		t.Errorf("Not sorted!")
	}
}
func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	var s = gens(QuickSIZE)
	b.StartTimer()
	quickSort(s, 0, len(s)-1)
	b.StopTimer()
}
func Benchmark2QuickSort(b *testing.B) {
	b.StopTimer()
	var s = gens(QuickSIZE)
	b.StartTimer()
	quickSort2(s, 0, len(s)-1)
	b.StopTimer()
}
func Benchmark3QuickSort(b *testing.B) {
	b.StopTimer()
	var s = gens(QuickSIZE)
	b.StartTimer()
	QSort(s)
	b.StopTimer()
}
func Benchmark_32_QuickSort(b *testing.B) {
	b.StopTimer()
	var s = gens(QuickSIZE)
	b.StartTimer()
	quickSort32(s, 0, len(s)-1)
	b.StopTimer()
}
