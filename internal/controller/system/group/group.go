package group

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {

	group := system.NewGroup()
	result := common.NewResult(c)

	if err := c.ShouldBind(&group); err != nil {
		result.Fail("参数格式错误，请检查")
		return
	}

	if err := group.Add(); err != nil {
		result.Error(err)
		return
	}

	result.Success()
}

func Update(c *gin.Context) {

	group := system.NewGroup()
	result := common.NewResult(c)

	if err := c.ShouldBind(&group); err != nil {
		result.Fail("参数格式错误，请检查")
		return
	}

	if err := group.Update(); err != nil {
		result.Error(err)
		return
	}

	result.Success()
}
