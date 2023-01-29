package permission

import (
	"fmt"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/system"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	ServicesPermission := system.NewServicesPermission()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&ServicesPermission); err != nil {
		response.Error(err)
	}

	ServicesPermission.OptUser = c.GetInt("userId")
	if err := ServicesPermission.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {

	ServicesPermission := system.NewServicesPermission()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&ServicesPermission); err != nil {
		response.Error(err)
	}

	fmt.Printf("%#v\n", ServicesPermission)
	ServicesPermission.OptUser = c.GetInt("userId")
	if err := ServicesPermission.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func List(c *gin.Context) {

	ServicesPermission := system.NewServicesPermission()
	response := common.NewResponse(c)
	params := common.NewParams()
	if err := c.ShouldBind(ServicesPermission); err != nil {
		response.Error(err)
	}
	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
	}

	list, total, err := ServicesPermission.List(params)
	if err != nil {
		response.Error(err)
	}

	var data = make(map[string]interface{})
	data["rows"] = list
	data["total"] = total

	response.SuccessData(data)
}

func Lists(c *gin.Context) {

	ServicesPermission := system.NewServicesPermission()
	response := common.NewResponse(c)
	list, err := ServicesPermission.Lists()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(list)
}
