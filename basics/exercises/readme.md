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
  1. Create a simple stack which can hold a fixed number of ints. It does not have to
     grow beyond this limit. Define push – put something on the stack – and pop –
     retrieve something from the stack – functions. The stack should be a LIFO (last
     in, first out) stack.
  2. Bonus. Write a String method which converts the stack to a string representation.
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
}

// [1 4 9 16 25]
```
