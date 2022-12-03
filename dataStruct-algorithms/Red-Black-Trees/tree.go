package main

type ordered interface {
	~int | ~float64 | ~string
}
type OrderedStringer interface {
	ordered
	String() string
}
type Node[T OrderedStringer] struct {
	value  T
	red    bool
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}
type RedBlackTree[T OrderedStringer] struct {
	size int
	root *Node[T]
}

func NewTree[T OrderedStringer](value T) *RedBlackTree[T] {
	return &RedBlackTree[T]{1, &Node[T]{value, false, nil, nil, nil}}
}
func redNode[T OrderedStringer](node *Node[T]) bool {
	if node == nil {
		return false
	}
	return node.red
}

func (node *Node[T]) uncle() *Node[T] {
	if node == nil || node.parent == nil || node.parent.parent == nil {
		return nil
	}
	return node.parent.sibling()
}
func (node *Node[T]) sibling() *Node[T] {
	if node == nil || node.parent == nil {
		return nil
	}
	if node == node.parent.left {
		return node.parent.right
	}
	return node.parent.left
}
func (node *Node[T]) grandparent() *Node[T] {
	if node != nil && node.parent != nil {
		return node.parent.parent
	}
	return nil
}
func (node *Node[T]) maximumNode() *Node[T] {
	if node == nil {
		return nil
	}
	if node.right != nil {
		node = node.right
	}
	return node
}

func (tree *RedBlackTree[T]) replaceNode(old *Node[T], new *Node[T]) {
	//用new节点替换old节点的位置
	if old.parent == nil {
		tree.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

// 这个rotate相当于分步骤的实现，没有之前那个实现连贯自然
func (tree *RedBlackTree[T]) rotateLeft(node *Node[T]) {
	right := node.right
	tree.replaceNode(node, right)
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
	node.parent = right
}
func (tree *RedBlackTree[T]) rotateRight(node *Node[T]) {
	left := node.left
	tree.replaceNode(node, left)
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
	node.parent = left
}

func (tree *RedBlackTree[T]) Insert2(val T) {
	var insertedNode *Node[T]
	if tree.root == nil {
		tree.root = &Node[T]{val, true, nil, nil, nil}
		insertedNode = tree.root
	} else {
		node := tree.root
		loop := true
		for loop {
			switch {
			case val == node.value:
				return
			case val < node.value:
				if node.left == nil {
					node.left = &Node[T]{val, true, nil, nil, nil}
					insertedNode = node.left
					loop = false
				} else {
					node = node.left
				}
			case val > node.value:
				if node.right == nil {
					node.right = &Node[T]{val, true, nil, nil, nil}
					insertedNode = node.right
					loop = false
				} else {
					node = node.right
				}
			}
		}
		insertedNode.parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++
}
func (tree *RedBlackTree[T]) insertCase1(node *Node[T]) {
	if node.parent == nil {
		node.red = false
	} else {
		tree.insertCase2(node)
	}
}
func (tree *RedBlackTree[T]) insertCase2(node *Node[T]) {
	if redNode(node.parent) == false {
		return
	}
	tree.insertCase3(node)
}
func (tree *RedBlackTree[T]) insertCase3(node *Node[T]) {
	uncle := node.uncle()
	if redNode(uncle) == true {
		node.parent.red = false
		uncle.red = false
		node.grandparent().red = true
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}
func (tree *RedBlackTree[T]) insertCase4(node *Node[T]) {
	grandparent := node.grandparent()
	if node == node.parent.right && node.parent == grandparent.left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandparent.right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.insertCase5(node)
}
func (tree *RedBlackTree[T]) insertCase5(node *Node[T]) {
	node.parent.red = false
	grandparent := node.grandparent()
	grandparent.red = true
	if node == node.parent.left && node.parent == grandparent.left {
		tree.rotateRight(grandparent)
	} else if node == node.parent.right && node.parent == grandparent.right {
		tree.rotateLeft(grandparent)
	}
}

// Remove remove the node from the tree
func (tree *RedBlackTree[T]) Remove(val T) {
	var child *Node[T]
	node := tree.lookup(val)
	if node == nil {
		return
	}
	if node.left != nil && node.right != nil {
		pred := node.left.maximumNode()
		node.value = pred.value
		node = pred
	}
	if node.left == nil || node.right == nil {
		if node.right == nil {
			child = node.left
		} else {
			child = node.right
		}
		if node.red == false {
			node.red = redNode(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.parent == nil && child != nil {
			child.red = false
		}
	}
	tree.size--
}
func (tree *RedBlackTree[T]) lookup(val T) *Node[T] {
	node := tree.root
	for node != nil {
		switch {
		case val == node.value:
			return node
		case val < node.value:
			node = node.left
		case val > node.value:
			node = node.right
		}
	}
	return nil
}
func (tree *RedBlackTree[T]) deleteCase1(node *Node[T]) {
	if node.parent == nil {
		return
	}
	tree.deleteCase2(node)
}
func (tree *RedBlackTree[T]) deleteCase2(node *Node[T]) {
	sibling := node.sibling()
	if redNode(sibling) == true {
		node.parent.red = true
		sibling.red = false
		if node == node.parent.left {
			tree.rotateLeft(node.parent)
		} else {
			tree.rotateRight(node.parent)
		}
	}
	tree.deleteCase3(node)
}
func (tree *RedBlackTree[T]) deleteCase3(node *Node[T]) {
	sibling := node.sibling()
	if redNode(node.parent) == false && redNode(sibling) == false &&
		redNode(sibling.left) == false && redNode(sibling.right) == false {
		sibling.red = true
		tree.deleteCase1(node.parent)
	} else {
		tree.deleteCase4(node)
	}
}
func (tree *RedBlackTree[T]) deleteCase4(node *Node[T]) {
	sibling := node.sibling()
	if redNode(node.parent) == true && redNode(sibling) == false &&
		redNode(sibling.left) == false && redNode(sibling.right) == false {
		sibling.red = true
		node.parent.red = false
	} else {
		tree.deleteCase5(node)
	}
}
func (tree *RedBlackTree[T]) deleteCase5(node *Node[T]) {
	sibling := node.sibling()
	if node == node.parent.left &&
		redNode(sibling) == false &&
		redNode(sibling.left) == true &&
		redNode(sibling.right) == false {
		sibling.red = true
		sibling.left.red = false
		tree.rotateRight(sibling)
	} else if node == node.parent.right &&
		redNode(sibling) == false &&
		redNode(sibling.right) == true &&
		redNode(sibling.left) == false {
		sibling.red = true
		sibling.right.red = false
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}
func (tree *RedBlackTree[T]) deleteCase6(node *Node[T]) {
	sibling := node.sibling()
	sibling.red = redNode(node.parent)
	node.parent.red = false
	if node == node.parent.left && redNode(sibling.right) == true {
		sibling.right.red = false
		tree.rotateLeft(node.parent)
	} else if redNode(sibling.left) == true {
		sibling.left.red = false
		tree.rotateRight(node.parent)
	}
}
