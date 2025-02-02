package model

// OrderItem 订单详情明细
type OrderItem struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
