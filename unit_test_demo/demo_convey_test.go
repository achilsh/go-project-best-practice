package unit_test_demo

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// TestDemoCallConvey 递归调用 Convey; 分组管理测试用例.
func TestDemoCallConvey(t *testing.T) {
	//最外层需要传入测试标准库 t 对象.
	convey.Convey("测试调用: DemoCall", t, func() {
		// 内部就不需要传入 测试标准库 t 对象.
		convey.Convey("正常情况: ", func() {
			result := DemoCall(1)
			convey.So(result, convey.ShouldEqual, 2)
		})
		//
		convey.Convey("异常情况: ", func() {
			result := DemoCall(2)
			convey.So(result, convey.ShouldNotEqual, 3)
		})
		//
		convey.Convey("边界测试场景: ", func() {
			convey.Convey("小于场景: ", func() {
				ret := DemoCall(4)
				convey.So(ret, convey.ShouldBeLessThan, 9)

			})
			convey.Convey("大于场景: ", func() {
				ret := DemoCall(4)
				convey.So(ret, convey.ShouldBeGreaterThan, 7)

			})

		})
	})
}
