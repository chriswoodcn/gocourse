package str

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
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

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)
		//fmt.Println(err)

		//if pathError, ok := err.(*os.PathError); !ok {
		//	panic(err)
		//} else {
		//	fmt.Println(pathError.Op)
		//	fmt.Println(pathError.Path)
		//	fmt.Println(pathError.Err)
		//}
		var path *os.PathError
		if errors.As(err, &path) {
			fmt.Println(path)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	fibonacci := Fibonacci()
	for i := 0; i < 20; i++ {
		//writer.WriteString(strconv.Itoa(fibonacci()))
		//writer.WriteString("\n")
		fmt.Fprintf(writer, "%d \n", fibonacci())
	}
}
func TestFibonacci1(t *testing.T) {
	fibonacci := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci())
	}
}
func TestFibonacci2(t *testing.T) {
	writeFile("fibonacci.txt")
}
