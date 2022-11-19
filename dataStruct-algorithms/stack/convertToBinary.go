package main

import "stack/slicestack"

// 使用辗转相除法
func convertToBinary(input int) (binary []int) {
	binary = []int{}
	sliceStack := slicestack.Stack[int]{}

	if input == 0 {
		sliceStack.Push(0)
	}
	for input > 1 {
		remainder := input % 2 //取余
		sliceStack.Push(remainder)

		input = input / 2 //取整重置
	}
	// 输入是1,或者商是1
	if input == 1 {
		sliceStack.Push(1)
	}
	for !sliceStack.IsEmpty() {
		binary = append(binary, sliceStack.Pop())
	}
	return binary
}
