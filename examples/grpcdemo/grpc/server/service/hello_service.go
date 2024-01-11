package service

import (
	"context"
	"fmt"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/proto/hello"
	"io"
	"net/http"
	"strconv"
	"time"
)

type HelloService struct {
	*hello.UnimplementedHelloServiceServer
}

// 一元RPC

func (h HelloService) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {

	fmt.Println("say hello")
	msg := "Hello! 8000" + request.GetName()

	resp := &hello.HelloResponse{
		Code: http.StatusOK,
		Msg:  msg,
	}

	return resp, nil
}

// 服务端流 服务端一直发Send

func (h HelloService) LotsOfReplies(request *hello.HelloRequest, server hello.HelloService_LotsOfRepliesServer) error {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		resp := &hello.HelloResponse{
			Code: http.StatusOK,
			Msg:  "Hello! " + request.GetName() + "@" + strconv.Itoa(i),
		}
		err := server.Send(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// 客户端流 客户端一直在发

func (h HelloService) LotsOfGreetings(server hello.HelloService_LotsOfGreetingsServer) error {

	for {

		helloRequest, err := server.Recv()

		// 客户端发送完成
		if err == io.EOF {

			// 响应客户端结果并关闭连接
			return server.SendAndClose(&hello.HelloResponse{
				Code: http.StatusOK,
				Msg:  "全部接收完成...",
			})

		}

		// 接收异常
		if err != nil {
			return err
		}

		fmt.Println(helloRequest)
	}

}

// 双向流

func (h HelloService) BidiHello(server hello.HelloService_BidiHelloServer) error {

	for {

		// 接收客户端消息
		helloRequest, err := server.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		resp := &hello.HelloResponse{
			Code: http.StatusOK,
			Msg:  "Hello!" + helloRequest.GetName(),
		}

		// 响应客户端消息
		err = server.Send(resp)

		if err != nil {
			return err
		}

	}
}
