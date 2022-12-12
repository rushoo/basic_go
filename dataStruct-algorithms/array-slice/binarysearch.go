package main

import (
	"fmt"
)

/*
todo===============================================================================
76.最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
输入：s = "ADOBECODEBANC", t = "ABC" 	输出："BANC"
输入：s = "a", t = "a"			输出："a"
输入: s = "a", t = "aa" 	输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，因此没有符合条件的子字符串，返回空字符串。
=========================================================================

704二分查找  	-SearchInsertPosition
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
你可以假设 nums 中的所有元素是不重复的,n 将在 [1, 10000]之间,nums 的每个元素都将在 [-9999, 9999]之间。

示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

思路:
这道题目的前提是数组为有序数组，同时题目还强调数组中无重复元素，因为一旦有重复元素，使用二分查找法返回的元素下标可能不是唯一的，这些都是使用二分法的前提条件.
注意最好不要使用递归，因为递归还得考虑跟踪变化的中间值。

35.搜索插入位置	-SearchInsertPosition
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
示例 2:
输入: nums = [1,3,5,6], target = 2
输出: 1

思路：
整个就是这四种情况，在二分查找基础上修改即可。
1、目标值在数组所有元素之前，结束时low=0,high=-1
2、目标值等于数组中某一个元素，直接返回mid
3、目标值插入数组中的位置，当low==high时候，

	如果此时指向的元素大于target，那么这个位置将插入target，而此时high还要做一次自减结束循环，
	否则就该插入它的右边，对应着的是low还要做一次自增来结束循环。

4、目标值在数组所有元素之后，结束是low=high+1
这四种情况下插入的位置都应该是high++

34. 在排序数组中查找元素的第一个和最后一个位置	-SearchRange
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。如果数组中不存在目标值 target，返回 [-1, -1]。
要求时间复杂度为O（logn）
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

思路：
寻找target在数组里的左右边界，有如下三种情况：
target 在数组范围的右边或者左边，例如数组{3, 4, 5}，target为2或者数组{3, 4, 5},target为6，此时应该返回{-1, -1}
target 在数组范围中，且数组中不存在target，例如数组{3,6,7},target为5，此时应该返回{-1, -1}
target 在数组范围中，且数组中存在target，例如数组{3,6,7},target为6，此时应该返回{1, 1}
接下来，再采用二分法去寻找左边界，和右边界。
就是在原来折半查找的前提下，没找到就输出-1，-1，找到了再左右遍历拓展边界再输出。

69. x 的平方根	-MySqrt
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
输入：x = 4 输出：2
输入：x = 8 输出：2 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

思路：
一个数的平方根怎么求：
二分法：以x^2=10为例，可能的范围是[0, 10]，取中间的值5，5*5>10,缩小范围为[0, 5]，接着取2.5，小了，取[2.5，5]……
牛顿法：(前提是构造函数有根且可导)就是不断地用切线方程去逼近
以x^2=10为例，构造方程f(x)=x^2-10,求f(x)=0时的正数解。将函数泰勒展开，舍弃高阶导，对于原来函数上的任意一点Xo，得到一个近似的切线方程为F(X)=f(Xo)+f'(Xo)(X-Xo)。
现在我们观察这个切线方程在是否可以近似的表示原方程，也就是在区间内当切线方程等于0时原函数有多接近等于0。令F(X)=f(Xo)+f'(Xo)(X-Xo)=0 ==>  X=Xo-f(Xo)/f'(Xo).
现在的误差为f(x)-F(x)=f(Xo)/f'(Xo),若大于误差范围则需要将F(x)继续逼近，取x1=Xo-f(Xo)/f'(Xo),下一次的误差就是f(x1)/f'(x1)...直到误差可忽略

367. 有效的完全平方数	-isPerfectSquare1
给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
思路：
1、二分法
2、利用 1+3+5+7+9+…+(2n-1)=n^2，即完全平方数肯定是前n个连续奇数的和

27. 移除元素	-removeElement
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
思路：
双指针，从零出发，一指针不断地向后移动遍历，若数据不为val，则将另一指针的数据cover，否则不处理。遍历结束相当于跳过了所有可能为val的数据。
由于两个指针总是相对同步的，其中一个并没有什么实际意义，可以用for range忽略。

26.删除排序数组中的重复项 	-removeDuplicates
思路：第一项默认保留，后一项与前一项相比较，跳过重复。

283. 移动零	-moveZeroes
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
思路：
遇到零值，计数器加一，再有非零值就与前面第一个零值交换。

844. 比较含退格的字符串	-backspaceCompare
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
思路：用rune slice模拟栈的操作分别处理两个字符串，非空遇到退格时pop，最后再比较两者的字符串。

977. 有序数组的平方	-sortedSquares
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
思路：平方后的数组从中到左和从中到右都是有序的，可以类似归并排序那样，相对顺序地比较两部分的值。

209.长度最小的子数组	-minSubArrayLen
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。
如果不存在符合条件的子数组，返回 0。
思路：某种形式的滑动窗口
1、右标开始右移直到寻找到第一个数组，或者遍历结束
2、当寻得第一个数组时，记录下当前数组的长度，并开始操作左标
3、左标每次向右移动直到数组失效，此时记录上一次数组有效长度，并继续右标、左标移动

904. 水果成篮	-totalFruit
给定一个混合类型元素数列，求它的子类型为2的最大子数列最多可包含元素的个数
输入：fruits = [1,2,1]		输出：3
输入：fruits = [0,1,2,2]		输出：3
输入：fruits = [1,2,3,2,2]	输出：4
输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]	输出：5
思路：
构造滑动窗口，有左右指针分别为窗口的边界，当类型数量少于2时，右指针带着窗口向右滑动。
当遇到第三种元素时，记录下此时窗口长度，然后左指针收缩到右指针 连续相同 最左端元素。
然后开始接受第三种元素。
问题：如何比较识别第三种不同的元素。

59.螺旋矩阵II  generateMatrix
思路：
分上右下左上四个阶段循环遍历对二维数组赋值即可。若是奇数数组可能最中间无法遍历到，需要额外检查并赋值。
*/
func generateMatrix(m, n int) [][]int {
	//	当n为偶数时，可以刚好遍历结束，n为奇数时还剩中间一个元素
	//遍历一个m*n矩阵,m行n列
	table := make([][]int, m)
	for t := 0; t < m; t++ {
		table[t] = make([]int, n)
	}
	i, j := 0, 0 //i行j列
	c := 1       //游标，螺旋前进
	for c < m*n {
		fmt.Printf("第%d行向右,从第%d列向%d列(不含)： ", i, j, n-j-1)
		for t := i; t < n-j-1; t++ {
			table[i][t] = c
			print(table[i][t], " ")
			c++
		}
		fmt.Println()

		fmt.Printf("第%d列向下,从第%d行向第%d行(不含)： ", n-j-1, i, m-i-1)
		for t := i; t < m-i-1; t++ {
			//第n-i-1列,第i行到最后一行
			table[t][n-j-1] = c
			print(table[t][n-j-1], " ")
			c++
		}
		fmt.Println()

		fmt.Printf("第%d行向左,从第%d列到第%d列(不含)： ", m-i-1, n-j-1, j)
		for t := n - j - 1; t > j; t-- {
			table[m-i-1][t] = c
			print(table[m-i-1][t], " ")
			c++
		}
		fmt.Println()

		fmt.Printf("第%d列向上,从第%d行到第%d行(不含)： ", j, m-i-1, i)
		for t := m - i - 1; t > i; t-- {
			table[t][j] = c
			print(table[t][j], " ")
			c++
		}
		fmt.Println()
		fmt.Println("c的值为: ", c)
		i++
		j++
	}
	if c == m*n {
		//说明总计遍历了m*n-1个元素
		table[m/2][n/2] = m * n
	}
	return table
}
func generateMatrix2(n int) [][]int {
	//	当n为偶数时，可以刚好遍历结束，n为奇数时还剩中间一个元素
	//遍历一个m*n矩阵,m行n列
	table := make([][]int, n)
	for t := 0; t < n; t++ {
		table[t] = make([]int, n)
	}
	i, j := 0, 0 //i行j列
	c := 1       //游标，螺旋前进
	for c < n*n {
		fmt.Printf("第%d行向右,从第%d列向%d列(不含)： ", i, j, n-j-1)
		for t := i; t < n-j-1; t++ {
			table[i][t] = c
			print(table[i][t], " ")
			c++
		}
		fmt.Println()

		fmt.Printf("第%d列向下,从第%d行向第%d行(不含)： ", n-j-1, i, n-i-1)
		for t := i; t < n-i-1; t++ {
			//第n-i-1列,第i行到最后一行
			table[t][n-j-1] = c
			print(table[t][n-j-1], " ")
			c++
		}
		fmt.Println()

		fmt.Printf("第%d行向左,从第%d列到第%d列(不含)： ", n-i-1, n-j-1, j)
		for t := n - j - 1; t > j; t-- {
			table[n-i-1][t] = c
			print(table[n-i-1][t], " ")
			c++
		}
		fmt.Println()

		fmt.Printf("第%d列向上,从第%d行到第%d行(不含)： ", j, n-i-1, i)
		for t := n - i - 1; t > i; t-- {
			table[t][j] = c
			print(table[t][j], " ")
			c++
		}
		fmt.Println()
		fmt.Println("c的值为: ", c)
		i++
		j++
	}
	if c == n*n {
		//说明总计遍历了m*n-1个元素
		table[n/2][n/2] = n * n
	}
	return table
}
func totalFruit(fruits []int) int {
	if len(fruits) < 2 {
		return len(fruits)
	}
	//窗口左右指针、最大长度、类型记录
	var L, R, MaxLen int
	var Ts []int
	for R < len(fruits) {
		/*
			类型变更讨论：
			1、直接后入，空 + 唯一不重叠
			3、交换顺序，等于[0]但不等于[1]
			3、前出后入，第三种元素
			4、直接忽略，default
		*/
		switch {
		case len(Ts) == 0 || len(Ts) == 1 && fruits[R] != Ts[0]:
			Ts = append(Ts, fruits[R])
		case len(Ts) == 2 && fruits[R] == Ts[0] && fruits[R] != Ts[1]:
			Ts[0], Ts[1] = Ts[1], Ts[0]
		case len(Ts) == 2 && fruits[R] != Ts[0] && fruits[R] != Ts[1]:
			//记录最新最长，调整左指针，指向最远连续相同元素
			if R-L > MaxLen {
				MaxLen = R - L
			}
			L = R - 1
			for L > 0 && fruits[L-1] == fruits[R-1] {
				L--
			}
			//	模拟出列,更新元素类型
			Ts = Ts[1:]
			Ts = append(Ts, fruits[R])
		default:
		}
		R++
	}
	//最后两种元素个数统计
	if R-L > MaxLen {
		MaxLen = R - L
	}
	return MaxLen
}
func minSubArrayLen(target int, nums []int) int {
	/*
		1、右标开始右移直到寻找到第一个数组，或者遍历结束
		2、当寻得第一个数组时，记录下当前数组的长度，并开始操作左标
		3、左标每次向右移动直到数组失效，此时记录上一次数组有效长度，并继续右标、左标移动
	*/
	i, sum := 0, 0
	result := len(nums) //用这个函数变量将for循环中临时的有效长度传递出来
	for j := 0; j < len(nums); {
		sum += nums[j]
		//外层for循环会在此等待内层循环结束
		for sum >= target {
			tmp := j - i + 1 //找到第一个数组,记录长度
			if tmp < result {
				result = tmp
			}
			sum -= nums[i]
			i++
		}
		j++
	}
	//未经for sum >= target
	if i == 0 {
		return 0
	}
	return result
}

func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	length := len(nums)
	var res = make([]int, length)
	var l, r = 0, length - 1
	for length > 0 {
		if nums[l] > nums[r] {
			res[length-1] = nums[l]
			l++
		} else {
			res[length-1] = nums[r]
			r--
		}
		length--
	}
	return res
}
func backspaceCompare(s string, t string) bool {
	var ts, tt []rune
	for _, v := range s {
		if v != '#' {
			ts = append(ts, v)
		} else if len(ts) > 0 {
			ts = ts[:len(ts)-1]
		}
	}
	for _, v := range t {
		if v != '#' {
			tt = append(tt, v)
		} else if len(tt) > 0 {
			tt = tt[:len(tt)-1]
		}
	}
	return string(ts) == string(tt)
}
func moveZeroes(nums []int) {
	var count int
	for i, v := range nums {
		if v == 0 {
			count++
		} else if count > 0 {
			// nums[i-count], nums[i] = nums[i], nums[i-count]  直接赋值避免开辟临时内存
			nums[i-count] = nums[i]
			nums[i] = 0
		}
	}
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//nums[0] = nums[0]
	var i, res = 1, 1
	for i < len(nums) {
		if nums[i] != nums[i-1] {
			nums[res] = nums[i]
			res++
		}
		i++
	}
	nums = nums[:res]
	return res
}
func removeElement(nums []int, val int) int {
	res := 0
	for _, v := range nums {
		if v != val {
			nums[res] = v
			res++
		}
	}
	nums = nums[:res]
	return res
}
func isPerfectSquare1(num int) bool {
	if num == 1 {
		return true
	}
	L, R := 1, (num+1)>>1
	for L <= R {
		mid := (L + R) >> 1
		s := mid * mid
		if s > num {
			R = mid - 1
		} else if s < num {
			L = mid + 1
		} else {
			return true
		}
	}
	return false
}
func isPerfectSquare2(num int) bool {
	i := 1
	for num > 0 {
		num = num - i
		i = i + 2
	}
	return num == 0
}
func MySqrt(m int) int {
	limit := 0.01                 //精度
	x := float64(m)               //初始值
	var offset = float64(m) / 2.0 //offset=f(Xo)/f'(Xo),初始值以m计
	for offset > limit {
		offset = (x*x - float64(m)) / (2 * x)
		x = x - offset //X=Xo-f(Xo)/f'(Xo)
		fmt.Println(x)
	}
	return int(x)
}
func MySqrt2(x int) int {
	//利用整数特性的取巧做法，时间复杂度高，毕竟要求返回的仅仅是整数值而已
	i := 1
	for i*i <= x {
		fmt.Println(i)
		i++
	}
	return i - 1
}

func SearchRange[T dataType](data []T, target T) []int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) >> 1
		num := data[mid]
		switch {
		case target < num:
			high = mid - 1
		case target > num:
			low = mid + 1
		default:
			left, right := mid, mid
			for left > 0 && data[left-1] == target {
				left--
			}
			for right < len(data)-1 && data[right+1] == target {
				right++
			}
			return []int{left, right}
		}
	}
	return []int{-1, -1}
}

func BinarySearch[T dataType](data []T, target T) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) >> 1
		num := data[mid]
		switch {
		case target < num:
			high = mid - 1
		case target > num:
			low = mid + 1
		default:
			return mid
		}
	}
	return -1
}
func SearchInsertPosition[T dataType](data []T, target T) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) >> 1
		switch {
		case target < data[mid]:
			high = mid - 1
		case target > data[mid]:
			low = mid + 1
		default:
			return mid
		}
	}
	return high + 1
}
