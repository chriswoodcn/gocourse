package main

import (
	"fmt"
	"gocourse/examples/error/errorHandle"
	"gocourse/examples/error/filelist"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func errPanic(http.ResponseWriter, *http.Request) error {
	panic("errPanic")
}
func TestErrorWrap(t *testing.T) {
	tests := []struct {
		handler errorHandle.AppHandler
		code    int
		message string
	}{
		{errPanic, 500, http.StatusText(500)},
		{filelist.HandleFileList, http.StatusOK, http.StatusText(http.StatusOK)},
	}
	for _, tt := range tests {
		tar := errorHandle.ErrorWrap(tt.handler)
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8888/list/aaa.txt", nil)
		tar(res, req)
		all, _ := io.ReadAll(res.Body)
		fmt.Printf("response code: %d  body: %s\n", res.Code, string(all))
	}
}
