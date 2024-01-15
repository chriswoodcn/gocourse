package main

import (
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/prometheus/interceptor"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/proto/hello"
	"github.com/chriswoodcn/gocourse/examples/grpcdemo/grpc/server/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {

	var opts []grpc.ServerOption
	opts = append(opts, grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor))
	opts = append(opts, grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor, interceptor.PrometheusInterceptor))

	server := grpc.NewServer(opts...)

	helloService := new(service.HelloService)
	hello.RegisterHelloServiceServer(server, helloService)

	grpc_prometheus.Register(server)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.EnableClientHandlingTimeHistogram()

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		_ = http.ListenAndServe("127.0.0.1:9001", nil)
	}()

	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		return
	}

	_ = server.Serve(listen)
}
