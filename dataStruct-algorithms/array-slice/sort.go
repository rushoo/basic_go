package main

/*
一、冒泡排序bubbleSort：
从第一个数开始，依次对每对相邻元素比较大小，若左大于边则交换两数。这样每轮结束时可以把当前最大数放到（相对）最后。
1、n个数只需要排序n-1轮 2、每轮都把最大数置最后，故每轮排序次数递减 3、最优情况是已排序的顺序数列，最差是逆序数列
复杂度分析：
空间：整个冒泡排序过程中，只用了一个可以重复利用的临时变量，用于交换需要调换位置的元素，所以空间复杂度为O(1)。
时间：外层n-1次，内层循环最多的时候执行n-1次，最少的时候执行1次，平均n/2次，时间复杂度是O(n^2)


二、插入排序：
类似于整理手上的扑克牌，开始时以最左边第一张牌为有序数列，用第二张牌与之比较，如果较小则插入左边，较大就插入右边。
每次整牌后左边的有序数列就向右扩充一位。为了代码方便，这里是从有序数列从右到左逐次比较的。
复杂度分析：
空间：整个冒泡排序过程中，只用了一个可以重复利用的临时变量，用于交换需要调换位置的元素，所以空间复杂度为O(1)。
时间：在最坏情况下，数组完全逆序，第i个元素插入时要比较i-1次，因此，最坏情况下的比较次数是 1 + 2 + 3 + ... + (N - 1)，时间复杂度是O(n^2)


三、归并排序
先不断将slice向下折半拆分,然后将最终拆分得到的每一对slice向上顺序合并，这样每次合并结束数列的元素都是顺序的，最终全部元素都是顺序的。
要点是设计拆分和合并。
*/

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0 //通过移动k来给新slice赋值
	//将两个顺序数列合并成为一个顺序数列
	//递归拆分最后被insertSort的子数列必然是顺序数列
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	//当一边遍历结束另一边还有
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

func MergeSort(items []int) []int {
	if len(items) == 0 || len(items) == 1 {
		return items
	}
	//先后对左右半归并排序，再合并一起
	left := MergeSort(items[:len(items)/2])
	right := MergeSort(items[len(items)/2:])
	return merge(left, right)
}
func InsertSort(items []int) {
	end := len(items)
	for i := 1; i < end; i++ {
		tmp := items[i] //临时记录此时参与排序的元素
		j := i - 1      //从有序数列的从右到左开始比较，也相当于每次都向右扩充了一位
		for j >= 0 && tmp < items[j] {
			items[j+1] = items[j] //逆序,则右移一位
			j--
		}
		items[j+1] = tmp
	}
}
func BubbleSort(items []int) {
	end := len(items)
	for i := 0; i < end-1; i++ {
		for j := 0; j < end-i-1; j++ {
			if items[j] >= items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}
