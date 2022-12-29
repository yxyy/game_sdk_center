package system

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	login := system.NewLoginParams()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&login); err != nil {
		response.Error(err)
	}

	var loginer system.Loginer
	if login.Phone >= 0 {
		// 手机登录
		loginer = system.NewMobile()
	} else {
		// 账号登录
		loginer = system.NewAccount()
	}

	if err := system.Login(loginer); err != nil {
		response.Error(err)
	}

	response.Success()
}
