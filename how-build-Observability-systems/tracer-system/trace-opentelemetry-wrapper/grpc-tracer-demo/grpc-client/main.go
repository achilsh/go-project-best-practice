package main

import (
	"context"
	"fmt"

	pb "tracer-opentelemetry-libs/gen/go/otal-tracer"
	grpcTracer "tracer-opentelemetry-libs/otal-tracer"
)
func main() {
	sp, err := grpcTracer.InitOtel("grpc-client", "otlp")
	if err != nil {
		fmt.Println("init otel fail, err: ", err)
		return 
	}
	//
	cli := grpcTracer.NewGrpcClientTracer("0.0.0.0:6567", sp)
	if cli == nil {
		fmt.Println("create grp client fail")
		return 
	}
	//
	echoIn := &pb.StringMessage {
		Value: "----",
	}
	//
	oMsg, err := cli.Echo(context.Background(), echoIn)
	if err != nil {
		fmt.Println("echo request fail, err: ", err)
		return 
	}
	fmt.Println("echo msg: ", oMsg.GetValue())

	select {}
}