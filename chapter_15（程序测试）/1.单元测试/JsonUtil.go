package __单元测试

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Person struct {
	Id     int    `json:"id,string"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func EncodeStructJson(p *Person, filename string) bool {
	dstfile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer dstfile.Close()

	encoder := json.NewEncoder(dstfile)
	err := encoder.Encode(p)
	if err != nil {
		fmt.Println("编码失败")
		return false
	} else {
		fmt.Println("编码成功")
		return true
	}

}

func DecodeJsonStruct(filename string) (*Person, error) {
	srcfile, _ := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer srcfile.Close()

	personPtr := new(Person)
	decoder := json.NewDecoder(srcfile)
	err := decoder.Decode(personPtr)
	if err != nil {
		fmt.Println("解码失败")
		return nil, errors.New("解码失败")
	} else {
		fmt.Println("解码成功")
		return personPtr, nil
	}
}
