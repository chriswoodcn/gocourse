package interceptor

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"time"
)

var (
	HTTPReqDuration *prometheus.HistogramVec

	HTTPReqTotal *prometheus.CounterVec

	TaskRunning *prometheus.GaugeVec
)

func init() {

	// 监控接口请求耗时
	// 指标类型是Histogram
	HTTPReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "http request latencies in seconds",
		Buckets: nil,
	}, []string{"method"})

	// "method"、"path" 是 label

	// 监控接口请求次数
	// 指标类型是 Counter
	HTTPReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "total number of http requests",
	}, []string{"method", "status"})
	// "method"、"path"、"status" 是 label

	// 监控当前在执行的task数量
	// 监控类型是Gauge
	TaskRunning = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "task_running",
		Help: "current count  of running task",
	}, []string{"type", "state"})
	// "type"、"state" 是 label

	prometheus.MustRegister(
		HTTPReqDuration,
		HTTPReqTotal,
		TaskRunning,
	)

}

func PrometheusInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	start := time.Now()

	duration := float64(time.Since(start)) / float64(time.Second)

	// 请求数加1
	HTTPReqTotal.With(prometheus.Labels{
		"method": info.FullMethod,
		"status": "200",
	}).Inc()

	//  记录本次请求处理时间
	HTTPReqDuration.With(prometheus.Labels{
		"method": info.FullMethod,
	}).Observe(duration)

	// 模拟新建任务
	TaskRunning.With(prometheus.Labels{
		"type":  info.FullMethod,
		"state": "200",
	}).Inc()

	// 模拟任务完成
	TaskRunning.With(prometheus.Labels{
		"type":  info.FullMethod,
		"state": "200",
	}).Dec()

	return handler(ctx, req)
}
