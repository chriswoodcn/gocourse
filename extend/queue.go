package extend

import (
	"fmt"
	"log"
	"runtime"
)

type BaseError struct {
	module string
	code   int
	msg    string
}

func (e *BaseError) Error() string {
	format := fmt.Sprintf(" [module] %s [code] %d [msg] %s", e.module, e.code, e.msg)
	log.Println(format)
	return e.msg
}
func (e *BaseError) RuntimeError() {
	format := fmt.Sprintf(" [module] %s [code] %d [msg] %s", e.module, e.code, e.msg)
	log.Println(format)
}

// Queue 使用别名来扩展
type Queue []int

func (q *Queue) Push(v int) {
	if *q == nil {
		*q = make([]int, 10)
	}
	*q = append(*q, v)
}

func (q *Queue) Pop() (int, runtime.Error) {
	if q.isEmpty() {
		return 0, &BaseError{module: "Queue", code: 500, msg: "queue is empty"}
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head, nil
}
func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}
