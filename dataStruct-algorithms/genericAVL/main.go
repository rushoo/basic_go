package main

import "fmt"

/*
平衡二叉树：
特点是其任意节点的两棵子树的最大高度差不超过1

AVL插入全过程：
   1、逐层递归以二叉搜索树方式插入节点，插入的节点一定是叶子节点
   2、沿路回退更新每一节点的平衡因子，并判断当前节点是否平衡
   3、将该节点平衡后回退上一层

影响平衡的插入：
   若在一棵平衡二叉树上插入新的节点导致其失衡，则新节点一定是插入到叶子节点上。插到非叶子节点不改变子树的度，不会影响原树的平衡性。
所有可能失衡的插入：
   1、左子树叶子左节点
   2、左子树叶子右节点
   3、右子树叶子左节点
   4、右子树叶子右节点

节点平衡过程：
   1、右旋
   2、左旋再右旋
   3、右旋再左旋
   4、右旋
*/

type Integer int

func (num Integer) String() string {
	return fmt.Sprintf("%d", num)
}
func main() {
	avlTree := AVLTree[Integer]{}
	avlTree.Insert(15)
	avlTree.Insert(12)
	avlTree.Insert(17)
	avlTree.Insert(6)
	avlTree.Insert(13)
	avlTree.Insert(4)
}
