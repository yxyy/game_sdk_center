package basics

import (
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
)

type Package struct {
	common.Model
	GameId     int    `json:"game_id"`
	Name       string `json:"name"`
	ChannelId  int    `json:"channel_id"`
	Sdk        int    `json:"sdk"`
	Campaign   string `json:"campaign"`
	Status     int8   `json:"status"`
	PackStatus int8   `json:"pack_status"`
	PackAt     int64  `json:"pack_at"`
}

func (p Package) List(params common.Params) (list []map[string]interface{}, total int64, err error) {
	tx := tool.MysqlDb.Model(&p).
		Select("lhc_packages.*",
			"lhc_games.name,lhc_games.app_id,lhc_games.os,lhc_games.icon",
			"lhc_apps.company_id,lhc_apps.name AS app_name,lhc_apps.type AS app_type",
			"lhc_companies.name AS company_name ").
		Joins("left join lhc_apps on lhc_games.app_id = lhc_apps.id").
		Joins("left join lhc_companies on lhc_apps.company_id = lhc_companies.id")
	if p.ChannelId > 0 {
		tx = tx.Where("lhc_packages.channel_id", p.ChannelId)
	}
	if p.GameId > 0 {
		tx = tx.Where("lhc_packages.game_id", p.GameId)
	}
	if p.Campaign != "" {
		tx = tx.Where("lhc_packages.campaign", p.Campaign)
	}
	if p.Status > 0 {
		tx = tx.Where("lhc_packages.status", p.Status)
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err = tx.Offset(params.Offse).Limit(params.Limit).Find(&list).Error

	return

}

func (p Package) Create() error {
	return tool.MysqlDb.Model(&p).Create(&p).Error
}

func (p Package) Update() error {
	return tool.MysqlDb.Model(&p).Where("id", p.Id).Updates(&p).Error
}
