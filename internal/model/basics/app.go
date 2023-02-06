package basics

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type App struct {
	common.Model
	Name      string `json:"name" form:"name"`
	CompanyId int    `json:"company_id"`
	Alias     int    `json:"alias"`
	Remark    string `json:"remark"`
}

func (c App) Create() error {
	return tool.MysqlDb.Model(&c).Create(&c).Error
}

func (c App) Update() error {
	return tool.MysqlDb.Model(&c).Where("id", c.Id).Updates(&c).Error
}

func (c App) List(params common.Params) (companys []*App, total int64, err error) {
	tx := tool.MysqlDb.Model(&c)
	if c.Id > 0 {
		tx = tx.Where("id", c.Id)
	}
	if c.CompanyId > 0 {
		tx = tx.Where("company_id", c.CompanyId)
	}
	if c.Name != "" {
		tx = tx.Where("name like ?", c.Name+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(params.Offse).Limit(params.Limit).Find(&companys).Error

	return

}
