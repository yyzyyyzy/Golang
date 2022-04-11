package __性能测试

import "testing"

func BenchmarkEncodeStructJson(b *testing.B) {
	path := "E:\\golandlearning\\chapter_15（程序测试）\\编码存放文件1.json"
	p1 := Person{Id: 1, Name: "LZK", Gender: "男"}
	b.Log("EncodeStructJson开始压力测试")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EncodeStructJson(&p1, path)
	}
}

func BenchmarkDecodeJsonStruct(b *testing.B) {
	path := "E:\\golandlearning\\chapter_15（程序测试）\\编码存放文件1.json"
	b.Log("DecodeJsonStruct开始压力测试")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		DecodeJsonStruct(path)
	}
}
