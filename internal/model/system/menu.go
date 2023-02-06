package system

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
	"time"
)

type Menu struct {
	common.Model
	Title      string `gorm:"column:title" form:"title" json:"title"`
	Name       string `gorm:"column:name" form:"name" json:"name"`
	Parent     int    `gorm:"column:parent" form:"parent" json:"parent"`
	Path       string `gorm:"column:path" form:"path" json:"path"`
	Redirect   string `gorm:"column:redirect" form:"redirect" json:"redirect"`
	Component  string `gorm:"column:component" form:"component" json:"component"`
	Icon       string `gorm:"column:icon" form:"icon" json:"icon"`
	Sort       int    `gorm:"column:sort" form:"sort" json:"sort"`
	AlwaysShow int    `gorm:"column:alwaysShow" form:"alwaysShow" json:"alwaysShow"`
}

type MenuTree struct {
	Id         int         `json:"id"`
	Parent     int         `json:"parent"`
	Path       string      `json:"path"`
	Redirect   string      `json:"redirect"`
	Component  string      `json:"component"`
	AlwaysShow int         `json:"alwaysShow"`
	Name       string      `json:"name"`
	Title      string      `json:"title"`
	Icon       string      `json:"icon"`
	Children   []*MenuTree `json:"children,omitempty"`
}

func (m Menu) Create() error {

	return tool.MysqlDb.Model(&m).Create(&m).Error
}

func (m Menu) Update() error {

	var data = make(map[string]interface{})
	data["id"] = m.Id
	data["title"] = m.Title
	data["name"] = m.Name
	data["parent"] = m.Parent
	data["path"] = m.Path
	data["redirect"] = m.Redirect
	data["component"] = m.Component
	data["icon"] = m.Icon
	data["sort"] = m.Sort
	data["alwaysShow"] = m.AlwaysShow
	data["opt_user"] = m.OptUser
	data["updated_at"] = time.Now().Unix()

	if err := tool.MysqlDb.Model(m).Where("id", m.Id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (m Menu) GetAll() (menus []*Menu, err error) {

	if err = tool.MysqlDb.Model(&m).Where("status", 0).Order("sort desc").Order("created_at desc").Find(&menus).Error; err != nil {
		return
	}

	return
}

func (m Menu) GetByIds(ids []int) (menus []*Menu, err error) {

	if err = tool.MysqlDb.Model(&m).Where("id in ?", ids).Find(&menus).Error; err != nil {
		return
	}

	return
}

func (m Menu) List(params common.Params) (menus []*Menu, total int64, err error) {
	tx := tool.MysqlDb.Model(&m)
	if m.Id > 0 {
		tx = tx.Where("id", m.Id)
	}
	if m.Title != "" {
		tx = tx.Where("title", m.Title)
	}
	if m.Name != "" {
		tx = tx.Where("id", m.Name)
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
