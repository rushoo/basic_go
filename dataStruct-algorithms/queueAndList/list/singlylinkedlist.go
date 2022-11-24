package main

import "fmt"

// 节点
type SinglyNode[T DataType] struct {
	item T
	next *SinglyNode[T]
}

// 单链表,由节点组成，并包含长度信息
type SinglyLinkedList[T DataType] struct {
	first  *SinglyNode[T]
	length int
}

// append，尾插
func (sl *SinglyLinkedList[T]) Append(item T) {
	newNode := &SinglyNode[T]{item, nil}
	//	插入最后一个节点后
	if sl.length == 0 {
		sl.first = newNode
	} else {
		last := sl.first
		for last.next != nil {
			last = last.next
		}
		last.next = newNode
	}
	sl.length++
}

// insertAt,插入到指定位置，失败返回错误
func (sl *SinglyLinkedList[T]) InsertAt(index int, item T) error {
	preNode := sl.first
	newNode := &SinglyNode[T]{item, nil}
	moveToDest := func() { //指针移动至插入位置前节点
		for index > 1 {
			preNode = preNode.next
			index--
		}
	}
	switch {
	case index < 0 || index > sl.length:
		return fmt.Errorf("index out of range")
	case index == 0:
		newNode.next = preNode
		sl.first = newNode
	case index == sl.length:
		moveToDest()
		preNode.next = newNode
	default:
		moveToDest()
		newNode.next = preNode.next
		preNode.next = newNode
	}
	sl.length++
	return nil
}

// removeAt
func (sl *SinglyLinkedList[T]) RemoveAt(index int) (T, error) {
	var zero T
	var node *SinglyNode[T]
	preNode := sl.first
	moveToDest := func() { //指针移动至删除位置前节点
		for index > 1 {
			preNode = preNode.next
			index--
		}
	}
	switch {
	case index < 0 || index > sl.length-1:
		return zero, fmt.Errorf("index out of range")
	case index == 0:
		node = sl.first
		sl.first = sl.first.next
	case index == sl.length-1:
		moveToDest()
		node = preNode.next
		preNode.next = nil
	default:
		moveToDest()
		node = preNode.next
		preNode.next = preNode.next.next
	}
	sl.length--
	return node.item, nil
}

// reverse,将单链表反转,
// 就是先把原来的头节点摘出来，然后把每个节点从余链表依次摘出来指向此头节点
func (sl *SinglyLinkedList[T]) Reverse() {
	first := sl.first
	p := first.next //初始摘出节点

	first.next = nil //把头节点从链表摘出来
	for p != nil {
		tmp := p.next  //暂存下一节点位置
		p.next = first //余链表最左边节点摘出来指向原头节点
		first = p      //头节点后移
		p = tmp        //准备下一个摘出节点
	}
	sl.first = first //将链表头节点指向当前事实头节点
}

// indexOf，查找item出现的位置
func (sl *SinglyLinkedList[T]) IndexOf(item T) int {
	node := sl.first
	//遍历结束未找到，返回-1，找到了就返回当前index值
	for i := 0; i < sl.length-1; i++ {
		if node.item == item {
			return i
		}
		node = node.next
	}
	return -1
}

// itemAfter
func (sl *SinglyLinkedList[T]) ItemAfter(item T) T {
	node := sl.first
	//遍历结束未找到，返回-1，找到了就返回当前index值
	for i := 0; i < sl.length-1; i++ {
		// panic if node.next==nil，thus
		if node.next == nil { //返回对应类型的零值
			var zero T
			return zero
		}
		if node.item == item {
			break
		}
		node = node.next
	}
	return node.next.item
}

// items，返回链表所有的元素
func (sl *SinglyLinkedList[T]) Items() []T {
	var items []T
	node := sl.first
	//for node != nil {
	//	items = append(items, node.item)
	//	node = node.next
	//}
	for i := 0; i < sl.length; i++ {
		items = append(items, node.item)
		node = node.next
	}
	return items
}

// first
func (sl *SinglyLinkedList[T]) First() *SinglyNode[T] {
	return sl.first
}

// size
func (sl *SinglyLinkedList[T]) Size() int {
	return sl.length
}
