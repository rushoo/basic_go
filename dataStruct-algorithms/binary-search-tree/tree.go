package main

import "fmt"

type dataType interface {
	~int | ~float64 | ~string
	String() string
}
type Float float64

func (num Float) String() string {
	return fmt.Sprintf("%0.1f", num)
}

type BSTree[T dataType] interface {
	Find(val T) bool
	Insert(val T)
	Delete(val T) *Node[T]
}
type Node[T dataType] struct {
	value T
	left  *Node[T]
	right *Node[T]
}
type Tree[T dataType] struct {
	root *Node[T]
	size int
}

func inorder[T dataType](node *Node[T]) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Print(node.value, " ")
	inorder(node.right)
}
func preorder[T dataType](node *Node[T]) {
	if node == nil {
		return
	}
	fmt.Print(node.value, " ")
	preorder(node.left)
	preorder(node.right)
}
func postorder[T dataType](node *Node[T]) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Print(node.value, " ")
}

func NewTree[T dataType]() *Tree[T] {
	return &Tree[T]{nil, 0}
}
func find[T dataType](node *Node[T], val T) bool {
	if node == nil {
		return false
	}
	switch {
	case val == node.value:
		return true
	case val < node.value:
		return find(node.left, val)
	default:
		return find(node.right, val)
	}
}
func insertNode[T dataType](node, newNode *Node[T]) {
	switch {
	case newNode.value < node.value:
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	case newNode.value > node.value:
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}
func deleteNode[T dataType](node *Node[T], val T) *Node[T] {
	switch {
	case node == nil:
		return nil
	case val < node.value:
		node.left = deleteNode(node.left, val)
	case val > node.value:
		node.right = deleteNode(node.right, val)
	default:
		switch {
		case node.left == nil && node.right == nil:
			node = nil
		case node.right == nil:
			node = node.left
		case node.left == nil:
			node = node.right
		default: //找到右子树的最小节点，用其数据替换删除节点，再删除该最小节点
			minRight := node.right
			for minRight.left != nil { //找到右子树的最小节点
				minRight = minRight.left
			}
			node.value = minRight.value
			node.right = deleteNode(node.right, minRight.value)
		}
	}
	return node
}

func (t *Tree[T]) Find(val T) bool {
	if t.root == nil {
		return false
	}
	switch {
	case val == t.root.value:
		return true
	case val < t.root.value:
		return find(t.root.left, val)
	case val > t.root.value:
		return find(t.root.right, val)
	default:
		return false
	}
}
func (t *Tree[T]) Insert(val T) {
	if t.Find(val) == false {
		newNode := &Node[T]{val, nil, nil}
		if t.root == nil {
			t.root = newNode
		} else {
			insertNode(t.root, newNode)
		}
		t.size++
	}
}
func (t *Tree[T]) Delete(val T) {
	if t.Find(val) == true {
		deleteNode(t.root, val)
		t.size--
	}
}

// Inorder 几种常用的遍历方法
func (t *Tree[T]) Inorder() {
	//  左中右
	inorder(t.root)
}
func (t *Tree[T]) Preorder() {
	//	中左右
	preorder(t.root)
}
func (t *Tree[T]) Postorder() {
	//	左右中
	postorder(t.root)
}
func (t *Tree[T]) LevelOrder() {
	if t.root == nil {
		return
	}
	//层序遍历
	q := &queue[T]{}
	q.Push(t.root)
	processor(t.root, q)
}

// 定义一个处理器，将队列中的元素取出的同时，将其子节点入列，
// 这样利用队列先进先出的特性实现逐层遍历
func processor[T dataType](node *Node[T], q *queue[T]) {
	if node == nil {
		return
	}
	for q.Len() > 0 {
		node = q.Pop()
		fmt.Print(node.value, " ")
		if node.left != nil {
			q.Push(node.left)
		}
		if node.right != nil {
			q.Push(node.right)
		}
	}
}

// 构造队列用以临时存放树的节点
type queue[T dataType] struct {
	items []*Node[T]
	size  int
}

func (q *queue[T]) Len() int {
	return q.size
}
func (q *queue[T]) Push(node *Node[T]) {
	q.items = append(q.items, node)
	q.size++
}
func (q *queue[T]) Pop() *Node[T] {
	node := q.items[0]
	q.items = q.items[1:]
	q.size--
	return node
}
