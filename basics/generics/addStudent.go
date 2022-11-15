package main

/*
-------package sort---------
//sort包部分源码，要使用sort的Sort、Reverse、IsSorted方法，需先实现Interface接口
//已经实现的类型包括[]int、[]float64、[]string，且通过语法糖包装可简便调用
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
func Sort(data Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}
	limit := bits.Len(uint(n))
	pdqsort(data, 0, n, limit)
}
func Reverse(data Interface) Interface {...}
func IsSorted(data Interface) bool {...}

type IntSlice []int
func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
// Sort is a convenience method: x.Sort() calls Sort(x).
func (x IntSlice) Sort() { Sort(x) }

type Float64Slice []float64
func (x Float64Slice) Len() int { return len(x) }
func (x Float64Slice) Less(i, j int) bool { return x[i] < x[j] || (isNaN(x[i]) && !isNaN(x[j])) }
func (x Float64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
// Sort is a convenience method: x.Sort() calls Sort(x).
func (x Float64Slice) Sort() { Sort(x) }

type StringSlice []string
func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
// Sort is a convenience method: x.Sort() calls Sort(x).
func (x StringSlice) Sort() { Sort(x) }

// 简便用法，sort.FuncName(typeSlice)
func Ints(x []int) { Sort(IntSlice(x)) }
func Float64s(x []float64) { Sort(Float64Slice(x)) }
func Strings(x []string) { Sort(StringSlice(x)) }
...
*/
import (
	"fmt"
	"sort"
)

/*
假定在输出每个学生的name、id、student{}信息集合之前要对其进行排序
golang默认intSlice、floatSlice、stringSlice实现了Interface
所以对于结构体，可以仅以Student结构体实现Interface即可。
---源码sort包里的三种typeSlice接口实现也可以用泛型重写
*/

type Student struct {
	Name string
	ID   int
	Age  float64
}
type stuInfoType interface {
	~int | ~float64 | ~string | Student
}
type StudentList []Student

// 对应的四种返回类型[]int、[]float64、[]string、[]Student
func addStudent[T stuInfoType](stuInfoSet []T, stuInfo T) []T {
	return append(stuInfoSet, stuInfo)
}
func (s StudentList) Len() int {
	return len(s)
}
func (s StudentList) Less(i, j int) bool {
	return s[i].ID < s[j].ID // 以ID从小到大排序
}
func (s StudentList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var (
		stuName  []string
		stuID    []int
		students []Student
	)

	// 调用时根据参数类型的约束，实现不同学生信息的添加
	result := addStudent[string](stuName, "Michael")
	result = addStudent[string](result, "Jennifer")
	result = addStudent[string](result, "Elaine")
	result1 := addStudent[int](stuID, 78)
	result1 = addStudent[int](result1, 45)
	result1 = addStudent[int](result1, 64)
	result2 := addStudent[Student](students, Student{"James", 111, 18.75})
	result2 = addStudent[Student](result2, Student{"John", 213, 17.5})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})

	sort.Strings(result)
	sort.Ints(result1)
	sort.Sort(StudentList(result2))

	fmt.Println(result)
	fmt.Println(result1)
	fmt.Println(result2)
}
