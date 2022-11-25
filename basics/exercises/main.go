package main

import "fmt"

/*
5. (1) Map function A map()-function is a function that takes a function and a
  list. The function is applied to each member in the list and a new list containing these
  calculated values is returned. Thus:
  map(f(),(a1, a2, . . . , an−1, an)) = (f(a1), f(a2), . . . , f(an−1), f(an))
*/

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
