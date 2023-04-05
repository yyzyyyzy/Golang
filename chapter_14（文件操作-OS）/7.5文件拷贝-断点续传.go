package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// 创建一个临时文件记录已经复制的字节数，复制完成之后删除这个临时文件
func ContinueSend(srcFileName, dstFileName, tempFileName string) {
	srcfile, err := os.Open(srcFileName)
	HandleError(err)
	dstfile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	HandleError(err)
	tempfile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_RDWR, 0666)
	HandleError(err)

	defer func() {
		srcfile.Close()
		dstfile.Close()
	}()

	//先读取临时文件中的数据，再seek
	tempfile.Seek(0, io.SeekStart)
	buffer := make([]byte, 100, 100)
	n1, err := tempfile.Read(buffer)
	HandleError(err)
	countStr := string(buffer[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64) //将字符串转换为 10进制 int64位的数值
	HandleError(err)

	//设置读，写的位置
	srcfile.Seek(count, io.SeekStart)
	dstfile.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	n2 := -1
	n3 := -1
	total := int(count)

	for {
		//3.读取数据
		n2, err = srcfile.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕:", total)
			tempfile.Close()
			//一旦复制完，就删除临时文件
			os.Remove(tempFileName)
			break
		}
		//将数据写入到目标文件
		n3, err = dstfile.Write(data[:n2])
		total += n3

		//将赋值的总量存储到临时文件中
		tempfile.Seek(0, io.SeekStart)
		tempfile.WriteString(strconv.Itoa(total))

		fmt.Println("已经复制了", total, "字节数据")

		//假装断电
		if total >= 350000 {
			panic("断电啦")
		}
	}
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

}
func main() {
	srcpath := "F:\\摄影软件\\pic\\丁香湖\\微信\\飞机.jpg"
	dstpath := "F:\\摄影软件\\pic\\丁香湖\\微信\\飞机(1).jpg"
	temppath := "F:\\摄影软件\\pic\\丁香湖\\微信\\飞机temp.txt"
	ContinueSend(srcpath, dstpath, temppath)
}
