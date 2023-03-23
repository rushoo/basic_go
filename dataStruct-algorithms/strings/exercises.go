package main

import (
	"fmt"
)

/*
344. 反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
要求：不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

思路：
定义两个指针（也可以说是索引下标），一个从字符串前面，一个从字符串后面，两个指针同时向中间移动，并交换元素。
*/
func reverseString(s []byte) {
	var l, r = 0, len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

/*
541. 反转字符串II
给定一个字符串 s 和一个整数 k，从字符串开头算起, 每计数至 2k 个字符，就反转这 2k 个字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
可以假设s仅由小写字母构成。

输入：s = "abcdefg", k = 2   输出："bacdfeg"
输入：s = "abcd", k = 2 	输出："bacd"

思路：
题目意思是依次遍历2k然后将前k个字符反转。问题在于边界，即最后一次2k遍历后还剩下多少字符。
当剩余字符>=k时，反转前k(这种情况包含了依次遍历2k然后将前k个字符反转)；
当剩余字符<k时，全部反转。
*/
// 做字符串翻转时如果要维持汉字的完整性，就需要按照rune来遍历翻转，不过这意味着更大的空间消耗
func reverse(rs []rune) {
	var l, r = 0, len(rs) - 1
	for l < r {
		rs[l], rs[r] = rs[r], rs[l]
		l++
		r--
	}
}
func reverseStr(s string, k int) string {
	// 不能在string上直接操作，需要先转为string-slice
	ss := []rune(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if length-i >= k {
			//剩余大于等于k，反转此起k个
			reverse(ss[i : i+k])
		} else {
			//剩余小于k，将剩余全部反转
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

/*
剑指Offer 05.替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1： 输入：s = "We are happy."		输出："We%20are%20happy."
思路：
string类型在golang中以utf-8的编码形式存在，按字节读取是byte，按字符读取是rune。
这里"%20"在string中应该是三个字符，应该分开处理。
可以按字节读取，也可以按字符读取，只是对应的容器不同。

还有种方式是将计算string的空格数m，然后将copy到新的容量为length+2m的rune/byte-slice中
指针1从尾部length-1向前遍历，指针2从新slice尾部向前，遇到空格就替换，否则就直接填上。
在go中因为string的可读性好，不需要这么麻烦。
*/
func replaceSpace(s string) string {
	var re []rune
	for _, v := range s {
		if v == ' ' {
			re = append(re, '%', '2', '0')
		} else {
			re = append(re, v)
		}
	}
	return string(re)
}

// 按字节读取
func replaceSpace2(s string) string {
	var re []byte
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			//re = append(re, "%20"...)
			re = append(re, '%', '2', '0')
		} else {
			re = append(re, s[i])
		}
	}
	return string(re)
}

/*
151. 反转字符串中的单词
给定一个字符串，逐个翻转字符串中的每个单词。

输入: "the sky is blue"	输出: "blue is sky the"

输入: "  hello world!  "		输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

输入: "a good   example"		输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

思路如下：
移除多余空格-->将整个字符串反转-->将每个单词反转
移除多余空格：第一个空格删除、前面是空格的删除、最后一个空格删除
*/
func reverseWords(s string) string {
	var re []rune
	for i, v := range s {
		if (i == 0 && v == ' ') || (i > 0 && v == ' ' && s[i-1] == ' ') {
			//移除多余空格：第一个空格删除、前面是空格的空格删除
		} else {
			re = append(re, v)
		}
	}
	//删除最后一个空格
	if re[len(re)-1] == ' ' {
		re = re[:len(re)-1]
	}
	fmt.Println(string(re))
	reverse(re)
	reverseByWord(re)
	return string(re)
}
func reverseByWord(rs []rune) {
	//暂存空格位置
	var tmp int
	for i, v := range rs {
		if v == ' ' {
			reverse(rs[tmp:i])
			tmp = i + 1
		}
	}
	//翻转最后一个word
	reverse(rs[tmp:])
}

/*
题目：剑指Offer58-II.左旋转字符串
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。

输入: s = "abcdefg ", k = 2		输出: "cdefg ab"
输入: s = "lrloseumgh ", k = 6	输出: "umgh lrlose"

思路：先将后b个字符读到slice，再将前a个字符读到slice
*/

/*
28. 实现 strStr()
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2
示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1
思路:
kmp算法
*/

/*
代码思路：
https://www.bilibili.com/video/BV1AY4y157yL/?spm_id_from=333.999.0.0

生成next数组的过程说明：
对于next[n]==m,意味对于已有的n+1个字符，前m个和后m个是完全一样的（也是最大等同）
那么对于新来的一个字符，可以直接将其与第m+1个字符也就是s[m]相比较，如果两者相等，就
意味着前m+1与后m+1个字符是一样的。如果不相等，就求出前m-1个字符和后m-1个字符的相同
子串，以便在此基础上扩充。问题即回到将新来的字符在前m-1个字符的共同子串上比较的问题。
这类似于一种递归过程。

例如 s="ABACABAB"
next=[0 0 1 0 1 2 3 2]
将最后一个字符s[7]=B加入时，此时的next数组最后一位元素next[6]=3, 这意味着已经处理的子串"ABACABA"
前三个元素s[0,1,2]与后三个元素s[4,5,6]组成的串是相等的(这里都是"ABA")，所以将s[7]与s[ next[6] ]
直接比较的意义就是利用已有的相同缀在其上直接扩充，这里s[7] != s[ next[6] ],也就是不能直接扩充，那么
继续检查是否可以在它的子序列上直接扩充，所以接着看m-1.前m-1=2个也就是前2是AB，对应next[2]=1,也就是说前2的BA
和后2的AB可以1个元素A，s[7]可以在此前提上比较，而s[7]=s[1],可以直接入串，故next[7]=2.
假设s[7]=C,则s[7]!=s[1]，可公用的基础上再-1，1-1=0，也就是无可公用，那么直接比较s[7]和s[0]即可。

例如	s="ABABC"
							next:
初始无前后缀，值当然为			0
可共用长度0，直接比较	A=!s[0]	0
可共用长度0，直接比较	A=s[0]	1
可共用长度1，补充比较	B=s[1]	2
可共用长度2，补充比较	B!=s[2]
	长度-1考虑可共用，2-1=1，next[1]=0，
	可共用长度0，直接比较	C=!s[0]	 0

另一个问题是模式串与文本串匹配时的跳转，例如：
	"abbabababcabb"    ____		 "abbabababcabb"
	"ababc"            			   "ababc"
当比较到模式串第3个元素时，出现了元素不相等，next=[0 0 1 2 0]，记下此时不相等的前一个位置next[1]=0,右移0

	"abbabababcabb"  ____   "abbabababcabb"  ____   "abbabababcabb"
	  "ababc"                 "ababc"                  "ababc"
当比较到模式串第5个元素时，出现了元素不相等，记下此时不相等的前一个位置next[3]=2,右移2
	 "abbabababcabb"
		  "ababc"       右移后继续比较发现模式串全部匹配，也就得到了第一次在文本中出现模式串

关于逻辑上右移的实现，比如右移2前，是从文本串中第4个字符开始做比对，因此得到next数组上冲突值为next[3]=2，
那么从文本串中第5个字符开始做比对时，就直接从模式串的第2+1个字符开始比对，逻辑上就相当于跳过了两个同缀字符的比对。
KMP算法能够实现在文本串上指针持续向后扫描， 比较的过程伴随着指针向后移动。
其中假设文本串长度n，模式串长度m，生成next数组时间复杂度是m，匹配是n，共计m+n  ---省略了（O）
*/
//例如 ABACABAB
func getNext(s string) []int {
	next := make([]int, len(s))
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := getNext(needle)
	for i := 0; i < len(haystack); i++ {
		// 将模式串想象成一个滑动窗口
		// 从第模式串第j+1个位置开始比较，相等于向右滑动了j位
		for j > 0 && haystack[i] != needle[j] {
			// 比对失败时记下next数组中前一个元素的值
			j = next[j-1]
		}
		// 将模式串与文本串逐一比对
		if haystack[i] == needle[j] {
			j++
		}
		// 将模式串比对完了元素都相等，返回本次比对模式串起始位置
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

/*
459.重复的子字符串
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。

示例 1:
输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。

示例 2:
输入: "aba"
输出: False

示例 3:
输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

思路：
假设字符串重复组成，元字符串为X，则有 X * n = Y，
可能是n=2，即X+X=Y,这样有next[last]=len(X),len(Y)-next[last]len(X)
可能是n=m>2,即X+X+...+X+X=Y，next[last]=len(X+...+X+X)，m-1个X，len(Y)-next[last]=len(X)
所以如为重构字符串，串长度就应该是元子串长度的2倍或者m倍，即:
len(Y) % (len(Y)-next[last])==0,当next[last]!=0时排整除结果为1的情况。
*/
func genNext(s string) []int {
	next := make([]int, len(s))
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		// 不可直接共用
		for j > 0 && s[i] != s[j] {
			//就缩减长度再得到可共用的部分
			j = next[j-1]
		}
		// 在可共用的基础上直接扩展
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}
func repeatedSubstringPattern(s string) bool {
	all := len(s)
	next := genNext(s)
	rest := all - next[all-1]
	if next[all-1] != 0 && all%rest == 0 {
		return true
	}
	return false

}
