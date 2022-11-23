package main

//模拟航班系统乘客优先级
//实现：定义一个slice，容量为capacity，其中slice的每个元素都代表一个优先级队列，capacity代表计划的优先级数
//乘客的属性有姓名和优先级序数，1优先于2优先于3，以此类推
//插入操作，根据乘客的优先级参数插入到对应的nodeQueue中，并将计数器加1
//删除，依据优先级从高到低从对应的nodeQueue中出列一个元素，并将计数器减1
//判断空，判断计数器是否为0即可，这里把所有删除、查找前的判空逻辑都放在调用前，我认为对于队列的调用理应先判断是否为空的。

type Passenger struct {
	name     string
	priority int
}
type PriorityQueue[T any] struct {
	q    []nodeQueue[T]
	size int
}

func NewPriorityQueue[T any](PriorityNumber int) (pq PriorityQueue[T]) {
	pq.q = make([]nodeQueue[T], PriorityNumber)
	return pq
}
func (pq *PriorityQueue[T]) Insert(item T, priority int) {
	pq.q[priority-1].Insert(item)
	pq.size++
}
func (pq *PriorityQueue[T]) Remove() T {
	pq.size--
	var item T
	for i := 0; i < len(pq.q); i++ {
		//要是当前nodeQueue不为空就出一个，否则就往后再找
		if pq.q[i].Size() > 0 {
			item = pq.q[i].Remove()
			break
		}
	}
	return item
}
func (pq *PriorityQueue[T]) First() T {
	var item T
	for i := 0; i < len(pq.q); i++ {
		//要是当前nodeQueue不为空就出一个，否则就往后再找
		if pq.q[i].Size() > 0 {
			item = pq.q[i].First()
			break
		}
	}
	return item
}
func (pq *PriorityQueue[T]) Last() T {
	var item T
	for i := len(pq.q) - 1; i > 0; i-- {
		//要是当前nodeQueue不为空就出一个，否则就往后再找
		if pq.q[i].Size() > 0 {
			item = pq.q[i].last.item
			break
		}
	}
	return item
}
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.size == 0
}
