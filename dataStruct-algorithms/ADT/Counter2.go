package main

/*
这里通过接口完成封装，
所有对count变量的操作仅能通过接口定义的方法来实现

若仅为了实现count变量不被外部直接访问，
还可通过counter package的非导出变量来实现
*/
type counter struct {
	count  int
	Count2 int
}
type Counter2 interface {
	Increment()
	Decrement()
	Reset()
	GetCount() int
}

func (c *counter) Increment() {
	c.count++
}
func (c *counter) Decrement() {
	c.count--
}
func (c *counter) Reset() {
	c.count = 0
}
func (c *counter) GetCount() int {
	return c.count
}
func main() {
	myCounter := Counter2(&counter{}) //对象封装
	for i := 1; i <= 10; i++ {
		myCounter.Increment()
	}
	myCounter.Decrement()
	final := myCounter.GetCount()
	println(final) //9

	//myCounter.count += 100        //无效的访问
	//myCounter.Count2 += 100       //无效的访问
}
