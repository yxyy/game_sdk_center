package group

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Add(c *gin.Context) {
	group := system.NewGroup()
	result := common.NewResult(c)
	if err := c.ShouldBind(&group); err != nil {
		result.Fail("参数格式错误，请检查")
		log.WithField("request_id", c.GetString("request_id")).Error(err)
		return
	}

	if err := group.Add(); err != nil {
		result.Error(err)
		log.WithField("request_id", c.GetString("request_id")).Error(err)
		return
	}

	result.Success()
}

func Update(c *gin.Context) {
	group := system.NewGroup()
	result := common.NewResult(c)
	if err := c.ShouldBind(&group); err != nil {
		result.Fail("参数格式错误，请检查")
		log.WithField("request_id", c.GetString("request_id")).Error(err)
		return
	}

	if err := group.Update(); err != nil {
		result.Error(err)
		log.WithField("request_id", c.GetString("request_id")).Error(err)
		return
	}

	result.Success()
}
