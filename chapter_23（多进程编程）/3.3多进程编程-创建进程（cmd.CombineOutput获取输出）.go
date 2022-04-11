package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
)

// UTF-8 转 GBK
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func main() {
	cmd := exec.Command("tasklist")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	} else {
		result, err := GbkToUtf8(out)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(result))
		}
	}

}
