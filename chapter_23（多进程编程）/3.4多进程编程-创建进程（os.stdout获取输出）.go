package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
)

func GbkToUtf8_1(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func main() {
	cmd := exec.Command("ping", "你麻痹")

	var stdout, stderr bytes.Buffer //创建二进制输入，stdout为输出，stderr为异常

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	res1, err := GbkToUtf8_1(stdout.Bytes())
	res2, err := GbkToUtf8_1(stderr.Bytes())

	outStr, errStr := string(res1), string(res2)
	fmt.Println(outStr)
	fmt.Println(errStr)

	if err != nil {
		fmt.Println(err)
	}
}
