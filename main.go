package main

import (
	"fmt"
	"game.sdk.center/lib"
)

func main() {

	fmt.Println(lib.MysqlDb)

	fmt.Println(lib.Mysql("CENTER"))
}
