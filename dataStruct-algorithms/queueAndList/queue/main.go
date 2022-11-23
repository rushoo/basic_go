package main

import "fmt"

// We compare the performance of sliceQueue and nodeQueue
const size = 1_000_000

//func main() {
//	sliceQueue := queue[int]{}
//	nodeQueue := nodeQueue[int]{}
//	start := time.Now()
//	for i := 0; i < size; i++ {
//		sliceQueue.Insert(i)
//	}
//	elapsed := time.Since(start)
//	fmt.Println("Time for inserting 1 million ints in sliceQueue is", elapsed)
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		nodeQueue.Insert(i)
//	}
//	elapsed = time.Since(start)
//	fmt.Println("Time for inserting 1 million ints in nodeQueue is", elapsed)
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		sliceQueue.Remove()
//	}
//	elapsed = time.Since(start)
//	fmt.Println("Time for removing 1 million ints from sliceQueue is", elapsed)
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		nodeQueue.Remove()
//	}
//	elapsed = time.Since(start)
//	fmt.Println("Time for removing 1 million ints from nodeQueue is", elapsed)
//}
//func main() {
//	input := []int{9, 8, 7, 6, 5, 0, 1, 0, 6, 8}
//	output1 := MaxSubarrayBruteForce(input, 3)
//	fmt.Println("Output = ", output1)
//	output2 := MaxSubarrayUsingDeque(input, 3)
//	fmt.Println("Output = ", output2)
//}

func main() {
	airlineQueue := NewPriorityQueue[Passenger](3)
	passengers := []Passenger{
		{"Erika", 3}, {"Robert", 3}, {"Danielle", 3},
		{"Madison", 1}, {"Frederik", 1}, {"James", 2},
		{"Dante", 2}, {"Shelley", 3}}
	fmt.Println("Passengers: ", passengers)
	for i := 0; i < len(passengers); i++ {
		//根据不同的优先级将乘客插入到不同的队列
		airlineQueue.Insert(passengers[i], passengers[i].priority)
	}
	fmt.Println("First passenger in line: ", airlineQueue.First())
	turn := 3
	for turn > 0 {
		if airlineQueue.size > 0 {
			airlineQueue.Remove()
		}
		turn--
	}
	fmt.Println("First passenger in line: ", airlineQueue.First())
	fmt.Println("Last passenger in line: ", airlineQueue.Last())
}
