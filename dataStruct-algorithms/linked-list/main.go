package main

//关于链表的操作最应该注意的是状态的连续性，这种状态的连续性是通过显式的显式的指针指向来实现的。
//golang里是值传递，隐式的指针变化可能无法得到预期结果。
/*
203. 移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 ListNode.val == val 的节点，并返回 新的头节点 。
思路：
删除一个节点，意味着原本指向该节点的指针指向了下一个节点，需要注意链表/节点状态是否为空，避免panic

707.设计链表
在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

206.反转链表
题意：反转一个单链表。
示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL
思路：将链表一分为二，不断地将第二部分链表的第一个节点头插入第一部分链表

24. 两两交换链表中的节点
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。（只能进行节点交换）。
思路：如果不要求节点交换，可以直接交换两个节点的数据，这里可以用一个递归程序完成从左到右依次的两两交换。

19.删除链表的倒数第N个节点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
思路：获取链表长度，然后根据要求遍历到对应的节点，并将指向该节点的指针指向下一个节点

160.链表相交
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
思路：
先求出两个链表的长度，假设长度分别为15、10，那么先让长度为15的链表走5步，然后两链表同时向后遍历并不断比较节点指针是否相等，
相等则返回当前的节点计数值，否则最终返回nil

142.环形链表II
题意： 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
思路：快慢指针法

*/
func main() {
	linkedList := Constructor()
	//linkedList.Insert(9)
	//linkedList.Insert(8)
	//linkedList.Insert(7)
	//linkedList.Insert(6)
	//linkedList.Insert(5)
	//linkedList.Insert(4)
	//linkedList.Insert(3)
	linkedList.Insert(2)
	linkedList.Insert(1)
	//linkedList.PrintList()
	//linkedList.ReverseList()
	//linkedList.PrintList()

	//head := reverseList(linkedList.head)
	head := removeNthFromEnd(linkedList.head, 1)
	head.PrintList()

}
