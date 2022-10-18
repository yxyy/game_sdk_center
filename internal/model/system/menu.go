package system

import (
	"errors"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type Menu struct {
	common.Model
	Title     string `gorm:"title" form:"title" json:"title"`
	Flag      string `gorm:"flag" from:"flag" json:"flag"`
	Parent    int    `gorm:"parent" from:"parent" json:"parent"`
	Path      string `gorm:"path" form:"path" json:"path"`
	Component string `gorm:"component" form:"component" json:"component"`
	Icon      string `gorm:"icon" form:"icon" json:"icon"`
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) Add() error {
	if m.Title == "" {
		return errors.New("标题不能空")
	}
	if m.Flag == "" {
		return errors.New("标识不能空")
	}
	if m.Path == "" {
		return errors.New("路径不能空")
	}
	if m.Component == "" {
		return errors.New("组件不能空")
	}

	if err := tool.MysqlDb.Model(m).Where("flag", m.Flag).First(&Menu{}).Error; err == nil {
		return errors.New("标识已存在")
	}

	if err := tool.MysqlDb.Model(m).Where("component", m.Component).First(&Menu{}).Error; err == nil {
		return errors.New("组件已存在")
	}

	if err := tool.MysqlDb.Model(m).Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (m *Menu) Update() error {
	if m.Id <= 0 {
		return errors.New("id 无效")
	}

	if err := tool.MysqlDb.Model(m).Where("id", m.Id).Updates(m).Error; err != nil {
		return err
	}

	return nil
}
