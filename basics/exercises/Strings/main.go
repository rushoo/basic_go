package main

import (
	"fmt"
	"unicode/utf8"
)

/*
prints the following (up to 100 characters):
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
*/
func printAs(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}

func countString(str string) {
	//golang的字符串实际上是字符数组，一个中文字符或者符号占3个字节长度
	//若想直观地计算字符串长度，须用utf8.RuneCount或者先转为[]rune
	fmt.Println("\nlen([]byte(str)) == len(str) ?", len([]byte(str)) == len(str))
	fmt.Printf("字符串 (12345678“你好”！) 的字符长度为：%d\n", len(str))
	fmt.Printf("字符串 (12345678“你好”！) 的直观长度为：%d\n", len([]rune(str)))
	fmt.Printf("字符串 (12345678“你好”！) 的直观长度为：%d", utf8.RuneCount([]byte(str)))
}
func reverseString(str string) {
	// 含中文字符串，这里使用rune，否则字符串转码时一个中文字符转码成了三个字符，会出现乱码
	stringSlice := []rune(str)
	var reverseSlice []rune
	for i := len(stringSlice); i > 0; i-- {
		reverseSlice = append(reverseSlice, stringSlice[i-1])
	}
	fmt.Println(string(reverseSlice))
}

func averageFloat(numbers []float64) {
	var avg, sum float64
	switch len(numbers) {
	case 0:
		avg = 0.0
	default:
		for _, num := range numbers {
			sum = +num
		}
		avg = sum / float64(len(numbers))
	}
	fmt.Println(avg)
}
func main() {
	printAs(5)

	str := "12345678“你好”！" //含汉字、中文冒号和感叹号
	countString(str)
	reverseString(str)

	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	averageFloat(numbers)

}
