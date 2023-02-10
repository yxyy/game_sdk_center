package system

import (
	"game.sdk.center/tool"
)

type Log struct {
	Id        int    `json:"id" form:"id" gorm:"primarykey"`
	Ip        string `json:"ip"`
	Path      string `json:"path"`
	OptUser   int    `json:"opt_user,omitempty" form:"opt_user"`
	CreatedAt int64  `json:"created_at" form:"created_at" `
	RequestId string `json:"request_id"`
}

func (l Log) Create() error {
	return tool.MysqlDb.Model(&l).Create(&l).Error
}
