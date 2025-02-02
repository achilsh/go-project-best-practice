package order_server

import "fmt"

// ProductionServerHandler 相对独立; 每个 handle 不会交叉调用
func ProductionServerHandler() {
	fmt.Println("call production server handle.")
}
