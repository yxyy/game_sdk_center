package company

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/company"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	response := common.NewResponse(c)
	serviceCompany := company.NewServiceCompany()
	params := common.NewParams()

	if err := c.ShouldBind(&serviceCompany); err != nil {
		response.Error(err)
	}

	if err := c.ShouldBind(&params); err != nil {
		response.Error(err)
	}

	sc, total, err := serviceCompany.List(params)
	if err != nil {
		response.Error(err)
	}

	data := make(map[string]interface{})
	data["rows"] = sc
	data["total"] = total

	response.SuccessData(data)
}

func Create(c *gin.Context) {
	response := common.NewResponse(c)
	serviceCompany := company.NewServiceCompany()

	if err := c.ShouldBind(&serviceCompany); err != nil {
		response.Error(err)
	}

	if err := serviceCompany.Create(); err != nil {
		response.Error(err)
	}

	response.Success()
}

func Update(c *gin.Context) {
	response := common.NewResponse(c)
	serviceCompany := company.NewServiceCompany()

	if err := c.ShouldBind(&serviceCompany); err != nil {
		response.Error(err)
	}

	if err := serviceCompany.Update(); err != nil {
		response.Error(err)
	}

	response.Success()
}
