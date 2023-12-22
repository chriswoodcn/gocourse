package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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
	d := "741852963"
	fmt.Printf("声明加赋值变量值： %d %q \n", c, d)
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
func main() {
	fmt.Println("hello")
	variables()
	euler()
	triangle()
	constants()
}
