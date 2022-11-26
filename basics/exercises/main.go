package main

import "fmt"

/*
6.Minimum and maximum
     1. Write a function that finds the maximum value in an int slice ([]int).
     2. Write a function that finds the minimum value in an int slice ([]int).
*/

type dataType interface {
	~int | ~float64 | ~string
}

func Map[T dataType](f func(T) T, list []T) []T {
	var result []T
	for _, v := range list {
		result = append(result, f(v))
	}
	return result
}

func Minimum(s []int) int {
	min := s[0]
	for _, v := range s {
		if min > v {
			min = v
		}
	}
	return min
}
func Maximum(s []int) int {
	max := s[0]
	for _, v := range s {
		if max < v {
			max = v
		}
	}
	return max
}
func Bubbleort(list []int) {
	for i := 0; i < len(list); i++ {
		//每轮扫描，依次将相邻元素比较大小并交换位置，确保每次的最大数都冒泡到最后
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	//	slice是引用类型，不需要返回新的slice
}

/*
8.Functions that return functions

	1.Write a function that returns a function that performs a +2 on integers.
	  Name the function plusTwo.
	2.Generalize the function from 1, and create a plusX(x) which returns functions
	  that add x to an integer.
*/
func plusX(x int) func(int) int {
	return func(a int) int {
		return a + x
	}
}
func plusTwo() func(int) int {
	return func(a int) int {
		return a + 2
	}
}
func main() {
	f1 := plusTwo()
	f2 := plusX(10)
	fmt.Println(f1(2), f2(2)) //4 12
}
