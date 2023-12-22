package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// if条件
func convert2bin(n int) string {
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
func loop() {
	//Go 只有一种循环结构：for 循环
	fmt.Println(">>>>>>>>>> for 循环 >>>>>>>>>>")
}
func main() {
	if file, err := os.ReadFile("aa.txt"); err == nil {
		panic("error")
	} else {
		fmt.Printf("%s \n", file)
		fmt.Println(file)
	}
	fmt.Println(convert2bin(25))
}
