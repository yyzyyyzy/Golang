package __性能测试

import "testing"

func BenchmarkAdd(b *testing.B) {
	b.Log("Add		开始压力测试")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		Add(4, 5)
	}
}
