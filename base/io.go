package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func TestIoRead() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 32)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
func PrintFile() {
	file, err := os.Open("aaa.txt")
	if err != nil {
		fmt.Println("OpenFile Err")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func Env() {
	environ := os.Environ()
	for _, v := range environ {
		fmt.Println(v)
	}
}
func ReadFile() {
	file, err := os.ReadFile("aaa.txt")
	if err != nil {
		fmt.Println("ReadFile Err")
	}
	for _, v := range file {
		println(string(v))
	}
}
