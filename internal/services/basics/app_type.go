package basics

import (
	"errors"
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/basics"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/conmon"
)

type ServiceAppType struct {
	basics.AppType
	conmon.Format
	CompanyName string `json:"company_name"`
	TypeName    string `json:"type_name"`
}

func NewServiceAppType() ServiceAppType {
	return ServiceAppType{}
}

func (c ServiceAppType) List(params common.Params) (sc []*ServiceAppType, total int64, err error) {
	params.Check()
	list, total, err := c.AppType.List(params)
	if err != nil {
		return nil, 0, err
	}
	userMap, err := mapping.User()
	if err != nil {
		return nil, 0, err
	}

	for _, v := range list {

		format := conmon.Formats(v.Model)
		format.OptUserName = userMap[v.OptUser]
		serviceAppType := &ServiceAppType{
			AppType: *v,
			Format:  format,
		}

		sc = append(sc, serviceAppType)
	}

	return
}

func (c ServiceAppType) Create() error {
	if c.Name == "" {
		return errors.New("名称不能为空")
	}

	return c.AppType.Create()
}

func (c ServiceAppType) Update() error {
	if c.Id <= 0 {
		return errors.New("id无效")
	}

	return c.AppType.Update()
}

func (c ServiceAppType) Lists() (sc []*basics.AppType, err error) {

	return c.AppType.GetAll()
}
