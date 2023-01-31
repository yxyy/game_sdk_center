package system

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
	"time"
)

type Permission struct {
	common.Model
	Name       string `gorm:"column:name" form:"name" json:"name"`
	Permission string `gorm:"column:permission" form:"-" json:"permission"`
	Router     string `gorm:"column:router" form:"-" json:"router"`
}

func (p Permission) Create() error {

	return tool.MysqlDb.Model(&p).Create(&p).Error
}

func (p Permission) Update() error {

	data := make(map[string]interface{})
	data["id"] = p.Id
	data["name"] = p.Name
	data["permission"] = p.Permission
	data["router"] = p.Router
	data["opt_user"] = p.OptUser
	data["updated_at"] = time.Now().Unix()

	return tool.MysqlDb.Model(&p).Where("id", p.Id).Updates(data).Error

}

func (p *Permission) Get() (err error) {

	return tool.MysqlDb.Model(&p).Where("id", p.Id).Find(&p).Error
}

func (p Permission) GetAll() (permissions []*Permission, err error) {

	if err = tool.MysqlDb.Model(&p).Find(&permissions).Error; err != nil {
		return
	}

	return
}

func (p Permission) List(params common.Params) (menus []*Permission, total int64, err error) {
	tx := tool.MysqlDb.Model(&p)
	if p.Id > 0 {
		tx = tx.Where("id", p.Id)
	}
	if p.Name != "" {
		tx = tx.Where("name", p.Name)
	}
	if err = tx.Count(&total).Error; err != nil {
		if err != nil {
			return
		}
	}
	if err = tx.Offset(params.Offse).Limit(params.Limit).Find(&menus).Error; err != nil {
		if err != nil {
			return
		}
	}

	return
}
