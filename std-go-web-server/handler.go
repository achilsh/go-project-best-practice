package main

import (
	"fmt"
	"std-go-web-server/handler"
)

// InitRouter include router api and register handle
// 注册 api 处理逻辑
func InitRouter() {
	fmt.Println("register api router.")
	handler.GetOrder()
	handler.GetProduct()
}
