package main

type stackNode[T any] struct {
	value T
	next  *stackNode[T]
}
type nodeStack[T any] struct {
	first *stackNode[T]
}

func (stack *nodeStack[T]) Push(item T) {
	newNode := &stackNode[T]{item, nil}
	newNode.next = stack.first
	stack.first = newNode
}

func (stack *nodeStack[T]) Top() T {
	return stack.first.value
}

func (stack *nodeStack[T]) Pop() T {
	if stack.IsEmpty() {
		var zero T
		return zero
	}
	result := stack.first
	stack.first = stack.first.next
	return result.value
}

func (stack *nodeStack[T]) IsEmpty() bool {
	return stack.first == nil
}

type sliceStack[T any] struct {
	items []T
}

func (stack *sliceStack[T]) Push(item T) {
	stack.items = append(stack.items, item)
}

func (stack *sliceStack[T]) Top() T {
	len := len(stack.items)
	return stack.items[len-1]
}

func (stack *sliceStack[T]) Pop() T {
	len := len(stack.items)
	result := stack.items[len-1]
	stack.items = stack.items[:len-1]
	return result
}

func (stack *sliceStack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}
