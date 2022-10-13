package group

import (
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	group := system.NewGroup()

	if err := c.ShouldBind(&group); err != nil {
		return
	}
}
