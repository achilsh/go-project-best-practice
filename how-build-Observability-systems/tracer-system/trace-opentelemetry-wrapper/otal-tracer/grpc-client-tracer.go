package otal_tracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "tracer-opentelemetry-libs/gen/go/otal-tracer"
)

type  GrpcClientTracer struct {
	tracer trace.TracerProvider
	conn *grpc.ClientConn
	serverHost string 
}

func NewGrpcClientTracer(host string, tracerItem trace.TracerProvider) *GrpcClientTracer{
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err !=nil {
		fmt.Println("new grpc client fail, err: ", err)
		return nil
	}
	ret := &GrpcClientTracer{
		conn:conn,
		tracer: tracerItem,
		serverHost: host,
	}

	return ret
}
func (g *GrpcClientTracer) Close() {
	if g.conn != nil {
		g.conn.Close()
		g.conn = nil
	}
}

func NewClient(c *grpc.ClientConn) pb.GrpcCallDemoClient{
	return pb.NewGrpcCallDemoClient(c)
}


func(c *GrpcClientTracer) Echo(ctx context.Context, in *pb.StringMessage, opts ...grpc.CallOption) (*pb.StringMessage, error) {
	 tr := c.tracer.Tracer("grpc-client")
	 ret := &pb.StringMessage{
		Value: in.GetValue() +",  1111",
	 }
	 ctx, sp := tr.Start(ctx, "echo request", trace.WithSpanKind(trace.SpanKindClient))
	 // 可以设置一些属性：比如： sp.SetAttributes(attribute.Key("param.driver.location").String(location))
	 defer sp.End()

	 return ret, nil
}
	
func(c *GrpcClientTracer) Hello(ctx context.Context, in *pb.StringMessage, opts ...grpc.CallOption) (*pb.StringMessage, error) {
	//TODO:
	return nil,nil
}
