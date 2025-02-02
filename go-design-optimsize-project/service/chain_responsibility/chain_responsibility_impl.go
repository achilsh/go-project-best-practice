package chain_responsibility

import (
	"fmt"
	"time"
)

type InputChainResponsibility struct {
	A int
}
type OutputChainResponsibility struct {
	B int
}

// HandleChainResponsibility 业务逻辑的统一抽象接口
type HandleChainResponsibility func(*InputChainResponsibility) *OutputChainResponsibility

// CostCalc 计算入参函数运行耗时统计. 返回和入参函数一样格式的 函数对象;
// 在入参函数的基础上，增加额外的处理逻辑。手动构造相同入参的函数格式，这种是装饰器的原理.
// 统一这种格式方便统一存储，统一处理和调用这类函数.
func CostCalc(handle HandleChainResponsibility) HandleChainResponsibility {
	return func(in *InputChainResponsibility) *OutputChainResponsibility {
		beginTime := time.Now()
		defer func() {
			cost := time.Now().Sub(beginTime)
			fmt.Println("cost ms: ", cost.Milliseconds())
		}()
		return handle(in)
	}
}

// CheckParameters 构造一个函数，该函数 在入参函数参数基础上，增加对参数校验.
func CheckParameters(handle HandleChainResponsibility) HandleChainResponsibility {
	return func(in *InputChainResponsibility) *OutputChainResponsibility {
		var out *OutputChainResponsibility = new(OutputChainResponsibility)
		if in == nil {
			return nil
		}
		if in.A <= 100 {
			out.B = -1
		} else {
			out.B = 1
		}
		o := handle(in)
		o.B = out.B
		return o
	}
}

// HandleFunc 是对上面职责链节点的抽象，接收一个处理节点，经过内部加工，返回一个新的同样式的节点.
type HandleFunc func(HandleChainResponsibility) HandleChainResponsibility

// ApplyHandles 在一个基础处理节点handle上,运用handles多个自定义职责链节点对象，不断给 handle上加上新的处理逻辑。
func ApplyHandles(handles []HandleFunc, handle HandleChainResponsibility) HandleChainResponsibility {
	for _, handleItem := range handles {
		handle = handleItem(handle)
	}
	return handle
}

func demoHandle(in *InputChainResponsibility) *OutputChainResponsibility {
	return &OutputChainResponsibility{}
}
func CallChainResponsibility() {
	var in *InputChainResponsibility = &InputChainResponsibility{
		A: 0,
	}
	h := ApplyHandles([]HandleFunc{
		CostCalc, CheckParameters,
	}, demoHandle)
	o := h(in)
	if o != nil {
		fmt.Println("call chain responsibility data: ", o.B)
	}

}
