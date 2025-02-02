package dal

import "fmt"

// CallOrderTabOP 订单数据库的操作封装
func CallOrderTabOP(id int64) {
	fmt.Println("call order table op: query, insert, del, update, op_id: ", id)

}
