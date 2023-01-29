package router

import (
	"fmt"
	"game.sdk.center/internal/controller/system"
	"game.sdk.center/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Log)
	// r.Use(middleware.Auto)

	r.POST("/system/login", system.Login)
	r.POST("/system/logout", system.Logout)

	r.Use(middleware.Auth)
	r.Use(middleware.Menu)

	r.GET("/home", func(c *gin.Context) {

		fmt.Println(666666)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "6666",
		})
	})

	// 加载系统路由
	InitSystemRouter(r)

	return r
}
