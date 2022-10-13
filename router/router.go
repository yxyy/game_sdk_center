package router

import (
	"game.sdk.center/internal/controller/system"
	"game.sdk.center/internal/controller/system/group"
	"game.sdk.center/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Log())

	r.POST("/login", system.Login)

	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "6666",
		})
	})

	r.POST("/system/group", group.Add)

	return r
}
