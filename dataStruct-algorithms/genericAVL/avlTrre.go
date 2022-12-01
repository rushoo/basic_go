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

func maxTreeNode[T OrderedStringer](node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}
	if node.Right == nil {
		return node
	}
	return maxTreeNode(node.Right)
}
func rotateDelete[T OrderedStringer](node *Node[T]) *Node[T] {
	node.updateHeight()
	bFactor := node.balanceFactor()
	switch { //类似插入时的旋转，分别讨论删除后四种可能的失衡情况
	case bFactor > 1 && node.Left.balanceFactor() >= 0:
		return rightRotate(node)
	case bFactor > 1 && node.Left.balanceFactor() < 0:
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	case bFactor < -1 && node.Right.balanceFactor() <= 0:
		return leftRotate(node)
	case bFactor < -1 && node.Right.balanceFactor() > 0:
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	default: //平衡，直接返回节点状态
		return node
	}
}
func deleteNode[T OrderedStringer](node *Node[T], val T) *Node[T] {
	if val < node.Value {
		node.Left = deleteNode(node.Left, val)
	} else if val > node.Value {
		node.Right = deleteNode(node.Right, val)
	} else {
		//val == node.Value
		switch {
		case node.Left != nil && node.Right != nil:
			//若删除的节点含双子树，将节点值替换为左子树最大值，并删除该最大值节点
			maxLeft := maxTreeNode(node.Left)
			node.Value = maxLeft.Value
			node.Left = deleteNode(node.Left, maxLeft.Value)
		case node.Left != nil:
			node = node.Left
		case node.Right != nil:
			node = node.Right
		default: //node.Left == nil && node.Right == nil,叶子节点直接删除（置为空）
			node = nil
		}
	}
	// 并将叶子节点的状态返回
	if node == nil {
		return nil
	}
	return rotateDelete(node)
}
func deleteNode2[T OrderedStringer](node *Node[T], val T) *Node[T] {
	// 这个判断的必要性是什么？
	if node == nil {
		return nil
	}

	if val > node.Value {
		node.Right = deleteNode(node.Right, val)
	}
	if val < node.Value {
		node.Left = deleteNode(node.Left, val)
	}
	//val == node.Value
	switch {
	case node.Left != nil && node.Right != nil:
		//若删除的节点含双子树，将节点值替换为左子树最大值，并删除该最大值节点
		maxLeft := maxTreeNode(node.Left)
		node.Value = maxLeft.Value
		node.Left = deleteNode(node.Left, maxLeft.Value)
	case node.Left != nil:
		node = node.Left
	case node.Right != nil:
		node = node.Right
	default: //叶子节点直接删除
		node = nil
	}
	if node == nil {
		return nil
	}
	return rotateDelete(node)
}
func (avl *AVLTree[T]) Delete(value T) {
	if avl.Search(value) == true {
		avl.Root = deleteNode(avl.Root, value)
		avl.NumNodes -= 1
	}
}
