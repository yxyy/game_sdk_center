package router

import (
	"game.sdk.center/internal/controller/system/group"
	"game.sdk.center/internal/controller/system/menu"
	"game.sdk.center/internal/controller/system/user"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(r *gin.Engine) {

	system := r.Group("system")
	{
		// 系统账号
		system.GET("/user/list", user.List)
		system.GET("/user/userInfo", user.Info)
		system.POST("/user/create", user.Create)
		system.POST("/user/update", user.Update)
		system.POST("/user/remove", user.Remove)

		// 系统账号分组
		system.POST("/group/create", group.Create)
		system.POST("/group/update", group.Update)
		system.POST("/group/list", group.List)

		// 菜单
		system.POST("/menu/add", menu.Add)
		system.POST("/menu/update", menu.Update)
	}

}
