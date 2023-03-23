package main

import (
	"sort"
	"strconv"
)

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

输入: s = "anagram", t = "nagaram"
输出: true
输入: s = "rat", t = "car"
输出: false
说明: 你可以假设字符串只包含小写字母

思路：
定义一个数组，来记录字符串s里字符出现的次数。
*/
func isAnagram(s, t string) bool {
	//type rune = int32,直接使用的rune是一个整数值
	//使用r-'a'是为了让数值范围在0~26之间
	record := [26]int{}

	for _, r := range s {
		record[r-'a']++
	}
	for _, r := range t {
		record[r-'a']--
	}
	// 如果两字符串的频次相等，最终的结果数组一定每项都为零
	return record == [26]int{}
}

/*
349. 两个数组的交集
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的

思路：
输出结果中的每个元素一定是唯一的，也就是说输出的结果的去重的， 同时可以不考虑输出结果的顺序
可以使用map去重存数，再对照另一个数组去读取出来
*/
func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]struct{})
	//将num1的元素作为key保存在map中，value为空结构体
	for _, v := range nums1 {
		numMap[v] = struct{}{}
	}
	var result []int
	//使用num2中的元素去查map，命中则保存在结果slice中
	for _, v := range nums2 {
		if _, ok := numMap[v]; ok {
			result = append(result, v)
			delete(numMap, v)
		}
	}
	return result
}

/*
第202题. 快乐数
编写一个算法来判断一个数 n 是不是快乐数。
快乐数 定义为：

	对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
	然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
	如果这个过程 结果为1，那么这个数就是快乐数。

输入：n = 19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
输入：n = 2
输出：false
*/
func isHappy(n int) bool {
	//将每次参与计算的n值记录在map中以设置循环终止条件(重复出现)
	numMap := make(map[int]bool)
	for n != 1 && !numMap[n] {
		numMap[n] = true
		n = getSum(n)
	}
	return n == 1
}

// 计算一个正整数的每位平方和：将整数逐次取10的余数再缩小10倍
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

/*
1. 两数之和
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
输入：nums = [3,2,4], target = 6
输出：[1,2]
输入：nums = [3,3], target = 6
输出：[0,1]

思路：
从前往后读数组的元素值v，依次将 target-v 为key拿到map中比对，如果没有命中就把
这个数v作为key存入map，对应的value值为这个数v在数组中的下标值；如果命中则说明
target-v这个元素已经此前记录在了map中，输出对应的数组下标即可。
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if pre, ok := m[target-v]; ok {
			return []int{pre, i}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

/*
第454题.四数相加II
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能同时满足：
1、0 <= i, j, k, l < n
2、nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
为了使问题简单化，所有数组元素长度 0 ≤ N ≤ 500 。所有元素的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。

输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释，两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1

思路：
将所有的a+b的和作为key存入map，value为和出现的次数
再将后两个数组所有元素的和按-(c+d)去查map，存在则说明出现a+b+c+d=0，统计其出现的次数即value值。
将所有出现的次数累加。
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 将nums1和nums2数组任意两元素之和放到map中，value值是它们出现的次数
	m := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	// 将nums3和nums4数组任意两数之和取反去查map，累计出现的次数
	count := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-(v3 + v4)]
		}
	}
	return count
}

/*
383. 赎金信
给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串 ransom 能不能由
第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。
(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)
你可以假设两个字符串均只含有小写字母。

输入：ransomNote = "a", magazine = "b"
输出：false
输入：ransomNote = "aa", magazine = "ab"
输出：false
输入：ransomNote = "aa", magazine = "aab"
输出：true

思路：
将magazine中的每个字符作为key存入map，对应的value是字符出现的次数
再逐个地遍历ransom字符，map每次命中则将对应的value值减一，
当字符出现次数为0(未命中/出现次数不够)就输出false。
当判断是否命中时可能会更有效率。
改进：对于rune值，它本身就是int32整数类型，可以r-'a'作为下标使用数组
*/
func canConstruct2(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, v := range magazine {
		magazineMap[v]++
	}
	for _, v := range ransomNote {
		//if t, _ := magazineMap[v]; t > 0 {
		if t, ok := magazineMap[v]; ok && t > 0 {
			magazineMap[v]--
		} else {
			return false
		}
	}
	return true
}
func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}

/*
第15题. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
输入：nums = [-1,0,1,2,-1,-4]  输出：[[-1,-1,2],[-1,0,1]]
输入：nums = [0,1,1]  输出：[]
输入：nums = [0,0,0]  输出：[[0,0,0]]

思路：
双指针法，从头到位遍历数组，同时用两个指针指向当前后一个元素以及最后一个元素
当三者的和为0时，记下此时的值。
结果里还要去除元素相同的数组。
*/
func threeSum(nums []int) [][]int {
	nums = insertSort(nums)
	var res [][]int
	//records := make(map[string]struct{})
	for i, v := range nums {
		if v > 0 {
			break
		}
		if i > 0 && v == nums[i-1] {
			continue
		}
		for l := i + 1; l < len(nums)-1; l++ {
			if v+nums[l] > 0 {
				break
			}
			for r := len(nums) - 1; l < r; r-- {
				if nums[l]+nums[r]+v == 0 {
					tmp := []int{v, nums[l], nums[r]}
					if len(res) > 0 {
						if tmp[0] == res[len(res)-1][0] && tmp[1] == res[len(res)-1][1] {
							continue
						}
					}
					res = append(res, tmp)
				}
			}
		}
	}
	return res
}
func insertSort(data []int) []int {
	for i, v := range data {
		pre := i - 1
		for pre >= 0 && data[pre] > v {
			data[pre+1] = data[pre]
			pre -= 1
		}
		data[pre+1] = v
	}
	return data
}
func arrToString(data []int) string {
	var as string
	for _, v := range data {
		as = as + strconv.Itoa(v) + "_"
	}
	return as
}

/*
第18题. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：答案中不可以包含重复的四元组。
输入：nums = [1,0,-1,0,-2,2], target = 0   输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
输入：nums = [2,2,2,2,2], target = 8  	输出：[[2,2,2,2]]
*/
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		// if n1 > target { // 不能这样写,因为可能是负数
		// 	break
		// }
		if i > 0 && n1 == nums[i-1] { // 对nums[i]去重
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] { // 对nums[j]去重
				continue
			}
			l := j + 1
			r := len(nums) - 1
			for l < r {
				n3 := nums[l]
				n4 := nums[r]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l+1] { // 去重
						l++
					}
					for l < r && n4 == nums[r-1] { // 去重
						r--
					}
					// 找到答案时,双指针同时靠近
					r--
					l++
				}
			}
		}
	}
	return res
}
