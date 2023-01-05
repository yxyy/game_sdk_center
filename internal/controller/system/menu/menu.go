package menu

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	menu := system.NewMenu()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&menu); err != nil {
		response.Error(err)
	}

	if err := menu.Create(); err != nil {
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

func List(c *gin.Context) {
	menu := system.NewMenu()
	response := common.NewResponse(c)
	params := common.NewParams()
	if err := c.ShouldBind(menu); err != nil {
		response.Error(err)
	}
	if err := c.ShouldBind(params); err != nil {
		response.Error(err)
	}
	params.Check()
	menus, err := menu.List(params)
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(menus)
}

func Tree(c *gin.Context) {

	menu := system.NewMenu()
	response := common.NewResponse(c)

	tree, err := menu.GetTree()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(tree)
}
