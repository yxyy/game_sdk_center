package middleware

import (
	"fmt"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
)

func Auto(c *gin.Context) {
	err := tool.MysqlDb.AutoMigrate(&system.Group{})
	if err != nil {
		fmt.Println(err)
	}
	c.Next()
}
