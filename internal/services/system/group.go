package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	// 先清空角色路由权限
	if err := g.removeRouterCache(); err != nil {
		return err
	}

	// 设置角色路由权限
	return g.SetRouterCache()
}

// SetRouterCache 设置角色权限缓存
func (g ServiceGroup) SetRouterCache() error {

	if g.Id <= 0 {
		return errors.New("角色Id无效")
	}

	if err := g.Group.Get(); err != nil {
		return err
	}

	if g.PermissionId <= 0 {
		return errors.New("权限Id无效")
	}

	var permission system.Permission
	permission.Id = g.PermissionId
	if err := permission.Get(); err != nil {
		return err
	}

	var router []string
	if err := json.Unmarshal([]byte(permission.Router), &router); err != nil {
		return err
	}

	for _, v := range router {
		go func(route string) {
			if err := tool.RedisClient.HSet(context.Background(), "menu_router:"+g.Flag, route, 1).Err(); err != nil {
				fmt.Println(err)
			}
		}(v)
	}

	return nil

}

// removeRouterCache 清空角色路由缓存
func (g ServiceGroup) removeRouterCache() error {

	if g.Id <= 0 {
		return errors.New("角色Id无效")
	}

	return tool.RedisClient.Del(context.Background(), "menu_router:"+g.Flag).Err()

}

func (g ServiceGroup) removeCache() error {

	return tool.RedisClient.Del(context.Background(), mapping.GroupCacheKey).Err()

}
