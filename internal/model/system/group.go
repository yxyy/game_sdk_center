package system

import (
	"errors"
	"game.sdk.center/tool"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name string `gorm:"name"`
}

func NewGroup() *Group {
	return &Group{}
}

func (this *Group) Add() error {
	if this.Name == "" {
		return errors.New("名称不能为空")
	}

	if err := tool.MysqlDb.Model(this).Create(this).Error; err != nil {
		return err
	}

	return nil

}
