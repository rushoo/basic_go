package main

import (
	"math/rand"
	"testing"
)

func bubbleSort(s []int) {
	//	每次依次将第一个元素与后一个相比较
	if len(s) < 2 {
		return
	}
	m := len(s)
	for m > 1 {
		for i := 0; i < m-1; i++ {
			//如果当前元素大于后一个，则将两者交换
			//这样每轮循环结束后，本轮最大数一定在本轮最后
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
		//最后的几个元素已经排序，无需重复比较
		m--
	}
}
func TestBubbleSort(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, rand.Int())
	}
	fmt.Println(s)
}
