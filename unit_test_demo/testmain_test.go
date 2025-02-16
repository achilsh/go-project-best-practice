package unit_test_demo

import (
	"fmt"
	"testing"
)

var demoLocalDemoValue int

// function: 在多个测试用例中，统一的初始化，统一的资源回收
func initDemo() {
	demoLocalDemoValue = 100
	fmt.Println("set demoLocalDemoValue init value is 100")
}
func destoryDemo() {
	demoLocalDemoValue = -1
	fmt.Println("set demoLocalDemoValue invalid value is  -1")
}

// TestMain； 主要介绍 TestMain 的使用。
func TestMain(m *testing.M) {
	fmt.Println("call before  all test cases...")
	initDemo()

	m.Run()

	destoryDemo()
	fmt.Println("call after all test cases....")
}

func TestCallOne(t *testing.T) {
	demoLocalDemoValue = 20
	t.Logf("call one test demo: %v", demoLocalDemoValue)
}
func TestCallTwo(t *testing.T) {
	demoLocalDemoValue = 30
	t.Logf("call two test demo: %v", demoLocalDemoValue)
}
