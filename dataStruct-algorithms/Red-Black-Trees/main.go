package main

import "strconv"

// References:https://github.com/emirpasic/gods/blob/master/trees/redblacktree/redblacktree.go
/*
红黑树：
红黑树在avl树之后十年之后提出，与avl树类似，是一种近似平衡的二叉搜索树。
因为红黑树并不显式地要求平衡，没有自平衡操作，所以它的插入删除通常涉及更少的旋转修正。
所以较多增删操作而查询较少适用红黑树，反之较少增删操作而较多查询的更适合用avl树来处理。

红黑树的特征:
1、全部都是红色或者黑色
2、根节点黑色
3、红色节点的孩子为黑色
4、每一条由根节点到叶子节点的路径上，所含的黑色节点数量相等

关于红黑树增删操作的算法描述：
插入：（新节点表示新插入的数据被实例化为red节点，节点颜色判断方法为，空节点为黑，非空节点为本色。本算法遵循先插入后调整的原则。）
1、将新节点插入空树，直接插入，同时根据父节点为空将节点变色为black，同时将节点计数器+1
2、将新节点插入的非空树，先依据二叉搜索树的方式将新节点插入，将新节点的parent指针指向其父节点。如果新节点的父节点为black，则无需调整插入完成。
3、如果新节点的父节点为red，先找到叔父节点，如果叔父节点为red，先将父节点、叔父节点变black，祖父节点变red，再将祖父节点作为新节点从（1）开始执行调整。
4、如果叔父节点为black，找到祖父节点。如果 祖父->父节点->新节点 为LR或者RL型，先以父节点为入口左旋或者右旋，这样就变成了单边型。
   再以这单边型的末节点为起始节点，变色父black祖red，以祖父节点为入口再旋转一次即可。

删除：
*/

type Integer int

func (i Integer) String() string {
	return strconv.Itoa(int(i))
}
func main() {
	myTree := NewTree[Integer](10)
	myTree.Insert(20)
	//myTree.Insert(4)
	//myTree.Insert(15)
	//myTree.Insert(17)
	//myTree.Insert(40)
	//myTree.Insert(50)
	//myTree.Insert(60)
	//myTree.Insert(70)
	//myTree.Insert(35)
	//myTree.Insert(38)
	//myTree.Insert(18)
	//myTree.Insert(19)
	//myTree.Insert(45)
	//myTree.Insert(30)
	//myTree.Insert(25)
	myTree.Insert2(20)
	myTree.Insert2(4)
	myTree.Insert2(15)
	myTree.Insert2(17)
	myTree.Insert2(40)
	myTree.Insert2(50)
	myTree.Insert2(60)
	myTree.Insert2(70)
	myTree.Insert2(35)
	myTree.Insert2(38)
	myTree.Insert2(18)
	myTree.Insert2(19)
	myTree.Insert2(45)
	myTree.Insert2(30)
	myTree.Insert2(25)
	myTree.Remove(60)

	//myTree.Insert2(8)
	//myTree.Insert2(20)
	//myTree.Insert2(6)
	//myTree.Insert2(4)
	ShowTreeGraph(*myTree)
}
