package system

import (
	"errors"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type Group struct {
	common.Model
	Name    string `gorm:"name" form:"name" json:"name"`
	Flag    string `gorm:"flag" form:"flag" json:"flag"`
	Remarks string `gorm:"remarks" form:"remarks" json:"remarks"`
}

func NewGroup() *Group {
	return &Group{}
}

func (g Group) Create() error {

	if err := tool.MysqlDb.Model(&g).Where("name", g.Name).Take(&Group{}).Error; err == nil {
		return errors.New("该名称已存在")
	}

	if err := tool.MysqlDb.Model(&g).Where("flag", g.Flag).Take(&Group{}).Error; err == nil {
		return errors.New("该标识已存在")
	}

	if err := tool.MysqlDb.Model(&g).Create(&g).Error; err != nil {
		return err
	}

	return nil

}

func (g Group) Update() error {

	if err := tool.MysqlDb.Model(g).Where("id", g.Id).Updates(&g).Error; err != nil {
		return err
	}

	return nil
}

func (g Group) List(params common.Params) (grous []*Group, total int64, err error) {

	tx := tool.MysqlDb.Model(&g)
	if g.Name != "" {
		tx = tx.Where("name like ?", "%"+g.Name+"%")
	}
	if g.Flag != "" {
		tx = tx.Where("flag like ?", "%"+g.Flag+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err = tx.Offset(params.Offse).Limit(params.Limit).Find(&grous).Error; err != nil {
		return
	}

	return
}

func (g Group) GetAll() (groups []*Group, err error) {
	if err = tool.MysqlDb.Model(&g).Find(&groups).Error; err != nil {
		return nil, err
	}

	return
}
