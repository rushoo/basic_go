package main

import (
	"fmt"
	"time"
)

// We compare the performance of sliceQueue and nodeQueue
const size = 1_000_000

func main() {
	sliceQueue := queue[int]{}
	nodeQueue := nodeQueue[int]{}
	start := time.Now()
	for i := 0; i < size; i++ {
		sliceQueue.Insert(i)
	}
	elapsed := time.Since(start)
	fmt.Println("Time for inserting 1 million ints in sliceQueue is", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		nodeQueue.Insert(i)
	}
	elapsed = time.Since(start)
	fmt.Println("Time for inserting 1 million ints in nodeQueue is", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceQueue.Remove()
	}
	elapsed = time.Since(start)
	fmt.Println("Time for removing 1 million ints from sliceQueue is", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		nodeQueue.Remove()
	}
	elapsed = time.Since(start)
	fmt.Println("Time for removing 1 million ints from nodeQueue is", elapsed)
}
