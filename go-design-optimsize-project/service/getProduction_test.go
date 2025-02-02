package service

import (
	"fmt"
	"testing"
)

type MockRedisGetProduction struct {
}

func (g *MockRedisGetProduction) GetProduction(id int64) {
	fmt.Println("mock get redis get production.")
}

type MockDbGetProduction struct {
}

func (g *MockDbGetProduction) GetProduction(id int64) {
	fmt.Println("mock get mysql get production.")
}

// TestGetProduction;增加测试用例，可以mock依赖的redis,mysql的访问接口.
func TestGetProduction(t *testing.T) {
	tt := NewProductionServerImpl(new(MockRedisGetProduction), new(MockDbGetProduction))
	tt.GetProduction(300)

}
