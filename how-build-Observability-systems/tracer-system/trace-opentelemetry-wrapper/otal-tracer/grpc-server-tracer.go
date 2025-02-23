package otal_tracer

import (
	"fmt"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type GrpcServerTracer struct {
	server   *grpc.Server
	hostPort string
	serverName string
	tracer  trace.TracerProvider
}

// 其中 pb 的接口
// var _ DriverServiceServer = (*GrpcServerTracer)(nil)

func NewGrpcServerTracer(srvName string, host string, otelExporter string)  *GrpcServerTracer{
	g := &GrpcServerTracer {
		serverName: srvName,
		hostPort: host,
	}

	sp, err := InitOtel(g.serverName, otelExporter)
	if err != nil {
		return nil
	}
	server := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler(otelgrpc.WithTracerProvider(sp))),
	)
	//
	g.tracer = sp
	g.server= server
	return g
}


func (g *GrpcServerTracer) Run() {
	lis, err := net.Listen("tcp", g.hostPort)
	if err != nil {
		fmt.Println("Unable to create http listener", err)
		return 
	}
	//
	// RegisterDriverServiceServer(g.server, g)
	fmt.Printf("Starting, %v, type grpc\n",g.hostPort)
	err = g.server.Serve(lis)
	if err != nil {
		fmt.Printf("Unable to start gRPC server, err: %v\n", err)
		return 
	}
	return
}

// 实现pb 中定义的接口
// FindNearest(context.Context, *DriverLocationRequest) (*DriverLocationResponse, error)
// func (g *GrpcServerTracer) FindNearest(ctx context.Context, req *DriverLocationRequest) (*DriverLocationResponse, error) {
// 	// 从请求中获取数据
// 	// customerID := req.CustomerId
// 	// fmt.Printf("customerID: %v\n", customerID
// }