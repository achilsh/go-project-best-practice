package main

import (
	grpcTracer "tracer-opentelemetry-libs/otal-tracer"
)

func main() {
	item := grpcTracer.NewGrpcServerTracer("grpc-server", "0.0.0.0:6567", "otlp")
	item.Run()
}