package main

import (
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/etcd/register"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/proto/hello"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/server/handler"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	_ "net/http/pprof"
)

func main() {

	server := grpc.NewServer(grpc.StatsHandler(&handler.StatsHandler{}), grpc.UnknownServiceHandler(handler.UnknownServiceHandler))

	// 健康检查
	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.Health", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(server, healthServer)

	// 注册HelloService rpc服务
	helloService := new(service.HelloService)
	hello.RegisterHelloServiceServer(server, helloService)

	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------------
	// 创建一个注册器
	etcdRegister, err := register.NewEtcdRegister()
	if err != nil {
		log.Println(err)
		return
	}

	defer etcdRegister.Close()

	serviceName := "order-service-2"

	addr := "127.0.0.1:8000"

	// 注册服务
	err = etcdRegister.RegisterServer("/etcd/"+serviceName, addr, 5)
	if err != nil {
		log.Printf("register error %v \n", err)
		return
	}

	server.Serve(listen)

}
