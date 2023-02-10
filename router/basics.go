package router

import (
	"game.sdk.center/internal/controller/basics/app"
	"game.sdk.center/internal/controller/basics/app_type"
	"game.sdk.center/internal/controller/basics/channel"
	"game.sdk.center/internal/controller/basics/company"
	"game.sdk.center/internal/controller/basics/game"
	"github.com/gin-gonic/gin"
)

func InitBasicsRouter(r *gin.Engine) {

	basics := r.Group("/basics")
	{
		// 研发公司
		basics.POST("/company/create", company.Create)
		basics.POST("/company/update", company.Update)
		basics.GET("/company/list", company.List)

		// 应用类型
		basics.POST("/app_type/create", app_type.Create)
		basics.POST("/app_type/update", app_type.Update)
		basics.GET("/app_type/list", app_type.List)

		// 应用
		basics.POST("/app/create", app.Create)
		basics.POST("/app/update", app.Update)
		basics.GET("/app/list", app.List)

		// 游戏
		basics.POST("/game/create", game.Create)
		basics.POST("/game/update", game.Update)
		basics.GET("/game/list", game.List)

		// 游戏
		basics.POST("/channel/create", channel.Create)
		basics.POST("/channel/update", channel.Update)
		basics.GET("/channel/list", channel.List)
	}
}
