package main

/*
最大公共子序列：
方法，将两个公共子序列类似背包问题那样构造二维数组，然后逐行从左到右与列元素相比对
*/
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func reverse(x []rune) []rune {
	var result []rune
	for index := len(x) - 1; index >= 0; index-- {
		result = append(result, x[index])
	}
	return result
}
func longestCommonSubsequenceTable(x, y []rune) (LCS [][]int) {
	n := len(x)
	m := len(y)
	LCS = make([][]int, n+1)
	for row := 0; row < n+1; row++ {
		LCS[row] = make([]int, m+1)
	}
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if x[row] == y[col] {
				//本次遇到相同的，当前值就等于斜上角的值+1，相当于不含当前各自元素时的公共子串长度
				LCS[row+1][col+1] = 1 + LCS[row][col]
			} else {
				// 否则就选取本行前列或者上行此列较大的数作为当前值
				LCS[row+1][col+1] = max(LCS[row][col+1], LCS[row+1][col])
			}
		}
	}
	return LCS
}

func LongestCommonSequence(x, y []rune) string {
	table := longestCommonSubsequenceTable(x, y)
	var result []rune
	//取字符串长度
	j, k := len(x), len(y)

	//从按照数据填入的方式从表尾回溯
	for {
		if table[j][k] == 0 {
			break
		}
		switch {
		case x[j-1] == y[k-1]:
			//元素相等，斜左上回溯
			result = append(result, x[j-1])
			j--
			k--
		case table[j-1][k] > table[j][k-1]:
			// 上>=左，向上回溯
			j--
		case table[j-1][k] <= table[j][k-1]:
			// 上小于左，向左回溯
			k--
		}
	}
	return string(reverse(result))
}
