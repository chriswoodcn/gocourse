package main

import (
	"fmt"
	"time"
)

func testGoroutinePanicRecover() {
	items := make([]int, 0, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover")
				fmt.Printf("go程中 defer recover的value： %v \n", r)
			}
		}()

		panic("go程中发生panic")
		// panic之后的代码会执行不到
		items = append(items, 1)
	}()
	fmt.Println("主go程中的输出")
}

var maxRetryTimes = 3
var retryTimes = 1

func testMainPanicRecover() {
	items := make([]int, 0, 1)
	// recover只能在defer中生效
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover")
			fmt.Printf("main程中 defer recover的value： %v \n", r) //recover只能获取到第一次panic
			if retryTimes <= maxRetryTimes {
				fmt.Println("retryTimes", retryTimes)
				retryTimes += 1
				testMainPanicRecover()
			}
		}
	}()
	//fmt.Println(items[0])
	items = append(items, 1)
	fmt.Println(items[0])

	i := 0
	for {
		i++
		time.Sleep(time.Millisecond)
		if i > 1000 {
			fmt.Println(i)
			panic("main程中发生panic")
		}
	}
}

// recover无效 panic之后的代码不执行 不能放到panic之后;panic之前也不行，执行到recover时未发生panic，得到的是nil
// 因此recover只能在defer中生效
func testMainPanicRecover0() {
	if r := recover(); r != nil {
		fmt.Println("recover0")
		fmt.Printf("main程中recover的value： %v \n", r)
	}
	panic("main程中发生panic")
}
func main() {
	//testMainPanicRecover0()
	testMainPanicRecover()
}
