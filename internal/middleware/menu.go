package middleware

import (
	"context"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
	"strings"
)

func Menu(c *gin.Context) {

	if c.GetString("groupId") == "supper" {
		c.Next()
		return
	}
	response := common.NewResponse(c)
	path := c.FullPath()
	index := strings.LastIndex(path, "/")
	if index <= 0 {
		response.SetResult(4004, "无效的请求路径", nil)
	}

	groupId := c.GetString("groupId")
	result, err := tool.RedisClient.HGet(context.Background(), "menu_router:"+groupId, path[:index]).Result()
	if err != nil {
		response.Error(err)
	}

	if result != "1" {
		response.SetResult(4003, "没有权限", nil)
	}

	c.Next()
}
