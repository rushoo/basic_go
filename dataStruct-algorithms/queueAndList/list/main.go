package main

import "fmt"

// DataType constraint
type DataType interface {
	~string | ~int | ~float64
}

func main() {
	cars := SinglyLinkedList[string]{}
	cars.Append("Honda")
	cars.InsertAt(0, "Nissan")
	cars.InsertAt(1, "Chevy")
	cars.InsertAt(2, "Ford")
	cars.InsertAt(3, "Tesla")
	cars.InsertAt(4, "Audi")
	cars.InsertAt(5, "Volkswagon")
	cars.Append("Volvo")
	fmt.Println(cars.Items())
	fmt.Println("Index of Tesla: ", cars.IndexOf("Tesla"))
	cars.Reverse()
	fmt.Println(cars.Items())
}
