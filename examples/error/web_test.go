package main

import (
	"fmt"
	"github.com/chriswoodcn/gocourse/examples/error/errorHandle"
	"github.com/chriswoodcn/gocourse/examples/error/filelist"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func errPanic(http.ResponseWriter, *http.Request) error {
	panic("errPanic")
}

var tests = []struct {
	handler errorHandle.AppHandler
	code    int
	message string
}{
	{errPanic, 500, http.StatusText(500)},
	{filelist.HandleFileList, http.StatusOK, http.StatusText(http.StatusOK)},
}

// 模拟response和request 测试handler
func TestErrorWrap(t *testing.T) {
	for _, tt := range tests {
		tar := errorHandle.ErrorWrap(tt.handler)
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
		tar := errorHandle.ErrorWrap(tt.handler)
		server := httptest.NewServer(http.HandlerFunc(tar))
		fmt.Println(server)
		//res, _ := http.Get(server.URL)
		res, _ := http.Get(server.URL + "/list/aaa.txt")
		all, _ := io.ReadAll(res.Body)
		fmt.Printf("response code: %d  body: %s\n", res.StatusCode, string(all))
	}

}
