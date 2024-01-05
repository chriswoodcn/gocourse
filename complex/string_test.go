package complex

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	str := "abc这是简单的测试"
	for _, s := range []byte(str) {
		fmt.Printf("%X ", s)
	}
	println("")
	for _, s := range str {
		fmt.Printf("%X ", s)
	}
	println("")
	fmt.Printf("rune count is str: %d", utf8.RuneCountInString(str))
	println("")
	bytes := []byte(str)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		fmt.Printf("rune: %c size: %d \n", r, size)
		bytes = bytes[size:]
	}
	fields := strings.Fields(str)
	fmt.Println(fields)
	split := strings.Split(str, "b")
	fmt.Println(split)
}
