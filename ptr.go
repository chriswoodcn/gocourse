package main

import (
	"fmt"
	"math"
)

type _vertex struct {
	X int
	Y int
}

func definePtr() {
	//类型 *T 是指向 T 类型值的指针。其零值为 nil。
	var p *int
	i, j := 42, 2701
	p = &i          // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
	//Go 没有指针运算。
}
func structFunc() {
	// 结构体
	v := _vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
	//结构体指针
	p := &v
	p.X = 1e9
	fmt.Println(v)
	//结构体文法通过直接列出字段的值来新分配一个结构体。
	//使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
	//特殊的前缀 & 返回一个指向结构体的指针。
	var (
		v1 = _vertex{1, 2}  // 创建一个 _vertex 类型的结构体
		v2 = _vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = _vertex{}      // X:0 Y:0
		p0 = &_vertex{1, 2} // 创建一个 *_vertex 类型的结构体（指针）
	)
	fmt.Println(p0, v1, v2, v3)
}
func arrayFunc() {
	//类型 [n]T 表示拥有 n 个 T 类型的值的数组。
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
	//类型 []T 表示一个元素类型为 T 的切片
	var s []int = primes[1:4]
	fmt.Println(s)
	//切片并不存储任何数据，它只是描述了底层数组中的一段。
	//更改切片的元素会修改其底层数组中对应的元素。
	//与它共享底层数组的切片都会观测到这些修改。
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	A := names[0:2]
	B := names[1:3]
	fmt.Println(A, B)
	B[0] = "XXX"
	fmt.Println(A, B)
	fmt.Println(names)
}
func sliceFunc() {
	//切片类型的写法是 []T ， T 是切片元素的类型。和数组不同的是，切片类型并没有给定固定的长度
	//切片的字面值和数组字面值很像，不过切片没有指定元素个数
	//切片可以使用内置函数 make 创建，函数签名为：func make([]T, len, cap) []T
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	//切片下界的默认值为 0，上界则是该切片的长度
	slice1 := s[:2]
	fmt.Println(slice1)
	slice1 = s[1:]
	fmt.Println(slice1)

	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)
	// 截取切片使其长度为 0
	s2 = s2[:0]
	printSlice(s2)
	// 拓展其长度
	s2 = s2[:4]
	printSlice(s2)
	// 舍弃前两个值
	s2 = s2[2:]
	printSlice(s2)

	//nil 切片 nil 切片的长度和容量为 0 且没有底层数组
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil!")
	}

	//用 make 创建切片 make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	a := make([]int, 5)
	printSlice2("a", a)
	b := make([]int, 0, 5)
	printSlice2("b", b)
	c := b[:2]
	printSlice2("c", c)
	d := c[2:5]
	printSlice2("d", d)

	//切片的切片
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	fmt.Println(board)
	board[2][2] = "O"
	fmt.Println(board)
	board[1][2] = "X"
	fmt.Println(board)
	board[1][0] = "O"
	fmt.Println(board)
	board[0][2] = "X"
	fmt.Println(board)

	var s4 []int
	printSlice(s4)
	// 添加一个空切片
	s4 = append(s4, 0)
	printSlice(s4)
	// 这个切片会按需增长
	s4 = append(s4, 1)
	printSlice(s4)
	// 可以一次性添加多个元素  append 函数将 x 追加到切片 s 的末尾，并且在必要的时候增加容量。
	// 下面有切片原理和自控追加代码演示
	s4 = append(s4, 2, 3, 4)
	printSlice(s4)
	//如果是要将一个切片追加到另一个切片尾部，需要使用 ... 语法将第2个参数展开为参数列表。
	s5 := []int{8, 9, 10}
	s4 = append(s4, s5...) // equivalent to "append(a, b[0], b[1], b[2])"

	//一个切片是一个数组片段的描述。它包含了指向数组的指针，片段的长度， 和容量（片段的最大长度）。
	//切片增长不能超出其容量。增长超出切片容量将会导致运行时异常，就像切片或数组的索引超 出范围引起异常一样。
	//同样，不能使用小于零的索引去访问切片之前的元素

	//要增加切片的容量必须创建一个新的、更大容量的切片，然后将原有切片的内容复制到新的切片。
	//整个技术是一些支持动态数组语言的常见实现。下面的例子将切片 s 容量翻倍，先创建一个2倍
	//容量的新切片 t ，复制 s 的元素到 t ，然后将 t 赋值给 s ：
	t := make([]int, len(s4), (cap(s4)+1)*2) // +1 in case cap(s) == 0
	copy(t, s4)
	s4 = t

	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)

	//for 循环的 range 形式可遍历切片
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

type vertex2 struct {
	Lat, Long float64
}

var m map[string]vertex2

func mapFunc() {
	//映射将键映射到值。
	//映射的零值为 nil 。nil 映射既没有键，也不能添加键
	//make 函数会返回给定类型的映射，并将其初始化备用
	m = make(map[string]vertex2)
	m["Bell Labs"] = vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var m2 = map[string]vertex2{
		"Bell Labs": vertex2{
			40.68433, -74.39967,
		},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m2)

	//修改映射
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 0
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
func funFunc() {
	//函数值可以用作函数的参数或返回值。
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func closureFunc() {
	//Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。
	//该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
	//例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	//definePtr()
	//structFunc()
	//arrayFunc()
	//sliceFunc()
	//mapFunc()
	//funFunc()
	closureFunc()
}
