package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"stack/slicestack"
)

/*
问题描述：
我们用 0 和 1 的二维列表表示一个迷宫。值为 1 的单元格表示阻挡迷宫路径的障碍物。值为 0 的单元格表示可能的迷宫路径位置。
给定这样的 0 和 1 矩阵文件，并给定起始位置和结束位置，目标是从开始位置找到结束位置的路径（如果存在一个或多个此类路径）。
以及使用暴力策略来枚举每个可能的起点到终点路径。

在迷宫路径上的任意位置，可以向八个相邻位置（东、南、西、北、东南、东北、西南、西北）为 0（表示打开）的地方移动。
我们让程序在开放的相邻位置中随机选择，每次运行时可能会生成不同的可行路径。
由于路径不能多次访问同一位置，因此将路径经过的每个单元格的值由 0 设置为 1 .
每次移动后，我们将当前位置以及移动的方向存入堆栈。
如果陷入死胡同，则回溯并访问最后一个安全位置，然后从那里继续。

迷宫算法将使用堆栈作为其核心控制机制：
1. 加载迷宫文件，行数，列数和开始和结束位置。
2. 初始化path-stack用于存放path（路径对象）。
3. 一个路径对象包含迷宫内的坐标、当前移动方向和可用移动方向的列表。
4. 从开放的邻近位置中选择一个初始移动方向。
5. 每次做方向尝试时，对应地将选择的方向从八个可能方向的列表中删除。
6. 从起点构造一个新的路径对象，一个初始移动方向和剩余移动方向的列表。
7. 将初始路径对象入栈。
8. 当堆栈不为空时，弹出堆栈获取路径对象。
9. 开启循环：
   当 当前的路径对象有更多可用的移动时，随机选择一个可用位置并将其值由 0 设置为 1。
   构造一个新的路径对象并将其压入堆栈。当堆栈不为空时，弹出堆栈获取路径对象。

1111111111111111111111111111111111111111
1011110111111111111111111111111111111111
1101101111111111111111111111111111111111
1000011111111111111110111111100001111111
1111011111111111111110111111111111111111
1111101111111111111110111111111111111111
1111110111111111111110111111111111111111
1111101011111111111110111111111111111111
1111011100011111111110111111111111111111
1111011111101111111110111111111111111111
1111011111101111111110111111111111111111
1111011111110111111110111111111111111111
1111011111110000011101111111111111111111
1111101111111111000011111111111111111111
1111100111111111101111111111111111111111
1111110011111111101111111111111111111111
1111101111111111101111111111111111111111
1111011111111111110000011110000000000111
1111101111111111111111000000111111111111
*/

type Direction int

const (
	N            Direction = iota
	NE                     = 1
	E                      = 2
	SE                     = 3
	S                      = 4
	SW                     = 5
	W                      = 6
	NW                     = 7
	NotAvailable           = 8
)

// 重写String方法，打印输出时自动转换code为相应的方向
func (d Direction) String() string {
	switch d {
	case 0:
		return "north"
	case NE:
		return "north-east"
	case E:
		return "east"
	case SE:
		return "south-east"
	case S:
		return "south"
	case SW:
		return "south-west"
	case W:
		return "west"
	case NW:
		return "north-west"
	case NotAvailable:
		return "not available"
	default:
		return "unknown"
	}
}

func (d Direction) PrintDirection() {
	fmt.Println("direction: ", d)
}

// ******************************

// Point abstraction
type Point struct {
	x, y int
}

func (p Point) Equals(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func (p Point) PrintPoint() {
	fmt.Printf("<%d, %d>\n", p.x, p.y)
}

var None = Point{-1, -1}

// *********************************

// Path abstraction
type Path struct {
	point          Point
	move           Direction
	movesAvailable []Direction
}

func NewPath(point Point) Path {
	path := Path{point, NotAvailable, []Direction{}}
	path.move = NotAvailable
	// Initially all directions available
	path.movesAvailable = []Direction{0, NE, E, SE, S, SW, W, NW}
	return path
}

func (path *Path) RandomMove() Direction {
	// Sets value of move
	indicesAvailable := []int{}
	for index := 0; index < 8; index++ {
		if path.movesAvailable[index] != NotAvailable {
			indicesAvailable = append(indicesAvailable, index)
		}
	}
	count := len(indicesAvailable)
	if count > 0 {
		randomIndex := rand.Intn(count)
		path.move = path.movesAvailable[indicesAvailable[randomIndex]]
		path.movesAvailable[indicesAvailable[randomIndex]] = NotAvailable
		return path.move
	} else {
		return NotAvailable
	}
}

// ********************************

type Maze struct {
	rows, cols int
	start, end Point
	mazefile   string
	barriers   [][]bool
	current    Path
	moveCount  int
	pathStack  slicestack.Stack[Path]
	gameOver   bool
}

func NewMaze(rows, cols int, start, end Point, mazefile string) (maze Maze) {
	maze.rows = rows
	maze.cols = cols
	maze.start = start
	maze.end = end

	// golang的二维slice，须先定义一个二维slice，再循环定义slice的slice
	maze.barriers = make([][]bool, rows)
	for i := range maze.barriers {
		maze.barriers[i] = make([]bool, cols)
	}
	//逐行读取迷宫文件mazefile
	file, err := os.Open(mazefile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textLines []string
	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}
	//将读取到的迷宫文件内容与maze.barriers相关联
	for row := 0; row < rows; row++ {
		line := textLines[row]
		for col := 0; col < cols; col++ {
			if string(line[col]) == "1" {
				maze.barriers[row][col] = true
			} else {
				maze.barriers[row][col] = false
			}
		}
	}
	maze.current = NewPath(start)
	maze.pathStack = slicestack.Stack[Path]{}
	maze.pathStack.Push(maze.current)
	//将起始点设置为障碍
	maze.barriers[start.x][start.y] = true
	return maze
}

func NewPosition(oldPosition Point, move Direction) Point {
	switch move {
	case N:
		return Point{oldPosition.x, oldPosition.y - 1}
	case NE:
		return Point{oldPosition.x + 1, oldPosition.y - 1}
	case E:
		return Point{oldPosition.x + 1, oldPosition.y}
	case SE:
		return Point{oldPosition.x + 1, oldPosition.y + 1}
	case S:
		return Point{oldPosition.x, oldPosition.y + 1}
	case SW:
		return Point{oldPosition.x - 1, oldPosition.y + 1}
	case W:
		return Point{oldPosition.x - 1, oldPosition.y}
	default: //NW
		return Point{oldPosition.x - 1, oldPosition.y - 1}
	}
}

func (m *Maze) StepAhead() (Point, Point) {
	validMove := false
	backTrackPoint, newPos := None, None
	for {
		if m.gameOver || validMove || m.pathStack.IsEmpty() {
			break
		}
		validMove = false
		m.current = m.pathStack.Pop()
		nextMove := m.current.RandomMove()
		m.moveCount += 1
		for {
			if validMove || nextMove == NotAvailable {
				break
			}
			newPos = NewPosition(m.current.point, m.current.move)
			if m.barriers[newPos.y][newPos.x] == false {
				validMove = true
				if newPos.Equals(m.end) {
					for {
						if m.pathStack.IsEmpty() {
							break
						}
						m.pathStack.Pop()
					}
					m.gameOver = true
				}
				m.barriers[newPos.y][newPos.x] = true
				m.pathStack.Push(m.current)
				newPathObject := NewPath(newPos)
				m.pathStack.Push(newPathObject)
			} else {
				nextMove = m.current.RandomMove()
			}
		}
		if !validMove && !m.pathStack.IsEmpty() {
			fmt.Printf("\nBacktrack from %v to %v\n",
				m.current.point,
				m.pathStack.Top().point)
			backTrackPoint = m.pathStack.Top().point
		}
	}
	if m.pathStack.IsEmpty() {
		fmt.Println("No solution is possible")
		return None, None
	}
	return newPos, backTrackPoint
}
