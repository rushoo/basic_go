package main

import "fmt"

/*
704二分查找
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

35.搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
*/

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
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	data := []int{5}
	fmt.Println(BinarySearch(nums, 9))
	fmt.Println(BinarySearch(data, 5))

}
