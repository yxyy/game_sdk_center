package group

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {

	group := system.NewGroup()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&group); err != nil {
		response.Fail("参数格式错误，请检查")
		return
	}

	if err := group.Add(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func Update(c *gin.Context) {

	group := system.NewGroup()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&group); err != nil {
		response.Fail("参数格式错误，请检查")
		return
	}

	if err := group.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func List(c *gin.Context) {
	group := system.NewGroup()
	response := common.NewResponse(c)
	params := common.NewParams()
	if err := c.ShouldBind(group); err != nil {
		response.Error(err)
		return
	}
	if err := c.ShouldBind(params); err != nil {
		response.Error(err)
		return
	}
	grous, err := group.List(params)
	if err != nil {
		response.Error(err)
		return
	}

	response.SuccessData(grous)
}
