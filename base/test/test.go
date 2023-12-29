package main

import (
	"fmt"
	"gocourse/base"
	"strconv"
)

func testHello() {
	fmt.Println(">>>>>>>>>>testHello>>>>>>>>>>")
	fmt.Println("hello")
	base.Variables()
	base.Euler()
	base.Triangle()
	base.Constants()
	fmt.Println(base.Add(12, 15))
	fmt.Println(base.Swap(strconv.Itoa(12), strconv.Itoa(12)))
	fmt.Println(base.Split(17))
	bytes := base.String2Bytes("testHello")
	bytes2String := base.Bytes2String(bytes)
	fmt.Println(bytes2String)
}
func testBranch() {
	fmt.Println(">>>>>>>>>>testBranch>>>>>>>>>>")
	fmt.Println(base.Convert2bin(25))
	base.Loop()
	base.IfFunc()
	fmt.Println(base.Sqrt(2))
	base.SwitchFunc()
	base.DeferFunc()
}
func testIo() {
	base.TestIoRead()
	base.PrintFile()
	base.Env()
	base.ReadFile()
}
func main() {
	testHello()
	testBranch()
	testIo()
}
