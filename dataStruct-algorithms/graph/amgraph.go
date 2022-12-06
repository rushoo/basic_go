package main

// Adjacency Matrix Graph  图的邻接矩阵表示法

import (
	"fmt"
	"math/rand"
	"time"
)

// 通过一个给定类型的数组/slice构造一个AMGraph图，其节点为给定数组的每一个元素，连接关系表示为二维数组里每个整数值
type dataType interface {
	~int | ~float64 | ~string
}
type AMGraph[T dataType] struct {
	nodes []T     //NodesMatrix[T]
	table [][]int //AdjacencyMatrix
}

func NewAMGraph[T dataType](nodes []T) *AMGraph[T] {
	length := len(nodes)
	table := make([][]int, length)
	for i := 0; i < length; i++ {
		table[i] = make([]int, length)
	}
	//对于有权图，可以先在此给每个元素一个较大的值表示为不可达
	return &AMGraph[T]{
		nodes: nodes,
		table: table,
	}
}

// 求有向图第k个节点的度
func sum[T dataType](amg *AMGraph[T], k int) int {
	if k >= len(amg.table) {
		//	数组越界
		return 0
	}
	sum := 0
	for i := 0; i < len(amg.table); i++ {
		if amg.table[k][i] == 1 {
			sum++
		}
		if amg.table[i][k] == 1 {
			sum++
		}
	}
	return sum
}
func main() {
	rand.Seed(time.Now().UnixNano())
	nodes := []string{"A", "B", "C", "D", "E"}
	amg := NewAMGraph[string](nodes)
	len := len(nodes)
	// 填充有向图
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			amg.table[i][j] = rand.Intn(2)
			if i == j {
				amg.table[i][j] = 0
			}
		}
	}
	// 求节点B的度,也就是第2行和第2列节点为1的数量
	fmt.Printf("第2个节点%s的度为%d", amg.nodes[1], sum(amg, 2))
}
