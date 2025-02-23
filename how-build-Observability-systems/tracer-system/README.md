jaeger 的作用：
    Distributed context propagation
    Distributed transaction monitoring
    Root cause analysis
    Service dependency analysis
    Performance / latency optimization （看每步骤的耗时）

* 参考源码： https://github.com/uber/jaeger ； https://github.com/jaegertracing/jaeger
* 文档： https://www.jaegertracing.io/ 
*  opentelemetry 文档： https://opentelemetry.io/ ； https://opentelemetry.io/zh/

*  jaeger 作用： receives tracing telemetry data and provides processing, aggregation, data mining, and visualizations of that data.
* 如何使用 jaeger? 建议使用: OpenTelemetry SDKs， 
* jaeger 的后端存储选型：
  1. 默认： Cassandra ， Elasticsearch
  2. 本地嵌入式 kv 数据库（go实现的）：https://github.com/hypermodeinc/badger
  3. grpc api 方式： TimescaleDB via Promscale, ClickHouse.
  4. 开源社区支持： InfluxDB, Amazon DynamoDB, YugabyteDB(YCQL).
   
* RocksDB 和 badger的区别; span的含义？
* 如何给应用程序添加监控：1） 检测应用程序 2) 配置一个导出器
* OpenTelemetry-Go is the Go implementation of OpenTelemetry. It provides a set of APIs to directly measure performance and behavior of your software and send this data to observability platforms.
* 使用官方提供的检测 instrumentation: https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation
* 如果要自定义增加检测，可参考：https://pkg.go.dev/go.opentelemetry.io/otel, 示例参考：https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/examples
*  needs an export pipeline to send that telemetry to an observability platform. 官方提供的包：https://github.com/open-telemetry/opentelemetry-go/tree/main/exporters
* 如何使用 open-telemetry sdk 来上报 tracer 数据， 下面以 opentelemetry-go sdk: https://github.com/open-telemetry/opentelemetry-go 为列解说： 
* 1 ) 主动请求三方库的的tracer 使用： 创建tracer 对象：
  ``` 
    import (
       "go.opentelemetry.io/otel/trace"
    )

    var demoTracer trace.Tracer = tracing.InitOTEL("mysql", otelExporter, metricsFactory, logger).Tracer("mysql") 
    var once sync.Once

    // InitOTEL initializes OpenTelemetry SDK, 其中 serviceName 是自定义的， exporterType 可以是： otlp or stdout
    // matricFactory 是 import  (
      "github.com/jaegertracing/jaeger/internal/metrics/prometheus"
      "github.com/jaegertracing/jaeger/internal/metrics/prometheus") 
       后
    //  prometheus.New().Namespace(metrics.NSOptions{Name: "hotrod", Tags: nil})
    func InitOTEL(serviceName string, exporterType string, metricsFactory metrics.Factory, logger log.Factory) trace.TracerProvider {
      once.Do(func() {
        otel.SetTextMapPropagator(
          propagation.NewCompositeTextMapPropagator(
            propagation.TraceContext{},
            propagation.Baggage{},
          ))
      })

      exp, err := createOtelExporter(exporterType)
      if err != nil {
        logger.Bg().Fatal("cannot create exporter", zap.String("exporterType", exporterType), zap.Error(err))
      }
      logger.Bg().Debug("using " + exporterType + " trace exporter")

      rpcmetricsObserver := rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)

      res, err := resource.New(
        context.Background(),
        resource.WithSchemaURL(otelsemconv.SchemaURL),
        resource.WithAttributes(otelsemconv.ServiceNameKey.String(serviceName)),
        resource.WithTelemetrySDK(),
        resource.WithHost(),
        resource.WithOSType(),
      )
      if err != nil {
        logger.Bg().Fatal("resource creation failed", zap.Error(err))
      }

      tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(1000*time.Millisecond)),
        sdktrace.WithSpanProcessor(rpcmetricsObserver),
        sdktrace.WithResource(res),
      )
      logger.Bg().Debug("Created OTEL tracer", zap.String("service-name", serviceName))
      return tp
    }

    // withSecure instructs the client to use HTTPS scheme, instead of hotrod's desired default HTTP
    func withSecure() bool {
      return strings.HasPrefix(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"), "https://") ||
        strings.ToLower(os.Getenv("OTEL_EXPORTER_OTLP_INSECURE")) == "false"
    }

    func createOtelExporter(exporterType string) (sdktrace.SpanExporter, error) {
      var exporter sdktrace.SpanExporter
      var err error
      switch exporterType {
      case "jaeger":
        return nil, errors.New("jaeger exporter is no longer supported, please use otlp")
      case "otlp":
        var opts []otlptracehttp.Option
        if !withSecure() {
          opts = []otlptracehttp.Option{otlptracehttp.WithInsecure()}
        }
        exporter, err = otlptrace.New(
          context.Background(),
          otlptracehttp.NewClient(opts...),
        )
      case "stdout":
        exporter, err = stdouttrace.New()
      default:
        return nil, fmt.Errorf("unrecognized exporter type %s", exporterType)
      }
      return exporter, err
    }


  ```

  ```
  使用上面生成的tracer 对象：
  ctx, span := d.tracer.Start(ctx, "SQL SELECT", trace.WithSpanKind(trace.SpanKindClient))
	span.SetAttributes(
		otelsemconv.PeerServiceKey.String("mysql"),
		attribute.
			Key("sql.query").
			String(fmt.Sprintf("SELECT * FROM customer WHERE customer_id=%d", customerID)),
	)
	defer span.End()
  ````



  ```
  3） 服务端的tracer,span的使用：
  http 服务端定义：
  type Server struct {
    hostPort string
    tracer   trace.TracerProvider
    logger   log.Factory
  }

  func NewServer(hostPort string, tracer trace.TracerProvider, logger log.Factory) *Server {
	return &Server{
		hostPort: hostPort,
		tracer:   tracer,
		logger:   logger,
	  }
  }
  上面函数第二个参数： 是由 tracing.InitOTEL("route", otelExporter, metricsFactory, logger) 创建，
   otelExporter Otlp 或者 stdout


  // Run starts the Route server
  func (s *Server) Run() error {
    mux := s.createServeMux()
    s.logger.Bg().Info("Starting", zap.String("address", "http://"+s.hostPort))
    server := &http.Server{
      Addr:              s.hostPort,
      Handler:           mux,
      ReadHeaderTimeout: 3 * time.Second,
    }
    return server.ListenAndServe()
  }
  func (s *Server) createServeMux() http.Handler {
      mux := tracing.NewServeMux(false, s.tracer, s.logger)
      mux.Handle("/route", http.HandlerFunc(s.route))
      mux.Handle("/debug/vars", http.HandlerFunc(movedToFrontend))
      mux.Handle("/metrics", http.HandlerFunc(movedToFrontend))
      return mux
  }

  
  // TracedServeMux is a wrapper around http.ServeMux that instruments handlers for tracing.
  type TracedServeMux struct {
    mux         *http.ServeMux
    copyBaggage bool
    tracer      trace.TracerProvider
    logger      log.Factory
  }

  // Handle implements http.ServeMux#Handle, which is used to register new handler.
  func (tm *TracedServeMux) Handle(pattern string, handler http.Handler) {
    tm.logger.Bg().Debug("registering traced handler", zap.String("endpoint", pattern))

    middleware := otelhttp.NewHandler(
      otelhttp.WithRouteTag(pattern, traceResponseHandler(handler)),
      pattern,
      otelhttp.WithTracerProvider(tm.tracer))

    tm.mux.Handle(pattern, middleware)
  }

  // ServeHTTP implements http.ServeMux#ServeHTTP.
  func (tm *TracedServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tm.mux.ServeHTTP(w, r)
  }

  // Returns a handler that generates a traceresponse header.
  // https://github.com/w3c/trace-context/blob/main/spec/21-http_response_header_format.md
  func traceResponseHandler(handler http.Handler) http.Handler {
    // We use the standard TraceContext propagator, since the formats are identical.
    // But the propagator uses "traceparent" header name, so we inject it into a map
    // `carrier` and then use the result to set the "tracereponse" header.
    var prop propagation.TraceContext
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      carrier := make(map[string]string)
      prop.Inject(r.Context(), propagation.MapCarrier(carrier))
      w.Header().Add("traceresponse", carrier["traceparent"])
      handler.ServeHTTP(w, r)
    })
  }
 ```


 ```
  4) grpc 的监控：
  
 ```
