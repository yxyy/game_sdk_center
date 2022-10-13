package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println(6666)
	}
}
