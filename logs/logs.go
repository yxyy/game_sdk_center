package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogs() {

	// 设置格式
	log.SetFormatter(&log.TextFormatter{})
	// 	设置输出位置
	log.SetOutput(os.Stdout)
	//  设置日志等级
	log.SetLevel(log.WarnLevel)

}
