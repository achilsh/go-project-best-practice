package unit_test_demo

import (
	"fmt"
	"reflect"
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

func monkeyDemoFunc2(in any) int {
	if v, ok := in.(int); ok {
		return v * 10
	} else {
		return -100
	}
}

// TestMockFuncDemo mock 一般函数的测试场景.
func TestMockFuncDemo(t *testing.T) {
	// create 打桩 object.
	mock := gomonkey.NewPatches()
	// 测试结束后，清理打桩
	defer mock.Reset()

	// 修改mock函数的内部实现; 其中ApplyFunc()的第一个参数是要mock的实际函数名，
	// 第二个参数是mock的成的函数，其样式和实际的函数样式保持一致。
	mock.ApplyFunc(monkeyDemoFunc, func(_ int) int {
		return 100
	})

	// 修改mock函数的内部实现，对原来函数进行mock, 主要是Mock 函数的入参函数
	mock.ApplyFunc(monkeyDemoFunc2, func(v any) int {
		if _, ok := v.(*int); ok {
			reflect.ValueOf(v).Elem().Set(reflect.ValueOf(200))
		}
		return -200
	})
	//

	// test for mock func; 开始测试 上面mock的函数.
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

		convey.Convey("mock func input parameters.", func() {
			in := 0
			ret := monkeyDemoFunc2(&in)
			assert.Equal(t, -200, ret)
			assert.Equal(t, 200, in)
		})
	})
}

// TestMockGlobalValDemo begin to mock global variable table.
// 测试 mock 全局变量测试场景.
func TestMockGlobalValDemo(t *testing.T) {
	convey.Convey("测试mock全局变量", t, func() {
		// create 打桩 object.
		mock := gomonkey.NewPatches()
		// 测试结束后，清理打桩
		defer mock.Reset()
		// mock全局变量值
		mock.ApplyGlobalVar(&GlobalVar, 100)
		//调用使用全局变量的代码
		ret := GetGlobalVar()
		convey.So(ret, convey.ShouldEqual, 1000)
	})
}

// TestMockStructMethod 测试 mock对象的方法场景
func TestMockStructMethod(t *testing.T) {
	convey.Convey("测试 mock struct 方法场景", t, func() {
		var p *Person = &Person{}
		// create 打桩 object.
		mock := gomonkey.NewPatches()
		//
		mockNameValue := "this is mock name for person, getName()"
		mock.ApplyMethod(reflect.TypeOf(p), "GetName", func(_ *Person) string {
			return mockNameValue
		})
		// 测试结束后，清理打桩
		defer mock.Reset()
		//
		ret := p.GetName()
		t.Logf("data: %v\n", ret)
		convey.So(ret, convey.ShouldEqual, mockNameValue)

		var setNameValue string = "this is mock name for Person SetName()"
		mock.ApplyMethod(reflect.TypeOf(p), "SetName", func(pp *Person, _ string) {
			pp.Name = setNameValue
		})
		p.SetName("----")
		ret = p.Name
		convey.So(ret, convey.ShouldEqual, setNameValue)
	})
}

type MockUserEr struct {
}

func (s *MockUserEr) SetName(string) {

}
func (s *MockUserEr) GetName(vType int) string {
	return ""
}

// TestMockInterface 对接口进行mock.主要是某些是依赖接口，测试时就需要mock这些依赖的接口.
func TestMockInterface(t *testing.T) {
	convey.Convey("mock 依赖的接口", t, func() {
		// create 打桩 object.
		mock := gomonkey.NewPatches()
		// 测试结束后，清理打桩
		defer mock.Reset()

		mockInterfaceDemo := &MockUserEr{}
		mock.ApplyMethod(reflect.TypeOf(mockInterfaceDemo), "SetName", func(_ *MockUserEr, _ string) {
			//
		})
		getNameRet := "this is mock interface test."
		mock.ApplyMethod(reflect.TypeOf(mockInterfaceDemo), "GetName", func(_ *MockUserEr, in int) string {
			return getNameRet
		})

		testItem := &UserWrapper{
			OPEr: mockInterfaceDemo,
		}

		ret := testItem.DoCall("sdfadf123")
		convey.So(ret, convey.ShouldEqual, getNameRet)
	})

}
