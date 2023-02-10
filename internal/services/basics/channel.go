package basics

import (
	"errors"
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/basics"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/services/conmon"
)

type ServiceChannel struct {
	basics.Channel
	conmon.Format
}

func NewServiceChannel() ServiceChannel {
	return ServiceChannel{}
}

func (c ServiceChannel) Create() error {
	if c.Name == "" {
		return errors.New("名称不能为空")
	}
	if c.Flag == "" {
		return errors.New("标识不能为空")
	}

	return c.Channel.Create()
}

func (c ServiceChannel) Update() error {
	if c.Id <= 0 {
		return errors.New("id无效")
	}

	return c.Channel.Update()
}

func (c ServiceChannel) List(params common.Params) (list []*ServiceChannel, total int64, err error) {
	params.Check()
	channels, total, err := c.Channel.List(params)
	if err != nil {
		return nil, 0, err
	}

	userMap, err := mapping.User()
	if err != nil {
		return nil, 0, err
	}

	for _, v := range channels {
		format := conmon.Formats(v.Model)
		format.OptUserName = userMap[v.OptUser]
		node := &ServiceChannel{
			Format:  format,
			Channel: *v,
		}
		list = append(list, node)
	}

	return
}

func (c ServiceChannel) Lists() (list []*basics.Channel, err error) {

	return c.GetAll()
}
