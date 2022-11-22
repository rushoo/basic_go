package main

type queue[T any] struct {
	items []T
}
type iterator[T any] struct {
	items []T
	next  int //index of item
}

func (q *queue[T]) Insert(item T) {
	//FIFO,尾插头取
	q.items = append(q.items, item)
}
func (q *queue[T]) Remove() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
func (q *queue[T]) First() T {
	return q.items[0]
}
func (q *queue[T]) Size() int {
	return len(q.items)
}

func (q *queue[T]) Range() iterator[T] {
	return iterator[T]{q.items, 0}
}
func (it *iterator[T]) Empty() bool {
	return it.next == len(it.items)
}
func (it *iterator[T]) Next() T {
	item := it.items[it.next]
	it.next++
	return item
}
