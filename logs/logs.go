package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogs() {

	// 设置格式
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               false,
		DisableColors:             true,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          true,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})

	fp, err := os.OpenFile("./storage/log/request.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer fp.Close()
	// 	设置输出位置
	log.SetOutput(fp)
	//  设置日志等级
	log.SetLevel(log.InfoLevel)

}
