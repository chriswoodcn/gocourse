// go build -o hello.so -buildmode=c-shared .

package main

/*
extern int helloFromC();
*/
import "C"
import "fmt"

//export HelloFromGo
func HelloFromGo() {
	fmt.Printf("Hello from Go!\n")
	//内部实现调用了C实现的函数helloFromC go->C
	C.helloFromC()
}

func main() {
}
