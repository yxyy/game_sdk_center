package game

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/basics"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	game := basics.NewServiceGame()
	params := common.NewParams()

	if err := c.ShouldBind(&game); err != nil {
		response.Error(err)
	}

	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
	}

	sc, total, err := game.List(params)
	if err != nil {
		response.Error(err)
	}

	data := make(map[string]interface{})
	data["rows"] = sc
	data["total"] = total

	response.SuccessData(data)
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	game := basics.NewServiceGame()

	if err := c.ShouldBind(&game); err != nil {
		response.Error(err)
	}

	game.OptUser = c.GetInt("userId")
	if err := game.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	game := basics.NewServiceGame()

	if err := c.ShouldBind(&game); err != nil {
		response.Error(err)
	}

	game.OptUser = c.GetInt("userId")
	if err := game.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Lists(c *gin.Context) {
	response := common.NewResponse(c)
	game := basics.NewServiceGame()

	if err := c.ShouldBind(&game); err != nil {
		response.Error(err)
	}

	list, err := game.Lists()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(list)
}
