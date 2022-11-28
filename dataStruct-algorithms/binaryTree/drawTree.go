package main

import (
	"fmt"
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

func prepareToDraw(tree BinaryTree) {
	inorderLevel(tree.Root, 1)   //level用以指定相对坐标
	setXValues()                 //根据节点在data slice从左到右相对位置给节点横坐标赋值
	getEndPoints(tree.Root, nil) //获取所有子节点-父节点对
}
func prepareDrawTree(tree BinaryTree) {
	prepareToDraw(tree)
	fmt.Println("\nslice of endPoints: %v", endPoints)
	fmt.Println("\nslice of data: %v", data)
}

// 画图
func drawGraph(a fyne.App, w fyne.Window) {
	image := canvas.NewImageFromResource(theme.FyneLogo())
	image = canvas.NewImageFromFile(path + "tree.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.Show()
}
func ShowTreeGraph(myTree BinaryTree) {
	prepareDrawTree(myTree)

	myApp := app.New()
	myWindow := myApp.NewWindow("Binary Tree")
	myWindow.Resize(fyne.NewSize(1000, 600))
	path, _ := homedir.Dir()
	//path += "/Desktop//"

	//坐标容器,将所有的节点坐标copy过来
	nodePts := make(plotter.XYs, myTree.NumNodes)
	for i := 0; i < len(data); i++ {
		nodePts[i].X = float64(data[i].XPos)
		nodePts[i].Y = float64(data[i].YPos)
	}

	//获取坐标集的副本
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

	// 根据遍历获取每对子-父节点的坐标，画图
	for index := 0; index < len(endPoints); index++ {
		val1 := endPoints[index].Val1
		x1, y1 := findXY(val1)
		val2 := endPoints[index].Val2
		x2, y2 := findXY(val2)
		pts := plotter.XYs{{X: float64(x1), Y: float64(y1)},
			{X: float64(x2), Y: float64(y2)}}
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

	// 给绘制的每个节点加标记（其实就是各节点的实际值）
	for index := 0; index < len(data); index++ {
		x := float64(data[index].XPos) - 0.05
		y := float64(data[index].YPos) - 0.02
		str := data[index].Val
		label, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    []plotter.XY{{X: x, Y: y}},
			Labels: []string{str},
		})
		if err != nil {
			log.Fatalf("Could not creates labels plotter: %+v", err)
		}
		p.Add(label)
	}
	path, _ = homedir.Dir()
	path += "/Desktop/GoDS/"
	err = p.Save(1000, 600, "tree.png")
	if err != nil {
		log.Panic(err)
	}
	drawGraph(myApp, myWindow)
	myWindow.ShowAndRun()
}
