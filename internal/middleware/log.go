package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Log(c *gin.Context) {

	var err error
	var body []byte

	uuid := tool.Uuid()

	c.Set("request_id", uuid)

	logger := log.
		WithField("request_id", uuid).
		WithField("ip", c.ClientIP()).
		WithField("method", c.Request.Method).
		WithField("url", fmt.Sprint(c.Request.URL)).
		WithField("Access-Token", c.Request.Header.Get("Access-Token"))

	if c.Request.Method == "POST" {

		// TODO ContentType 和 body 无法添加，以后研究
		logger.WithField("ContentType", c.ContentType())

		switch c.ContentType() {
		case "application/x-www-form-urlencoded":
			if err = c.Request.ParseForm(); err != nil {
				logger.Error(err)
				return
			}
			body, err = json.Marshal(c.Request.Form)
			if err != nil {
				logger.Error(err)
				return
			}

		case "application/json":
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Error(err)
				return
			}
			// 重写回去
			c.Request.Body = io.NopCloser(bytes.NewReader(body))

		}

		logger.WithField("body", string(body))
	}

	logger.Info("请求日志")

	c.Next()

}
