package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int32
	}{
		{1, 2, 3},
		{5, 6, 11},
		{2, 2, 4},
		{math.MaxInt32, 1, math.MaxInt32},
	}
	for _, tt := range tests {
		if actual := Add(tt.a, tt.b); actual != tt.c {
			//t.Errorf("test error v: %v \n", tt)
		}
	}
}
func TestAny(t *testing.T) {
	fmt.Println("test any func exec")
}
func TestAll(t *testing.T) {
	TestAdd(t)
	TestAny(t)
	TestFib1(t)
	TestFib2(t)
}

var times = 30

func TestFib1(t *testing.T) {
	fmt.Printf("fib1计算第%d个斐波那契数：%d\n", times, fib1(times-1))
}
func TestFib2(t *testing.T) {
	fmt.Printf("fib2计算第%d个斐波那契数：%d\n", times, fib2(times-1))
}

// 模糊测试或随机测试
//函数必须是Fuzz开头，唯一的参数只有*testing.F，没有返回值
//Fuzz tests必须在名为*_test.go的文件下才能执行
//fuzz target是个方法，它调用(*testing.F).Fuzz，第一个参数是 *testing.T，之后的参数就是称之为fuzzing arguments的参数，方法没有返回值
//每个fuzz test中只能有一个fuzz target
//调用f.Add()的时候需要参数类型跟fuzzing arguments顺序和类型都保持一致
//fuzzing arguments只支持以下类型：
//1、int, int8, int16, int32/rune, int64
//2、uint, uint8/byte, uint16, uint32, uint64
//3、string, []byte
//4、float32, float64
//5、bool

func FuzzFoo(f *testing.F) {
	f.Add(5, "hello")
	f.Fuzz(func(t *testing.T, i int, s string) {
		out, err := Foo(i, s)
		if err != nil && out != "" {
			t.Errorf("%q, %v", out, err)
		}
	})
}
func FuzzSliceSum(f *testing.F) {
	// 初始化随机数种子
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// 语料
	f.Add(10)

	f.Fuzz(func(t *testing.T, n int) {
		n %= 20

		var arr []int64
		var expect int64 // 期望值

		for i := 0; i < n; i++ {
			val := rand.Int63() % 1000000
			arr = append(arr, val)
			expect += val
		}

		// 自己求和的结果和调用函数求和的结果比对
		assert.Equal(t, expect, SliceSum(arr))
	})
}

// 测试的代码覆盖率
// cd .
// go test .
// go test -coverprofile c.out
// go tool cover -html c.out

// 性能测试
// cd .
// go test -bench .
// 结尾是Fib+数字的正则测试
// go test -bench 'Fib\d$' .
// 指定cpu核数
// go test -bench 'Fib\d$' -cpu 2,4,8,16 .
// 测试5s
// go test -bench 'Fib\d$' -benchtime 5s .
// 3000次
// go test -bench 'Fib\d$' -benchtime 3000x .
// 3轮
// go test -bench 'Fib\d$' -benchtime 2s -count 3 .
// -benchmem 参数看到内存分配的情况
// go test -bench 'Fib\d$' -benchtime 2s -count 3 -benchmem .

// bench可以生成 profile 文件，使用 go tool pprof 分析
// -cpuprofile=$FILE
// -memprofile=$FILE, -memprofilerate=N 调整记录速率为原来的 1/N。
// -blockprofile=$FILE
// go test -bench 'Fib1$' -cpuprofile cpu.out .
// go tool pprof cpu.out   ->  help  -> command options(web/text/tree/png...)
// go tool pprof -text cpu.out
// pprof 支持多种输出格式（图片、文本、Web等），直接在命令行中运行 go tool pprof 即可看到所有支持的选项：

// 进行性能调优
// -cpuprofile cpu.out ->  go tool pprof cpu.out -> 分析慢的地方（最大的框和最大的箭头） -> 进行优化 -> -cpuprofile cpu.out
func BenchmarkFib1(b *testing.B) {
	time.Sleep(time.Second * 1) // 模拟耗时准备任务
	b.ResetTimer()              // 重置定时器 把之前的耗时排除在计算之外
	for n := 0; n < b.N; n++ {
		fib1(times) // run fib(n) b.N times
	}
}
func BenchmarkFib2(b *testing.B) {
	b.StopTimer() //使用 StopTimer 和 StartTimer 避免将中间部分时间计算在内
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		fib2(times) // run fib(n) b.N times
	}
}
