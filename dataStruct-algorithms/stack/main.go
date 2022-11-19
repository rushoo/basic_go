package main

import (
	"fmt"
	"stack/nodestack"
	"stack/slicestack"
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
//遇到左括号直接进栈
//遇到右括号时需要一直弹出栈顶元素直到弹出对应的左括号为止

var (
	values   map[string]float64
	opWeight map[string]int //运算符优先级权重表
)

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

func evaluate(postfix string) float64 {
	//扫描后缀表达式，将操作数入栈，如遇到运算符，取出栈顶两个操作数运算，并将运算结果入栈
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

// 使用辗转相除法
func convertToBinary(input int) (binary []int) {
	binary = []int{}
	sliceStack := slicestack.Stack[int]{}

	if input == 0 {
		sliceStack.Push(0)
	}
	for input > 1 {
		remainder := input % 2 //取余
		sliceStack.Push(remainder)

		input = input / 2 //取整重置
	}
	// 输入是1,或者商是1
	if input == 1 {
		sliceStack.Push(1)
	}
	for !sliceStack.IsEmpty() {
		binary = append(binary, sliceStack.Pop())
	}
	return binary
}
func main() {
	fmt.Println(opWeight)
	postfix := infixToPostfix("a + (b - c) / (d * e)")
	postfix2 := infixToPostfix("a * (b - c) / d * ( b + c - e)")
	fmt.Println(postfix)
	fmt.Println(postfix2)
	values = make(map[string]float64)
	values["a"] = 10
	values["b"] = 5
	values["c"] = 2
	values["d"] = 4
	values["e"] = 3
	result := evaluate(postfix)
	result2 := evaluate(postfix2)
	fmt.Println("function evaluates to: ", result)
	fmt.Println("function evaluates to: ", result2)

	b := convertToBinary(1000000)
	fmt.Println(b)
}
