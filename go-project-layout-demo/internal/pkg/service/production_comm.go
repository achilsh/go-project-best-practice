package service

import (
	"fmt"
	"go-project-layout-demo/internal/pkg/model"
	"go-project-layout-demo/pkg/utils"
)

func CallCommonONProduction() {
	fmt.Println("call common production.")
}
func CallOpOnRedisAndDB() {
	var item = model.ProductionDBModel{
		ID: 123213,
	}

	fmt.Printf("call call redis and db op: %+v\n", item)

	utils.MethDemo()
}
