package main

/*
27. 移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

思路：
使用双指针，从头到尾遍历到底，遇到非值替换
可以直接使用for range，这样可以专注于值交换，且不用考虑指针越界的问题。
*/
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

/*
344.反转字符串
编写一个函数，其作用是将输入的字符串反转过来。
*/
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

/*
题目：剑指Offer 05.替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/
func replaceSpace(s string) string {
	// 字符串上不可以直接修改，这里使用for range按rune读取
	var ns []rune
	//var ns []byte，也可以按字节读取，空间节省一些
	for _, v := range s {
		if v == ' ' {
			ns = append(ns, '%', '2', '0')
		} else {
			ns = append(ns, v)
		}
	}
	return s
}

/*
151.翻转字符串里的单词
*/
//去除头尾及中间多余空格
//[]rune和byte[]空间占用是一样的，len()得到的长度是其储存的元素个数
func rmSpace(s string) []rune {
	// 字符串上不可以直接修改，这里使用for range按rune读取
	var rs []rune
	var tmp = ' ' //这样相当于跳过了开头的全部空格
	for _, v := range s {
		if v == tmp && v == ' ' {
		} else {
			rs = append(rs, v)
		}
		tmp = v
	}
	//去除最后一个空格
	if rs[len(rs)-1] == ' ' {
		rs = rs[:len(rs)-1]
	}
	return rs
}

// 字符串整个按字符翻转
func rvByRune(rs []rune) {
	l, r := 0, len(rs)-1
	for l < r {
		rs[l], rs[r] = rs[r], rs[l]
		l++
		r--
	}
}

func rvByWord(rs []rune) {
	i := 0
	for j, v := range rs {
		if v == ' ' {
			rvByRune(rs[i:j])
			i = j + 1
		}
	}
	// 翻转最后一个word
	rvByRune(rs[i:])
}
func reverseWords(s string) string {
	rs := rmSpace(s)
	rvByRune(rs)
	rvByWord(rs)
	return string(rs)
}

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }

思路：
1->2->3->4->5->O
1->O  2->3->4->5->O  ==>    2->1->O  3->4->5->O  ==>    3->2->1->O  4->5->O
==>    4->3->2->1->O  5->O  ==>    5->4->3->2->1->O

func reverseList(head *ListNode) *ListNode {
	p := head.Next
	head.Next = nil
	for p != nil {
		tmp := p
		p = tmp.Next
		tmp.Next = head
		head = tmp
	}
	return head
}
*/
