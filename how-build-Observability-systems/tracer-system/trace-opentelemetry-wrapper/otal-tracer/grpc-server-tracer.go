package otal_tracer

import (
	"context"
	"fmt"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	pb "tracer-opentelemetry-libs/gen/go/otal-tracer"
)

type GrpcServerTracer struct {
	server   *grpc.Server
	hostPort string
	serverName string
	tp  trace.TracerProvider
	tracer trace.Tracer
	pb.UnimplementedGrpcCallDemoServer
}

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
	g.tp = sp
	g.server= server
	g.tracer = sp.Tracer("grpc-server")
	return g
}


func (g *GrpcServerTracer) Run() {
	lis, err := net.Listen("tcp", g.hostPort)
	if err != nil {
		fmt.Println("Unable to create http listener", err)
		return 
	}
	//
	pb.RegisterGrpcCallDemoServer(g.server, g)
	fmt.Printf("Starting, %v, type grpc\n",g.hostPort)
	err = g.server.Serve(lis)
	if err != nil {
		fmt.Printf("Unable to start gRPC server, err: %v\n", err)
		return 
	}
	return
}

func(g *GrpcServerTracer)Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error){
	ctx, span := g.tracer.Start(ctx, "Echo", trace.WithAttributes(attribute.String("extra.key", "extra.value")))
	defer  span.End() 
	fmt.Println("[Echo] receive in message: ", in.Value)

	ret := &pb.StringMessage{
		Value: in.Value +", out data.",
	}
	return ret,nil
}


func(g *GrpcServerTracer)Hello(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	ctx, span := g.tracer.Start(ctx, "Hello", trace.WithAttributes(attribute.String("extra.key", "extra.value")))
	defer  span.End() 
	fmt.Println("[Hello] receive in message: ", in.Value)

	ret := &pb.StringMessage{
		Value: in.Value +", out data.",
	}
	
	return ret,nil
}

// 实现pb 中定义的接口
// FindNearest(context.Context, *DriverLocationRequest) (*DriverLocationResponse, error)
// func (g *GrpcServerTracer) FindNearest(ctx context.Context, req *DriverLocationRequest) (*DriverLocationResponse, error) {
// 	// 从请求中获取数据
// 	// customerID := req.CustomerId
// 	// fmt.Printf("customerID: %v\n", customerID
// }