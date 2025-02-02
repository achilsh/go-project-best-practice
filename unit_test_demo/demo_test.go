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
