package middleware

import (
	"context"
	"encoding/json"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
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
	user := system.NewUser()
	if err = json.Unmarshal([]byte(result), &user); err != nil {
		response.SetResult(5000, err.Error(), nil)
		c.Abort()
		return
	}

	c.Set("userInfo", user)

	c.Next()
}
