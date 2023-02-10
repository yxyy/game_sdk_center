package basics

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type Channel struct {
	common.Model
	Name string `json:"name" form:"name"`
	Flag string `json:"flag" form:"flag"`
}

func (c Channel) Create() error {
	return tool.MysqlDb.Model(&c).Create(&c).Error
}

func (c Channel) Update() error {
	return tool.MysqlDb.Model(&c).Updates(&c).Error
}

func (c Channel) List(params common.Params) (list []*Channel, total int64, err error) {
	tx := tool.MysqlDb.Model(&c)
	if c.Id > 0 {
		tx = tx.Where("id", c.Id)
	}
	if c.Name != "" {
		tx = tx.Where("name like ?", c.Name+"%")
	}
	if c.Flag != "" {
		tx = tx.Where("flag like ?", c.Flag+"%")
	}
	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err = tx.Offset(params.Offse).Limit(params.Limit).Find(&list).Error

	return
}

func (c Channel) GetAll() (list []*Channel, err error) {
	err = tool.MysqlDb.Model(&c).Find(&list).Error
	return
}
