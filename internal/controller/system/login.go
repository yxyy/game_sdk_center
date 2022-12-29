package system

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	params := system.NewLoginParams()
	response := common.NewResponse(c)

	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
	}

	var login system.Loginer
	if params.Phone >= 0 {
		// 手机登录
		login = system.NewMobile()
	} else {
		// 账号登录
		login = system.NewAccount()
	}

	if err := system.Login(login); err != nil {
		response.Error(err)
	}

	response.Success()
}
