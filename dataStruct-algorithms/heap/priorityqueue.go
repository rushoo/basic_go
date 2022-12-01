package main

/*
优先级队列：
因为元素入堆时完成排序，一个基于堆的姓名队列则可视为一个优先级队列，其中
排列依据就是字符串的顺序。
*/

type PriorityQueue[T Ordered] struct {
	infoHeap Heap[T]
}

func (queue *PriorityQueue[T]) Push(item T) {
	queue.infoHeap.Insert(item)
}
func (queue *PriorityQueue[T]) Pop() T {
	returnValue := queue.infoHeap.Largest()
	queue.infoHeap.Remove()
	return returnValue
}
