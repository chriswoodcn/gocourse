package main

import (
	"context"
	"fmt"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/proto/hello"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	clientConn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor), grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
	if err != nil {
		log.Println(err)
		return
	}

	helloServiceClient := hello.NewHelloServiceClient(clientConn)

	helloResponse, err := helloServiceClient.SayHello(context.Background(), &hello.HelloRequest{})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(helloResponse)

}
