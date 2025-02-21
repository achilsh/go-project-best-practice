* jaeger 的作用：
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