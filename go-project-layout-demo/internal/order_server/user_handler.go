package order_server

import "fmt"

// UserHandler 相对独立;  每个 handle 不会交叉调用;
func UserHandler() {
	fmt.Println("call user handler.")
}
