package concurrency

import (
	"fmt"
	"math/rand"
	"runtime"
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

func TestStartMultiRoutine(t *testing.T) {
	var a [10]int
	for i := 0; i < 100; i++ {
		go func(n int) {
			for {
				//fmt.Printf("goroutine: #%d\n", n)// io操作会切换 交出控制权
				//a[i%10]++ //一直++操作，交不出控制权，会一直停留在这里
				a[n%10]++
				if a[n%10] >= 10 {
					runtime.Goexit() //退出当前协程
				}
				runtime.Gosched() //手动交出控制权
			}
		}(i)
	}
	// 主死从随  主程序一旦结束 所有开的协程结束
	time.Sleep(time.Second)
	fmt.Println(a)
}
func TestRoutineWithTopOberveThreads(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func(n int) {
			for {
				fmt.Printf("goroutine: #%d\n", n) // io操作会切换 交出控制权
			}
		}(i)
	}
	time.Sleep(time.Minute)
}

// goroutine可能切换的点
// 1.IO  Select
// 2.channel
// 3.等待锁
// 4.函数调用（有时）
// 5.runtime.Gosched()
func TestChanWorker(t *testing.T) {
	var channels [10]chan int
	//c := make(chan int)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Second)
}
func worker(id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if ok {
	//		fmt.Printf("worker %d receive from main : %c\n", id, n)
	//	} else {
	//		fmt.Printf("worker %d closed\n", id)
	//		runtime.Goexit()
	//	}
	//}
	// 可以使用更加简便的方式
	for n := range c {
		fmt.Printf("worker %d receive from main : %c\n", id, n)
	}
}
func TestChanCreateWorker(t *testing.T) {
	var channels [10]chan<- int
	//c := make(chan int)
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Second)
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		fmt.Printf("worker %d receive from main : %c\n", id, <-c)
	}()
	return c
}
func TestBufferedChan(t *testing.T) {
	//带缓冲的信道
	bufferCh := make(chan int, 3)
	//可以发送三次不会deadlock
	bufferCh <- 'a' + 1
	bufferCh <- 'a' + 2
	bufferCh <- 'a' + 3
	go worker(0, bufferCh)
	channelClose(bufferCh)
	time.Sleep(time.Second)
}

// 注意 发送方close
func channelClose(c chan int) {
	close(c)
}

func TestChanWorker2(t *testing.T) {
	c := make(chan int)
	done := make(chan bool)
	go worker2(0, c, done)
	for i := 0; i < 10; i++ {
		c <- 'a' + i
	}
	<-done
}
func worker2(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d receive from main : %c\n", id, n)
	}
	done <- true
}
func generate() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}
func createConsumer(id int) chan<- int {
	c := make(chan int)
	go func() {
		for v := range c {
			time.Sleep(time.Second)
			fmt.Printf("worker %d receive from main : %d\n", id, v)
		}
	}()
	return c
}
func TestSelectSample(t *testing.T) {
	var g1, g2 = generate(), generate()
	var consumer = createConsumer(0)
	var values []int
	tm := time.After(time.Second * 10)
	tick := time.Tick(time.Second)
	for {
		var activeConsumer chan<- int //注意此处 nil的chan可以作为select的case之后
		var activeValue int
		if len(values) > 0 {
			activeConsumer = consumer
			activeValue = values[0]
		}
		select {
		case n := <-g1:
			values = append(values, n)
		case n := <-g2:
			values = append(values, n)
		case activeConsumer <- activeValue: //select 中可以使用nil channel
			values = values[1:]
		case <-time.After(time.Millisecond * 800):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len is: ", len(values)) //定时器
		case <-tm:
			fmt.Println("byebye")
			return
		}
	}
}
