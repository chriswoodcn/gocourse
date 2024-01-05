package base

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println(">>>>>>>>>>testHello>>>>>>>>>>")
	fmt.Println("hello")
	Variables()
	Euler()
	Triangle()
	Constants()
	fmt.Println(Add(12, 15))
	fmt.Println(Swap(strconv.Itoa(12), strconv.Itoa(12)))
	fmt.Println(Split(17))
	bytes := String2Bytes("testHello")
	bytes2String := Bytes2String(bytes)
	fmt.Println(bytes2String)
}
func TestBranch(t *testing.T) {
	fmt.Println(">>>>>>>>>>testBranch>>>>>>>>>>")
	fmt.Println(Convert2bin(25))
	Loop()
	IfFunc()
	fmt.Println(Sqrt(2))
	SwitchFunc()
	DeferFunc()
}
func TestIo(t *testing.T) {
	TestIoRead()
	PrintFile()
	Env()
	ReadFile()
}

type Animal interface {
	Say() string
	Walk()
}

type dog struct{}

func (d *dog) Say() string {
	return "我是一只小狗,旺旺旺..."
}

func (d *dog) Walk() {
	fmt.Println("我走起路来非常可爱")
}
func TestTypeAssert(t *testing.T) {
	fmt.Println(">>>>>>>>>>类型断言>>>>>>>>>>")
	//注意：
	//如果我们试图将一个非指针类型的变量转换成一个指针类型，那么.()会导致panic；
	//如果我们试图将一个非切片类型的变量转换成一个切片类型，那么.()会导致panic；
	//value, ok = interface{}.(Type)
	var i interface{} = "hello"
	s, ok := i.(string)
	fmt.Println("\"hello\" 类型断言 string ：", ok)
	if ok {
		fmt.Println(s)
	}
	i = nil
	_, ok = i.(bool)
	fmt.Println("nil 类型断言 bool ：", ok)

	var iface interface{} = &dog{}

	switch iface.(type) {
	case *dog: // 具体类型的指针
		fmt.Println("iface is *dog")
	case dog: //具体类型
		fmt.Println("iface is dog")
	case Animal: // interface
		fmt.Println("iface is animal")
	case interface{}: // 空interface
		fmt.Println("iface is interface")
	}

	//内置类型是不能定义方法的，接口类型不能作为方法的接收者。所以内置类型和接口不存在方法。
	//不管是内置类型还是自定义类型，都有描述他们的信息，称之为类型的元 数据信息，
	//每种数据类型的元数据信息都是全局唯一的，每种数据类型都是在_type字段的基础上，添加一些额外的字段进行管理的
	//type arraytype struct {
	//    typ _type
	//    elem *_type
	//    slice *_type
	//    len uintptr
	//}
	//type chantype struct {
	//    typ _type
	//    elem *_type
	//    dir uintptr
	//}
	//type slicetype struct {
	//    typ _type
	//    elem *_type
	//}

	//自定义类型，除了了类型共有的字段信息，还有一个uncommontype 信息,
	//记录了包路径，方法的个数，可导出方法的个数，
	//moff记录的是这些方法组成的相对于uncommon结构体偏移了多少字节
	//type uncommontype struct {
	//    pkgpath nameOff
	//    mcount  uint16 // number of methods
	//    xcount  uint16 // number of exported methods
	//    moff    uint32 // offset from this uncommontype to [mcount]method
	//    _       uint32 // unused
	//}

	//type myType1 = int32
	//type myType2 int32
	//type myType1 = int32 和 type myType2 int32有什么区别
	//myType1和int32底层元数据信息都是同一个， rune和int32就是这样的关系
	//myType2属于已有数据创建的新类型，它和int32是不同的类型，底层的数据元数据信息也是不同的
}
