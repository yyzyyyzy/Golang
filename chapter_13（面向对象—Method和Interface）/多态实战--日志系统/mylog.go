package main

// 日志写入接口
type LogWriter interface {
	Write(data interface{}) error
}

// 日志器
type Logger struct {
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// 将data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

// 新建的日志实例
func NewLogger() *Logger {
	return &Logger{}
}
