package router

import (
	"game.sdk.center/internal/controller/basics/company"
	"github.com/gin-gonic/gin"
)

func InitBasicsRouter(r *gin.Engine) {

	basics := r.Group("/basics")
	{
		// 研发公司
		basics.POST("/company/create", company.Create)
		basics.POST("/company/update", company.Update)
		basics.GET("/company/list", company.List)
	}
}
