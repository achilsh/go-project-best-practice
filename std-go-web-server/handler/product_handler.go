package handler

import (
	"fmt"
	"std-go-web-server/dal"
	"std-go-web-server/service"
)

func GetProduct() {
	// do logic...
	fmt.Println("all product in handler. include logic process and db processor.")
	// call db op
	dal.CallProductionTabOP(10)
	service.GetProduction()
}
