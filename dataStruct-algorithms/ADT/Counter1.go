package main

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}
func (c *Counter) Decrement() {
	c.count--
}
func (c *Counter) Reset() {
	c.count = 0
}
func (c *Counter) GetCount() int {
	return c.count
}
func main() {
	myCounter := new(Counter)
	for i := 1; i <= 10; i++ {
		myCounter.Increment()
	}
	myCounter.Decrement()
	final := myCounter.GetCount()
	println(final) //9

	myCounter.count += 100        //这表明当前ADT的封装是无效的,将在Counter2中改造
	println(myCounter.GetCount()) //109
}
