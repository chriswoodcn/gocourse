package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
// 是非抢占式多任务处理，由协程主动交出控制权
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func TestGoroutine1(t *testing.T) {
	//go f(x, y, z)会启动一个新的 Go 程并执行f(x, y, z)
	go say("world")
	//Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。
	//sync 包提供了这种能力，不过在 Go 中并不经常用到
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}
func TestChan1(t *testing.T) {
	// 信道 无缓冲的
	ch := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6}
	//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
	//这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步
	go sum(s, ch)
	r := <-ch
	fmt.Println(r)

}
func TestChan2(t *testing.T) {
	//带缓冲的信道
	bufferCh := make(chan int, 100)
	//仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞
	bufferCh <- 1
	bufferCh <- 2
	close(bufferCh)
	//fmt.Println(<-bufferCh)
	//fmt.Println(<-bufferCh)

	//发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
	//接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：
	//若没有值可以接收且信道已被关闭，那么在执行完之后 ok 会被设置为 false
	for {
		v, ok := <-bufferCh
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func TestChan3(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c2, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c2 <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func TestChan4(t *testing.T) {
	//select 语句使一个 Go 程可以等待多个通信操作
	//select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。
	//当多个分支都准备好时会随机选择一个执行
	c2 := make(chan int, 10)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()
	fibonacci2(c2, quit)
}

func TestStartMillionRoutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go func(n int) {
			for {
				fmt.Printf("goroutine: #%d\n", n)
				time.Sleep(time.Second)
			}
		}(i)
	}
	// 主死从随
	time.Sleep(time.Second * 10)
}
