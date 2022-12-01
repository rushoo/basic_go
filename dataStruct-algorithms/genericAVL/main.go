package main

import (
	"fmt"
)

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
对于两次旋转的情况，操作前在破坏节点先补出一个反向空节点再旋转可能更容易理解，
比如7 4 6，4的右子是6，则给6补一个左子tmp(标记的空节点)，这样全部非空指针包括
7的左指4、4的右指6、6的左指tmp，将4、6、tmp左旋（4->6反向）一个节点，此时
7的左指6、4的右指tmp、6的左指4，原各元素的非空指针都没有变化，指向的元素变了。

现在7的左指6、6的左指4，给入口节点的下一节点这里是加上6的右指tmp(反向)，右旋一个节点后，
此时7的左指tmp、6的左指4、6的右指7，检查指针的完整性就可以确认旋转。
*/

type Integer int

func (num Integer) String() string {
	return fmt.Sprintf("%d", num)
}
func main() {
	avlTree := AVLTree[Integer]{}
	avlTree.Insert(12)
	avlTree.Insert(4)
	avlTree.Insert(1)
	avlTree.Insert(3)
	avlTree.Insert(7)
	avlTree.Insert(8)
	avlTree.Insert(10)
	avlTree.Insert(9)
	avlTree.Insert(2)
	avlTree.Insert(11)
	avlTree.Insert(6)
	avlTree.Insert(5)

	avlTree.Delete(2)
	avlTree.Delete(3)
	avlTree.Delete(1)
	avlTree.Delete(10)
	ShowTreeGraph(avlTree, 1000, 600)
}
