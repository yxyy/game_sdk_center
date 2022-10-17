package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestId(cxt *gin.Context) {
	fmt.Println("------------------------")
	cxt.Next()

	fmt.Println("*******************************")
}
