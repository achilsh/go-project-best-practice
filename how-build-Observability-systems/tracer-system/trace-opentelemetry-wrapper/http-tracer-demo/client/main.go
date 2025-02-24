package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	httpTracer "tracer-opentelemetry-libs/otal-tracer"
)


func main() {
	trp, err := httpTracer.InitOtel("http-client", "otlp")
	if err != nil {
		fmt.Println("init otel tracer provide fail. err: ", err)
		return 
	}
	
	for i:=0; i< 100; i++ {
		httpCli := httpTracer.NewHTTPClient(trp,"http-client-request" )
		if httpCli == nil {
			fmt.Println("create http client fail")
			return 
		}
		//

		var reqData =&httpTracer.RequestData{
			CustomerId: 10000 +i,
		}
		reqDataBin, _ := json.Marshal(reqData)

		err, retData := httpCli.Request(context.Background(), fmt.Sprintf("request-%d", i), 2 /* 1: get, 2: post*/,  "http://0.0.0.0:5656/dispatch", reqDataBin)
		if err != nil {
			fmt.Println("get request fail, err: ", err)
			return
		}
		_ = retData
		time.Sleep(10*time.Microsecond)
	}

	select {}
}