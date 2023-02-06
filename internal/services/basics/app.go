package basics

import (
	"errors"
	"game.sdk.center/internal/model/basics"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/conmon"
	"time"
)

type ServiceApp struct {
	basics.App
	conmon.Format
}

func NewServiceApp() ServiceApp {
	return ServiceApp{}
}

func (c ServiceApp) List(params common.Params) (sc []*ServiceApp, total int64, err error) {
	params.Check()
	companys, total, err := c.App.List(params)
	if err != nil {
		return nil, 0, err
	}
	for _, v := range companys {
		company := &ServiceApp{
			App: *v,
			Format: conmon.Format{
				CreateDate: time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05"),
				UpdateDate: time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
			},
		}
		sc = append(sc, company)
	}

	return
}

func (c ServiceApp) Create() error {
	if c.Name == "" {
		return errors.New("名称不能为空")
	}

	if c.CompanyId <= 0 {
		return errors.New("研发公司不能为空")
	}

	return c.App.Create()
}

func (c ServiceApp) Update() error {
	if c.Id <= 0 {
		return errors.New("id无效")
	}

	return c.App.Update()
}
