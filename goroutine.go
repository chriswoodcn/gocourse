package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func testGoroutine() {
	//go f(x, y, z)会启动一个新的 Go 程并执行f(x, y, z)
	go say("world")
	//Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。
	//sync 包提供了这种能力，不过在 Go 中并不经常用到

	say("hello")
}

func testChan() {
	// 信道 无缓冲的
	ch := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6}
	//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
	//这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步
	go sum(s, ch)
	r := <-ch
	fmt.Println(r)

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

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

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
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}
func testChanDefault() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case i := <-tick:
			fmt.Print(i, "  ")
			fmt.Println("tick.")
		case i := <-boom:
			fmt.Print(i, "  ")
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// 等价二叉查找树示例
func testTree() {

}

func oneSenderAndMultiReceivers() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())
	log.SetFlags(1 + 4)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)

	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// the only sender can close the channel safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
func multiSendersAndOneReceiver() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	log.SetFlags(5)

	// ...
	const MaxRandomNumber = 100000
	const NumSenders = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the receiver of channel dataCh.
	// Its reveivers are the senders of channel dataCh.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				value := rand.Intn(MaxRandomNumber)

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}()
	}

	// the receiver
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				// the receiver of the dataCh channel is
				// also the sender of the stopCh cahnnel.
				// It is safe to close the stop channel here.
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}
func multiSendersAndMultiReceivers() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	log.SetFlags(5)

	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 10
	const NumSenders = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an additional signal channel.
	// Its sender is the moderator goroutine shown below.
	// Its reveivers are all senders and receivers of dataCh.
	toStop := make(chan string, 1)
	// the channel toStop is used to notify the moderator
	// to close the additional signal channel (stopCh).
	// Its senders are any senders and receivers of dataCh.
	// Its reveiver is the moderator goroutine shown below.

	var stoppedBy string

	// moderator 仲裁者 启动一个仲裁者协程
	go func() {
		// 接收方一直阻塞，直到有数据过来执行后续操作
		stoppedBy = <-toStop // part of the trick used to notify the moderator
		// to close the additional signal channel.
		//  >>>>>>>>>>>>>>> 接收到toStop信道的参数后执行关闭stopCh信道操作 >>>>>>>>>>>>>>>
		close(stopCh)
	}()

	// senders 发送者 启动多个发送者协程
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			//无限循环
			for {
				value := rand.Intn(MaxRandomNumber)
				// 随机到value符合条件，发送消息到toStop信道,唤起仲裁者执行逻辑
				if value%101 == 0 {
					// here, a trick is used to notify the moderator
					// to close the additional signal channel.
					//select 语句使一个 Go 程可以等待多个通信操作
					select {
					case toStop <- "sender#" + id:
					default:
					}
					// 退出发送者go程
					return
				}

				// the first select here is to try to exit the
				// goroutine as early as possible.
				//select 语句使一个 Go 程可以等待多个通信操作
				select {
				case <-stopCh:
					// 接收到stopCh信道消息，退出发送者go程
					return
				default:
				}
				//select 语句使一个 Go 程可以等待多个通信操作
				select {
				case <-stopCh:
					// 接收到stopCh信道消息，退出发送者go程
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers 接受者
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()
			//无限循环
			for {
				// same as senders, the first select here is to
				// try to exit the goroutine as early as possible.
				select {
				case <-stopCh:
					// 接收到stopCh信道消息，退出接受者go程
					return
				default:
				}

				select {
				case <-stopCh:
					// 接收到stopCh信道消息，退出接受者go程
					return
				case value := <-dataCh:
					// 接收到dataCh信道消息，拿到传递过来的数据 value
					// 随机到value符合条件，发送消息到toStop信道,唤起仲裁者执行逻辑
					if value%111 == 0 {
						// the same trick is used to notify the moderator
						// to close the additional signal channel.
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	// ...
	wgReceivers.Wait()
	log.Println("stopped by", stoppedBy)
}
func startMillionRoutine() {
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

var total int
var wg sync.WaitGroup

// 互斥锁
var lockMutex sync.Mutex

// 读写锁  绝大多数都是读多写少 读的时候不用上锁 写的时候上锁
var lockRW sync.RWMutex

func read() {
	defer wg.Done()
	lockRW.RLock()
	time.Sleep(time.Second)
	lockRW.RUnlock()
	fmt.Println("读取成功 " + strconv.Itoa(total))
}
func write() {
	defer wg.Done()
	lockRW.Lock()
	total += 999
	time.Sleep(time.Second)
	lockRW.Unlock()
	fmt.Println("修改成功 " + strconv.Itoa(total))
}

func _add() {
	defer wg.Done()
	for i := 1; i < 10000; i++ {
		lockMutex.Lock()
		total += 1
		lockMutex.Unlock()
	}
}
func _minus() {
	defer wg.Done()
	for i := 1; i < 10000; i++ {
		lockMutex.Lock()
		total -= 1
		lockMutex.Unlock()
	}
}
func useTotal() {
	wg.Add(2)
	go _add()
	go _minus()
	wg.Wait()
	fmt.Println(total)
}
func useRwLock() {
	wg.Add(10)
	for i := 0; i < 8; i++ {
		go read()
	}
	for i := 0; i < 2; i++ {
		go write()
	}
	wg.Wait()
}
func main() {
	//testGoroutine()
	//testChan()
	//testChanDefault()
	//fmt.Println("cpus:", runtime.NumCPU())
	//fmt.Println("goroot:", runtime.GOROOT())
	//fmt.Println("archive:", runtime.GOOS)
	//oneSenderAndMultiReceivers()
	//multiSendersAndOneReceiver()
	//multiSendersAndMultiReceivers()
	//startMillionRoutine()
	//useTotal()
	useRwLock()
}
