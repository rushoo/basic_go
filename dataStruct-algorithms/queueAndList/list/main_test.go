package main

import "testing"

/*
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
*/
func TestAppend(t *testing.T) {
	/*
		测试方法：
		构造一个doublyLinked_list,然后添加两个元素，
		验证元素是否正确添加
	*/
	list := new(DoublyLinkedList[string])
	list.Append("hello")
	list.Append("aloha")
	
}
