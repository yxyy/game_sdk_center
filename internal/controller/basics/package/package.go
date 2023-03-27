package packages

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/basics"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	servicePackage := basics.NewServicePackage()

	if err := c.ShouldBind(&servicePackage); err != nil {
		response.Error(err)
	}

	if err := servicePackage.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	servicePackage := basics.NewServicePackage()

	if err := c.ShouldBind(&servicePackage); err != nil {
		response.Error(err)
	}

	servicePackage.OptUser = c.GetInt("userId")
	if err := servicePackage.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	servicePackage := basics.NewServicePackage()

	if err := c.ShouldBind(&servicePackage); err != nil {
		response.Error(err)
	}

	servicePackage.OptUser = c.GetInt("userId")
	if err := servicePackage.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}
