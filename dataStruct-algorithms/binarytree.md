### # Tree.树的基本概念
- 根节点 root   --- 没有父节点，是树的源头
- 节点 node   --- 略   
- 边 edge   --- 描述两节点的连接关系，若A->B，就称AB之间有一条边
- 路径 path   --- 边的组合，比如A->B，B->C，那么就有一条路径：A->B->C
- 叶子 leaf   --- 没有子节点的节点
- 树的深度 depth   --- 从根节点到叶子节点的最长路径
- 节点高度 height   --- 从某个节点出发，到叶子节点的最长路径，根节点的高度就是树的深度。
- 节点阶度 level   --- 节点位于树中第几代或者第几层就是该节点的阶度
- 节点的度 degree   --- 节点的子节点数量，二叉树的任一节点，度最大是2.
- 子节点 children   --- 节点指向的下层节点，二叉树有左子节点、右子节点
- 父节点 parent   --- A是B的子节点，那么B就是A的父节点 
- 兄弟节点 sibling   --- 相同父节点的一组节点
- 祖先节点与子孙节点 ancestor && descendant --- 父子节点是有直接指向关系的祖孙关系节点 

### Binary Tree
每个节点至多含两个子节点的树，根据[wiki](https://simple.wikipedia.org/wiki/Binary_tree)，二叉树有几种基本类型：
- complete binary tree --- every level, except possibly the last, is completely filled, and all items in the last level are as far left as possible.
  - 完全二叉树，除了最下层不一定，其它层都完全填充，最下层从左往右依次分布，满布时就是完美二叉树。
- full binary tree --- every item has either 0 or 2 children.
  - 满二叉树，每个节点都有0或2个子节点。
- perfect binary tree --- all interior items have two children and all leaves have the same level.
  - 完美二叉树，每一层都完全填充。
- balanced binary tree --- the left and right branches of every item differ in height by no more than 1.
  - 对于任一节点，其左子树和右子树的高度差最多是1，例如AVL树和红黑树(非严格).
- X skewed binary tree ---  either each node is have a X (left/right) child or no child (leaf)
  - 左/右偏二叉树，除叶子节点外，每个节点都只有左/右一种子节点
```
// 二叉树本质上是一种特殊的链表结构，这里以int类型节点值为例，每个节点包含一个节点值和指向两个子节点的指针
// 以根节点来表示树，树的大小为此二叉树所有节点数量总和
type Node struct {
    value int
    left  *Node
    right *Node
}
type Tree struct {
    root *Node
    size int
}
```
#### #创建完全二叉树
以符合直觉的层序遍历视角来看，假设第一层的根节点为完全二叉树的第0个节点，那么它的左、右子节点分别是第1个、第2个子节点。     
对于第1个节点而言，它的左、右子节点分别是第3个、第4个节点，对于第2个节点而言，它的左、右子节点分别是第5个、第6个节点.     
不难发现，对于完全二叉树的第k个节点而言，第2k+1个节点为其左子节点，第2k+2个节点为其右子节点。     
(第0个的说法并不妥当，这里将二叉树的节点从0开始计数是为了与数组保持一致)
二叉树本质上是一种特殊的链表，所以其创建过程和链表类似：先创建一个节点，再将(树)root指向这个节点。用递归将所有节点相关联。   
```
func newTree(s []int) *Tree {
    //数组第0个元素创建的节点，这里应该作为根节点。
    node := genTreeNode(0, s)
    tree := &Tree{}
    tree.root = node
    tree.size = len(s)
    return tree
}
// genNode根据数组中第k个元素创建节点，左子节点是第2k+1个元素，右子节点是第2k+2个元素
func genTreeNode(k int, s []int) *TreeNode {
    //越界判断
    if k >= len(s) {
        return nil
    }
    return &TreeNode{s[k], genTreeNode(2*k+1, s), genTreeNode(2*k+2, s)}
}
```
#### #遍历二叉树
上面根据直观，层序创建了完全二叉树，现在使用层数遍历检验是否正确。
层序遍历二叉树，就是从二叉树的第一层开始，依次从左往右遍历每个节点，每层遍历结束后开始下一层节点的遍历。层序遍历的可以借助队列FIFO特性    
将树节点不断入列出列来实现。  
```
/*
    设计数据结构：
    这里队列频繁进行单个元素的增加删除,所以应该使用单链表结构实现队列。而出队的结果需要统计，所以出队应该有返回值。
    队列应用到的方法包括：将一个树节点入列、出列得到一个树节点、队列判空
    因此对应单链表应实现的方法：链尾添加元素、链头删除元素并返回值、判长
    链表节点的数据域为入列出列的对象，即树节点指针
*/
type Queue struct {
    list *SinglyList  // 单链表应实现队列 
}
type SinglyList struct {
    first *Node
    last  *Node
    size  int
}
type Node struct {
    next  *Node
    value *TreeNode  // 将指针类型 *TreeNode 作为value以便直接对树节点入列 
}

// 队尾入列，队首出列，队列判空
func (queue *Queue) enqueue(tNode *TreeNode) {
    queue.list.add(tNode)
}
func (queue *Queue) dequeue() *TreeNode {
    if queue.empty() {
    //这里用-1值来表示对空列出列的异常
        return &TreeNode{value: -1}
    }
    return queue.list.remove()
}
func (queue *Queue) empty() bool {
    return queue.list.size == 0
}

// 向链表中添加元素，从链表中删除元素
func (list *SinglyList) add(tNode *TreeNode) {
    node := &Node{value: tNode}
    //	当链表为空时
    if list.size == 0 {
        list.first = node
        list.last = node
    } else {
        list.last.next = node
        list.last = node
    }
    list.size++
}
func (list *SinglyList) remove() *TreeNode {
    //从头节点开始删除,并返回节点的值
    list.size--
    tNode := list.first.value
    if list.size == 0 {
        list.first = nil
        list.last = nil
    } else {
        list.first = list.first.next
    }
    return tNode
}
```
利用上面的数据结构实现层序遍历,经过测试，层序遍历的结果和插入时所用的数列元素一致。
```
func levelTraverse(tree *Tree) []int {
    if tree.root == nil {
        //切片是引用类型，这里判空直接返回nil
        return nil
    }
    //以恰当空间开辟数组，避免自动扩容的时空开销
    var res = make([]int, tree.size)
    que := newQueue()
    //初始时先将root节点入列
    que.enqueue(tree.root)
    for i := 0; !que.empty(); i++ {
        tNode := que.dequeue()
        res[i] = tNode.value
        //树节点出列后，再分别将左右子节点入列
        if tNode.left != nil {
            que.enqueue(tNode.left)
        }
        if tNode.right != nil {
            que.enqueue(tNode.right)
        }
    }
    return res
}
```
前中后序遍历
```
func (tree *Tree) PrintInOrder() {
    fmt.Println("\n中序遍历：")
    printInOrder(tree.root)
}
func (tree *Tree) PrintPostOrder() {
    fmt.Println("\n后序遍历：")
    printPostOrder(tree.root)
}
func (tree *Tree) PrintPreOrder() {
    fmt.Println("\n前序遍历：")
    printPreOrder(tree.root)
}
func printPreOrder(tNode *TreeNode) {
    //前序遍历，中左右
    if tNode == nil {
    	return
    }
    fmt.Printf(" %d ", tNode.value)
    printPreOrder(tNode.left)
    printPreOrder(tNode.right)
}
func printPostOrder(tNode *TreeNode) {
    //后序遍历，左右中
    if tNode == nil {
    	return
    }
    printPostOrder(tNode.left)
    printPostOrder(tNode.right)
    fmt.Printf(" %d ", tNode.value)
}
func printInOrder(tNode *TreeNode) {
    //中序遍历，左中右
    if tNode == nil {
    	return
    }
    printInOrder(tNode.left)
    fmt.Printf(" %d ", tNode.value)
    printInOrder(tNode.right)
}
```