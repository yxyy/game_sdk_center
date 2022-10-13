package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/login", func(context *gin.Context) {

	})

	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "6666",
		})
	})

	return r
}
