package router

import (
	"game.sdk.center/internal/controller/system/group"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(r *gin.Engine) {

	system := r.Group("system")
	{
		system.POST("/group/add", group.Add)
		system.POST("/group/update", group.Update)
	}

}
