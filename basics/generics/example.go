package main

import (
	"fmt"
	"sort"
)

/*
泛型变量与实例化

		slice：type MySlice [T int|float64|string] []T
					//type SliceInt []int
					//type SliceFloat []float64
					//type SliceInt []string
	           var mySlice1 = MySlice[int]{1, 2, 3, 4}
			   mySlice2 := MySlice[string]{"hello", "small"}

		map: type MyMap [KEY int|string, VALUE string| float64] map[KEY]VALUE
				//type MyMap1 map[int]string
				//type MyMap2 map[int]float64
				//type MyMap3 map[string]string
				//type MyMap4 map[string]float64
			myMap := MyMap[string, string]{ "one": "hello",  "two": "small" }

		  	struct: type MyStruct [T string|int|float64] struct {  Title string  Content  T}
					//type Struct3 struct {  Title string  Content  string}
					//type Struct4 struct {  Title string  Content  int}
					//type Struct5 struct {  Title string  Content  float64}
			var MyStruct2 = Struct1[string]{  Title:   "hello",  Content: "small"}

泛型函数定义与调用
	//自定义一个类型约束名，包含以下四种类型
	type TypeConstraint interface {
		// ~表示一个类型的超集,也就是只要底层类型满足即可用，方便灵活地使用自定义类型别名
		~int | ~float64 | ~string | Student
	}
	//定义时指定泛型参数类型约束，无约束就是 '[T any]'或 '[T interface{}]'
	func addStudent[T TypeConstraint](stuInfoSet []T, stuInfo T) []T {
		return append(stuInfoSet, stuInfo)
	}

	//调用
	var (
		stuName  []string
		stuID    []int
		students []Student
	)
	//调用时指定参数类型
	result := addStudent[string](stuName, "Michael")
	result1 := addStudent[int](stuID, 78)
	result2 := addStudent[Student](students, Student{"James", 111, 18.75})

	//泛型函数的定义和调用可以基于普通函数改写参数类型，再加上泛型约束即可。
*/

func myMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}

// 可以给参数不同的类型
func myMap2[T1, T2 int | int32 | int64 | float32 | float64](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}
func myFilter[T int | float64 | float32 | string](input []T, f func(T) bool) []T {
	var result []T
	for _, value := range input {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
func main() {
	input := []int{2, 3, 5, 7, 11}
	input2 := []float64{2.2, 3.3, 5.5, 7.7, 11.11}
	input3 := []float32{2.2, 3.3, 5.5, 7.7, 11.11}
	input5 := []string{"screens", "here", "aloha", "alphabetize", "sunday", "cat", "Ri"}
	result := myMap(input, func(i int) int {
		return i * i
	})
	result2 := myMap2[float64](input2, func(i float64) float64 {
		return i * i
	})
	result3 := myMap2[float32, float64](input3, func(i float32) float64 {
		//会有精度损失，精度计算可以参考https://github.com/shopspring/decimal
		return float64(i) * float64(i)
	})
	result4 := myFilter[float64](input2, func(i float64) bool {
		return i < 5
	})
	result5 := myFilter[string](input5, func(i string) bool {
		return len(i) > 2
	})
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	//字符串排序
	sort.Strings(result5)
	fmt.Println(result5)
}
