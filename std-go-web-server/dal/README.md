* 具体 数据结构数据库, 具体数据结构redis实际访问统一封装. eg: order, production.
* 会根据组件的不同新建不同的子目录(eg: mysql, redis)
* response to op table of order table, production table.
* then other server can call those interfaces.