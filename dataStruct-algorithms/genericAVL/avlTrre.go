package main

type ordered interface {
	~int | ~float64 | ~string
}
type OrderedStringer interface {
	ordered
	String() string
}

type Node[T OrderedStringer] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
	Ht    int // 含节点最大子树深度
}
type AVLTree[T OrderedStringer] struct {
	Root     *Node[T]
	NumNodes int
}

func newNode[T OrderedStringer](val T) *Node[T] {
	return &Node[T]{
		Value: val,
		Left:  nil,
		Right: nil,
		Ht:    1,
	}
}
func search[T OrderedStringer](n *Node[T], value T) bool {
	if n == nil {
		return false
	}
	if value < n.Value {
		return search(n.Left, value)
	}
	if value > n.Value {
		return search(n.Right, value)
	}
	return true
}
func (avl *AVLTree[T]) Search(value T) bool {
	return search(avl.Root, value)
}

func (avl *AVLTree[T]) Height() int {
	return avl.Root.Height()
}

func (n *Node[T]) Height() int {
	if n == nil {
		return 0
	} else {
		return n.Ht
	}
}
func (n *Node[T]) updateHeight() {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n.Ht = max(n.Left.Height(), n.Right.Height()) + 1
}
func (n *Node[T]) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.Height() - n.Right.Height()
}
func leftRotate[T OrderedStringer](x *Node[T]) *Node[T] {
	y := x.Right
	t := y.Left
	y.Left = x
	x.Right = t
	x.updateHeight()
	y.updateHeight()
	return y
}
func rightRotate[T OrderedStringer](x *Node[T]) *Node[T] {
	y := x.Left
	t := y.Right
	y.Right = x
	x.Left = t
	x.updateHeight()
	y.updateHeight()
	return y
}
func rotateInsert[T OrderedStringer](node *Node[T], val T) *Node[T] {
	node.updateHeight()
	bFactor := node.balanceFactor()
	if bFactor > 1 && val < node.Left.Value {
		return rightRotate(node)
	}
	if bFactor > 1 && val > node.Left.Value {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if bFactor < -1 && val > node.Right.Value {
		return leftRotate(node)
	}
	if bFactor < -1 && val < node.Right.Value {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}
	return node
}

func insertNode[T OrderedStringer](node *Node[T], val T) *Node[T] {
	if node == nil {
		// 结束末层递归函数调用，返回构造的叶子节点
		return newNode(val)
	}
	if val < node.Value {
		node.Left = insertNode(node.Left, val)
	}
	if val > node.Value {
		node.Right = insertNode(node.Right, val)
	}
	//返回最终结果前调整二叉树结构
	return rotateInsert(node, val)
}
func (avl *AVLTree[T]) Insert(newValue T) {
	if avl.Search(newValue) == false {
		avl.Root = insertNode(avl.Root, newValue)
		avl.NumNodes += 1
	}
}
