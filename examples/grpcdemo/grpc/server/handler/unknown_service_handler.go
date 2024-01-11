package handler

import (
	"fmt"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/proto/hello"
	"google.golang.org/grpc"
	"net/http"
)

func UnknownServiceHandler(srv interface{}, stream grpc.ServerStream) error {

	fmt.Println("服务未找到...")
	resp := &hello.HelloResponse{
		Code: http.StatusOK,
		Msg:  "not found",
	}

	err := stream.SendMsg(resp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
