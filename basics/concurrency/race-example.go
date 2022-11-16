package main

import (
	"fmt"
	"sync"
)

// 写数据时不控制，这样可能会有1000个goroutine同时写操作，最终结果会很奇怪
// 可以在数据操作前后加锁，确保每次仅一个goroutine写入
// 可以使用 go run -race example2.go在运行时做race检测

const number = 1000

// var mux sync.Mutex
var countValue int

func main() {
	var wg sync.WaitGroup
	wg.Add(number)

	for i := 0; i < number; i++ {
		go func() {
			//mux.Lock()
			countValue++
			//mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\ncountValue = %d", countValue)
}
