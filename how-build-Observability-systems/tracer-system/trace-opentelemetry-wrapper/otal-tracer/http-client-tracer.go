package otal_tracer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

type HTTPClient struct {
	TracerProvider trace.TracerProvider
	Client         *http.Client
	TracerName string
	Tracer trace.Tracer
}

func NewHTTPClient(tp trace.TracerProvider, tracerName string) *HTTPClient {
	return &HTTPClient{
		TracerProvider: tp,
		Client: &http.Client{
			Transport: otelhttp.NewTransport(
				http.DefaultTransport,
				otelhttp.WithTracerProvider(tp),
			),
		},
		TracerName: tracerName,
		Tracer: tp.Tracer(tracerName),
	}
}

// get: url is: http://0.0.0.0:5657/hello-world
// post  data is request body
func (h *HTTPClient) Request(ctx context.Context, spaName string, method int, url string, data []byte)(error, []byte) {
	var reqMethod string 
	var req *http.Request = nil 
	
	ctx, span := h.Tracer.Start(ctx, spaName,trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(semconv.PeerService("ExampleService")))
		// span.SetAttributes(
			// otelsemconv.PeerServiceKey.String("mysql"),
			// attribute.
			// 	Key("sql.query").
			// 	String(fmt.Sprintf("SELECT * FROM customer WHERE customer_id=%d", customerID)))
	fmt.Println("get span: ", span)
	defer span.End()

	if method == 1 {
		reqMethod =  http.MethodGet
		req, _ = http.NewRequestWithContext(ctx, reqMethod, url, nil)
	} else {
		reqMethod = http.MethodPost
		req, _ = http.NewRequestWithContext(ctx, reqMethod, url, bytes.NewBuffer(data))
	}

	req.Header.Set("Content-Type", "application/json")
	//////////////////////////////////////////////////
	fmt.Printf("Sending request...\n")
	res, err := h.Client.Do(req)
	if err != nil {
		fmt.Printf("request http fail, err: %v\n", err)
		return err, nil
	}
	defer  res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("send %v fail, err: %v\n", reqMethod, err)
		return err, nil
	} 
	
	fmt.Println("receive response data: ", string(responseBody))
	return nil, responseBody
}