package main

func main() {
	root := Node{"A", nil, nil}
	nodeB := Node{"B", nil, nil}
	nodeC := Node{"C", nil, nil}
	nodeD := Node{"D", nil, nil}
	nodeE := Node{"E", nil, nil}
	nodeF := Node{"F", nil, nil}
	nodeG := Node{"G", nil, nil}
	nodeH := Node{"H", nil, nil}
	nodeI := Node{"I", nil, nil}
	nodeJ := Node{"J", nil, nil}
	nodeK := Node{"K", nil, nil}
	nodeL := Node{"L", nil, nil}
	nodeM := Node{"M", nil, nil}
	nodeN := Node{"N", nil, nil}
	nodeO := Node{"O", nil, nil}
	nodeP := Node{"P", nil, nil}
	nodeQ := Node{"Q", nil, nil}
	nodeR := Node{"R", nil, nil}
	root.Left = &nodeB
	root.Right = &nodeC
	nodeB.Left = &nodeD
	nodeD.Right = &nodeH
	nodeD.Left = &nodeE
	nodeE.Left = &nodeF
	nodeE.Right = &nodeG
	nodeC.Right = &nodeI
	nodeC.Left = &nodeJ
	nodeI.Right = &nodeK
	nodeK.Left = &nodeL
	nodeL.Left = &nodeM
	nodeL.Right = &nodeN
	nodeN.Right = &nodeO
	nodeO.Left = &nodeP
	nodeO.Right = &nodeQ
	nodeM.Left = &nodeR
	myTree := BinaryTree{&root, 18}
	ShowTreeGraph(myTree)
}
