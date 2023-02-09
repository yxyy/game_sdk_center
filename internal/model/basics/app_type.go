package basics

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type AppType struct {
	common.Model
	Name string `json:"name" form:"name"`
}

func (t AppType) Create() error {
	return tool.MysqlDb.Model(&t).Create(&t).Error
}

func (t AppType) Update() error {
	return tool.MysqlDb.Model(&t).Where("id", t.Id).Updates(&t).Error
}

func (t AppType) List(params common.Params) (companys []*AppType, total int64, err error) {
	tx := tool.MysqlDb.Model(&t)
	if t.Id > 0 {
		tx = tx.Where("id", t.Id)
	}
	if t.Name != "" {
		tx = tx.Where("name like ?", t.Name+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(params.Offse).Limit(params.Limit).Find(&companys).Error

	return

}

func (t AppType) GetAll() (appType []*AppType, err error) {

	err = tool.MysqlDb.Model(&t).Find(&appType).Error
	return
}
