package main

import (
	"fmt"

	httpTracer "tracer-opentelemetry-libs/otal-tracer"
)

func main() {
	trp, err := httpTracer.InitOtel("http-server", "otlp")
	if err != nil {
		fmt.Println("init otel tracer provide fail. err: ", err)
		return 
	}

	httpSrv := httpTracer.NewHttpServerTracer(trp)
	if httpSrv == nil {
		fmt.Println("init http server fail, obj is nil")
		return 
	}
	if err := httpSrv.Run(); err != nil {
		fmt.Println("run http server fail, err: ", err)
		return 
	}

}