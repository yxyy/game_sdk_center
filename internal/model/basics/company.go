package basics

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type Company struct {
	common.Model
	Name    string `json:"name" form:"name"`
	Address string `json:"address,omitempty"`
	Phone   int    `json:"phone,omitempty"`
}

func (c Company) Create() error {
	return tool.MysqlDb.Model(&c).Create(&c).Error
}

func (c Company) Update() error {
	return tool.MysqlDb.Model(&c).Where("id", c.Id).Updates(&c).Error
}

func (c Company) List(params common.Params) (companys []*Company, total int64, err error) {
	tx := tool.MysqlDb.Model(&c)
	if c.Id > 0 {
		tx = tx.Where("id", c.Id)
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

func (c Company) GetAll() (companys []*Company, err error) {
	err = tool.MysqlDb.Model(&c).Select("id,name").Find(&companys).Error
	return
}
