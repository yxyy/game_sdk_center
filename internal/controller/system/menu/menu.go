package menu

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {

	menu := system.NewMenu()
	result := common.NewResult(c)

	if err := c.ShouldBind(&menu); err != nil {
		result.Error(err)
	}

	if err := menu.Add(); err != nil {
		result.Error(err)
	}

	result.Success()
}

func Update(c *gin.Context) {

	menu := system.NewMenu()
	result := common.NewResult(c)

	if err := c.ShouldBind(&menu); err != nil {
		result.Error(err)
	}

	if err := menu.Update(); err != nil {
		result.Error(err)
	}

	result.Success()
}
