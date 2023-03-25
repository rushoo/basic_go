### 快速排序

##### #基本快速排序
```
func quickSort(s []int, low, high int) {
    if low < high {
    	pivot := partition(s, low, high)
    	quickSort(s, low, pivot-1)
    	quickSort(s, pivot+1, high)
    }
}
func partition(s []int, low, high int) int {
    if high-low < 2 {
    	return low
    }
    //通过交换，选定中间元素为基准值
    s[low], s[(high+low)/2] = s[(high+low)/2], s[low]
    var pivot = s[low]
    //双指针分别从数组最左边与最右边出发，指针停止条件：指向数与基准数逆序。
    var l, r = low, high
    for l < r {
    	//先移动右指针，所以最后两指针则会合于较小的数
    	for r > l && s[r] > pivot {
    		r--
    	}
    	for l < r && s[l] <= pivot {
    		l++
    	}
    	if l < r {
    		s[l], s[r] = s[r], s[l]
    	}
    }
    //考虑终止时的情况：
    //	1、l没动，r一路畅通与l重合  2、r和l合于中间某一处  3、r没动，l畅通至与r重合
    // 这三种情况下重合位置都有分界属性，除了第一种实质原地交换，后两种情况都指向了较小值。
    // 所以直接直接交换选定基准值和终止元素值即可，终止索引为期望分界标。
    s[low] = s[r]
    s[r] = pivot
    return r
}
```

##### #改进方法1
优先处理较长数列，以及数组元素较少时直接插入，可以避免递归栈过深。取`insertSort(100/50/30)`,改进前后递归次数对比为:   
`[1333265,32767],[1333569,65535],[1333207,68523]`,但执行完成时间基本没有区别，只是增加了内存安全。
```
func quickSort2(s []int, low, high int) {
    //当元素较少是时，直接做插入排序
    if high-low < 50 {
    	//slice左闭右开
    	insertSort(s[low : high+1])
    	return
    }
    if low < high {
    	p := partition(s, low, high)
    	if p-low > high-p {
    		quickSort3(s, low, p-1)
    		quickSort3(s, p+1, high)
    	} else {
    		quickSort3(s, p+1, high)
    		quickSort3(s, low, p-1)
    	}
    }
}
```

##### #改进方法2
快速排序的子过程中，处理数组的分段是相互独立的过程，可以并发处理。
```
func quickSort3(s []int, low, high int, wg *sync.WaitGroup) {
    //当元素较少时，直接做插入排序
    if high-low < 50 {
    	insertSort(s)
    	return
    }
    if low < high {
    	p := partition(s, low, high)
    	//先处理较长部分，再处理较短部分，大数组并发处理
    	if p-low > high-p {
    		//先处理前半段
    		if p-low > 10000 {
    			wg.Add(1)
    			go func(data []int) {
    				defer wg.Done()
    				quickSort4(data, low, p-1, wg)
    			}(s)
    		} else {
    			quickSort3(s, low, p-1)
    		}
    		//再处理后半段
    		if high-p > 10000 {
    			wg.Add(1)
    			go func(data []int) {
    				defer wg.Done()
    				quickSort4(data, p+1, high, wg)
    			}(s)
    		} else {
    			quickSort3(s, p+1, high)
    		}
    	} else {
    		//先处理后半段
    		if high-p > 10000 {
    			wg.Add(1)
    			go func(data []int) {
    				defer wg.Done()
    				quickSort4(data, p+1, high, wg)
    			}(s)
    		} else {
    			quickSort3(s, p+1, high)
    		}
    		//再处理前半段
    		if p-low > 10000 {
    			wg.Add(1)
    			go func(data []int) {
    				defer wg.Done()
    				quickSort4(data, low, p-1, wg)
    			}(s)
    		} else {
    			quickSort3(s, low, p-1)
    		}
    	}
    }
}
func QSort(s []int) {
    // 再套一层函数外壳方便传入waitGroup
    var wg sync.WaitGroup
    quickSort3(s, 0, len(s)-1, &wg)
    wg.Wait()
}
```
基于样本范围0~100_000_000、样本容量为10_000_000的随机数数列，取`insertSort(50)`，比较上述这几种快速排序的性能，使用并发后提升很明显：   
```
$ go test -bench='QuickSort$' -benchtime=4s -count=2 .
goos: windows
goarch: amd64
pkg: tmp
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkQuickSort-16           1000000000               0.7196 ns/op
BenchmarkQuickSort-16           1000000000               0.7120 ns/op

Benchmark_2_QuickSort-16        1000000000               0.6610 ns/op
Benchmark_2_QuickSort-16        1000000000               0.6502 ns/op

Benchmark_3_QuickSort-16        1000000000               0.1212 ns/op
Benchmark_3_QuickSort-16        1000000000               0.1124 ns/op
PASS
ok      tmp     40.095s
```

##### #改进方法3
对于元素重复出现频率的较高的样本，可以使用三分法在每次partition时，将数组分成三部分，大于基准数，等于基准数，小于基准数。
```
func partition3(s []int, begin, end int) (int, int) {
    if end-begin < 2 {
    	return begin, begin
    }
    //通过交换，选定中间元素为基准值
    s[begin], s[(begin+end)/2] = s[(begin+end)/2], s[begin]
    var pivot = s[begin]
    //l,r分别是两段-开-边界：小部分的右边界和大部分的左边界,也是中间实际开始和最后终止的地方
    //mid是中间组的右边界，实际指向是待处理的数或最终的 大-1st
    l, r, mid := begin, end, begin+1
    for mid <= r {
    	if s[mid] > pivot {
    		//将大于pivot的数换到末尾，同时尾指针左移一位，紧接着继续处理换过来的这个数
    		s[mid], s[r] = s[r], s[mid]
    		r--
    	} else if s[mid] < pivot {
    		//将最左端的pivot数换到此时中间组的最右端
    		s[mid], s[l] = s[l], s[mid]
    		mid++
    		l++
    	} else {
    		// 相等数直接扩充
    		mid++
    	}
    }
    return l, r
}
func quickSort32(s []int, begin, end int) {
    //当元素较少是时，直接做插入排序
    if end-begin < 30 {
    	//slice左闭右开
    	insertSort(s[begin : end+1])
    	return
    }
    if begin < end {
    	l, r := partition3(s, begin, end)
    	quickSort32(s, begin, l-1)
    	quickSort32(s, r+1, end)
    }
}
```
举个稍极端的例子，取样本范围0~10_000、样本容量为10_000_000的随机数数列，取`insertSort(50)`，横向对比结果如下：
```
$ go test -bench='QuickSort$' -benchtime=4s -count=2 .
goos: windows
goarch: amd64
pkg: tmp
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_2_QuickSort-16        1000000000               2.974 ns/op
Benchmark_2_QuickSort-16        1000000000               2.947 ns/op

Benchmark_32_QuickSort-16       1000000000               0.4153 ns/op
Benchmark_32_QuickSort-16       1000000000               0.4040 ns/op
PASS
ok      tmp     278.746s
```
可见在数据重复度比较高的情况下，这种三分改进的意义是很大的。    

###### 以上所有结果的测试用例和详细过程在[这里](test/sort/main_test.go)


