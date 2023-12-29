package main

import (
	"fmt"
	"math"
	"time"
)

// Abser 接口类型 是由一组方法签名定义的集合 接口类型的变量可以保存任何实现了这些方法的值。
type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func testInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())

	var i I = &T{"hello"}
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

type I interface {
	M()
}

type T struct {
	S string
}

// M 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t *T) M() {
	// 即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
	// 在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 空接口
// 指定了零个方法的接口值被称为 *空接口：*
// interface{}
func testEmptyInterface() {
	var i interface{}
	i = 42
	describe2(i)
	i = "hello"
	describe2(i)
}

func typeAssert() {
	var i interface{} = "hello"
	//该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // 报错(panic)
	//fmt.Println(f)
}
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func testSwitchType() {
	do(21)
	do("hello")
	do(true)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func testPersonStringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type CommonError struct {
	When time.Time
	What string
}

// 实现error接口
func (e *CommonError) Error() string {
	return fmt.Sprintf("%v, %s",
		e.When, e.What)
}
func run() error {
	return &CommonError{
		time.Now(),
		"it didn't work",
	}
}
func testError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
func main() {
	testInterface()
	testEmptyInterface()
	typeAssert()
	testSwitchType()
	testPersonStringer()
	testError()
}
