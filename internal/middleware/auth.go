package middleware

import (
	"context"
	"encoding/json"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func Auth(c *gin.Context) {
	accessToken := c.Request.Header.Get("Access-Token")
	response := common.NewResponse(c)
	if accessToken == "" {
		response.SetResult(403, "Access-Token is not empty", nil)
		c.Abort()
		return
	}

	result, err := tool.RedisClient.Get(context.Background(), accessToken).Result()
	if err != nil {
		response.SetResult(403, "Access-Token is invalid", nil)
		c.Abort()
		return
	}
	user := system.NewTokenInfo()
	if err = json.Unmarshal([]byte(result), &user); err != nil {
		response.SetResult(5000, err.Error(), nil)
		c.Abort()
		return
	}

	if err = tool.RedisClient.Expire(context.Background(), accessToken, time.Second*2*3600).Err(); err != nil {
		log.WithField("request_id", c.GetString("request_id")).
			WithField("key", accessToken).
			WithField("message", "更新tokenInfo key失败").Error(err)
	}

	if err = tool.RedisClient.Expire(context.Background(), "access_token"+user.Account, time.Second*2*3600).Err(); err != nil {
		log.WithField("request_id", c.GetString("request_id")).
			WithField("key", "access_token"+user.Account).
			WithField("message", "更新access_token key失败").Error(err)
	}

	c.Set("userId", user.Id)
	c.Set("groupId", user.GroupId)
	c.Set("groupName", user.GroupName)
	c.Set("userAccount", user.Account)
	c.Set("userInfo", user)

	c.Next()
}
