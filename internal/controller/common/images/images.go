package images

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

func Uploads(c *gin.Context) {

	response := common.NewResponse(c)
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(err)
	}

	format := time.Now().Format("20060102")
	path := viper.GetString("oss.images") + "/" + format
	if err = tool.Directory("./" + path); err != nil {
		response.Error(err)
	}
	filepath := path + "/" + tool.Salt() + ".jpg"
	if err = c.SaveUploadedFile(file, "./"+filepath); err != nil {
		response.Error(err)
	}

	response.SuccessData(filepath)

}
