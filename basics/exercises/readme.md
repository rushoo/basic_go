 1.FizzBuzz   
     Write a program that prints the numbers from 1 to 100. But for multiples   
     of three print “Fizz” instead of the number ,and for the multiples of five   
     print “Buzz”. For numbers which are multiples of both three and five print   
     “FizzBuzz”.   
    
 2.Strings   
     1. Create a Go program that prints the following (up to 100 characters):   
         A   
         AA   
         AAA   
         AAAA   
         AAAAA   
         AAAAAA   
         AAAAAAA   
         ...   
     2. Create a program that counts the number of characters in this string:  12345678“你好”！   
        In addition, make it output the number of bytes in that string.   
     3. Write a Go program that reverses a string, so “foobar” is printed as “raboof”. Hint:   
        You will need to know about conversion.   
    
3.Average   
  Write code to calculate the average of a float64 slice.   
   
4.Stack    
  1.Create a simple stack which can hold a fixed number of ints. It does not have to   
     grow beyond this limit. Define push – put something on the stack – and pop –   
     retrieve something from the stack – functions. The stack should be a LIFO (last   
     in, first out) stack.   
  2.Bonus. Write a String method which converts the stack to a string representation.   
     This way you can print the stack using: fmt.Printf("My stack %v\n", stack)   
     The stack in the figure could be represented as: [0:m] [1:l] [2:k]   
    
5.Map function    
    A map()-function is a function that takes a function and a list.       
    The function is applied to each member in the list and a new list    
    containing these calculated values is returned. Thus:   
    map(f(),(a1, a2, . . . , an−1, an)) = (f(a1), f(a2), . . . , f(an−1), f(an))    
```
type dataType interface {
	~int | ~float64 | ~string
}

func Map[T dataType](f func(T) T, list []T) []T {
	var result []T
	for _, v := range list {
		result = append(result, f(v))
	}
	return result
}
func main() {
	list := []int{1, 2, 3, 4, 5}
	f := func(i int) int {
		return i * i
	}
	result := Map[int](f, list)
	fmt.Println(result)

	list2 := []string{"HELLO", "ALOHA", "HI", "YEP", "GREAT"}
	f2 := func(i string) string {
		return strings.ToLower(i)
	}
	result2 := Map[string](f2, list2)
	fmt.Println(result2)
}

// [1 4 9 16 25]
// [hello aloha hi yep great]
```
    
6.Minimum and maximum    
     1. Write a function that finds the maximum value in an int slice ([]int).   
     2. Write a function that finds the minimum value in an int slice ([]int).    
```
func Minimum(s []int) int {
	min := s[0]
	for _, v := range s {
		if min > v {
			min = v
		}
	}
	return min
}
func Maximum(s []int) int {
	max := s[0]
	for _, v := range s {
		if max < v {
			max = v
		}
	}
	return max
}
func main() {
	list := []int{1, 2, -1, 32, 9, 3, 4, 5}
	min := Minimum(list)
	max := Maximum(list)
	fmt.Println("list:", list)
	fmt.Println("min:", min, "\tmax:", max)
}
// list: [1 2 -1 32 9 3 4 5]
// min: -1         max: 32
```
   
7.Bubble sort   
```
//	slice是引用类型，不需要返回新的slice
func Bubbleort(list []int) {
	for i := 0; i < len(list); i++ {
		//每轮扫描，依次将相邻元素比较大小并交换位置，确保每次的最大数都冒泡到最后
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
func main() {
	list := []int{1, 2, -1, 32, 9, 3, 4, 5}
	fmt.Println("list:", list)
	Bubbleort(list)
	fmt.Println("list after Bubbleort:", list)
}
// list: [1 2 -1 32 9 3 4 5]
// list after Bubbleort: [-1 1 2 3 4 5 9 32]
```   
   
8.Functions that return functions    
    1.Write a function that returns a function that performs a +2 on integers.   
      Name the function plusTwo.   
    2.Generalize the function from 1, and create a plusX(x) which returns functions    
      that add x to an integer.   
```
func plusX(x int) func(int) int {
	return func(a int) int {
		return a + x
	}
}
func plusTwo() func(int) int {
	return func(a int) int {
		return a + 2
	}
}
func main() {
	f1 := plusTwo()
	f2 := plusX(10)
	fmt.Println(f1(2), f2(2)) //4 12
}
```