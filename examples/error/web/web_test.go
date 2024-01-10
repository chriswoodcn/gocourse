package web

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	_ "net/http/pprof"
	"testing"
)

func ejectError() {
	err1 := errors.New("this is customer error")
	fmt.Println(err1)
}

var server *http.Server

func startWeb() {
	http.HandleFunc("/", ErrorWrap(HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}

type R[E any] struct {
	Code int
	Msg  string
	Data []E
}

func TestWeb(t *testing.T) {
	ejectError()
	startWeb()
}

// debug/pprof 在线查看pprof
// go tool pprof 分析性能

func errPanic(http.ResponseWriter, *http.Request) error {
	panic("errPanic")
}

var tests = []struct {
	handler AppHandler
	code    int
	message string
}{
	{errPanic, 500, http.StatusText(500)},
	{HandleFileList, http.StatusOK, http.StatusText(http.StatusOK)},
}

// 模拟response和request 测试handler
func TestErrorWrap(t *testing.T) {
	for _, tt := range tests {
		tar := ErrorWrap(tt.handler)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8888/list/aaa.txt", nil)
		tar(res, req)
		all, _ := io.ReadAll(res.Body)
		fmt.Printf("response code: %d  body: %s\n", res.Code, string(all))
	}
}

// 模拟开启一个server测试
func TestServer(t *testing.T) {
	for _, tt := range tests {
		tar := ErrorWrap(tt.handler)
		server := httptest.NewServer(http.HandlerFunc(tar))
		fmt.Println(server)
		//res, _ := http.Get(server.URL)
		res, _ := http.Get(server.URL + "/list/aaa.txt")
		all, _ := io.ReadAll(res.Body)
		fmt.Printf("response code: %d  body: %s\n", res.StatusCode, string(all))
	}

}
