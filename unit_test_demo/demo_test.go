package unit_test_demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDemoCall 单元测试集合
func TestDemoCall(t *testing.T) {
	var xy = struct {
		a int
		y float32
	}{
		100,
		2.0,
	}
	t.Logf("data: %v, %v", xy.a, xy.y)

	testCases := []struct {
		in     int
		expect int
	}{
		{1, 2},
		{2, 4},
		{3, 6},
		{4, 8},
	}

	for _, item := range testCases {
		ret := DemoCall(item.in)
		assert.Equal(t, ret, item.expect)
	}
}

func TestRunDemo(t *testing.T) {
	 testCase := []struct {
		name string 
		in int 
		expect int
	} {
		{"aaa", 1, 1},
		{"bbb", 2, 2},
	}
	for	i := 0; i < len(testCase); i++ {
		t.Run(testCase[i].name, func(t *testing.T) {
			assert.Equal(t, testCase[i].in, testCase[i].expect)
		})
	}
}
