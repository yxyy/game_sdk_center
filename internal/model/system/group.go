package system

import (
	"errors"
	"game.sdk.center/lib"
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

	if err := lib.MysqlDb.Model(this).Create(this).Error; err != nil {
		return err
	}

	return nil

}
