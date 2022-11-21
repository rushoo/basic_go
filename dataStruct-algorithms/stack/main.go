package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func main() {
//	//Time for 10 million Push() operations on nodeStack:  416.4546ms
//	//
//	//Time for 10 million Pop() operations on nodeStack:  20.7493ms
//	//
//	//Time for 10 million Push() operations on sliceStack:  137.2793ms
//	//
//	//Time for 10 million Pop() operations on sliceStack:  15.0125ms
//
//  const size = 10_000_000
//	nodeStack := nodestack.Stack[int]{}
//	sliceStack := slicestack.Stack[int]{}
//
//	start := time.Now()
//	for i := 0; i < size; i++ {
//		nodeStack.Push(i)
//	}
//	elapsed := time.Since(start)
//	fmt.Println("\nTime for 10 million Push() operations on nodeStack: ",
//		elapsed)
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		nodeStack.Pop()
//	}
//	elapsed = time.Since(start)
//	fmt.Println("\nTime for 10 million Pop() operations on nodeStack: ",
//		elapsed)
//	// Benchmark sliceStack
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		sliceStack.Push(i)
//	}
//	elapsed = time.Since(start)
//	fmt.Println("\nTime for 10 million Push() operations on sliceStack: ", elapsed)
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		sliceStack.Pop()
//	}
//	elapsed = time.Since(start)
//	fmt.Println("\nTime for 10 million Pop() operations on sliceStack: ", elapsed)
//}

//func main() {
//	postfix := infixToPostfix("a + (b - c) / (d * e)")
//	postfix2 := infixToPostfix("a * (b - c) / d * ( b + c - e)")
//	operandSlice := []float64{10, 5, 2, 4, 3}
//	assignFixValues(operandSlice)
//	result := evaluate(postfix)
//	result2 := evaluate(postfix2)
//	fmt.Println("postfix", postfix, "\tfunction evaluates to: ", result)
//	fmt.Println("postfix", postfix2, "\tfunction evaluates to: ", result2)
//
//	num := 1000000
//	b := convertToBinary(num)
//	fmt.Println("binary of ", num, "\t", b)
//}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := Point{1, 1}
	end := Point{38, 38}
	maze := NewMaze(40, 40, start, end, "maze.txt")
	newPos, _ := maze.StepAhead()
	time.Sleep(1 * time.Second)
	if newPos != None {
		fmt.Println(newPos)
	}
	for {
		if newPos == None || newPos.Equals(end) {
			break
		}
		newPos, _ = maze.StepAhead()
		time.Sleep(100 * time.Millisecond)
		if newPos != None {
			fmt.Println(newPos)
		}
	}
	if newPos.Equals(end) {
		fmt.Println("SUCCESS!  Reached ", end)
	}
}
