package main

import "fmt"

// 节点
type DoublyNode[T DataType] struct {
	item T
	next *DoublyNode[T]
	pre  *DoublyNode[T]
}

// 双链表
type DoublyLinkedList[T DataType] struct {
	first  *DoublyNode[T]
	last   *DoublyNode[T]
	length int
}

// append
func (dl *DoublyLinkedList[T]) Append(item T) {
	newNode := &DoublyNode[T]{item, nil, nil}
	//无中生有
	if dl.first == nil {
		dl.first = newNode
		dl.last = newNode
	} else {
		dl.last.next = newNode //接，newNode接入
		newNode.pre = dl.last  //化，化原来的last变newNode前节点
		dl.last = newNode      //   newNode变尾节点
	}
	dl.length++ //发，节点长度计数器加1
}

// insertAt,插入到指定位置，失败返回错误
func (dl *DoublyLinkedList[T]) InsertAt(index int, item T) error {
	//插入失败因为越界
	//其次讨论的是index要插入起始位置还是末尾还是之间
	newNode := &DoublyNode[T]{item, nil, nil}
	first := dl.first
	last := dl.last
	switch {
	case index < 0 || index > dl.length:
		return fmt.Errorf("index out of range")
	case index == 0:
		dl.first = newNode
		newNode.next = first
		if dl.length == 0 {
			dl.last = newNode
		} else {
			newNode = first.pre
		}
	case index == dl.length-1:
		last.next = newNode
		newNode.pre = last
		dl.last = newNode
	default:
		node := dl.first
		for index > 1 {
			node = node.next
			index--
		}
		node.next.pre = newNode
		newNode.next = node.next
		newNode.pre = node
		node.next = newNode
	}
	dl.length++
	return nil
}

// removeAt，删除指定位置节点，失败就返回错误
func (dl *DoublyLinkedList[T]) RemoveAt(index int) (T, error) {
	var zero T
	var node *DoublyNode[T]
	switch {
	case index < 0 || index > dl.length-1:
		return zero, fmt.Errorf("index out of range")
	case index == dl.length-1:
		node = dl.first
		dl.last = node.pre
		if index == 0 {
			//考虑了删除仅单节点的情况
			dl.first = dl.last
		}
	case index == 0:
		node = dl.first
		dl.first = node.next //不可能为nil
	default:
		preNode := dl.first
		for index > 1 {
			preNode = preNode.next
			index--
		}
		node = preNode.next
		preNode.next = preNode.next.next
		node.next.pre = preNode
	}
	dl.length--
	return node.item, nil
}

// indexOf，根据参数内容查找链表，命中时返回对应的下标
func (dl *DoublyLinkedList[T]) IndexOf(item T) int {
	node := dl.first
	for i := 0; i < dl.length-1; i++ {
		if node.item == item {
			return i
		}
	}
	return -1
}

// itemAfter，返回给定内容节点的后一个节点内容
func (dl *DoublyLinkedList[T]) ItemAfter(item T) T {
	node := dl.first
	for i := 0; i < dl.length-1; i++ {
		if node.item == item {
			break
		}
	}
	var zero T
	if node.next == nil {
		return zero
	}
	return node.next.item
}

// itemBefore，返回给定内容节点的前一个节点内容
func (dl *DoublyLinkedList[T]) ItemBefore(item T) T {
	node := dl.first
	for i := 0; i < dl.length-1; i++ {
		if node.item == item {
			break
		}
	}
	var zero T
	if node == dl.first {
		return zero
	}
	return node.pre.item
}

// items，返回链表所有内容(slice)
func (dl *DoublyLinkedList[T]) Item() []T {
	var items []T
	node := dl.first
	for i := 0; i < dl.length-1; i++ {
		items = append(items, node.item)
		node = node.next
	}
	return items
}

// ReverseList 翻转链表结构，依次摘出重连即可，也可以保留一份值再创建一个链表
func (dl *DoublyLinkedList[T]) ReverseList() {
	first := dl.first
	p := first.next

	first.next = nil
	for p != nil {
		tmp := p.next
		p.next = first
		first.pre = p
		first = p
		p = tmp
	}
	tmp := dl.first
	dl.first = dl.last
	dl.last = tmp
}

// Reverse 获取链表翻转后的数据，其实就是从后往前直接遍历获取item，而单链表需要逆转和取数两个步骤，
// 属于牺牲空间换时间的做法
func (dl *DoublyLinkedList[T]) Reverse() []T {
	var items []T
	node := dl.last
	for node != nil {
		items = append(items, node.item)
		node = node.pre
	}
	return items
}

// first，获取第一个节点
func (dl *DoublyLinkedList[T]) First() *DoublyNode[T] {
	return dl.first
}

// last，获取最后一个节点
func (dl *DoublyLinkedList[T]) Last() *DoublyNode[T] {
	return dl.last
}

// size，获取链表长度
func (dl *DoublyLinkedList[T]) Size() int {
	return dl.length
}
