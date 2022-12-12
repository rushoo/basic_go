package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}
type MyLinkedList struct {
	head   *ListNode
	length int
}

func NewList() *MyLinkedList {
	return &MyLinkedList{}
}
func Constructor() MyLinkedList {
	return MyLinkedList{}
}
func (l *MyLinkedList) Insert(val int) {
	newNode := &ListNode{val: val}
	if l.length == 0 {
		l.head = newNode
	} else {
		res := l.head
		l.head = newNode
		newNode.Next = res
	}
	l.length++
}
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.length {
		return -1
	}
	node := l.head
	for index > 0 {
		node = node.Next
		index--
	}
	return node.val
}
func (l *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{val: val}
	res := l.head
	l.head = newNode
	newNode.Next = res
	l.length++
}
func (l *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{val: val}
	if l.length == 0 {
		l.head = newNode
	} else {
		node := l.head
		for node.Next != nil {
			node = node.Next
		}
		node.Next = newNode
	}
	l.length++
}
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.length {
		return
	}
	node := l.head
	newNode := &ListNode{val: val}
	if index == 0 {
		res := l.head
		l.head = newNode
		newNode.Next = res
	} else {
		for index > 1 {
			node = node.Next
			index--
		}
		res := node.Next
		node.Next = newNode
		newNode.Next = res
	}
	l.length++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > l.length-1 {
		return
	}
	node := l.head
	if index == 0 {
		res := node.Next
		l.head = res
	} else {
		for index > 1 {
			node = node.Next
			index--
		}
		node.Next = node.Next.Next
	}
	l.length--
}
func (l *MyLinkedList) ReverseList() {
	//将链表一分为二，不断地将第二部分链表的第一个节点头插入第一部分链表
	if l.length == 0 {
		return
	}
	head := l.head
	p := head.Next
	head.Next = nil
	for p != nil {
		tmp := p.Next
		p.Next = head
		head = p
		p = tmp
	}
	l.head = head
}
func (l *MyLinkedList) PrintList() {
	node := l.head
	for node != nil {
		fmt.Println(node.val)
		node = node.Next
	}
}
func (nd *ListNode) PrintList() {
	for nd != nil {
		fmt.Println(nd.val)
		nd = nd.Next
	}
}
func reverseList(head *ListNode) *ListNode {
	//将链表一分为二，不断地将第二部分链表的第一个节点头插入第一部分链表
	if head == nil {
		return nil
	}
	p := head.Next
	head.Next = nil
	for p != nil {
		tmp := p.Next
		p.Next = head
		head = p
		p = tmp
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	head = swapFrom(head)
	return head
}
func swapFrom(node *ListNode) *ListNode {
	if node == nil {
		return nil
	} else if node.Next == nil {
		return node
	}
	tmp := node.Next
	node.Next = tmp.Next
	tmp.Next = node
	//需要把状态传出来
	node.Next = swapFrom(node.Next)
	return tmp
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	vh := ListNode{Next: head}
	var count int
	node := vh.Next
	for node != nil {
		count++
		node = node.Next
	}
	m := count - n
	if m < 0 {
		return nil //越界
	}
	if m == 0 { //删除头节点
		vh.Next = vh.Next.Next
		return vh.Next
	} else {
		node = vh.Next
		for m > 1 {
			node = node.Next
			m--
		}
		node.Next = node.Next.Next
		return vh.Next
	}
}
