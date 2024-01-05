package str

import (
	"fmt"
	"testing"
)

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
func TestNonRepeatingSubStr(t *testing.T) {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabc"))
	fmt.Println(lengthOfNonRepeatingSubStrRune("一二三二一"))
}
