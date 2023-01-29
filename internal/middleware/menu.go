package middleware

import (
	"fmt"
	"game.sdk.center/internal/model/common"
	"github.com/gin-gonic/gin"
	"strings"
)

func Menu(c *gin.Context) {

	response := common.NewResponse(c)
	path := c.FullPath()
	index := strings.LastIndex(path, "/")
	fmt.Println("--------------")
	fmt.Println(path, index)
	fmt.Println(c.Request.URL.Path)
	if index <= 0 {
		response.SetResult(4003, "无效的请求路径", nil)
	}

	// groupId := c.GetString("groupId")
	// result, err := tool.RedisClient.Get(context.Background(), "router_"+groupId).Result()
	// if err != nil {
	// 	response.Error(err)
	// }

	c.Next()
}
