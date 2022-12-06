package main
var c = 0
func inorderLevel(node *Node, level int) {
	if node != nil {
		inorderLevel(node.left, level + 1)
		c += 1
		node.ch += string(c)
		data = append(data, NodePos{node.ch, 100 -
			level, -1})
		inorderLevel(node.right, level + 1)
	}
}
// Add Labels
for index := 0; index < len(data); index++ {
x := float64(data[index].XPos) - 0.1
y := float64(data[index].YPos) - 0.02
str := data[index].Val
label, err := plotter.NewLabels(plotter.XYLabels {
XYs: []plotter.XY {
{X: x ,Y: y},
},
Labels: []string{string(str[0])},
},)
if err != nil {
log.Fatalf("could not creates labels
plotter: %+v", err)
}
p.Add(label)
}