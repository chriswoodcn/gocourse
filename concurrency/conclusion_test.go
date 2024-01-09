package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 生成器模式 返回一个chan
func messageGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
			i++
			c <- fmt.Sprintf("message name %s content %d", name, i)
		}
	}()
	return c
}

// 同时等待多个任务
func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)
	// 第一种
	//go func() {
	//    for {
	//        c <-  <-c1
	//    }
	//}()
	//go func() {
	//    for {
	//        c <-  <-c2
	//    }
	//}()
	// 第二种
	go func() {
		for {
			select {
			case n := <-c1:
				c <- n
			case n := <-c2:
				c <- n
			}
		}
	}()
	return c
}
func fanIn2(cs ...chan string) chan string {
	c := make(chan string)
	// 此部分代码会踩坑  cc 会一直是最后一个chan
	// cc只保存了一份 go开启多个协程 但是cc循环只保存了一份
	// 协程是异步的还没有运行到不知道cc是什么值，等到运行的时候才发现都是同一个cc
	//for _, cc := range cs {
	//	go func() {
	//		for {
	//			c <- <-cc
	//		}
	//	}()
	//}
	// 传入协程中  或者{}中使用闭包保存一份拷贝
	for _, cc := range cs {
		go func(ch chan string) {
			for {
				c <- <-ch
			}
		}(cc)
	}
	return c
}
func TestConclusion(t *testing.T) {
	m1 := messageGen("service1")
	m2 := messageGen("service2")
	//for {
	//	fmt.Println( <-m1)
	//	fmt.Println( <-m2)
	//}
	// 每个messageGen是独立的，可以抽象为一个服务或一个任务
	// 拿到chan之后可以源源不断的从其中收消息

	// 同时等待多个任务
	//a := fanIn(m1, m2)
	//for {
	//    fmt.Println("receive message :   ", <-a)
	//}
	m3 := messageGen("service3")
	a2 := fanIn2(m1, m2, m3)
	for {
		fmt.Println("receive message :   ", <-a2)
	}
}

// 任务的控制
// 非阻塞等待
func noBlockWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

// 超时等待
func timeoutWait(c chan string, timeout chan time.Time) (string, bool) {
	select {
	case <-timeout:
		return "", false
	case m := <-c:
		return m, true
	}
}

// 任务中断退出
func messageGenWithGracefulExit(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(3000)) * time.Millisecond):
				c <- fmt.Sprintf("message name %s content %d", name, i)
			case <-done:
				return
			}
			i++
		}
	}()
	return c
}
