package handler

import (
	"fmt"
	"std-go-web-server/dal"
	"std-go-web-server/model"
)

// GetOrder include logic process and db processor.
func GetOrder() {
	//do logic
	fmt.Print("call logic order process. include logic process and db processor \n")
	//从 dal 层获取订单信息和商品信息。
	dal.CallProductionTabOP(100)
	//从 dal 层获取订单信息和商品信息。
	dal.CallOrderTabOP(-10)
	var orderItem = model.OrderItem{
		ID:    123123,
		Name:  "dfiadfa",
		Price: 123.123,
	}

	fmt.Printf("%+v\n", orderItem)

}
