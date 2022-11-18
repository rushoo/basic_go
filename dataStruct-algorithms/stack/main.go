package main

import (
	"fmt"
	"stack/nodestack"
)

const size = 10_000_000

//func main() {
//	//Time for 10 million Push() operations on nodeStack:  416.4546ms
//	//
//	//Time for 10 million Pop() operations on nodeStack:  20.7493ms
//	//
//	//Time for 10 million Push() operations on sliceStack:  137.2793ms
//	//
//	//Time for 10 million Pop() operations on sliceStack:  15.0125ms
//
//	nodeStack := nodestack.Stack[int]{}
//	sliceStack := slicestack.Stack[int]{}
//
//	start := time.Now()
//	for i := 0; i < size; i++ {
//		nodeStack.Push(i)
//	}
//	elapsed := time.Since(start)
//	fmt.Println("\nTime for 10 million Push() operations on nodeStack: ",
//		elapsed)
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		nodeStack.Pop()
//	}
//	elapsed = time.Since(start)
//	fmt.Println("\nTime for 10 million Pop() operations on nodeStack: ",
//		elapsed)
//	// Benchmark sliceStack
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		sliceStack.Push(i)
//	}
//	elapsed = time.Since(start)
//	fmt.Println("\nTime for 10 million Push() operations on sliceStack: ", elapsed)
//	start = time.Now()
//	for i := 0; i < size; i++ {
//		sliceStack.Pop()
//	}
//	elapsed = time.Since(start)
//	fmt.Println("\nTime for 10 million Pop() operations on sliceStack: ", elapsed)
//}

//中缀转后缀表达式：
// 从左到右遍历中缀表达式，遇到操作数，存到后缀表达式，
// 遇到操作符，当前操作符的优先级大于栈顶元素优先级，进栈，
// 否则，依次弹出栈顶优先级大于等于当前操作符的元素，然后当前操作符进栈。
// 当中缀表达式遍历结束，将操作符栈所有元素pop出到后缀表达式
//优先级：
//右括号 > *,/ > +,- > 左括号
//遇到右括号时需要一直弹出栈顶元素直到弹出对应的左括号为止

var opWeight map[string]int

func init() {
	opWeight := make(map[string]int)
	opWeight["("] = 0
	opWeight["+"] = 1
	opWeight["-"] = 1
	opWeight["*"] = 2
	opWeight["/"] = 2
}

// 优先级检查,将以precedence(top,new)表示栈顶元素优先级大于等于新来元素
func precedence(s1, s2 string) bool {
	return opWeight[s1] >= opWeight[s2]
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

// 转后缀表达式
func infixToPostfix(infix string) (postfix string) {
	postfix = "" //后缀表达式，通过字符串连接扩展
	nodeStack := nodestack.Stack[string]{}
	//遍历表达式
	for index := 0; index < len(infix); index++ {
		new := string(infix[index])
		//跳过空格
		if new == " " {
			continue
		}
		//操作数直接并入后缀表达式
		if new >= "a" && new <= "z" {
			postfix += new
		}
		//	操作符
		if isOperator(new) {
			//	空栈直接入
			if nodeStack.IsEmpty() {
				nodeStack.Push(new)
			} else {
				//遇到右括号时需要一直弹出栈顶元素直到弹出对应的左括号为止
				if new == ")" {
					for nodeStack.Top() != "(" {
						postfix += nodeStack.Pop()
					}
					nodeStack.Pop() //弹出"("
				}
				// 当前操作符的优先级大于栈顶元素优先级，进栈，
				if !precedence(nodeStack.Top(), new) {
					nodeStack.Push(new)
				} else {
					// 否则，依次弹出栈顶优先级大于等于当前操作符的元素，然后当前操作符进栈。
					for precedence(nodeStack.Top(), new) {
						postfix += nodeStack.Pop()
					}
					nodeStack.Push(new)
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

// -
// -
// -
// -
// -
// -
func precedence1(symbol1, symbol2 string) bool {
	//precedence(topSymbol, newSymbol)
	// Returns true if symbol1 has a higher precedence than symbol2
	if (symbol1 == "+" || symbol1 == "-") && (symbol2 == "(" || symbol2 == "/") {
		return false
	} else if (symbol1 == "(" && symbol2 != ")") || symbol2 == "(" {
		return false
	} else {
		return true
	}
}

// 判断是否为规定的运算符
func isPresent(symbol string, operators []string) bool {
	for i := 0; i < len(operators); i++ {
		if symbol == string(operators[i]) {
			return true
		}
	}
	return false
}

// 中缀变后缀表达式
func infixToPostfix2(infix string) (postfix string) {
	operators := []string{"+", "-", "*", "/", "(", ")"}
	postfix = ""
	nodeStack := nodestack.Stack[string]{}

	//从左到右遍历中缀表达式
	for index := 0; index < len(infix); index++ {
		newSymbol := string(infix[index])
		if newSymbol == " " || newSymbol == "\n" {
			// 跳过空格和换行
			continue
		}
		if newSymbol >= "a" && newSymbol <= "z" {
			//遇到操作数，存到后缀表达式
			postfix += newSymbol
		}
		//遇到操作符
		if isPresent(newSymbol, operators) {
			if !nodeStack.IsEmpty() {
				topSymbol := nodeStack.Top()
				//当前操作符的优先级大于栈顶元素优先级，进栈
				if precedence(topSymbol, newSymbol) == true {
					//若栈顶非"("，栈顶元素pop并入后缀表达式，否则直接pop
					if topSymbol != "(" {
						postfix += topSymbol
					}
					nodeStack.Pop()
				}
			}
			if newSymbol != ")" {
				nodeStack.Push(newSymbol)
			} else { // 那么这里要处理的就是newSymbol == ")"
				// Pop nodeStack down to first left parenthesis(括号)
				for {
					if nodeStack.IsEmpty() == true {
						break
					}
					ch := nodeStack.Top()
					if ch != "(" {
						postfix += ch
						nodeStack.Pop()
					} else {
						nodeStack.Pop()
						break
					}
				}
			}
		}
	}
	for {
		if nodeStack.IsEmpty() == true {
			break
		}
		if nodeStack.Top() != "(" {
			postfix += nodeStack.Top()
			nodeStack.Pop()
		}
	}
	return postfix
}

var values map[string]float64

func evaluate(postfix string) float64 {
	operandStack := nodestack.Stack[float64]{}
	for index := 0; index < len(postfix); index++ {
		ch := string(postfix[index])
		if ch >= "a" && ch <= "z" {
			operandStack.Push(values[ch])
		} else { // ch is an operator
			operand1 := operandStack.Pop()
			operand2 := operandStack.Pop()
			if ch == "+" {
				operandStack.Push(operand1 + operand2)
			} else if ch == "-" {
				operandStack.Push(operand2 - operand1)
			} else if ch == "*" {
				operandStack.Push(operand1 * operand2)
			} else if ch == "/" {
				operandStack.Push(operand2 / operand1)
			}
		}
	}
	return operandStack.Top()
}
func main() {
	postfix := infixToPostfix("a + (b - c) / (d * e)")
	fmt.Println(postfix)
	values = make(map[string]float64)
	values["a"] = 10
	values["b"] = 5
	values["c"] = 2
	values["d"] = 4
	values["e"] = 3
	result := evaluate(postfix)
	fmt.Println("function evaluates to: ", result)
}
