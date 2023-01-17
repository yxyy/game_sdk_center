package group

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/system"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	group := system.NewServiceGroup()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&group); err != nil {
		response.Fail("参数格式错误，请检查")
	}

	group.OptUser = c.GetInt("userId")
	if err := group.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {

	serviceGroup := system.NewServiceGroup()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&serviceGroup); err != nil {
		response.Fail("参数格式错误，请检查")
		return
	}

	serviceGroup.OptUser = c.GetInt("userId")
	if err := serviceGroup.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func List(c *gin.Context) {

	group := system.NewServiceGroup()
	response := common.NewResponse(c)
	params := common.NewParams()
	if err := c.ShouldBind(&group); err != nil {
		response.Error(err)
		return
	}
	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
		return
	}

	groups, total, err := group.List(params)
	if err != nil {
		response.Error(err)
		return
	}

	data := make(map[string]interface{})
	data["rows"] = groups
	data["total"] = total

	response.SuccessData(data)
}

func Lists(c *gin.Context) {

	group := system.NewServiceGroup()
	response := common.NewResponse(c)
	if err := c.ShouldBind(&group); err != nil {
		response.Error(err)
	}
	groups, err := group.Lists()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(groups)
}
