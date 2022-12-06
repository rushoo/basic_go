package main

/*
0-1背包问题
f[i][j]表示只看前i个物品，总体积是j的情况下，背包内物品的最大价值
考虑f[i][j]怎么计算：
1、不选第i件物品，那么f[i][j]==f[i-1][j]
2、选第i件物品，f[i][j]==f[i-1][j-v[i]]+p[i],
也就是去掉物品i的体积后，背包可用容量j-v[i]在i-1件物品可选时的最大价值 + 第i件物品的价值p[i]

f[i][j]=max{ f[i-1][j], f[i-1][j-v[i]]+p[i] }
*/

func maxWorthGoodsOnce(volume, worth []int, pkgVol int) int {
	len := len(volume)
	table := make([][]int, len+1)
	for i := 0; i <= len; i++ {
		table[i] = make([]int, pkgVol+1)
	}
	//	背包装物品
	for i := 1; i <= len; i++ {
		for j := 0; j <= pkgVol; j++ {
			if j == 0 {
				table[i][j] = 0
				continue
			}
			table[i][j] = table[i-1][j] //第i个元素不能选
			///列表里第i个元素在商品列volume中就是第i-1个元素,对应的体积vi
			vi := volume[i-1]
			if j >= vi {
				// 当有空间可用时，表示可选，但是否会选择要根据价值而定，也就是对两个取max
				table[i][j] = max(table[i][j], table[i-1][j-vi]+worth[i-1])
			}
		}
	}
	return table[len][pkgVol]
}
