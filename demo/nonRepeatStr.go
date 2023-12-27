package main

import (
	"fmt"
	"log"
)

type BaseError struct {
	module string
	code   int
	msg    string
}

func (e *BaseError) Error() string {
	format := fmt.Sprintf(" [module] %s [code] %d [msg] %s", e.module, e.code, e.msg)
	log.Println(format)
	return e.msg
}

func BuildBaseError(module string, code int, msg string) error {
	return &BaseError{
		module,
		code,
		msg,
	}
}

type SystemError struct {
	*BaseError
}

func (e *SystemError) Error() string {
	var b = BaseError{
		"system",
		e.code,
		e.msg,
	}
	return b.Error()
}
func BuildSystemError(code int, msg string) error {
	return &SystemError{
		&BaseError{
			"system",
			code,
			msg,
		},
	}
}
func lengthOfNonRepeatingSubStr(s string) (int, error) {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLen := 0
	for i, v := range []byte(s) {
		if last, ok := lastOccurred[v]; ok && last >= start {
			start = last + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[v] = i
	}
	return maxLen, nil
}
func lengthOfNonRepeatingSubStrRune(s string) (int, error) {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLen := 0
	for i, v := range []byte(s) {
		if last, ok := lastOccurred[v]; ok && last >= start {
			start = last + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[v] = i
	}
	return maxLen, nil
}
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabc"))
}
