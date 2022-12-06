package main

func main() {

	randomSearchTree := NewTree[Float]()
	randomSearchTree.Insert(Float(4))
	randomSearchTree.Insert(Float(2))
	randomSearchTree.Insert(Float(15))
	randomSearchTree.Insert(Float(1))
	//randomSearchTree.Delete(Float(1))
	//randomSearchTree.ShowTreeGraph()
	randomSearchTree.Insert(Float(3))
	randomSearchTree.Insert(Float(13))
	randomSearchTree.Insert(Float(17))
	randomSearchTree.Insert(Float(11))
	randomSearchTree.Insert(Float(12))
	randomSearchTree.Insert(Float(14))
	randomSearchTree.Insert(Float(9))
	randomSearchTree.Insert(Float(16.5))
	randomSearchTree.Insert(Float(17.5))
	randomSearchTree.Insert(Float(16.2))
	randomSearchTree.Insert(Float(16.8))
	randomSearchTree.Insert(Float(8))
	randomSearchTree.Insert(Float(13.5))
	randomSearchTree.Insert(Float(3.5))
	randomSearchTree.Insert(Float(12.5))
	//randomSearchTree.Delete(Float(8))
	//randomSearchTree.ShowTreeGraph()

	//randomSearchTree.Inorder()
	//fmt.Println()
	//randomSearchTree.Preorder()
	//fmt.Println()
	//randomSearchTree.Postorder()
	randomSearchTree.LevelOrder()
	randomSearchTree.ShowTreeGraph()
}
