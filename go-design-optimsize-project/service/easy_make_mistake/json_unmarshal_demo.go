package easy_make_mistake

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJsonDemo 在解析 JSON 数据时，如果使用 map[string]any 类型来做反序列化，我们就需要用 float64 去解析 int。
func UnmarshalJsonDemo() {
	var originStr string = `{"a":"afadf", "b":123}` //b value type is int.
	var dstMap map[string]any
	err := json.Unmarshal([]byte(originStr), &dstMap)
	if err != nil {
		fmt.Println("unmarshal fail, err: ", err)
	} else {
		fmt.Printf("unmarshal data: %+v\n", dstMap)
		for k, v := range dstMap {
			fmt.Println(k, v)
		}
		// json 的b字段对应的是float64,而不是int
		age := dstMap["b"].(float64) // interface {} is float64, not int
		fmt.Println(age)
	}

}
