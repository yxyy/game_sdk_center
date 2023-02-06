package router

import (
	"game.sdk.center/internal/controller/basics/app"
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

		// 应用
		basics.POST("/app/create", app.Create)
		basics.POST("/app/update", app.Update)
		basics.GET("/app/list", app.List)
	}
}
