package system

import (
	"context"
	"errors"
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/internal/services/conmon"
	"game.sdk.center/tool"
	"time"
)

type ServiceGroup struct {
	system.Group
	conmon.Format
}

func NewServiceGroup() ServiceGroup {
	return ServiceGroup{}
}

func (g ServiceGroup) List(params common.Params) (groups []*ServiceGroup, total int64, err error) {

	params.Check()
	group, total, err := g.Group.List(params)
	if err != nil {
		return nil, 0, err
	}

	userMap, err := mapping.User()
	for _, v := range group {
		tmp := &ServiceGroup{
			Group: *v,
			Format: conmon.Format{
				CreateDate:  time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05"),
				UpdateDate:  time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
				OptUserName: userMap[int(v.OptUser)],
			},
		}

		groups = append(groups, tmp)
	}

	return
}

func (g ServiceGroup) Lists() (map[int]string, error) {

	return mapping.Group()
}

func (g ServiceGroup) Create() error {
	if g.Name == "" {
		return errors.New("名称不能为空")
	}
	if g.Flag == "" {
		return errors.New("标识不能为空不能为空")
	}

	if err := g.Group.Create(); err != nil {
		return err
	}

	return g.removeCache()
}

func (g ServiceGroup) Update() error {
	if g.Id <= 0 {
		return errors.New("id 无效")
	}

	if err := g.Group.Update(); err != nil {
		return err
	}

	return g.removeCache()
}

func (g ServiceGroup) Permission() error {
	if g.Id <= 0 {
		return errors.New("id 无效")
	}

	if err := g.Group.Update(); err != nil {
		return err
	}

	return g.removeCache()
}

func (g ServiceGroup) removeCache() error {

	return tool.RedisClient.Del(context.Background(), mapping.GroupCacheKey).Err()

}
