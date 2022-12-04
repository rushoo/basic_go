package main

import (
	"fmt"
	"strings"
)

/*
---表达式树---
后缀表达式的运算规则是，将后缀表达式元素依次取出进栈，当遇到运算符元素时，
从栈顶pop()出最近两个元素，并用该运算符计算结果，再将结果入栈。
表达式树可以根据这个思路，构造双指针节点，其中操作数的指针为空，
每一个运算符节点的指针就近指向栈顶两个操作数并出栈，再将新节点入栈。
这样最终栈里仅剩一个元素，并以该元素为root节点形成表达式树，表达式的运算是节点值和左右两个指针。
(显然的特点是，表达式树的任意节点，若有子，就同时有双子)

中缀转后缀表达式：
从左到右遍历中缀表达式，遇到操作数，存到后缀表达式，
遇到操作符，当前操作符的优先级大于栈顶元素优先级，进栈，
否则，依次弹出栈顶优先级大于等于当前操作符的元素，然后当前操作符进栈。
当中缀表达式遍历结束，将操作符栈所有元素pop出到后缀表达式
优先级：
右括号 > *,/ > +,- > 左括号
遇到左括号直接进栈
遇到右括号时需要一直弹出栈顶元素直到弹出对应的左括号为止

运算规则：
扫描后缀表达式，将操作数入栈，如遇到运算符，取出栈顶两个操作数运算，并将运算结果入栈
*/

type Node struct {
	ch    string
	left  *Node
	right *Node
}

type ExpressionTree struct {
	postfix string
	root    *Node
}

var opWeight map[string]int //运算符优先级权重表
func init() {
	opWeight = make(map[string]int)
	opWeight["("] = 0
	opWeight["+"] = 1
	opWeight["-"] = 1
	opWeight["*"] = 2
	opWeight["/"] = 2
}
func priority(ch string) int {
	return opWeight[ch]
}

// 判断是否期望操作数
func isOperand(s string) bool {
	if s >= "a" && s <= "z" {
		return true
	}
	return false
}

// 判断是否为规定的运算符
func isOperator(s string) bool {
	operators := []string{"+", "-", "*", "/", "(", ")"}
	for _, v := range operators {
		if v == s {
			return true
		}
	}
	return false
}
func infixToPostfix(infix string) (postfix string) {
	postfix = "" //后缀表达式，通过字符串连接扩展
	nodeStack := sliceStack[string]{}
	//遍历表达式
	for index := 0; index < len(infix); index++ {
		new := string(infix[index])
		switch {
		case new == " " || new == "\n": //无操作，跳过空格与换行符
		case isOperand(new):
			postfix += new
		case isOperator(new):
			{
				if nodeStack.IsEmpty() || new == "(" {
					nodeStack.Push(new)
				} else { //非空栈
					switch {
					case new == ")": //遇到右括号时，需要一直弹出栈顶元素直到弹出对应的左括号为止
						for !nodeStack.IsEmpty() && nodeStack.Top() != "(" {
							postfix += nodeStack.Pop()
						}
						nodeStack.Pop() //弹出"("
					case priority(new) > priority(nodeStack.Top()):
						nodeStack.Push(new)
					default: //new优先级不高于栈顶元素，依次弹出栈顶优先级大于等于当前操作符的元素，然后当前操作符进栈。
						for !nodeStack.IsEmpty() && priority(new) <= priority(nodeStack.Top()) {
							postfix += nodeStack.Pop()
						}
						nodeStack.Push(new)
					}
				}
			}
		}
	}
	// 遍历结束后将所有的操作符出栈
	for !nodeStack.IsEmpty() {
		postfix += nodeStack.Pop()
	}
	return postfix
}

func NewTree(infix string) (tree *ExpressionTree) {
	infix = strings.ToLower(infix)
	tree = &ExpressionTree{"", nil}
	tree.postfix = infixToPostfix(infix)
	fmt.Println(tree.postfix)
	stack := nodeStack[*Node]{}
	str := strings.Split(tree.postfix, "")
	for index := 0; index < len(str); index++ {
		item := str[index]
		if isOperand(item) {
			node := &Node{item, nil, nil}
			stack.Push(node)
		} else if isOperator(item) {
			left := stack.Pop()
			right := stack.Pop()
			//想着遇到空余符号时候前边补0，其实这样没什么意义，并不保障运算结果
			//zero := &Node{"0", nil, nil}
			//if left == nil {
			//	left = zero
			//}
			//if right == nil {
			//	right = zero
			//	//交换两数，相当于在孤独运算符前面补0
			//	left, right = right, left
			//}
			node := &Node{item, nil, nil}
			node.left = left
			node.right = right
			stack.Push(node)
		}
	}
	tree.root = stack.Top()
	//stack.Pop()					//因为栈中仅有一个元素
	//fmt.Println(stack.Top().ch)   //所以这样会nil pointer panic
	return tree
}
func (tree *ExpressionTree) Evaluate(node *Node, operandValues map[string]float64) float64 {
	if node == nil {
		return 0.0
	}
	if node.left == nil && node.right == nil {
		value := operandValues[node.ch]
		return value
	}
	leftValue := tree.Evaluate(node.left, operandValues)
	rightValue := tree.Evaluate(node.right, operandValues)
	switch node.ch {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	case "*":
		return leftValue * rightValue
	default:
		return leftValue / rightValue
	}
}
