package main

import (
	"fmt"
	"game.sdk.center/confs"
	"game.sdk.center/lib"
	"game.sdk.center/logs"
	"game.sdk.center/router"
)

func main() {

	// 初始化日志
	logs.InitLogs()

	if err := confs.InitConf(); err != nil {
		panic(fmt.Errorf("配置初始失败： %w", err))
	}

	if err := lib.InitMysql(); err != nil {
		panic(fmt.Errorf("配置初始失败： %w", err))
	}

	r := router.InitRouter()
	if err := r.Run(":80"); r != nil {
		panic(fmt.Errorf("服务启动失败： %w", err))
	}

}
