package router

import (
	"game.sdk.center/internal/controller/system/group"
	"game.sdk.center/internal/controller/system/menu"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(r *gin.Engine) {

	system := r.Group("system")
	{
		system.POST("/group/add", group.Add)
		system.POST("/group/update", group.Update)
		system.POST("/group/list", group.List)

		system.POST("/menu/add", menu.Add)
		system.POST("/menu/update", menu.Update)
	}

}
