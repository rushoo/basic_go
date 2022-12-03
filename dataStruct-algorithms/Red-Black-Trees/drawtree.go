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

type NodePair struct {
	Val1, Val2 string
}
type NodePos struct {
	Val  string
	Red  bool
	YPos int
	XPos int
}

var data []NodePos
var endPoints []NodePair // Used to plot lines
func PrepareDrawTree[T OrderedStringer](tree RedBlackTree[T]) {
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
func prepareToDraw[T OrderedStringer](tree RedBlackTree[T]) {
	inorderLevel(tree.root, 1)
	SetXValues()
	getEndPoints(tree.root, nil)
}
func inorderLevel[T OrderedStringer](node *Node[T], level int) {
	if node != nil {
		inorderLevel(node.left, level+1)
		data = append(data,
			NodePos{node.value.String(), node.red,
				100 - level, -1})
		inorderLevel(node.right, level+1)
	}
}
func getEndPoints[T OrderedStringer](node *Node[T], parent *Node[T]) {
	if node != nil {
		if parent != nil {
			endPoints = append(endPoints,
				NodePair{node.value.String(),
					parent.value.String()})
		}
		getEndPoints(node.left, node)
		getEndPoints(node.right, node)
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
func ShowTreeGraph[T OrderedStringer](myTree RedBlackTree[T]) {
	PrepareDrawTree(myTree)
	myApp := app.New()
	myWindow := myApp.NewWindow("Tree")
	myWindow.Resize(fyne.NewSize(1000, 600))
	path, _ := homedir.Dir()
	path += "/Desktop//"
	nodePts := make(plotter.XYs, myTree.size)
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
	nodePoints.Color = color.RGBA{R: 255, G: 255, B: 250, A: 255} // White fill
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
		x := float64(data[index].XPos) - 0.10
		y := float64(data[index].YPos) - 0.02
		str := data[index].Val
		if data[index].Red == true {
			str += "(RED)"
		} else {
			str += "(BLACK)"
		}
		label, err :=
			plotter.NewLabels(plotter.XYLabels{
				XYs: []plotter.XY{
					{X: x, Y: y},
				},
				Labels: []string{str},
			})
		if err != nil {
			log.Fatalf("could not creates labels	plotter: %+v", err)
		}
		p.Add(label)
	}
	path, _ = homedir.Dir()
	path += "/Desktop/GoDS/"
	err = p.Save(1000, 600, "tree.png")
	if err != nil {
		log.Panic(err)
	}
	DrawGraph(myApp, myWindow)
	myWindow.ShowAndRun()
}
