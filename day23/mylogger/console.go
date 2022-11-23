package mylogger

import (
	"fmt"
	"time"
)

//console	控制台
//往终端写日志相关内容

// Logger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

// 写一个方法来进行比较
func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
	//一个enable方法，参数是LogLevel类型的logLeve

}

// 定义方法
/*	方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
 		函数体
}

	接收者变量：接收者中的参数变量名在命名时，
		官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
		例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
	接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
	方法名、参数列表、返回参数：具体格式与函数定义相同。
*/

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		str := now.Format("2006-01-02||15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s]【%s\t[%s\t%s\t%d] %s\n", str, getLogString(lv), funcName, fileName, lineNo, msg)
	}
}
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)
}
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
