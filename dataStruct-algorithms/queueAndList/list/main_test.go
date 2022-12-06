package main

import "testing"

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
