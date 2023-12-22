package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convert2bin(n int) string {
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
func main() {
	if file, err := os.ReadFile("aa.txt"); err == nil {
		panic("error")
	} else {
		fmt.Printf("%s \n", file)
		fmt.Println(file)
	}
	fmt.Println(convert2bin(25))
}
