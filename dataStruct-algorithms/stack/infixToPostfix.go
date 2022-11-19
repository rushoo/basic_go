package main

import (
	"fmt"
	"stack/nodestack"
)

// 中缀转后缀表达式：
// 从左到右遍历中缀表达式，遇到操作数，存到后缀表达式，
// 遇到操作符，当前操作符的优先级大于栈顶元素优先级，进栈，
// 否则，依次弹出栈顶优先级大于等于当前操作符的元素，然后当前操作符进栈。
// 当中缀表达式遍历结束，将操作符栈所有元素pop出到后缀表达式
// 优先级：
// 右括号 > *,/ > +,- > 左括号
// 遇到左括号直接进栈
// 遇到右括号时需要一直弹出栈顶元素直到弹出对应的左括号为止
//
// 运算规则：
// 扫描后缀表达式，将操作数入栈，如遇到运算符，取出栈顶两个操作数运算，并将运算结果入栈

var opWeight map[string]int   //运算符优先级权重表
var values map[string]float64 //操作数映射表

func init() {
	opWeight = make(map[string]int)
	opWeight["("] = 0
	opWeight["+"] = 1
	opWeight["-"] = 1
	opWeight["*"] = 2
	opWeight["/"] = 2
}

// 用以表示当前操作符的优先级大于栈顶元素优先级precedence(new,top)
func precedence(s1, s2 string) bool {
	return opWeight[s1] > opWeight[s2]
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

// 判断是否期望操作数
func isOperand(s string) bool {
	if s >= "a" && s <= "z" {
		return true
	}
	return false
}

// 转后缀表达式
func infixToPostfix(infix string) (postfix string) {
	postfix = "" //后缀表达式，通过字符串连接扩展
	nodeStack := nodestack.Stack[string]{}
	//遍历表达式
	for index := 0; index < len(infix); index++ {
		new := string(infix[index])
		switch {
		case new == " ": //无操作，跳过空格
		case new >= "a" && new <= "z": //操作数入postfix
			postfix += new
		case isOperator(new): //运算符号
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
					case precedence(new, nodeStack.Top()): // 当前操作符的优先级大于栈顶元素优先级，进栈，
						nodeStack.Push(new)
					default: //new优先级不高于栈顶元素，依次弹出栈顶优先级大于等于当前操作符的元素，然后当前操作符进栈。
						for !nodeStack.IsEmpty() && !precedence(new, nodeStack.Top()) {
							postfix += nodeStack.Pop()
						}
						nodeStack.Push(new)
					}
				}
			}
		}
	}
	// 遍历结束后将所有的操作符入栈
	for !nodeStack.IsEmpty() {
		postfix += nodeStack.Pop()
	}
	return postfix
}

// 运算，扫描后缀表达式，将操作数入栈，如遇到运算符，取出栈顶两个操作数运算，并将运算结果入栈
func evaluate(postfix string) float64 {
	operandStack := nodestack.Stack[float64]{} //操作数栈
	for index := 0; index < len(postfix); index++ {
		ch := string(postfix[index])
		if ch >= "a" && ch <= "z" {
			operandStack.Push(values[ch])
		} else { //遇到运算符
			operand1 := operandStack.Pop()
			operand2 := operandStack.Pop()
			switch ch {
			case "+":
				operandStack.Push(operand1 + operand2)
			case "-":
				operandStack.Push(operand2 - operand1)
			case "*":
				operandStack.Push(operand1 * operand2)
			case "/":
				operandStack.Push(operand2 / operand1)
			default:
				fmt.Println("Invalid operator")
			}
		}
	}
	return operandStack.Top()
}

// 用以建立操作数参数和具体数字的映射表
func assignFixValues(valueSlice []float64) {
	values = make(map[string]float64)
	values["a"] = valueSlice[0]
	values["b"] = valueSlice[1]
	values["c"] = valueSlice[2]
	values["d"] = valueSlice[3]
	values["e"] = valueSlice[4]
}
