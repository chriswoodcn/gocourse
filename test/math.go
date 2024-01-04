package test

func Add(a, b int32) int32 {
	return a + b
}
func Foo(i int, s string) (string, error) {
	return "", nil
}
func SliceSum(arr []int64) int64 {
	var sum int64

	for _, val := range arr {
		if val%100000 != 0 {
			sum += val
		}
	}
	return sum
}
func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib1(n-2) + fib1(n-1)
}
func _fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	fibFunc := _fib()
	for i := 0; i < n-1; i++ {
		fibFunc()
	}
	return fibFunc()
}
