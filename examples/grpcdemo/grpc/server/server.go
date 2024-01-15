package main

import (
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/etcd"
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

	server := grpc.NewServer(
		grpc.StatsHandler(&handler.StatsHandler{}),
		grpc.UnknownServiceHandler(handler.UnknownServiceHandler))

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
	point := etcd.ServerEndPoint{}
	// 创建一个注册器
	etcdRegister, err := register.NewEtcdRegister(point)
	if err != nil {
		log.Println(err)
		return
	}

	defer func(etcdRegister *register.EtcdRegister) {
		_ = etcdRegister.Close()
	}(etcdRegister)

	serviceName := "order-service-2"

	addr := "127.0.0.1:8000"

	// 注册服务
	err = etcdRegister.RegisterServer("/etcd/"+serviceName, addr, 30)
	// 源码 func (l *lessor) recvKeepAlive(resp *pb.LeaseKeepAliveResponse) 中
	// nextKeepAlive := time.Now().Add((time.Duration(karesp.TTL) * time.Second) / 3.0)
	// 因此这个过期时间(s) / 3 的间隔为每次keepAlive发出请求的时间
	if err != nil {
		log.Printf("register error %v \n", err)
		return
	}

	_ = server.Serve(listen)

}
