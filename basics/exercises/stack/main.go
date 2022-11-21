package main

import (
	"fmt"
	"strings"
)

type myStack[T any] struct {
	items []T
}

func (ms *myStack[T]) Push(item T) {
	ms.items = append(ms.items, item)
}
func (ms *myStack[T]) Pop() T {
	item := ms.items[len(ms.items)-1]
	ms.items = ms.items[:len(ms.items)-1]
	return item
}
func (ms myStack[T]) String() string {
	var str []string
	for i := 0; i < len(ms.items); i++ {
		s := fmt.Sprintf("myStack第%d个元素为：%v\n", i, ms.items[i])
		str = append(str, s)
	}
	return strings.Join(str, "")
}

// 多参数值
func printSome(values ...int) {
	for _, value := range values {
		fmt.Println(value)
	}
}
func main() {
	fmt.Println("Starting...")
	myStack := myStack[int]{}
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack)

	printSome(1, 1, 2, 3, 5, 8, 13)
}
