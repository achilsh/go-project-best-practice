package service

import (
	"fmt"
	"std-go-web-server/dal/mysql"
	"std-go-web-server/dal/redis"
)

// GetProduction 获取商品详情信息
func GetProduction() {
	fmt.Println("call production details.")
	mysql.GetProductionInDB()
	redis.SetProductionToRedis()
}
