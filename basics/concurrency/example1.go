package main

/*
在main函数中使用go funcName开启一个goroutine，它将和main-goroutine并发执行
有时候需要考虑main-goroutine执行结束退出程序而funcName-goroutine未执行结束的情况
1、一种使用channel，当funcName-goroutine中产生chan而main-goroutine未读取到就会使得程序成功阻塞。
2、若chan不匹配可能导致程序死锁，可以使用select优雅跳过问题goroutine的执行
3、可以使用WaitGroup,在main-goroutine中设置计数器，每个func-goroutine执行结束使得计数器减一
*/
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var QUIT chan string

func outputStrings() {
	defer wg.Done()
	strings := [5]string{"One", "Two", "Three", "Four", "Five"}
	for i := 0; i < 5; i++ {
		fmt.Println(strings[i])
	}
}
func outputInts() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
func outputFloats() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(float64(i*i) + 0.5)
	}
}

// sent to the channel
func pingGenerator(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- "ping"
		time.Sleep(time.Millisecond * 500)
	}
}
func pongGenerator(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- "pong"
		time.Sleep(time.Millisecond * 500)
	}
}

// receive from the channel
func output(c <-chan string) {
	for {
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(2 * time.Second):
			QUIT <- "Program timed out."
			break
		}
	}
}
func main() {
	wg.Add(3)
	go outputStrings()
	go outputInts()
	go outputFloats()

	c := make(chan string)
	QUIT = make(chan string)
	go pingGenerator(c)
	go pongGenerator(c)
	go output(c)
	fmt.Println(<-QUIT)
	wg.Wait()
}
