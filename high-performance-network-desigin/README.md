* 网络服务设计模式： 
* 1） 采用事件异步设计，reactor 模式， 比如 主从多 reactor模式（多个协程监听多个 epoll 池，每个 epoll 池放一部分需要监听的文件描述符（fd）。主 Reactor 监听连接事件，从 Reactor 监听读写等网络事件）。
* 2） 从协程池中的协程划分两种协程：a. 负责连接上的数据读写； b.负责处理业务逻辑；设计两种独立的协程解决业务不会被读写给阻塞（如果读写io被阻塞）
* 3） golang 实现的高性能网络参考： bytedance 的 netpoll: https://github.com/cloudwego/netpoll 

* 4） 对于golang 标准库的 net, 使用netpoll 适用于 linux系统的，接入网关，或者代理；可解决高并发，海量设备接入的场景（类似 nginx 能力）。