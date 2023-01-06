package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	v1 := strings.EqualFold("hello", "Hello")
	v2 := bytes.EqualFold([]byte("hello"), []byte("Hello"))
	fmt.Println(v1, v2)
	fmt.Println(strings.ToTitle("Hello"))
	fmt.Println(strings.Index("Hello", "l"))
	fmt.Println(unicode.ToLower('H'))

	description := "A boat for one person"
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB)) //输出b/B在字符串中首次出现的位置
}
