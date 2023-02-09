package app_type

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/basics"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	app := basics.NewServiceAppType()
	params := common.NewParams()

	if err := c.ShouldBind(&app); err != nil {
		response.Error(err)
	}

	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
	}

	sc, total, err := app.List(params)
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
	app := basics.NewServiceAppType()

	if err := c.ShouldBind(&app); err != nil {
		response.Error(err)
	}

	app.OptUser = c.GetInt("userId")
	if err := app.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	app := basics.NewServiceAppType()

	if err := c.ShouldBind(&app); err != nil {
		response.Error(err)
	}

	app.OptUser = c.GetInt("userId")
	if err := app.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Lists(c *gin.Context) {
	response := common.NewResponse(c)
	app := basics.NewServiceAppType()

	if err := c.ShouldBind(&app); err != nil {
		response.Error(err)
	}

	list, err := app.Lists()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(list)
}
