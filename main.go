package main

import (
	"fmt"
	"game.sdk.center/confs"
	"game.sdk.center/logs"
	"game.sdk.center/router"
	"game.sdk.center/tool"
)

func main() {

	// 初始化日志
	logs.InitLogs()

	if err := confs.InitConf(); err != nil {
		panic(fmt.Errorf("配置初始失败： %w", err))
	}

	if err := tool.InitMysql(); err != nil {
		panic(fmt.Errorf("MYSQL初始失败： %w", err))
	}

	if err := tool.InitRedis(); err != nil {
		panic(fmt.Errorf("redis初始失败： %w", err))
	}

	r := router.InitRouter()
	if err := r.Run(":80"); r != nil {
		panic(fmt.Errorf("服务启动失败： %w", err))
	}

}
