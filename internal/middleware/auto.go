package middleware

import (
	"fmt"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
)

func Auto(c *gin.Context) {
	if err := tool.MysqlDb.AutoMigrate(
		&system.User{},
		&system.Group{},
		&system.Menu{},
	); err != nil {
		fmt.Println(err)
	}
	c.Next()
}
