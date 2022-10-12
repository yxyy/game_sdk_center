package main

import (
	"fmt"
	"game.sdk.center/confs"
	"game.sdk.center/lib"
)

func main() {

	if err := confs.InitConf(); err != nil {
		panic(fmt.Errorf("配置初始失败： %w", err))
	}

	if err := lib.InitMysql(); err != nil {
		panic(fmt.Errorf("配置初始失败： %w", err))
	}

}
