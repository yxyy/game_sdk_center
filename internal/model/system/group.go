package system

import (
	"errors"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type Group struct {
	common.Model
	Name string `gorm:"name" form:"name"`
	Flag string `gorm:"flag" form:"flag"`
}

func NewGroup() *Group {
	return &Group{}
}

func (g *Group) Add() error {
	if g.Name == "" {
		return errors.New("名称不能为空")
	}
	if g.Flag == "" {
		return errors.New("标识不能为空不能为空")
	}

	if err := tool.MysqlDb.Model(g).Where("name", g.Name).Take(&Group{}).Error; err == nil {
		return errors.New("该名称已存在")
	}

	if err := tool.MysqlDb.Model(g).Where("flag", g.Flag).Take(&Group{}).Error; err == nil {
		return errors.New("该标识已存在")
	}

	if err := tool.MysqlDb.Model(g).Create(g).Error; err != nil {
		return err
	}

	return nil

}

func (g *Group) Update() error {

	if g.ID == 0 {
		return errors.New("id 无效")
	}

	if err := tool.MysqlDb.Model(g).Where("id", g.ID).Updates(g).Error; err != nil {
		return err
	}

	return nil
}
