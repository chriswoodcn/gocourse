package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str) {
  printf("go 简单调用内嵌型c代码\n");
  printf("%s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
	s := "Hello Cgo"
	cs := C.CString(s)
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}
