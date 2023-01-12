package menu

import (
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
	"time"
)

func Create(c *gin.Context) {

	menu := system.NewMenu()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&menu); err != nil {
		response.Error(err)
	}

	menu.OptUser = c.GetInt64("userId")
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
	menu.OptUser = c.GetInt64("userId")
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
	list, total, err := menu.List(params)
	if err != nil {
		response.Error(err)
	}

	userMap, err := mapping.User()
	if err != nil {
		return
	}
	for _, v := range list {
		v.UpdateDate = time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05")
		v.OptUserName = userMap[int(v.Id)]
	}

	var data = make(map[string]interface{})
	data["rows"] = list
	data["total"] = total

	response.SuccessData(data)
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
