package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
var aaa = 123
var bbb = true
var (
	ccc = map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	ddd = []string{
		"一个丁老头",
		"一只手套",
	}
	person = struct {
		Name string
		Age  int
	}{"Kim", 22}
)

// 定义变量
func variables() {
	fmt.Println(">>>>>>>>>> 定义变量 >>>>>>>>>>")
	fmt.Printf("全局变量： %d %t %v \n", aaa, bbb, bbb)
	fmt.Printf("map %v %#v\n", ccc, ccc)
	for k, v := range ccc {
		fmt.Printf("key: %v value: %v", k, v)
	}
	fmt.Printf("array %v %#v\n", ddd, ddd)
	fmt.Printf("struct %v %+v %#v\n", person, person, person)
	pointer := &person
	fmt.Printf("pointer %v %d %p\n", pointer, pointer, (*int)(nil))
	var a int
	var b string
	fmt.Printf("未初始化的变量值： %d %q \n", a, b)
	a = 10
	b = "123456"
	fmt.Println("初始化之后的变量值： ", a, b)
	c := 100 //type deduction类型推导
	// 类型推导
	// 在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出
	d := "741852963"
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("声明加赋值变量值： %d %q \n", c, d)
	fmt.Println(i, f, g)
	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)
}

// 验证欧拉公式
func euler() {
	fmt.Println(">>>>>>>>>> 验证欧拉公式 >>>>>>>>>>")
	c := 3 + 4i
	abs := cmplx.Abs(c)
	fmt.Println("复数取模的值： ", abs)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 类型转换 强制类型转换
// 表达式 T(v) 将值 v 转换为类型 T
func triangle() {
	fmt.Println(">>>>>>>>>> 强制类型转换 >>>>>>>>>>")
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2)))
	fmt.Println(c)
}

// 常量
func constants() {
	fmt.Println(">>>>>>>>>> 定义常量 >>>>>>>>>>")
	//常量不能用 := 语法声明。
	const (
		a, b     = "123456", "e23fd9jg91ngkgf01j91utg"
		filename = "hello.go"
	)
	// const数值可作为各种类型使用
	fmt.Println(a, b)
	fmt.Println(filename)
	// 特殊的常量 自增枚举类型
	const (
		//普通枚举类型
		//cpp    = 1
		//java   = 2
		//golang = 3
		//python = 4
		cpp = iota
		java
		golang
		python
		_
		typescript
		_
		javascript
	)
	fmt.Println(cpp, java, golang, python, typescript, javascript)
	// 特殊的常量 自增枚举类型
	const (
		byte = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
}

// 函数
func add(x, y int) int {
	fmt.Println(">>>>>>>>>> 函数形参类型相同时，除最后一个类型以外，其它都可以省略 >>>>>>>>>>")
	return x + y
}

// 函数
func swap(x, y string) (string, string) {
	fmt.Println(">>>>>>>>>> 函数可以返回任意数量的返回值 >>>>>>>>>>")
	return y, x
}

// 函数 命明返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// go 基本类型
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // uint8 的别名
// rune // int32 的别名
// // 表示一个 Unicode 码点
// float32 float64
// complex64 complex128
// int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽

// 没有明确初始值的变量声明会被赋予它们的 零值。
// 零值是：
// 数值类型为 0，
// 布尔类型为 false，
// 字符串为 ""（空字符串）。
func main() {
	fmt.Println("hello")
	variables()
	euler()
	triangle()
	constants()
	fmt.Println(add(12, 15))
	fmt.Println(swap(strconv.Itoa(12), strconv.Itoa(12)))
	fmt.Println(split(17))
}
