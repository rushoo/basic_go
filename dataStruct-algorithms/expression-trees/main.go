package main

import "fmt"

func main() {
	operandValues := map[string]float64{"a": 5.0, "b": 2.0, "c": 3.0, "d": 2.0, "f": 4.0, "g": 8, "h": 17, "y": 20, "x": 14, "z": 3}
	infix := "((a+b)+( d)/(f+g)+ h)+ y / (x - z)"
	expressionTree := NewTree(infix)
	fmt.Println("Expression tree evaluates to: ", expressionTree.Evaluate(expressionTree.root, operandValues))
}
