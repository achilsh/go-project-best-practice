package main

// "context"
// "log"
// "net/http"

// "github.com/gin-gonic/gin"

// "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
// "go.opentelemetry.io/otel"
// "go.opentelemetry.io/otel/attribute"
// stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
// "go.opentelemetry.io/otel/propagation"
// sdktrace "go.opentelemetry.io/otel/sdk/trace"
// oteltrace "go.opentelemetry.io/otel/trace"

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	serverTracer "tracer-opentelemetry-libs/otal-tracer"
)

func NewGinServer() {
	// r := gin.Default()

	// r.Use(otelgin.Middleware("my-server"))

	// r.GET("/ping", func(c *gin.Context) {
	// 	tracer := otel.GetTracerProvider().Tracer("my-server")
	// 	ctx := c.Request.Context()
	// 	ctx, span := tracer.Start(ctx, "ping")
	// 	defer span.End()

	// 	span.SetAttributes(attribute.String("custom", "value"))

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run(":8080")
}
func main() {
	tp, err  := serverTracer.InitOtel("gin-server", "otlp")
	if err != nil {
		fmt.Println("init tracer provider fail, err: ", err)
		return 
	}
	// var tracer   trace.Tracer

	r := gin.Default()
	r.Use(otelgin.Middleware("my-server", otelgin.WithTracerProvider(tp)))
	//
	r.GET("/user/demo", func(c *gin.Context){
		//
		tracer := tp.Tracer("my-gin-server")
		_, span := tracer.Start(c.Request.Context(), "user-demo-1",
			oteltrace.WithSpanKind(oteltrace.SpanKindServer),
			 oteltrace.WithAttributes(attribute.String("id", "assss")))
		defer span.End()


		c.JSON(200, gin.H{
			"message": "user demo",
		})
		fmt.Print("send response to client.")
	})

	_ = r.Run(":5650")
}