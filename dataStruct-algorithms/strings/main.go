package main

import "fmt"

func main() {
	//s := "abcabcabcabc"
	s := "abcabcabcabc"
	text := "abbbb_ababcabb"
	fmt.Println(strStr(text, s))
	fmt.Println(getNext(s))

}
