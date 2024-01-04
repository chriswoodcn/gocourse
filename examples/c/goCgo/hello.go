package main

import "fmt"
import "C"

//export HelloFromGo
func HelloFromGo() {
	fmt.Printf("Hello from Go!\n")
}
