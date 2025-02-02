package redis

import "fmt"

type GetProductionEr interface {
	GetProduction(id int64)
}

type GetProductionImpl struct {
}

func (g *GetProductionImpl) GetProduction(id int64) {
	fmt.Println("call get production on redis, id: ", id)
}
