package base

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Convert2bin if条件
func Convert2bin(n int) string {
	fmt.Println(">>>>>>>>>> if条件 >>>>>>>>>>")
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	if length := len(result); length < 8 {
		result = strings.Repeat("0", 8-length) + result
	}
	return result
}

// Loop for循环
func Loop() {
	//Go 只有一种循环结构：for 循环
	fmt.Println(">>>>>>>>>> for 循环 >>>>>>>>>>")
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//去掉分号，因为 C 的 while 在 Go 中叫做 for
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	//如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑
	//for {
	//}
}

// if条件
func IfFunc() {
	//Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
	//同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。 该语句声明的变量作用域仅在 if 之内
	if file, err := os.ReadFile("aa.txt"); err == nil {
		panic("error")
	} else {
		fmt.Printf("%s \n", file)
		fmt.Println(file)
	}
}

// Sqrt 牛顿法计算平方根
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// SwitchFunc switch函数
func SwitchFunc() {
	fmt.Println(">>>>>>>>>> switch函数 >>>>>>>>>>")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s\n", os)
	}
	nowDateTime := time.Now().Format(time.DateTime)
	fmt.Printf("now DateTime: %s\n", nowDateTime)
	location := time.Now().Location()
	fmt.Printf("location: %v\n", *location)
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//没有条件的 switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// DeferFunc defer函数
func DeferFunc() {
	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其  参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	defer fmt.Println("world")
	fmt.Println("hello")
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
