package main

import "github.com/alecthomas/log4go"

func main() {
	defer log4go.Close()
	log4go.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())           // DEBUG级别+打印到控制台
	log4go.AddFilter("file", log4go.INFO, log4go.NewFileLogWriter("test.log", true)) // INFO级别+输出到文件，并开启rotate
	log4go.Debug("这是DEBUG日志")                                                        // 输出测试
	log4go.Info("这是INFO日志")                                                          // 输出测试
	log4go.Info("这是INFO日志")                                                          // 输出测试
}
