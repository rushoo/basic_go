package main

type BinaryTree struct {
	Root     *Node
	NumNodes int
}
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// 节点对，包含一对子节点与父节点
type nodePair struct {
	Val1, Val2 string
}

// 节点信息，含节点数据、坐标
type nodePos struct {
	Val  string
	XPos int
	YPos int
}

var data []nodePos
var endPoints []nodePair // Used to plot lines
var path string

// 以参数节点为起点中序遍历遍历二叉树，顺序是左子树、中、右子树
func inorderLevel(node *Node, level int) {
	if node != nil {
		inorderLevel(node.Left, level+1)                          //遍历左子树，将节点数据和坐标保存到slice
		data = append(data, nodePos{node.Value, -1, 100 - level}) //保存起点信息
		inorderLevel(node.Right, level+1)                         //遍历保存右子树节点信息
	}
}

// 根据节点的值返回其坐标
func findXY(val string) (int, int) {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return data[i].XPos, data[i].YPos
		}
	}
	return -1, -1
}

// 找出val在data[]中的位置，比如val是第7个元素，则返回值是6
func findX(val string) int {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return i
		}
	}
	return -1
}

// 根据中序遍历的结果中元素的索引位置建立相对横坐标
func setXValues() {
	for index := 0; index < len(data); index++ {
		xValue := findX(data[index].Val)
		data[index].XPos = xValue
	}
}

// 获取节点对的信息，也即节点关系集
func getEndPoints(node *Node, parent *Node) {
	if node != nil {
		if parent != nil {
			endPoints = append(endPoints, nodePair{node.Value, parent.Value})
		}
		getEndPoints(node.Left, node)
		getEndPoints(node.Right, node)
	}
}
