package complex

import (
	"fmt"
	"math"
	"testing"
)

type __vertex struct {
	X, Y float64
}

// Abs 方法只是个带接收者参数的函数
func (v __vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs 这是正常的函数
func Abs(v __vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodFunc() {
	v := __vertex{3, 4}
	//v.Scale(10)
	p := &v
	p.Scale(10)
	fmt.Println(v.Abs())

	Scale(&v, 0.5)
	fmt.Println(Abs(v))

	f := myFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// 也可以为非结构体类型声明方法
// 接收者的类型定义和方法声明必须在同一包内

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Scale 为指针接收者声明方法
// 以指针为接收者的方法被调用时，接收者既能为值又能为指针
func (v *__vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 可变参数列表
func sum2(numbers ...int) int {
	s := 0
	for _, v := range numbers {
		s += v
	}
	return s
}

// Scale
// 带指针参数的函数必须接受一个指针
func Scale(v *__vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethod(t *testing.T) {
	methodFunc()
	println(sum2(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
