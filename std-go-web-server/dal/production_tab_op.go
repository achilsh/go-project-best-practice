package dal

import "fmt"

// CallProductionTabOP 商品数据库的操作封装
func CallProductionTabOP(id int64) {
	fmt.Println("call production op: query, insert. del, update ops, index: ", id)
}
