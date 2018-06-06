package main
import (
	"os"
	"github.com/sirupsen/logrus"
)
// 你可以创建很多instance
var log = logrus.New()
func main() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件失败")
}