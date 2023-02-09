package router

import (
	"game.sdk.center/internal/controller/basics/app"
	"game.sdk.center/internal/controller/basics/app_type"
	"game.sdk.center/internal/controller/basics/company"
	"game.sdk.center/internal/controller/system/group"
	"game.sdk.center/internal/controller/system/permission"
	"github.com/gin-gonic/gin"
)

func InitCommonRouter(r *gin.Engine) {

	common := r.Group("/common")
	{
		// 研发公司下拉框
		common.GET("/company", company.Lists)
		// 应用类型下拉
		common.GET("/app_type", app_type.Lists)
		// 应用下拉
		common.GET("/app", app.Lists)

		// 游戏下拉
		common.GET("/game", app_type.List)
		// 账号分组
		common.GET("/group", group.Lists)
		// 权限
		common.GET("/permission", permission.Lists)
	}
}
