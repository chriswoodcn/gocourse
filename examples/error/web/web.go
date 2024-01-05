package main

import (
	"errors"
	"fmt"
	"gocourse/examples/error/errorHandle"
	"gocourse/examples/error/filelist"
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
