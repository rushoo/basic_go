package main

import (
	"fmt"
	"strconv"
)

// adjacency list graph 图的邻接表表示法

type ALGraph[T any] struct {
	headNodes []*Node[T]
}
type Node[T any] struct {
	data T
	next *Node[T]
}

func NewALGraph[T any](items []T) *ALGraph[T] {
	alg := &ALGraph[T]{}
	for _, v := range items {
		node := &Node[T]{v, nil}
		alg.headNodes = append(alg.headNodes, node)
	}
	return alg
}
func main() {
	items := []string{"A", "B", "C", "D"}
	alg := NewALGraph[string](items)

	//	假设现在构造有向图，增加两条边A-->B，A-->C
	var Ai, Bi, Ci int
	for i, item := range items {
		if item == "A" {
			Ai = i
		}
		if item == "B" {
			Bi = i
		}
		if item == "C" {
			Ci = i
		}
	}

	Bnode := &Node[string]{strconv.Itoa(Bi), nil}
	alg.headNodes[Ai].next = Bnode

	//新增一条边，相当于头插法插入链表操作
	Cnode := &Node[string]{strconv.Itoa(Ci), nil}
	Cnode.next = alg.headNodes[Ai].next
	alg.headNodes[Ai].next = Cnode
	//求A节点的连接边
	node := alg.headNodes[Ai].next
	for node != nil {
		index, _ := strconv.Atoi(node.data)
		fmt.Println(alg.headNodes[index].data)
		node = node.next
	}

}
