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
	switch true {
	case params.Mobile != nil:
		// 手机登录
		mobile := system.NewMobile()
		mobile = params.Mobile
		login = mobile
		break
	case params.Account != nil:
		// 账号登录
		account := system.NewAccount()
		account = params.Account
		login = account
	default:
		response.Fail("无效的登录方式")
	}

	user, err := system.Login(login)
	if err != nil {
		response.Error(err)
	}

	var data = make(map[string]interface{})
	data["userInfo"] = user
	data["token"] = "xxxx"

	response.SuccessData(data)
}
