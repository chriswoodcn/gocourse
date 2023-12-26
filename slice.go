package main

import "fmt"

func _swap(a, b int) {
	a, b = b, a
}
func _swapRef(a, b *int) {
	*a, *b = *b, *a
}
func printArray(arr [5]int) {
	for i := range arr {
		print(arr[i], " ")
	}
	arr[0] = 100
	print(" \n")
}
func printArray2(arr *[5]int) {
	for i := range arr {
		print(arr[i], " ")
	}
	arr[0] = 100
	print(" \n")
}
func testArray() {
	a, b := 3, 4
	_swap(a, b)
	println(a, b)
	_swapRef(&a, &b)
	println(a, b)
	//数组是值类型
	arr1 := [5]int{}
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArray2(&arr1)
	printArray2(&arr1)
	printArray2(&arr2)
	printArray2(&arr2)
	//go语言中一般不直接使用数组，使用切片  slice
}
func testSlice() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[2:6]
	fmt.Println(s)
}
func main() {
	//testArray()
	testSlice()
}
