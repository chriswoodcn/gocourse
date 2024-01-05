package main

import (
	"cn.chriswood/gocourse/examples/error/errorHandle"
	"cn.chriswood/gocourse/examples/error/filelist"
	"errors"
	"fmt"
	"net/http"
)

func ejectError() {
	err1 := errors.New("this is customer error")
	fmt.Println(err1)
}

var server *http.Server

func startWeb() {
	http.HandleFunc("/", errorHandle.ErrorWrap(filelist.HandleFileList))
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

func main() {
	ejectError()
	startWeb()
}
