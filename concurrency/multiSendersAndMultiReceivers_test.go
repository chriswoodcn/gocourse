package concurrency

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 多个发送者和多个接受者  增加一个仲裁者关闭管道
func TestMultiSendersAndMultiReceivers(t *testing.T) {
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
