package main

import (
	"fmt"

	"go-project-layout-demo/internal/order_server"
)

func main() {
	fmt.Println("call user entry begin")
	order_server.OrderServerHandle()
}
