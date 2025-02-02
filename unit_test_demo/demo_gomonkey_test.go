package unit_test_demo

import (
	"fmt"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

// 下面可以mock go中数据，比如：
// 1. 函数
// 2. 全局变量
// 3. 接口
// 4. struct 的方法
////////////////////////////////////////////////
////////////////////////////////////////////////
// 下面将结合convey，使用 gomonkey 来打桩 mock上面场景.
// notice: need add option: -gcflags=all=-l for go test -gcflags=all=-l

// mock func
func monkeyDemoFunc(x int) int {
	return x * 2
}

// wrappermonkeyFunc indirect use mock func
func wrappermonkeyFunc(f float32) string {
	x := monkeyDemoFunc(int(f))
	return fmt.Sprintf("dd:%v", x)
}

func TestMockFuncDemo(t *testing.T) {
	// create 打桩 object.
	mock := gomonkey.NewPatches()
	// 测试结束后，清理打桩
	defer mock.Reset()

	// begin to 打桩
	mock.ApplyFunc(monkeyDemoFunc, func(_ int) int {
		return 100
	})

	convey.Convey("mock monkeyDemoFunc", t, func() {
		convey.Convey("mock func directly", func() {

			// 调用实际的函数，函数底层内部是采用mock的逻辑，而非函数的实际内部逻辑。
			ret := monkeyDemoFunc(1)
			assert.Equal(t, 100, ret)
			// run testing.
			convey.So(ret, convey.ShouldEqual, 100)
		})
		//
		convey.Convey("mock func indirectly", func() {
			// 调用实际的函数，函数底层内部是采用mock的逻辑，而非函数的实际内部逻辑。
			ret := wrappermonkeyFunc(1.1)
			assert.Equal(t, "dd:100", ret)
			// run testing.
			convey.So(ret, convey.ShouldEqual, "dd:100")
		})

	})
}
