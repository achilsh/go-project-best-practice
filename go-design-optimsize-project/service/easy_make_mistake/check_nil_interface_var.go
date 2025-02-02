package easy_make_mistake

import "fmt"

// ErrorDemo 是自定义错误信息的结构体
type ErrorDemo struct {
	Code    int
	Message string
}

// Error 定义 Error()接口的实现.
func (g *ErrorDemo) Error() string {
	if g == nil {
		return "is nil"
	}
	return fmt.Sprintf("code: %v, message: %s", g.Code, g.Message)
}

// getError()返回一个error 接口，返回值不管什么场景返回都不会为nil
// 因为值如果为nil,那么值的类型和值都是nil，而当前返回值的类型是ErrorDemo 指针，值有可能存在nil。所以返回值怎么也不会为nil.
func getError(typeVal int) error {
	var ret *ErrorDemo = nil
	//var ret error = nil //this declare is ok.
	if typeVal <= 1 {
		ret = &ErrorDemo{
			Code:    1,
			Message: "less than 1",
		}
	}
	return ret
}

// getErrorInRight 返回一个实际的接口类型或者是nil对象。
func getErrorInRight(typeVal int) error {
	if typeVal <= 1 {
		return &ErrorDemo{
			Code:    100,
			Message: "code is less than 1, is right",
		}
	}
	return nil
}

// CallCheckNilInterfaceVar 通过传入不同值判断返回的接口对象是否为nil
func CallCheckNilInterfaceVar() {
	var demoValue = 1
	e := getError(demoValue) //返回为一个不为nil的对象.
	if e != nil {
		fmt.Printf("type 1, %v\n", e)
	}

	e = getErrorInRight(demoValue)
	if e != nil {
		fmt.Printf("right type 1: %v\n", e)
	}

	e = getError(2) // 返回一个不为Nil的对象，因为类型不为nil, 值为nil。
	if e != nil {
		fmt.Printf("type 2: %v\n", e)
	} else {
		fmt.Printf("type 2 is nil\n")
	}

	e = getErrorInRight(2) //返回一个为nil的对象，因为类型和值都是Nil
	if e != nil {
		fmt.Printf("right type 2: %v\n", e)
	} else {
		fmt.Printf("right type 2 is nil, %v\n", e)
	}
}
