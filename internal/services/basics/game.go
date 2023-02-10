package basics

import (
	"errors"
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/basics"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/conmon"
	"game.sdk.center/tool"
	"github.com/spf13/viper"
)

type ServiceGame struct {
	basics.GameInfo
	conmon.Format
	GameName   string `json:"game_name"`
	TypeName   string `json:"type_name"`
	IconFormat string `json:"icon_format"`
}

type ServiceGameInfo struct {
	basics.GameInfo
	conmon.Format
}

func NewServiceGame() ServiceGame {
	return ServiceGame{}
}

func (g ServiceGame) List(params common.Params) (sc []*ServiceGame, total int64, err error) {
	params.Check()
	list, total, err := g.Game.List(params)
	if err != nil {
		return nil, 0, err
	}
	userMap, err := mapping.User()
	if err != nil {
		return nil, 0, err
	}

	appType, err := mapping.AppType()
	if err != nil {
		return nil, 0, err
	}

	ossDomain := viper.GetString("oss.host")
	for _, v := range list {
		format := conmon.Formats(v.Model)
		format.OptUserName = userMap[v.OptUser]
		if len(v.Icon) > 0 && v.Icon[0] != '/' {
			v.Icon = "/" + v.Icon
		}
		serviceGame := &ServiceGame{
			GameInfo:   *v,
			Format:     format,
			TypeName:   appType[v.AppType],
			IconFormat: ossDomain + v.Icon,
		}
		sc = append(sc, serviceGame)
	}

	return
}

func (g ServiceGame) Create() error {
	if g.Name == "" {
		return errors.New("名称不能为空")
	}

	if g.AppId <= 0 {
		return errors.New("应用不能为空")
	}

	if g.Os <= 0 {
		return errors.New("操作系统不能为空")
	}

	if g.Status <= 0 {
		return errors.New("状态不能为空")
	}

	if g.CallbackUrl == "" {
		return errors.New("回调地址不能为空")
	}
	if err := conmon.ParseUrl(g.CallbackUrl); err != nil {
		return errors.New("回调地址：" + err.Error())
	}

	if g.CallBackTestUrl != "" {
		if err := conmon.ParseUrl(g.CallBackTestUrl); err != nil {
			return errors.New("测试回调地址：" + err.Error())
		}
	}

	g.ClientKey = tool.Salt()
	g.ServerKey = tool.Salt()

	return g.Game.Create()
}

func (g ServiceGame) Update() error {
	if g.Id <= 0 {
		return errors.New("id无效")
	}

	if g.CallbackUrl != "" {
		if err := conmon.ParseUrl(g.CallbackUrl); err != nil {
			return errors.New("回调地址：" + err.Error())
		}
	}

	if g.CallBackTestUrl != "" {
		if err := conmon.ParseUrl(g.CallBackTestUrl); err != nil {
			return errors.New("测试回调地址：" + err.Error())
		}
	}

	// 不更新key
	g.ClientKey = ""
	g.ServerKey = ""

	return g.Game.Update()
}

func (g ServiceGame) Lists() (sc []*basics.App, err error) {

	return g.Game.GetAll()
}
