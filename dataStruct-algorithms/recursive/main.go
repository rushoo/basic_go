package main

import (
	"fmt"
)

/*
使用递归求解：
1.求n！
4.给定一个字符串,将其翻转,例如输入: "abcd" , 输出"dcba"
5.求第n 个斐波那契数列元素
6.求解最大公约数
7.插入排序的递归形式，对前n个元素做插入排序
8.汉诺塔问题,将 1 ~ N 从A 移动到B, C作为辅助,要求一次只能移一个,全过程小的不能在大的下面 (最下面一个为N)
9、青蛙上楼梯，楼梯有n个台阶, 一个青蛙一次可以上1 , 2 或3 阶 , 实现一个方法, 计算该青蛙有多少种上完楼梯的方法
10.例如给出正整数 n=12345，希望以各位数的逆序形式输出，即输出54321。

*/

func PrintDigit(n int) {
	fmt.Print(n % 10)
	if n > 10 {
		PrintDigit(n / 10)
	}
}
func Upstairs(n int) int {
	//考虑边界，这里定义原地跳为1次
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	//青蛙只可能从这前三种情况跳上来的
	return Upstairs(n-1) + Upstairs(n-2) + Upstairs(n-3)
}
func HanoTower(n int, src, mid, des string) {
	//n块从A上面以B为辅移到C，
	//需要把前n-1块从以C为辅助从A移到B，再将第n块从A移到C，
	//此时问题变为以A为辅助把前n-1块从B移到C
	//问题边界，当n=1时，直接从src移动到des（此时实际的src可能是A，也可能是B，但都不重要）
	if n == 1 {
		fmt.Printf("Move %dst from %s to %s\n", n, src, des)
		return
	}
	HanoTower(n-1, src, des, mid)
	fmt.Printf("Move %dst from %s to %s\n", n, src, des)
	HanoTower(n-1, mid, src, des)
}
func InsertSort(items []int, n int) {
	len := len(items)
	if len < n || n < 2 {
		return
	}
	InsertSort(items, n-1) //前n-1个元素排序
	tmp := items[n-1]
	for n > 1 && tmp < items[n-2] {
		items[n-1] = items[n-2]
		n--
	}
	items[n-1] = tmp
}
func GreatestCommonDivisor(m, n int) int {
	if m*n == 0 {
		return 0
	}
	if m%n == 0 {
		return n
	}
	return GreatestCommonDivisor(m, m%n)
}
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func ReverseString(s string) string {
	length := len(s)
	if length == 0 || length == 1 {
		return s
	}
	end := length - 1
	// golang里string的底层是unit8,不接受直接赋值
	//比如这个fmt.Println("abcd"[3])输出字符串第四个字节的ASCII码
	return string(s[end]) + ReverseString(s[:end])
}
func Factorial(n int) int {
	//处理异常
	if n <= 0 {
		return 0
	}
	//边界
	if n == 1 {
		return 1
	}
	//缩小问题规模
	return n * Factorial(n-1)
}
func main() {
	//数字反向
	PrintDigit(12345)
	//跳楼梯
	fmt.Println("\n", Upstairs(4))
	//汉诺塔
	HanoTower(6, "A", "B", "C")
	//阶乘
	fmt.Println(Factorial(5))
	fmt.Println(ReverseString("abcdefghi"))
	fmt.Println(GreatestCommonDivisor(6, 4))

	num2 := []int{5, 12, 3, 4, -5}
	InsertSort(num2, 5)
	fmt.Println(num2)

}
