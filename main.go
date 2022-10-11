package main

import (
	"fmt"
	"game.sdk.center/confs"
	"game.sdk.center/lib"
	"github.com/spf13/viper"
)

func main() {

	if err := confs.InitConf(); err != nil {
		panic(fmt.Errorf("配置初始失败： %w", err))
	}
	fmt.Println(viper.GetStringMap("mysql.master"))
	//MysqlMap := viper.GetStringMap("mysql")
	//fmt.Println(MysqlMap["master"].(map[string]interface{})["host"])

	//fmt.Println(lib.MysqlDb)
	//
	mysql := lib.Mysql("master")
	fmt.Println(mysql)
}
