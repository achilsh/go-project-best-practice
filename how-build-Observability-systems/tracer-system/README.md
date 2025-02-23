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
* [过期使用： opentrace, demo： https://github.com/yurishkuro/opentracing-tutorial/tree/master/go]
* #

* span 跨度 介绍：表示一个工作单元或操作。
A span contains name, time-related data, structured log messages, and other metadata (that is, Attributes) to provide information about the operation it tracks. span内有的东西：
```
  Name
  Parent span ID (empty for root spans)
  Start and End Timestamps
  Span Context
  Attributes
  Span Events
  Span Links
  Span Status
```

```
在tracer中 根 span , 其中 Span Context contains the Trace ID, it is used when creating Span Links.
{
  "name": "hello",
  "parent_id": null,
  "context": {
    "trace_id": "5b8aa5a2d2c872e8321cf37308d69df2",
    "span_id": "051581bf3cb55c13"
  },
  "start_time": "2022-04-29T18:52:58.114201Z",
  "end_time": "2022-04-29T18:52:58.114687Z",
  "attributes": {
    "http.route": "some_route1"
  },
  "events": [
    {
      "name": "Guten Tag!",
      "timestamp": "2022-04-29T18:52:58.114561Z",
      "attributes": {
        "event_attributes": 1
      }
    }
  ]
}


在tracer中子span:

{
  "name": "hello-greetings",
  "context": {
    "trace_id": "5b8aa5a2d2c872e8321cf37308d69df2",
    "span_id": "5fb397be34d26b51"
  },
  "parent_id": "051581bf3cb55c13",
  "start_time": "2022-04-29T18:52:58.114304Z",
  "end_time": "2022-04-29T22:52:58.114561Z",
  "attributes": {
    "http.route": "some_route2"
  },
  "events": [
    {
      "name": "hey there!",
      "timestamp": "2022-04-29T18:52:58.114561Z",
      "attributes": {
        "event_attributes": 1
      }
    },
    {
      "name": "bye now!",
      "timestamp": "2022-04-29T18:52:58.114585Z",
      "attributes": {
        "event_attributes": 1
      }
    }
  ]
}

上面 多个 span 的 tracer_id 相同,不同的是 span_id,通过 praent_id 关联其他的 span_id 来建立 父子链路。
```
* 上面 span 的 attributes 属性字段介绍：
* Span attributes are metadata attached to a span
*  优先在创建跨度时添加属性，以使这些属性可用于 SDK 采样。如果必须在创建跨度后添加值，请使用该值更新跨度。
* span 属性定义中的元素：
```
  Keys must be non-null string values
  Values must be a non-null string, boolean, floating point value ,integer, or an array of these values
```
* 比如属性例子：
```
    Key	                      Value
    http.request.method	      "GET"
    network.protocol.version	"1.1"
    url.path	                "/webshop/articles/4"
    url.query	                  "?s=1"
    server.address            	"example.com"
    server.port              	8080
    url.scheme	               "https"
    http.route	               "/webshop/articles/:article_id"
    http.response.status_code	  200
    client.address	           "192.0.2.4"
    client.socket.address	      "192.0.2.5" (the client goes through a proxy)
    user_agent.original      	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0"
```
* span link 是 span 内容之一；用于多个tracer 间的连接，比如：可以将第一个tracer 记录中的最后一个跨度链接到第二个跟踪记录中的第一个跨度.、
* span event： 作为 structured log message (or annotation) on a Span； 通常用于表示跨度持续时间内有意义的单个时间点。

  如何区分是使用 span attribute 还是 span event? 考虑  consider whether a specific timestamp is meaningful.

* span kind: 
  ```
    1. SERVER indicates that the span covers server-side handling of a remote request while the client awaits a response.

    2. CLIENT indicates that the span describes a request to a remote service where the client awaits a response. When the context of a CLIENT span is propagated, CLIENT span usually becomes a parent of a remote SERVER span.

    3. PRODUCER indicates that the span describes the initiation or scheduling of a local or remote operation. This initiating span often ends before the correlated CONSUMER span, possibly even before the CONSUMER span starts.

    4. In messaging scenarios with batching, tracing individual messages requires a new PRODUCER span per message to be created.

    5. CONSUMER indicates that the span represents the processing of an operation initiated by a producer, where the producer does not wait for the outcome.

    6. INTERNAL Default value. Indicates that the span represents an internal operation within an application, as opposed to an operations with remote parents or children.

  ```
  
* # 
Baggage： 最好用于包含通常仅在请求开始时才可用的信息，并在下游进一步传播。
他是独立的 kv 存储，和 spans属性的没有绝对的关系；可以手动从 baggage 读取数据，写到 span的属性中。

* #

信号：Metrics, logs, traces, and baggage are examples of signals.


* tracer provider: => is a factory for Tracers.
* tracer exporter: =>  send traces to a consumer. This consumer can be standard output for debugging and development-time, the OpenTelemetry Collector.

#
* 上下文传播：context propagation， 作用：可以使任何地方产生的信号相互关联。 包括 context 和 propagation.下面分别介绍：

1) context: 上下文是一个对象，包含用于发送和接收服务或执行单元的信息，以便将一个信号与另一个信号相关联。
2) propagation： 传播是在服务和进程之间移动上下文的机制。它对上下文对象进行序列化或反序列化，并提供相关信息以便从一个服务传播到另一个服务。
通常是由检测库提供，对上层用户透明。默认传播器正在使用W3C(https://www.w3.org/TR/trace-context/) 跟踪上下文规范中指定的标头。
