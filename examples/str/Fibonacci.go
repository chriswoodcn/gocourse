package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type fibonacciInt func() int

func (fi fibonacciInt) Write(p []byte) (n int, err error) {
	//TODO implement me
	f := fi()
	s := strconv.Itoa(f)
	b := strings.Builder{}
	b.Grow(len(s) + utf8.UTFMax)
	return b.WriteString(s)
}

func main() {
	fibonacci := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci())
	}
}
