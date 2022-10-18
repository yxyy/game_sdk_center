package menu

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {

	menu := system.NewMenu()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&menu); err != nil {
		response.Error(err)
	}

	if err := menu.Add(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {

	menu := system.NewMenu()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&menu); err != nil {
		response.Error(err)
	}

	if err := menu.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}
