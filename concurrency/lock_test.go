package concurrency

import (
	context2 "context"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

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

// 使用全局变量加锁
func useTotal() {
	wg.Add(2)
	go _add()
	go _minus()
	wg.Wait()
	fmt.Println(total)
}

// 使用读写锁
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

// 使用context WithCancel
func useContextCancel() {
	context, cancel := context2.WithCancel(context2.Background())
	wg.Add(1)
	go func(context context2.Context) {
		defer wg.Done()
		for {
			select {
			case <-context.Done():
				fmt.Println("context done,return goroutine")
				return
			default:
				fmt.Println("continue goroutine trick")
				time.Sleep(time.Second)
			}
		}
	}(context)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}

// 使用context WithTimeout
func useContextTimeout() {
	context, _ := context2.WithTimeout(context2.Background(), time.Second*5)
	wg.Add(1)
	go func(context context2.Context) {
		defer wg.Done()
		for {
			select {
			case <-context.Done():
				fmt.Println("context done,return goroutine")
				return
			default:
				fmt.Println("continue goroutine trick")
				time.Sleep(time.Second)
			}
		}
	}(context)
	wg.Wait()
}
func TestLock1(t *testing.T) {
	useTotal()
}
func TestLock2(t *testing.T) {
	useRwLock()
}
func TestLock3(t *testing.T) {
	useContextCancel()
}
func TestLock4(t *testing.T) {
	useContextTimeout()
}
