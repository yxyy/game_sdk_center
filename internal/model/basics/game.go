package basics

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type Game struct {
	common.Model
	Name            string `json:"name" form:"name"`
	Alias           string `json:"alias"`
	AppId           int    `json:"app_id"`
	Os              int    `json:"os"`
	Status          int    `json:"status"`
	Version         string `json:"version"`
	ClientKey       string `json:"client_key"`
	ServerKey       string `json:"server_key"`
	CallbackUrl     string `json:"callback_url"`
	CallBackTestUrl string `json:"callback_test_url" gorm:"column:callback_test_url"`
	Remark          string `json:"remark"`
}

type GameInfo struct {
	Game
	AppType     int    `json:"app_type"`
	AppName     string `json:"app_name"`
	CompanyId   string `json:"company_id"`
	CompanyName string `json:"company_name"`
}

func (g Game) Create() error {
	return tool.MysqlDb.Model(&g).Create(&g).Error
}

func (g Game) Update() error {
	return tool.MysqlDb.Model(&g).Where("id", g.Id).Updates(&g).Error
}

func (g Game) List(params common.Params) (games []*GameInfo, total int64, err error) {
	tx := tool.MysqlDb.Model(&g).
		Select("lhc_games.*,lhc_apps.company_id,lhc_apps.name AS app_name,lhc_apps.type AS app_type,lhc_companies.name AS company_name ").
		Joins("left join lhc_apps on lhc_games.app_id = lhc_apps.id").
		Joins("left join lhc_companies on lhc_apps.company_id = lhc_companies.id")
	if g.Id > 0 {
		tx = tx.Where("lhc_games.id", g.Id)
	}
	if g.AppId > 0 {
		tx = tx.Where("lhc_games.app_id", g.AppId)
	}
	if g.Name != "" {
		tx = tx.Where("lhc_games.name like ?", g.Name+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(params.Offse).Limit(params.Limit).Find(&games).Error

	return

}

func (g Game) GetAll() (app []*App, err error) {

	err = tool.MysqlDb.Model(&g).Find(&app).Error
	return
}
