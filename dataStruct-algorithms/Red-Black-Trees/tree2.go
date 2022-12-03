package main

// reference: 《Generic Data Structures and Algorithms in Go》

func search[T OrderedStringer](value T, node *Node[T]) (*Node[T], string) {
	if value == node.value {
		return nil, ""
	} else if value > node.value {
		if node.right == nil {
			return node, "R"
		}
		return search(value, node.right)
	} else if value < node.value {
		if node.left == nil {
			return node, "L"
		}
		return search(value, node.left)
	}
	return nil, ""
}
func (tree *RedBlackTree[T]) findParent(value T) (*Node[T], string) {
	return search(value, tree.root)
}
func (tree *RedBlackTree[T]) rightRotate(node, parent, grandfather *Node[T], modifyColor bool) {
	greatgrandfather := grandfather.parent
	tree.updateParent(parent, grandfather, greatgrandfather)
	oldRight := parent.right
	parent.right = grandfather
	grandfather.parent = parent
	grandfather.left = oldRight
	if oldRight != nil {
		oldRight.parent = grandfather
	}
	if modifyColor == true {
		parent.red = false
		node.red = true
		grandfather.red = true
	}
}
func (tree *RedBlackTree[T]) updateParent(node, parentOldChild, newParent *Node[T]) {
	node.parent = newParent
	if newParent != nil {
		if newParent.value > parentOldChild.value {
			newParent.left = node
		} else {
			newParent.right = node
		}
	} else {
		tree.root = node
	}
}
func (tree *RedBlackTree[T]) leftRotate(node, parent, grandfather *Node[T], modifyColor bool) {
	greatgrandfather := grandfather.parent
	tree.updateParent(parent, grandfather, greatgrandfather)
	oldLeft := parent.left
	parent.left = grandfather
	grandfather.parent = parent
	grandfather.right = oldLeft
	if oldLeft != nil {
		oldLeft.parent = grandfather
	}
	if modifyColor == true {
		parent.red = false
		node.red = true
		grandfather.red = true
	}
}
func (tree *RedBlackTree[T]) modifyColor(grandfather *Node[T]) {
	grandfather.right.red = false
	grandfather.left.red = false
	if grandfather != tree.root {
		grandfather.red = true
	}
	tree.checkReconfigure(grandfather)
}
func (tree *RedBlackTree[T]) checkReconfigure(node *Node[T]) {
	var nodeDirection, parentDirection, rotation string
	var uncle *Node[T]
	parent := node.parent
	value := node.value
	if parent == nil || parent.parent == nil || node.red == false || parent.red == false {
		return
	}
	grandfather := parent.parent
	if value < parent.value {
		nodeDirection = "L"
	} else {
		nodeDirection = "R"
	}
	if grandfather.value > parent.value {
		parentDirection = "L"
	} else {
		parentDirection = "R"
	}
	if parentDirection == "L" {
		uncle = grandfather.right
	} else {
		uncle = grandfather.left
	}
	rotation = nodeDirection + parentDirection
	if uncle == nil || uncle.red == false {
		if rotation == "LL" {
			tree.rightRotate(node, parent, grandfather, true)
		} else if rotation == "RR" {
			tree.leftRotate(node, parent, grandfather, true)
		} else if rotation == "LR" {
			tree.rightRotate(nil, node, parent, false)
			tree.leftRotate(parent, node, grandfather, true)
			node, parent = parent, node
		} else if rotation == "RL" {
			tree.leftRotate(nil, node, parent, false)
			tree.rightRotate(parent, node, grandfather, true)
		}
	} else {
		tree.modifyColor(grandfather)
	}
}

func (tree *RedBlackTree[T]) Insert(value T) {
	if tree.root == nil { // Empty tree
		tree.root = &Node[T]{value, false, nil, nil, nil}
		tree.size += 1
		return
	}

	parent, nodeDirection := tree.findParent(value)
	if nodeDirection == "" {
		return
	}
	newNode := Node[T]{value, true, parent, nil, nil}
	if nodeDirection == "L" {
		parent.left = &newNode
	} else {
		parent.right = &newNode
	}
	tree.checkReconfigure(&newNode)
	tree.size += 1
}
