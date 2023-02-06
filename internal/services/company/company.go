package company

import (
	"errors"
	"game.sdk.center/internal/model/basics"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/conmon"
	"time"
)

type ServiceCompany struct {
	basics.Company
	conmon.Format
}

func NewServiceCompany() ServiceCompany {
	return ServiceCompany{}
}

func (c ServiceCompany) List(params common.Params) (sc []*ServiceCompany, total int64, err error) {
	params.Check()
	companys, total, err := c.Company.List(params)
	if err != nil {
		return nil, 0, err
	}
	for _, v := range companys {
		company := &ServiceCompany{
			Company: *v,
			Format: conmon.Format{
				CreateDate: time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05"),
				UpdateDate: time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
			},
		}
		sc = append(sc, company)
	}

	return
}

func (c ServiceCompany) Create() error {
	if c.Name == "" {
		return errors.New("名称不能为空")
	}

	return c.Company.Create()
}

func (c ServiceCompany) Update() error {
	if c.Id <= 0 {
		return errors.New("id无效")
	}

	return c.Company.Update()
}
