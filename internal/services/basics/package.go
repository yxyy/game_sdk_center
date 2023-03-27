package basics

import (
	"errors"
	"game.sdk.center/internal/model/basics"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/conmon"
	"game.sdk.center/tool"
)

type ServicePackage struct {
	basics.Package
	conmon.Format
}

func NewServicePackage() ServicePackage {
	return ServicePackage{}
}

func (p ServicePackage) List(params common.Params) (list []map[string]interface{}, total int64, err error) {
	params.Check()
	// list, total, err = p.Package.List(params)
	return p.Package.List(params)
}

func (p ServicePackage) Create() error {
	if p.Name == "" {
		return errors.New("渠道包名称不能为空")
	}

	if p.ChannelId <= 0 {
		return errors.New("渠道不能为空")
	}

	if p.GameId <= 0 {
		return errors.New("游戏不能为空")
	}

	if p.Status < 0 {
		return errors.New("状态无效")
	}

	p.Campaign = tool.Random(8)
	if p.Campaign == "" {
		return errors.New("渠道包标识生成错误")
	}

	return p.Package.Create()
}

func (p ServicePackage) Update() error {

	if p.Id <= 0 {
		return errors.New("渠道不能为空")
	}

	if p.Status < 0 {
		return errors.New("状态无效")
	}

	return p.Package.Update()
}
