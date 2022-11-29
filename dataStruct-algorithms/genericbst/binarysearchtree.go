package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/mitchellh/go-homedir"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"log"
)

type ordered interface {
	~int | ~float64 | ~string
}

type BinarySearchTree[T OrderedStringer] struct {
	Root     *Node[T]
	NumNodes int
}

type Node[T OrderedStringer] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type OrderedStringer interface {
	ordered
	String() string
}

func insertNode[T OrderedStringer](node, newNode *Node[T]) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}
func (bst *BinarySearchTree[T]) Insert(newValue T) {
	if bst.Search(newValue) == false {
		newNode := &Node[T]{newValue, nil, nil}
		if bst.Root == nil {
			bst.Root = newNode
		} else {
			insertNode(bst.Root, newNode)
		}
		bst.NumNodes += 1
	}
}

/*
删除节点：
1、被删除的是叶子节点，直接删除即可
2、被删除的节点仅含左子节点，将指向该节点的指针直接指向其子节点
3、被删除的节点仅含右子节点，将指向该节点的指针直接指向其子节点
4、被删除的节点含有双子节点，找出右子节点最小值替换，然后删除该最小值节点
*/
func deleteNode[T OrderedStringer](node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}
	// 递归直到找到对应的元素
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	}
	if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	}
	//此时已经找到要删除的元素，后面对删除元素的位置进行讨论，分别是叶子节点，仅左子节点，仅右子节点，双节点
	switch {
	case node.Left == nil && node.Right == nil:
		node = nil
	case node.Right == nil:
		node = node.Left
	case node.Left == nil:
		node = node.Right
	default: //找到右子树的最小节点，用其数据替换删除节点，再删除该最小节点
		minRight := node.Right
		for { //找到右子树的最小节点
			if minRight != nil && minRight.Left != nil {
				minRight = minRight.Left
			} else {
				break // 结束for循环
			}
		}
		node.Value = minRight.Value
		node.Right = deleteNode(node.Right, minRight.Value)
	}
	return node
}
func (bst *BinarySearchTree[T]) Delete(value T) {
	if bst.Search(value) == true {
		deleteNode(bst.Root, value)
		bst.NumNodes -= 1
	}
}

func search[T OrderedStringer](root *Node[T], value T) bool {
	if root == nil {
		return false
	}
	if value < root.Value {
		return search(root.Left, value)
	}
	if value > root.Value {
		return search(root.Right, value)
	}
	return true
}
func (bst *BinarySearchTree[T]) Search(value T) bool {
	return search(bst.Root, value)
}

func (bst *BinarySearchTree[T]) InOrderTraverse(op func(T)) {
	inOrderTraverse(bst.Root, op)
}

// 二叉搜索树的最小值应该是最左边的节点元素
func (bst *BinarySearchTree[T]) Min() *T {
	node := bst.Root
	if node == nil {
		return nil
	}
	for {
		if node.Left == nil {
			return &node.Value
		}
		node = node.Left
	}
}

// 二叉搜索树的最大值应该是最右边的节点元素
func (bst *BinarySearchTree[T]) Max() (*T, int) {
	node := bst.Root
	height := 0 //搜索深度
	if node == nil {
		return nil, height
	}
	height++
	for {
		if node.Right == nil {
			return &node.Value, height
		}
		height += 1
		node = node.Right
	}
}

func inOrderTraverse[T OrderedStringer](node *Node[T], op func(T)) {
	if node != nil {
		inOrderTraverse(node.Left, op)
		op(node.Value)
		inOrderTraverse(node.Right, op)
	}
}

// Logic for drawing tree
type NodePair struct {
	Val1, Val2 string
}

type NodePos struct {
	Val  string
	YPos int
	XPos int
}

var data []NodePos
var endPoints []NodePair // Used to plot lines

func PrepareDrawTree[T OrderedStringer](tree BinarySearchTree[T]) {
	prepareToDraw(tree)
}

func FindXY(val interface{}) (int, int) {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return data[i].XPos, data[i].YPos
		}
	}
	return -1, -1
}

func FindX(val interface{}) int {
	for i := 0; i < len(data); i++ {
		if data[i].Val == val {
			return i
		}
	}
	return -1
}

func SetXValues() {
	for index := 0; index < len(data); index++ {
		xValue := FindX(data[index].Val)
		data[index].XPos = xValue
	}
}

func prepareToDraw[T OrderedStringer](tree BinarySearchTree[T]) {
	inorderLevel(tree.Root, 1)
	SetXValues()
	getEndPoints(tree.Root, nil)
}

func inorderLevel[T OrderedStringer](node *Node[T], level int) {
	if node != nil {
		inorderLevel(node.Left, level+1)
		data = append(data, NodePos{node.Value.String(), 100 - level, -1})
		inorderLevel(node.Right, level+1)
	}
}

func getEndPoints[T OrderedStringer](node *Node[T], parent *Node[T]) {
	if node != nil {
		if parent != nil {
			endPoints = append(endPoints, NodePair{node.Value.String(), parent.Value.String()})
		}
		getEndPoints(node.Left, node)
		getEndPoints(node.Right, node)
	}
}

var path string

func DrawGraph(a fyne.App, w fyne.Window) {
	image := canvas.NewImageFromResource(theme.FyneLogo())
	image = canvas.NewImageFromFile(path + "tree.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.Close()
	w.Show()
}

func ShowTreeGraph[T OrderedStringer](myTree BinarySearchTree[T]) {
	PrepareDrawTree(myTree)
	myApp := app.New()
	myWindow := myApp.NewWindow("Tree")
	myWindow.Resize(fyne.NewSize(600, 360))
	path, _ := homedir.Dir()
	path += "/Desktop//"

	nodePts := make(plotter.XYs, myTree.NumNodes)
	for i := 0; i < len(data); i++ {
		nodePts[i].Y = float64(data[i].YPos)
		nodePts[i].X = float64(data[i].XPos)
	}
	nodePtsData := nodePts
	p := plot.New()
	p.Add(plotter.NewGrid())
	nodePoints, err := plotter.NewScatter(nodePtsData)
	if err != nil {
		log.Panic(err)
	}
	nodePoints.Shape = draw.CircleGlyph{}
	nodePoints.Color = color.RGBA{G: 255, A: 255}
	nodePoints.Radius = vg.Points(12)

	// Plot lines
	for index := 0; index < len(endPoints); index++ {
		val1 := endPoints[index].Val1
		x1, y1 := FindXY(val1)
		val2 := endPoints[index].Val2
		x2, y2 := FindXY(val2)
		pts := plotter.XYs{{X: float64(x1), Y: float64(y1)}, {X: float64(x2), Y: float64(y2)}}
		line, err := plotter.NewLine(pts)
		if err != nil {
			log.Panic(err)
		}
		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			log.Panic(err)
		}
		p.Add(line, scatter)
	}

	p.Add(nodePoints)

	// Add Labels
	for index := 0; index < len(data); index++ {
		x := float64(data[index].XPos) - 0.2
		y := float64(data[index].YPos) - 0.02
		str := data[index].Val
		label, err := plotter.NewLabels(plotter.XYLabels{
			XYs: []plotter.XY{
				{X: x, Y: y},
			},
			Labels: []string{str},
		})
		if err != nil {
			log.Fatalf("could not creates labels plotter: %+v", err)
		}
		p.Add(label)
	}

	path, _ = homedir.Dir()
	path += "/Desktop/GoDS/"
	err = p.Save(600, 360, "tree.png")
	if err != nil {
		log.Panic(err)
	}

	DrawGraph(myApp, myWindow)

	myWindow.ShowAndRun()
}
