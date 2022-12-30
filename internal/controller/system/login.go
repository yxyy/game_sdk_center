package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
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

	token := CreateToken(user)
	var data = make(map[string]interface{})
	data["userInfo"] = user
	data["token"] = token

	if err = updateToken(token, user); err != nil {
		response.Error(err)
	}

	user.LastLoginIp = c.ClientIP()
	user.LastLoginTime = time.Now().Unix()
	if err = user.Update(); err != nil {
		response.Error(err)
	}

	response.SuccessData(data)
}

func Logout(c *gin.Context) {
	access_token := c.Request.Header.Get("Access-Token")
	account := c.GetString("account")
	response := common.NewResponse(c)
	// 删除token
	if err := tool.RedisClient.Del(context.Background(), "access_token:"+account).Err(); err != nil {
		response.Fail("退出失败")
	}
	// 删除token信息
	if err := tool.RedisClient.Del(context.Background(), access_token).Err(); err != nil {
		response.Fail("退出失败")
	}
	response.Success()
}

func CreateToken(user system.User) string {

	return fmt.Sprintf("%d_%s_%s", user.GroupId, user.Account, tool.Salt())

}

func updateToken(token string, user system.User) error {
	// 先删除可能存在的token
	result, err := tool.RedisClient.Get(context.Background(), "access_token:"+user.Account).Result()
	if !errors.Is(err, redis.Nil) {
		if err != nil {
			return err
		}
		if err = tool.RedisClient.Del(context.Background(), result).Err(); err != nil {
			return err
		}

	}

	marshal, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	// 缓存用户token
	if err = tool.RedisClient.Set(context.Background(), "access_token:"+user.Account, token, time.Second*2*3600).Err(); err != nil {
		return err
	}
	// 缓存 token-user信息
	if err = tool.RedisClient.Set(context.Background(), token, string(marshal), time.Second*2*3600).Err(); err != nil {
		return err
	}

	return nil
}
