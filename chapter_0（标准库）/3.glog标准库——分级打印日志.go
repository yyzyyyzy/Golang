package main

import (
	"flag"
	"github.com/golang/glog"
)

func main01() {

	//用户传递的命令行参数解析为对应的变量
	flag.Parse()

	//退出时调用，确保日志写入文件中
	defer glog.Flush()

	glog.Info("This is info message")
	glog.Infof("This is info message: %v", 12345)
	glog.InfoDepth(1, "This is info message", 12345)

	glog.Warning("This is warning message")
	glog.Warningf("This is warning message: %v", 12345)
	glog.WarningDepth(1, "This is warning message", 12345)

	glog.Error("This is error message")
	glog.Errorf("This is error message: %v", 12345)
	glog.ErrorDepth(1, "This is error message", 12345)

	glog.Fatal("This is fatal message")
	glog.Fatalf("This is fatal message: %v", 12345)
	glog.FatalDepth(1, "This is fatal message", 12345)
}

func main() {
	flag.Parse()
	//确保glog的操作写入I/O
	defer glog.Flush()

	glog.V(3).Info("LEVEL 3 message") // 使用日志级别 3
	glog.V(4).Info("LEVEL 4 message") // 使用日志级别 4
	glog.V(5).Info("LEVEL 5 message") // 使用日志级别 5
	glog.V(8).Info("LEVEL 8 message") // 使用日志级别 8

	//go run xxx -log_dir=log -alsologtostderr

	//mkdir -p log		在当前目录下创建log文件
	//-log_dir=log		在log目录下会生成相应的日志文件
	//-logtostderr		打印在标准输出中
	//-alsologtostderr  同时打印在 log/ 目录下和标准输出中
}
