package main

import (
	"fmt"
	"time"
)

/*
传统方式递归计算一个斐波那契数列的某项值是比较低效的，
若能记住之前的值，求解第n个元素值只需要将前两个元素值相加即可。
动态规划最核心的思想，就在于拆分子问题，记住过往，减少重复计算。
*/
func fib(n int64) int64 {
	//0、1、1、2、3、5、8、13、21、34
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
func computeFromCache(n int, cache map[int]int64) int64 {
	if val, found := cache[n]; found {
		return val
	}
	cache[n] = computeFromCache(n-1, cache) + computeFromCache(n-2, cache)
	return cache[n]
}

// 自顶向下，根据参数递归向下去计算并存值，再取值
func fibonacciTopDown(n int) int64 {
	earlyCases := map[int]int64{
		0: 0,
		1: 1,
	}
	return computeFromCache(n, earlyCases)
}

// 自底向上，循环计算前两个值保存到数组，直到参数位置
func fibonacciBottomUp(n int) int64 {
	table := []int64{0, 1}
	for i := 2; i <= n; i++ {
		table = append(table, table[i-1]+table[i-2])
	}
	return table[n]
}

func main() {
	fmt.Println("fib(7) = ", fibonacciTopDown(7))
	start := time.Now()
	fib45 := fibonacciTopDown(45)
	elapsed := time.Since(start)
	fmt.Println("Value of fibonacciTopDown(45): ", fib45)
	fmt.Println("Computation time: ", elapsed)

	fmt.Println("fib(7) = ", fibonacciBottomUp(7))
	start = time.Now()
	fib45 = fibonacciBottomUp(45)
	elapsed = time.Since(start)
	fmt.Println("\nValue of fibonacciBottomUp(45): ", fib45)
	fmt.Println("Computation time: ", elapsed)

	fmt.Println("fib(7) = ", fib(7))
	start = time.Now()
	fib45 = fib(45)
	elapsed = time.Since(start)
	fmt.Println("\nValue of Fib(45): ", fib45)
	fmt.Println("Computation time: ", elapsed)

	var volume = []int{4, 1, 2, 3, 5} //物品体积和价值
	var worth = []int{5, 2, 4, 4, 5}
	var pkgVol = 5 //背包容量
	res := maxWorthGoodsOnce(volume, worth, pkgVol)
	println(res)

	x := "algorithms"
	y := "alchemist"
	lcs := LongestCommonSequence([]rune(x), []rune(y))
	fmt.Println(lcs)
}
