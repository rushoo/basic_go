### 链表
太基础的东西太久不看就容易忘记，这里简单重写下链表的实现，包括单链表、双链表

#### #单链表
单链表及节点的定义。链表的头指针和尾指针分别指向链表的第一个节点和最后一个节点，链表长度属性指明链表所含的元素个数。     
和数组类似，从0开始计数，在0位置插入一个节点就意味着插入头节点位置。链表的每一个节点含两个属性，一个表示当前承载的    
元素的值，另一个指针指向下一个节点。空链表的头指针与尾指针指向都为空。   
```
type List struct {
	first  *Node
	last   *Node
	length int
}
type Node struct {
	next  *Node
	value int
}
```
创建链表：根据给定的值创建链表，没有传值则返回一个空链表
```
// 创建链表
func newList(values ...int) *List {
	list := &List{}
	if len(values) > 0 {
		list.add(values...)
	}
	return list
}
```
向链表中追加元素，在这里定义为，在链表的末尾插入
```
// 从末尾添加若干节点(元素)
func (list *List) add(values ...int) {
	for _, v := range values {
		newNode := &Node{value: v}
		if list.length == 0 {
			list.first = newNode
			list.last = newNode
		} else {
			list.last.next = newNode
			list.last = newNode
		}
		list.length++
	}
}
```
在链表的指定位置插入元素，如果给出的位置不对，就直接返回
```
// 在指定位置插入
func (list *List) insert(m int, value int) {
	// 越界判断
	if m < 0 || m > list.length {
		return
	}
	newNode := &Node{value: value}
	list.length++
	//插入空链表第一个位置
	if m == 0 && list.length == 0 {
		list.first = newNode
		list.last = newNode
		return
	}
	//插入非空链表第一个、最后一个、或者中间某个位置
	switch m {
	case 0:
		newNode.next = list.first
		list.first = newNode
	case list.length:
		list.last.next = newNode
		list.last = newNode
	default:
		before := list.first
		//从第二个节点开始遍历，找到插入位置
		for i := 1; i < m; i++ {
			before = before.next
		}
		newNode.next = before.next
		before.next = newNode
	}
}
```
根据给定位置删除链表中的元素，如果给出的位置不对就直接返回，如果删除的是头节点或者尾节点，还应注意调整相应指针重新指向。
```
// 删除第m个节点
func (list *List) remove(m int) {
	//越界判断
	if m < 0 || m >= list.length {
		return
	}
	list.length--
	//删除的是链表唯一节点
	if list.length == 0 {
		list.first = nil
		list.last = nil
		return
	}
	//删除位置可能是第一个、最后一个、或者中间某个位置
	if m == 0 {
		list.first = list.first.next
	} else {
		//从第二个节点开始遍历，找到删除位置的前一个节点
		before := list.first
		for i := 1; i < m; i++ {
			before = before.next
		}
		//根据中间或末尾情况将节点删除
		if m == list.length-1 {
			before.next = nil
			list.last = before
		} else {
			before.next = before.next.next
		}
	}
}
```

#### #双链表
和单链表结构类似，但每个节点拥有双指针，用以记录前后两个节点的位置。    
```
type List struct {
	first  *Node
	last   *Node
	length int
}
type Node struct {
	pre   *Node
	next  *Node
	value int
}
```
向双链表中添加元素：    
先将元素构造节点，然后插入链表末尾，这意味着每个节点构造时，per节点为链表尾节点。若初始链表为空，则新接入的节点
同时为头节点和尾节点，若不为空，新节点加入仅影响尾节点。
```
// 向链表中添加若干元素,按需更新头节点/尾节点
func (list *List) add(values ...int) {
	for _, v := range values {
		//向末尾添加，新节点的pre=list.last
		node := &Node{value: v, pre: list.last}
		if list.length == 0 {
			list.first = node
			list.last = node
		} else {
			list.last.next = node
			list.last = node
		}
		list.length++
	}
}
```
根据给定的值创建链表。写这里的时候，先考虑实现上面的如何将元素插入一个可能为空的链表，可能会更具象一些。
```
// 根据给定的值创建双链表
func newList(values ...int) *List {
	list := &List{}
	if len(values) > 0 {
		list.add(values...)
	}
	return list
}
```
在给定的位置插入元素。双链表有双指针，所以要根据插入位置选择顺序遍历还是逆序遍历获得插入位置。另外插入就插入动作代码实现而言，       
可以想象在目的位置先将新节点两个触角连到目的位置的前后两个节点，再调整前节点的next和后节点的pre指向这个新节点。
```
// 在给定位置插入元素，这里是双链表，所以直接找到目标位置即可
func (list *List) insert(m int, value int) {
	//越界判断
	if m < 0 || m > list.length {
		return
	}
	node := &Node{value: value}
	//链表判空
	if list.length == 0 {
		list.first = node
		list.last = node
		return
	}
	list.length++ //最容易忘记的先写上
	if m == 0 {
		//插入头节点位置，则原来的头节点pre要指向这个新节点，并更新头节点
		list.first.pre = node
		node.next = list.first
		list.first = node
	} else {
		if m == list.length {
			//末尾追加
			node.pre = list.last
			list.last.next = node
			list.last = node
		} else {
			var current *Node
			// 插入中间的某个位置，对于一个较大的双链表，可能必要考虑插入顺序，比如插入位置为90，长度100和长度10000，
			// 显然应该采取的遍历顺序是不一样的。
			if list.length-m >= m {
				//	插入位置靠前，顺序遍历
				current = list.first
				for i := 0; i < m; i++ {
					current = current.next
				}
			} else {
				//	插入位置靠后，逆序遍历
				current = list.last
				for i := 1; i < list.length-m; i++ {
					current = current.pre
				}
			}
			//找到插入位置的前置节点后，将新节点接入链表
			node.pre = current.pre
			node.next = current
			current.pre.next = node
			current.pre = node
		}
	}
}
```
删除双链表给定位置的节点。因为是双链表，可以直接找到目标位置使其失去前后引用即可。
```
// 删除 第m 个节点
func (list *List) remove(m int) {
	//越界判断
	if m < 0 || m >= list.length {
		return
	}
	list.length-- //最容易忘记的先写上
	//删除后为空链表
	if list.length == 0 {
		list.first = nil
		list.last = nil
		return
	}
	if m == 0 {
		//删除头节点，则将第二个节点的pre置空，然后调整头节点，这样原来的头节点就失去引用了
		list.first.next.pre = nil
		list.first = list.first.next
	} else {
		if m == list.length-1 {
			//删除尾节点
			list.last.pre.next = nil
			list.last = list.last.pre
		} else {
			var current *Node
			// 删除中间的某个位置，对于一个较大的双链表，可能必要考虑遍历顺序，比如位置为90，长度100和长度10000，
			// 显然应该采取的遍历顺序是不一样的。
			if list.length-m >= m {
				//位置靠前，顺序遍历
				current = list.first
				for i := 0; i < m; i++ {
					current = current.next
				}
			} else {
				//位置靠后，逆序遍历
				current = list.last
				for i := 1; i < list.length-m; i++ {
					current = current.pre
				}
			}
			//current为要删除位置的节点,调整前后指针使其失去引用即可
			current.pre.next = current.next
			current.next.pre = current.pre
		}
	}
}
```
