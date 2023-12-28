package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func testIoRead() {
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
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("printFile os.OpenFile error ")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func env() {
	environ := os.Environ()
	for _, v := range environ {
		fmt.Println(v)
	}
}
func readFile() {
	file, err := os.ReadFile("aaa.txt")
	if err != nil {
		fmt.Println("ReadFile Err")
	}
	for _, v := range file {
		println(string(v))
	}
}
func main() {
	//testIoRead()
	//printFile("aaa.txt")
	//env()
	readFile()
}
