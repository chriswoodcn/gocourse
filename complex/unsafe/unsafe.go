package main

import (
	"fmt"
	"unsafe"
)

func testPointer() {
	//（1）将*T1转换为指向*T2的指针。
	i := 30
	ptr1 := &i
	fmt.Printf("ptr1: %T %#x %d \n", ptr1, ptr1, *ptr1)
	ptr2 := (*int64)(unsafe.Pointer(ptr1))
	i = 8
	fmt.Printf("ptr2: %T %#x %d \n", ptr2, ptr2, *ptr2)
	//（2）将指针转换为uintptr（但不转换回指针）
	ptr3 := uintptr(unsafe.Pointer(ptr1))
	fmt.Printf("ptr3: %T %#x, %d \n", ptr3, ptr3, ptr3)
	//（3）通过算术将指针转换为uintptr并返回
	arr := [...]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pointer := unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])))
	fmt.Printf("pointer: %T %#x, %d  size: %d \n", pointer, pointer, pointer, unsafe.Sizeof(pointer))
	ptr4 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0]))))
	fmt.Printf("ptr4: %T %#x, %d \n", ptr4, ptr4, ptr4)
	fmt.Printf("ptr4: %T %d \n", *ptr4, *ptr4)
	ptr5 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 2*unsafe.Sizeof(arr[0])))
	fmt.Printf("ptr5: %T %d  \n", *ptr5, *ptr5)
	// 超出了声明的范围 无效的 end points outside allocated space.
	ptr6 := unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + unsafe.Sizeof(arr))
	val6 := (*int64)(ptr6)
	fmt.Printf("ptr6: %T %#x %d \n", ptr6, ptr6, *val6)
	//（4）调用syscall时将指针转换为uintptr
	//syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))

	var x struct {
		a bool
		b int16
		c []int
	}
	fmt.Printf("x sizeof %d alginof %d  \n", unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Printf("x.a sizeof %d alginof %d  offsetof %d \n", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Printf("x.b sizeof %d alginof %d  offsetof %d \n", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Printf("x.c sizeof %d alginof %d  offsetof %d \n", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))

	str := "abc 123 hello world"
	data := unsafe.StringData(str)
	// Slice(ptr, len) is equivalent to  (*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
	slice := unsafe.Slice(data, 2)
	fmt.Println(slice) //expect 97 98
}
func main() {
	testPointer()
}
