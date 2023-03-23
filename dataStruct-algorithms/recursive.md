### 递归
从一个简单的爬楼梯问题开始说起。     
假设要计算到达x层可能的方法，那么所有可能的方法的总数是楼梯层数的函数。可表示为：           
f(x) = num ，x为楼梯层数,num是可达方法总数。    
```
func f(x int) (num int) { return }
```
又有给定条件，青蛙一次可以上1 , 2 或3 阶。那么跳达x层时，一定是由x-1、x-2、x-3这三者的某一个出发的，所以应该将它们加总。         
根据这个条件可以完善方法：   
```
func f(x int) int {
	return f(x-1) + f(x-2) + f(x-3)
}
```
很显然，根据这个条件，可以将问题不断简化。比如青蛙到达第10层时，一定是从 9|8|7 上来的，那么，f(10)=f(9)+f(8)+f(7),       
进而 f(7)=f(6)+f(5)+f(4) ... f(4)=f(3)+f(2)+f(1) ... f(3)=f(2)+f(1)+f(0) ... f(2)=f(1)+f(0)+f(-1)，       
f(1)=f(0)+f(-1)+f(-2)          
根据实际，f(0)表示从地面起跳，所以对于任何x<0,f(x)是没有意义的，那么应该约定f(x)=0(x<0).         
而f(1)表示跳达第一层可能的方法，显然只可能 0->1 从地面跳达这一种情况，故f(1)=f(0)=1.现在我们计算方法变成：        
```
func f(x int) int {
	if x < 0 {
		return 0
	}
	if x == 0 || x == 1 {
		return 1
	}
	return f(x-1) + f(x-2) + f(x-3)
} 
```
根据这个计算方法，跳达第10层时有274种方法。         
回顾上述，先是明确了问题，即给定一楼梯层级，计算可达方法,由此明确了可达方法总数是楼梯层级数的函数。           
为了计算这个问题，可以根据条件缩小问题规模：f(x) = f(x-1) + f(x-2) + f(x-3)。    
以及问题边界，最基本地，f(0)、f(-x)表达的实际意义(值)是什么，这样也就确定了问题边界。           
以上就是递归问题三要素： 1、明确问题函数  2、缩小问题规模的方法  3、问题边界限定         

### #例：     
###### #求n！  
输入输出：求的是一个数的阶乘，结果也应该是一个数，f(num1) = num2      
缩小问题的方法：n！=n * (n-1)！,即 f(x) = xf(x-1)    
边界条件：阶乘的意义就是全排列，1！也就表示一个元素的全排列，所以 1！=1，计算少于一个元素的全排列没有任何意义，     
        不应该对乘法结果产生影响，所以m!=1(m<1).    
```
// 求n！
func f2(x int) int {
	if x < 1 {
		return 1
	}
	if x == 1 {
		return 1
	}
	return x * f2(x-1)
}
```

###### #求第n 个斐波那契数列元素     
问题定义：     
斐波那契数列又因数学家莱昂纳多·斐波那契以兔子繁殖为例子而引入，故又称为“兔子数列”。    
兔子在出生两个月后，就有繁殖能力，一对成年而有繁殖力的兔子每个月能生出一对小兔子来。    
假设一年以后所有兔子都不死，那么一对小兔子一年以后可以繁殖多少对兔子？    
1、兔子总数是时间的函数 f(x)=num    
2、本月的兔子数 = 上月存活至今兔数 + 本月新生兔数，而本月新生兔数由两月前兔数一对一产生。    
 那么即：f(x) = f(x-1) + f(x-2)     
3、初始时为f(0),第一个月为f(1),根据意义，显然f(1)=f(0)=1,所以：    
```
func f3(x int) int {
	if x <= 1 {
		return 1
	}
	return f3(x-1) + f3(x-2)
}
```

###### #汉诺塔问题:     
有三根杆子A，B，C。A杆上有 N 个 (N>1) 穿孔圆盘，盘的尺寸由下到上依次变小。要求按下列规则将所有圆盘移至 C 杆：    
每次只能移动一个圆盘；    
大盘不能叠在小盘上面。    
提示：可将圆盘临时置于 B 杆，也可将从 A 杆移出的圆盘重新移回 A 杆，但都必须遵循上述两条规则。    
问：如何移？最少要移动多少次？    
问题描述：f(n,A,B,C)表示：”以B为辅助，将n块从A移到C“，而就此而言，需要先做：f(n-1,A,C,B)，       
最后再想办法将此n-1块全移到C，即f(n-1,B,A,C)，也就是：    
```
func f(n int, A, B, C string) {
	f(n-1, A, C, B)                 //先做
	fmt.Println("外层函数目的解释")    //”以B为辅助，将n块从A移到C“
	f(n-1, B, A, C)                //最后
}
```
再来研究问题的边界： 假设只有一块，那么应该“将第一块从src直接移到des”   
```
var counter int
func f(n int, A, B, C string) {
	if n == 1 {
		counter++
		fmt.Printf("%d: 将第%d块从%s直接移到%s\n", counter, n, A, C)
		return
	}
	f(n-1, A, C, B)
	counter++
	fmt.Printf("%d: 以%s为辅助，将第%d块从%s移到%s\n", counter, B, n, A, C)
	f(n-1, B, A, C)
}
```
上述可以得到"合理"的结果，但其实偏离了问题的本质。问题是，”问：如何移？最少要移动多少次？“      
所以问题函数应该是f(n,A,B,C)=num，移动次数是给定A,B,C下圆盘数量n的函数。进而有：    
f(n,A,B,C)=f(n-1, A, C, B) + 1 + f(n-1, B, A, C)，即：   
```
func Hanoi(n int, A, B, C string) int {
	if n == 1 {
		fmt.Printf("%d: 将第%d块从%s直接移到%s\n", counter, n, A, C)
		return 1
	}
	a := Hanoi(n-1, A, C, B)
	fmt.Printf("%d: 以%s为辅助，将第%d块从%s移到%s\n", counter, B, n, A, C)
	b := Hanoi(n-1, B, A, C)
	return a + 1 + b
}
```
这会输出同样的过程步骤和总移动次数。明确问题定义和缩减问题规模的过程，在递归中是至关重要的。       

###### #插入排序的递归形式，对前n个元素做插入排序
其实就是对前n-1个元素调用插入排序，再将第n个元素按顺序插入排序
```
func InsertSort(items []int, n int) {
	l := len(items)
	if l < n || n < 2 {
		return
	}
	InsertSort(items, n-1) //前n-1个元素排序
	tmp := items[n-1]
	for n > 1 && tmp < items[n-2] {
		items[n-1] = items[n-2]
		n--
	}
	items[n-1] = tmp
}
```