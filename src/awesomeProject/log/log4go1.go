package main

import (
	"github.com/alecthomas/log4go"
	"time"
)

func main() {
	log4go.LoadConfiguration("./log/log4go.xml")
	log4go.Debug("这是DEBUG日志")
	log4go.Info("这是INFO日志")
	panic("")
	time.Sleep(time.Second)
}
