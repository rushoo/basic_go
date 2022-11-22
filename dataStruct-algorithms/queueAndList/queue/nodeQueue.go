package main

// 链式队列，应该有头尾双指针而实现FIFO，其中出列头指针后移一位，入列尾指针后扩一位

type nodeQueue[T any] struct {
	first, last *node[T]
	length      int
}
type node[T any] struct {
	item T
	next *node[T]
}
type nodeIterator[T any] struct {
	next *node[T]
}

func newNode[T any](item T) *node[T] {
	return &node[T]{item, nil}
}

func (nq *nodeQueue[T]) Insert(item T) {
	nn := newNode(item)
	if nq.length == 0 {
		nq.first = nn
		nq.last = nq.first
	} else {
		nq.last.next = nn
		nq.last = nn
	}
	nq.length++
}
func (nq *nodeQueue[T]) Remove() T {
	//判空放到代码逻辑里
	outNode := nq.first
	nq.first = nq.first.next
	nq.length--
	return outNode.item
}
func (nq *nodeQueue[T]) Size() int {
	return nq.length
}
func (nq *nodeQueue[T]) First() T {
	return nq.first.item
}
func (nq *nodeQueue[T]) Range() nodeIterator[T] {
	return nodeIterator[T]{nq.first}
}
func (iterator *nodeIterator[T]) Empty() bool {
	return iterator.next == nil
}
func (iterator *nodeIterator[T]) Next() T {
	item := iterator.next.item
	if iterator.next != nil {
		iterator.next = iterator.next.next
	}
	return item
}
