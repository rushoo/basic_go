package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "你好！ the sky is blue. 看 看"
	s := "  the sky is blue"
	fmt.Println(reverseWords2(s))
}

func rmSpace2(s string) []string {
	var rs []string
	var tmp = ' ' //这样相当于跳过了开头的全部空格
	i, runeLen := 0, 0
	for j, v := range s {
		if v == tmp && v == ' ' {
			i++
		} else if v == ' ' {
			rs = append(rs, string(s[i:j]))
			rs = append(rs, " ")
			i = j + 1
		}
		tmp = v
		runeLen = j
	}
	if i <= runeLen {
		rs = append(rs, string(s[i:]))
	}
	//去除最后一个空格
	if rs[len(rs)-1] == " " {
		rs = rs[:len(rs)-1]
	}
	return rs
}

// 字符串整个按字符翻转
func rvByString(rs []string) {
	l, r := 0, len(rs)-1
	for l < r {
		rs[l], rs[r] = rs[r], rs[l]
		l++
		r--
	}
}
func reverseWords2(s string) string {
	rs := rmSpace2(s)
	fmt.Println(rs)
	rvByString(rs)
	return strings.Join(rs, "")
}
