package main

type Deque[T any] struct {
	items []T
}

// InsertFront
func (dq *Deque[T]) InsertFront(item T) {
	dq.items = append(dq.items, item) //expand itemSlice
	for i := len(dq.items) - 1; i > 0; i-- {
		dq.items[i] = dq.items[i-1]
	}
	dq.items[0] = item
}

// InsertBack
func (dq *Deque[T]) InsertBack(item T) {
	dq.items = append(dq.items, item)
}

func (dq *Deque[T]) First() T {
	return dq.items[0]
}
func (dq *Deque[T]) Last() T {
	return dq.items[len(dq.items)-1]
}
func (dq *Deque[T]) RemoveFirst() T {
	item := dq.items[0]
	dq.items = dq.items[1:]
	return item
}
func (dq *Deque[T]) RemoveLast() T {
	len := len(dq.items)
	item := dq.items[len-1]
	dq.items = dq.items[:len-1]
	return item
}

func (dq *Deque[T]) Empty() bool {
	return len(dq.items) == 0
}

//func main() {
//	myDeque := Deque[int]{}
//	myDeque.InsertFront(5)
//	myDeque.InsertBack(10)
//	myDeque.InsertFront(2)
//	myDeque.InsertBack(12) // 2 5 10 12
//	fmt.Println("myDeque.First() = ", myDeque.First())
//	fmt.Println("myDeque.Last() = ", myDeque.Last())
//	myDeque.RemoveLast()
//	myDeque.RemoveFirst()
//	fmt.Println("myDeque.First() = ", myDeque.First())
//	fmt.Println("myDeque.Last() = ", myDeque.Last())
//
//	input := []int{3, 1, 6, 4, 2, 10, 5, 9}
//	output := MaxSubarray(input, 3)
//	fmt.Println("Output = ", output)
//}

func MaxSubarrayBruteForce(input []int, k int) (output []int) {
	//暴力算法，相当于每次都遍历比较找出最大值
	for first := 0; first <= len(input)-k; first++ {
		max := input[first]
		for second := 0; second < k; second++ {
			if input[first+second] > max {
				max = input[first+second]
			}
		}
		output = append(output, max)
	}
	return output
}

// 滑动窗口取最大值，左出右进
func MaxSubarrayUsingDeque(input []int, k int) (output []int) {
	//从前k个数中的最大值往后，取到一个最大的降序数列
	//这样下一个数组的最大数仅需拿后移时新的元素去更新之前的降序数列即可获取
	//但要确保第i次取子数组最大值时子数组窗口向右滑动了i-1次
	deque := Deque[int]{}
	var index int

	// 将窗口子数组的最大降序数列的索引依次入列，通过匿名函数封装函数内重复的逻辑代码
	LOOP := func() {
		for {
			if deque.Empty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}
	//初始，取第一个子数组窗口，构造数组降序队列索引
	for index = 0; index < k; index++ {
		LOOP()
	}

	for index < len(input) {
		//将第一个子数组中最大值放入output
		output = append(output, input[deque.First()])
		for {
			// 子数组窗口（逻辑上）向右滑动一位
			// 例如[4,3,2,1]的子数组窗口由[4,3,2]滑动到[3,2,1],
			// 那么逻辑上就是子数组删除了最左边第一个元素4，然后通过下面的loop向右扩展一位.
			if deque.Empty() || deque.First() > index-k {
				break
			}
			deque.RemoveFirst()
		}
		LOOP()  //完成本次最大数入列
		index++ //准备下一次
	}
	output = append(output, input[deque.First()])
	return output
}
