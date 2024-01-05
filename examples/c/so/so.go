// go build -o so.so -buildmode=c-shared so.go

package main

import "C"
import (
	"fmt"
	"math"
	"sort"
	"sync"
)

// 编写Golang共享库
// 包必须为main打包。编译器会将包及其所有依赖项构建到单个共享对象二进制文件中
// 源必须import "C"
// 使用//export注释以注释希望其他语言可以访问的函数。
// 一个空的main必须声明函数

var count int
var mtx sync.Mutex

//export Add
func Add(a, b int) int {
	return a + b
}

//export Cosine
func Cosine(x float64) float64 {
	return math.Cos(x)
}

//export Sort
func Sort(vals []int) {
	sort.Ints(vals)
}

//export SortPtr
func SortPtr(vals *[]int) {
	Sort(*vals)
}

//export Log
func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

//export LogPtr
func LogPtr(msg *string) int {
	return Log(*msg)
}
func main() {

}
