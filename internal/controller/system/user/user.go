package user

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	system2 "game.sdk.center/internal/services/system"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	user := system2.NewServiceUser()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&user); err != nil {
		response.Error(err)
		return
	}

	params := common.NewParams()
	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
		return
	}
	params.Check()
	list, total, err := user.List(params)
	if err != nil {
		response.Error(err)
		return
	}

	data := make(map[string]interface{})
	data["rows"] = list
	data["total"] = total

	response.SuccessData(data)

}

func Info(c *gin.Context) {
	user := system2.NewServiceUser()
	response := common.NewResponse(c)
	user.Id = c.GetInt("userId")

	data, err := user.UserInfos()
	if err != nil {
		response.Error(err)
	}

	response.SuccessData(data)
}

func Create(c *gin.Context) {
	user := system.NewUser()
	response := common.NewResponse(c)
	if err := c.ShouldBind(&user); err != nil {
		response.Error(err)
		return
	}
	user.OptUser = c.GetInt("user_id")
	if err := user.Create(); err != nil {
		response.Error(err)
		return
	}
	response.Success()
}

func Update(c *gin.Context) {
	user := system.NewUser()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&user); err != nil {
		response.Error(err)
		return
	}

	user.OptUser = c.GetInt("user_id")
	if err := user.Update(); err != nil {
		response.Error(err)
		return
	}

	response.Success()
}

func Remove(c *gin.Context) {
	user := system.NewUser()
	response := common.NewResponse(c)
	if err := c.ShouldBind(user); err != nil {
		response.Error(err)
		return
	}
	if err := user.Remove(); err != nil {
		response.Error(err)
		return
	}
	response.Success()
}
