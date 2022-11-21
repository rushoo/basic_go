package main

import "fmt"

// 增加代码可读性
const (
	FIZZ = 3
	BUZZ = 5
)

func fizzBuzz() {
	for i := 0; i < 100; i++ {
		switch {
		//switch case是顺序执行的，如果第一个条件满足，后续就不会执行了
		case i%FIZZ == 0 && i%BUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
func main() {
	fizzBuzz()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// change for-loop to goto-loop
	i := 0
loop:
	fmt.Printf("%d\n", i)
	i++
	if i < 10 {
		goto loop
	}

}
