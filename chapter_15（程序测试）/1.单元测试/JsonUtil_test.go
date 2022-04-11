package __单元测试

import "testing"

func TestEncodeStructJson(t *testing.T) {
	path := "E:\\golandlearning\\chapter_15（程序测试）\\编码存放文件1.json"
	p1 := Person{Id: 1, Name: "LZK", Gender: "男"}
	ok := EncodeStructJson(&p1, path)
	if ok {
		Pstruct, _ := DecodeJsonStruct(path)
		if Pstruct.Name == p1.Name && Pstruct.Gender == p1.Gender {
			t.Log("EncodeStructJson函数测试成功")
		} else {
			t.Fatal("EncodeStructJson函数测试失败，编码前后数据不符")
		}
	} else {
		t.Fatal("编码失败")
	}
}

func TestDecodeJsonStruct(t *testing.T) {
	path := "E:\\golandlearning\\chapter_15（程序测试）\\编码存放文件2.json"
	p1 := Person{Id: 1, Name: "LZK", Gender: "男"}
	ok := EncodeStructJson(&p1, path)
	if ok {
		Pstruct, _ := DecodeJsonStruct(path)
		if Pstruct.Name == p1.Name && Pstruct.Gender == p1.Gender {
			t.Log("EncodeStructJson函数测试成功")
		} else {
			t.Fatal("EncodeStructJson函数测试失败，编码前后数据不符")
		}
	} else {
		t.Fatal("编码失败")
	}
}
