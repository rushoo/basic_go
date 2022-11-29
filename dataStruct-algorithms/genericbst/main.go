package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Number int

func (num Number) String() string {
	return fmt.Sprintf("%d", num)
}

type Float float64

func (num Float) String() string {
	return fmt.Sprintf("%0.1f", num)
}

func inorderOperator(val Float) {
	fmt.Println(val.String())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Generate a random search tree
	randomSearchTree := BinarySearchTree[Float]{nil, 0}
	//for i := 0; i < 30; i++ {
	//	rn := 1.0 + 99.0*rand.Float64()
	//	randomSearchTree.Insert(Float(rn))
	//}
	randomSearchTree.Insert(Float(4))
	randomSearchTree.Insert(Float(2))
	randomSearchTree.Insert(Float(15))
	randomSearchTree.Insert(Float(1))
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

	//randomSearchTree.Delete(1.0)
	//randomSearchTree.Delete(9.0)
	//randomSearchTree.Delete(14.0)
	//randomSearchTree.Delete(3.0)
	//randomSearchTree.Delete(12.0)
	//randomSearchTree.Delete(13.0)
	//randomSearchTree.Delete(15.0)
	//randomSearchTree.Delete(4.0)
	ShowTreeGraph(randomSearchTree)
	//randomSearchTree.InOrderTraverse(inorderOperator)

	//min := randomSearchTree.Min()
	//max, h := randomSearchTree.Max()
	//fmt.Printf("\nMinimum value in random search tree is %0.1f  \nMaximum value in random search tree is %0.1f \nSearch height of tree: %d", *min, *max, h)

	//start := time.Now()
	//tree := BinarySearchTree[Number]{nil, 0}
	//for val := 0; val < 100_000; val++ {
	//	tree.Insert(Number(val))
	//}
	//elapsed := time.Since(start)
	//_, ht := tree.Max()
	//fmt.Printf("\nTime to build BST tree with 100,000 nodes in sequential order: %s. Search height of tree: %d", elapsed, ht)
}
