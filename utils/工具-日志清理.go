package main

import (
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

/*
 * clearLog
 *
 目标：
 设计一个package，可以被任何golang程序使用，一旦被main start，它会间隔性的扫描指定目录里的*.log文件，然后依据一些“清理规则”清理他们

 可以写死的清理规则：
 1、文件夹所有log文件加起来超过100MB了吗，如果超过，尝试从最早时间的日志文件开始删起，一直删到小于100MB
 2、删除10天前的所有log文件，顺序是尝试从最早时间的日志文件开始。
 3、删不掉就算了skip。
 *
*/

const (
	SUM_SIZE_THRESHOLD = 104857600 // log总和(100MB)
	SECOND_TEN_DAY     = 864000    // 时间差常量
	LOG_ACCUM_NUM      = 10        // 堆积文件个数
	LOG_LEVEL_INFO     = "INFO"
	LOG_LEVEL_WARNING  = "WARNING"
	LOG_LEVEL_ERROR    = "ERROR"
	LOG_LEVEL_FATAL    = "FATAL"
)

type LogClear struct {
	path        string
	timeinter   int
	file_prefix string
}

type InitOption struct {
	Path       string
	TimeInter  int
	FilePrefix string
}
type LogInfo struct {
	logtime int64
	Path    string
	Size    int64
}
type LogInfos []LogInfo

func (u LogInfos) Len() int {
	return len(u)
}
func (u LogInfos) Less(i, j int) bool {
	return u[i].logtime < u[j].logtime
}
func (u LogInfos) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func NewLogClear(option InitOption) *LogClear {
	lc := new(LogClear)
	lc.path = option.Path
	lc.timeinter = option.TimeInter
	lc.file_prefix = option.FilePrefix
	/*
	  防止因主程序不断重启，出现大量错误日志，而下面的gorotine又来不及清理就挂掉了，
	  所以这里初始化的时候就清理下日志，协助下gorotine
	*/
	lc.initClearLog()
	/*
	  gorotine 定时清理日志
	*/
	go lc.clearLogRegular()
	return lc
}

func (lc *LogClear) initClearLog() {
	var sizeSum int64
	var location int = -1
	var log LogInfos
	// 建立一个数组，存放4个级别的log文件数量
	// INFO->WARNING->ERROR->FATAL
	logCount := [4]int{0, 0, 0, 0}

	//STEP1:搜集所有
	//获取指定目录下的log文件，同时记录每一个log文件的信息
	//存放在一个LogInfos型的切片中，记录是信息有（名称，大小和创建时间）
	files, err := ioutil.ReadDir(lc.path)
	if err != nil {
		glog.Errorln(err)
		return
	}
	if len(files) == 0 {
		glog.Infoln(`this dir is empty ... `)
		return
	}
	for _, file := range files {
		ok_suffix := strings.HasSuffix(file.Name(), ".log")
		ok_prefix := strings.HasPrefix(file.Name(), lc.file_prefix)
		if ok_suffix && ok_prefix {
			modTime := file.ModTime().Unix()
			size := file.Size()
			name := file.Name()
			// 判断日志文件类型
			if strings.Contains(name, LOG_LEVEL_INFO) {
				logCount[0]++
			} else if strings.Contains(name, LOG_LEVEL_WARNING) {
				logCount[1]++
			} else if strings.Contains(name, LOG_LEVEL_ERROR) {
				logCount[2]++
			} else if strings.Contains(name, LOG_LEVEL_FATAL) {
				logCount[3]++
			}
			var tmpPathSize LogInfo
			tmpPathSize.Path = string(os.PathSeparator) + name
			tmpPathSize.Size = size
			tmpPathSize.logtime = modTime
			log = append(log, tmpPathSize)
			sizeSum = sizeSum + size
		}
	}
	//STEP2:排序&找出清理点
	//对切片按照创建时间进行排序，同时按照清理策略定位到需要清理的日志
	sort.Sort(log)
	nowtime := time.Now().Unix()
	// STEP2.1：如果堆积的文件个数超过30个,先清理超出的，按照生成日志的时间顺序杀
	var accuNumber = len(log)
	if accuNumber > LOG_ACCUM_NUM {
		glog.V(2).Infoln(`Find accmulated file number greater than 1024, Need to Clear ... `)
		removeIndex := 0
		for i := 0; i < accuNumber-LOG_ACCUM_NUM+3; i++ {
			// 清理时检查文件名
			fPath := log[removeIndex].Path
			// 避开各类型最后一个file
			if !avoidLastFile(fPath, logCount) {
				removeIndex++
				continue
			}
			// 更新总文件大小
			sizeSum -= log[removeIndex].Size
			if err := os.Remove(lc.path + log[removeIndex].Path); err != nil {
				glog.Errorln(err)
			}
			log = append(log[:removeIndex], log[removeIndex+1:]...)
		}
	}
	glog.V(2).Infoln(`Check File Size and Logtime ... `)
	// STEP2.2:此时文件个数一定在30之下了
	for key, value := range log { // 定位
		timeDistance := nowtime - value.logtime
		if sizeSum > SUM_SIZE_THRESHOLD {
			// 避免最后一个文件被删除
			if avoidLastFile(value.Path, logCount) {
				sizeSum = sizeSum - value.Size
				location = key
			}
			continue
		}
		if timeDistance > SECOND_TEN_DAY {
			if avoidLastFile(value.Path, logCount) {
				sizeSum = sizeSum - value.Size
				location = key
			}
			continue
		}
		if sizeSum <= SUM_SIZE_THRESHOLD && timeDistance <= SECOND_TEN_DAY {
			break
		}
	}
	//STEP3:清理
	//清理掉规则内的log文件
	removeIndex2 := 0
	for i := 0; i <= location; i++ {
		f := log[removeIndex2]
		// 判断是否最后一个
		if avoidLastFile(f.Path, logCount) {
			glog.V(2).Infoln(`Clear File:`, lc.path+log[removeIndex2].Path)
			err := os.Remove(lc.path + log[removeIndex2].Path)
			if err != nil {
				glog.Errorln(err)
			}
		}
		removeIndex2++
	}
}

func (lc *LogClear) clearLogRegular() {
	for {
		lc.initClearLog()
		//定时执行日志清理，间隔由main指定
		time.Sleep(time.Duration(lc.timeinter) * time.Second)
	}
}

func avoidLastFile(fPath string, logCount [4]int) bool {
	if strings.Contains(fPath, LOG_LEVEL_INFO) {
		if logCount[0] <= 1 {
			return false
		} else {
			logCount[0]--
		}

	}
	if strings.Contains(fPath, LOG_LEVEL_WARNING) {
		if logCount[1] <= 1 {
			return false
		} else {
			logCount[1]--
		}
	}
	if strings.Contains(fPath, LOG_LEVEL_ERROR) {
		if logCount[2] <= 1 {
			return false
		} else {
			logCount[2]--
		}
	}
	if strings.Contains(fPath, LOG_LEVEL_FATAL) {
		if logCount[3] <= 1 {
			return false
		} else {
			logCount[3]--
		}
	}
	return true
}
