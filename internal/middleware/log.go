package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {

		var body []byte
		var err error
		logger := log.
			WithField("ip", c.ClientIP()).
			WithField("method", c.Request.Method).
			WithField("url", fmt.Sprint(c.Request.URL)).
			WithField("Access-Token", c.Request.Header.Get("Access-Token"))

		if c.Request.Method == "POST" {

			logger.WithField("ContentType", c.ContentType())

			switch c.ContentType() {
			case "application/x-www-form-urlencoded":
				if err := c.Request.ParseForm(); err != nil {
					log.Fatal(err)
				}
				body, err = json.Marshal(c.Request.Form)
				if err != nil {
					log.Fatal(err)
				}

			case "application/json":
				body, err = io.ReadAll(c.Request.Body)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		logger.Info(string(body), "第二个参数")

	}
}
