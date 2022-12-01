package main

import (
	"fmt"
)

/*
heap:
堆是一种完全二叉树，对于其任意节点的值都大于或等于它的两个节点值(大顶堆，还有小顶堆正好相反)，
堆中最大的节点为根节点。完全二叉树的叶子节点从左到右依次填充。
需要注意的是，对于一个slice实现的堆，它的元素并不是有序的，仅有items[i] >= items[2i+1],items[2i+2]

用slice实现大顶堆：
参照完全二叉树的逻辑，对于任意位置节点i，它的子节点为2i+1、2i+2,父节点为(i-1)/2(奇数节点或者偶数节点)
*/

const size = 50_000_000

// 堆排序
func heapSort[T Ordered](input []T) []T {
	heap1 := NewHeap[T](input)
	descending := []T{} //降序数列
	ascending := []T{}  //升序数列
	for len(heap1.Items) > 0 {
		descending = append(descending, heap1.Largest())
		heap1.Remove()
	}
	for i := len(descending) - 1; i >= 0; i-- {
		ascending = append(ascending, descending[i])
	}
	//return descending		//按需返回想要的数列即可
	return ascending
}
func isSorted[T Ordered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
func main() {
	//slice1 := []int{100, 60, 80, 50, 30, 75, 40, 10, 35}
	//heap1 := NewHeap[int](slice1)
	//heap1.Insert(90)
	//fmt.Println("heap1 after inserting 90")
	//fmt.Println(heap1.Items)
	//fmt.Println("Largest item in heap: ", heap1.Largest())
	//
	//heap1.Remove()
	//fmt.Println("Removing largest item from heap yielding the heap: ")
	//fmt.Println(heap1.Items)
	//fmt.Println("Largest item in heap: ", heap1.Largest())
	//slice2 := []int{10, 35, 100, 80, 30, 75, 40, 50, 60}
	//heap2 := NewHeap[int](slice2)
	//heap2.Insert(90)
	//fmt.Println("heap2 with rearranged slice2 after inserting 90")
	//fmt.Println(heap2.Items)

	//data := make([]float64, size)
	//for i := 0; i < size; i++ {
	//	data[i] = 100.0 * rand.Float64()
	//}
	//start := time.Now()
	//largeSorted := heapSort[float64](data)
	//elapsed := time.Since(start)
	//fmt.Println("Time for heapSort of 50 million floats: ", elapsed)
	//if !isSorted[float64](largeSorted) {
	//	fmt.Println("largeSorted is not sorted.")
	//}

	myQueue := PriorityQueue[string]{}
	myQueue.Push("Helen")
	myQueue.Push("Apollo")
	myQueue.Push("Richard")
	myQueue.Push("Barbara")
	fmt.Println(myQueue)
	myQueue.Pop()
	fmt.Println(myQueue)
	myQueue.Push("Arlene")
	fmt.Println(myQueue)
	myQueue.Pop()
	myQueue.Pop()
	fmt.Println(myQueue)
}
