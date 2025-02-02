package service

import (
	"fmt"
	"go-design-optimsize-project/dal/mysql"
	"go-design-optimsize-project/dal/redis"
)

// GetProduction : 第一种调用方式
func GetProduction() {
	fmt.Println("call production message.")
	//
	new(redis.GetProductionImpl).GetProduction(123)
	new(mysql.GetProductionImpl).GetProduction(100)
}

// ProductionServerEr 使用一个接口，
type ProductionServerEr interface {
	GetProduction(id int64)
}

// ProductionServerImpl use impl for interface,
// 内部不直接依赖于实际的实现。只依赖于接口.这样service层就不直接依赖于dal的实现。
// 而是依赖于dal层的接口。该接口成员变量，通过NewProductionServerImpl()注入.
// 第一步： 把对实际的依赖变成对接口的依赖.
type ProductionServerImpl struct {
	redisOp redis.GetProductionEr //dal 接口
	dbOp    mysql.GetProductionEr //dal 接口
}

func (g *ProductionServerImpl) GetProduction(id int64) {
	if g.redisOp != nil {
		g.redisOp.GetProduction(id)
	}
	if g.dbOp != nil {
		g.dbOp.GetProduction(id)
	}
	fmt.Println("call production server impl logic")

}

// NewProductionServerImpl 创建一个接口，接口的实现是通过 NewProductService 方法注入，符合依赖反转原则.
func NewProductionServerImpl(r redis.GetProductionEr, d mysql.GetProductionEr) ProductionServerEr {
	return &ProductionServerImpl{
		redisOp: r,
		dbOp:    d,
	}
}

// GetProduction2 封装缓存和db的访问.
func GetProduction2(id int64) {
	NewProductionServerImpl(new(redis.GetProductionImpl), new(mysql.GetProductionImpl)).GetProduction(400)
}

// ProductionServerDBImpl 只定义实现接口的db实现
type ProductionServerDBImpl struct {
	dbOp mysql.GetProductionEr //dal 接口
}

// GetProduction  只用db实现方式来实现接口，用于只有db访问的场景.
func (g *ProductionServerDBImpl) GetProduction(id int64) {
	fmt.Println("db implement production server on db, id: ", id)
}
func NewProductionServerDBImpl(dbEr mysql.GetProductionEr) ProductionServerEr {
	return &ProductionServerDBImpl{
		dbOp: dbEr,
	}

}

// GetProduction3 根据传入参数来确定如何使用缓存和db操作. 采用策略模式来，1）定义策略接口. 2) 实现策略的具体类型.
func GetProduction3(id int64, sense int) {
	switch sense {
	case 1:
		// 根据类型来选择对象
		NewProductionServerDBImpl(new(mysql.GetProductionImpl)).GetProduction(id)
	case 2:
		// 根据对象类型选择
		NewProductionServerImpl(new(redis.GetProductionImpl), new(mysql.GetProductionImpl)).GetProduction(id)
	}
}
