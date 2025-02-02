package order_server

import (
	"fmt"
	"go-project-layout-demo/internal/pkg/service"
)

// OrderServerHandle 相对独立; 每个 handle 不会交叉调用
func OrderServerHandle() {
	fmt.Println("call order server handler, include logic process.")
	service.CallOpOnRedisAndDB()
}
