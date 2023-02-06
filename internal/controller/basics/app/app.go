package app

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/basics"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	app := basics.NewServiceApp()
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
	app := basics.NewServiceApp()

	if err := c.ShouldBind(&app); err != nil {
		response.Error(err)
	}

	if err := app.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	app := basics.NewServiceApp()

	if err := c.ShouldBind(&app); err != nil {
		response.Error(err)
	}

	if err := app.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}
