package unit_test_demo

import "testing"

type Item struct {
	id  int
	val [2048]byte
}

// BenchmarkSliceStructByFor-8   	 1856461	       589.5 ns/op
// BenchmarkSliceStructByFor-8   	 2002306	       668.4 ns/op
func BenchmarkSliceStructByFor(b *testing.B) {
	var item [2048]Item
	for i := 0; i < b.N; i++ {
		var temp int
		for j := 0; j < len(item); j++ {
			temp = item[j].id
		}
		_ = temp

	}
}

// 用for range下标遍历[]struct{}
// BenchmarkSliceStructByRangeIndex-8   	 1557676	       650.7 ns/op
func BenchmarkSliceStructByRangeIndex(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// 用for range值遍历[]struct{}的元素
func BenchmarkSliceStructByRangeValue(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		// for range value in slice, which is expensive, as happen copy.
		for _, item := range items { //each copy one item in items to item
			tmp = item.id
		}
		_ = tmp
	}
}
