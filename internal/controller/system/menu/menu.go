package menu

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/system"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	servicesMenu := system.NewServicesMenu()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&servicesMenu); err != nil {
		response.Error(err)
	}

	servicesMenu.OptUser = c.GetInt64("userId")
	if err := servicesMenu.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {

	servicesMenu := system.NewServicesMenu()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&servicesMenu); err != nil {
		response.Error(err)
	}
	servicesMenu.OptUser = c.GetInt64("userId")
	if err := servicesMenu.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func List(c *gin.Context) {

	servicesMenu := system.NewServicesMenu()
	response := common.NewResponse(c)
	params := common.NewParams()
	if err := c.ShouldBind(servicesMenu); err != nil {
		response.Error(err)
	}
	if err := c.ShouldBind(params); err != nil {
		response.Error(err)
	}
	params.Check()
	list, total, err := servicesMenu.List(params)
	if err != nil {
		response.Error(err)
	}

	var data = make(map[string]interface{})
	data["rows"] = list
	data["total"] = total

	response.SuccessData(data)
}

func Tree(c *gin.Context) {

	servicesMenu := system.NewServicesMenu()
	response := common.NewResponse(c)

	tree, err := servicesMenu.GetTree()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(tree)
}
