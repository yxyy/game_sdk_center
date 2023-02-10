package channel

import (
	"fmt"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/basics"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	channel := basics.NewServiceChannel()
	params := common.NewParams()

	if err := c.ShouldBind(&channel); err != nil {
		response.Error(err)
	}
	fmt.Printf("%#v\n", channel)
	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
	}

	sc, total, err := channel.List(params)
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
	channel := basics.NewServiceChannel()

	if err := c.ShouldBind(&channel); err != nil {
		response.Error(err)
	}

	channel.OptUser = c.GetInt("userId")
	if err := channel.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	channel := basics.NewServiceChannel()

	if err := c.ShouldBind(&channel); err != nil {
		response.Error(err)
	}

	channel.OptUser = c.GetInt("userId")
	if err := channel.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Lists(c *gin.Context) {
	response := common.NewResponse(c)
	channel := basics.NewServiceChannel()

	if err := c.ShouldBind(&channel); err != nil {
		response.Error(err)
	}

	list, err := channel.Lists()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(list)
}
