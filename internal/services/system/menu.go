package system

import (
	"context"
	"encoding/json"
	"errors"
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/internal/services/conmon"
	"game.sdk.center/tool"
	"github.com/go-redis/redis/v8"
	"time"
)

type ServicesMenu struct {
	system.Menu
	conmon.Format
}

func NewServicesMenu() ServicesMenu {
	return ServicesMenu{}
}

func (m ServicesMenu) Create() error {
	if m.Title == "" {
		return errors.New("标题不能空")
	}
	if m.Name == "" {
		return errors.New("标识不能空")
	}
	if m.Path == "" {
		return errors.New("路径不能空")
	}
	if m.Component == "" {
		return errors.New("组件不能空")
	}

	if err := m.Menu.Create(); err != nil {
		return err
	}

	return m.removeCache()
}

func (m ServicesMenu) Update() error {
	if m.Id <= 0 {
		return errors.New("id 无效")
	}

	if err := m.Menu.Update(); err != nil {
		return err
	}

	return m.removeCache()
}

func (m ServicesMenu) List(params common.Params) (servicesMenuList []*ServicesMenu, total int64, err error) {

	list, total, err := m.Menu.List(params)
	if err != nil {
		return nil, 0, err
	}

	userMap, err := mapping.User()
	if err != nil {
		return
	}
	for _, v := range list {
		tmp := &ServicesMenu{
			Menu: *v,
			Format: conmon.Format{
				UpdateDate:  time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
				OptUserName: userMap[v.OptUser],
			},
		}
		servicesMenuList = append(servicesMenuList, tmp)
	}

	return servicesMenuList, total, nil

}

func (m ServicesMenu) GetTree() (trees []*system.MenuTree, err error) {

	// 查缓存
	result, err := tool.RedisClient.Get(context.Background(), "menus").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if result != "" {
		if err = json.Unmarshal([]byte(result), &trees); err != nil {
			return nil, err
		}

		return
	}

	// 无缓存，查数据库
	menus, err := m.GetAll()
	if err != nil {
		return nil, err
	}

	trees = tree(menus, 0)
	jsonTress, err := json.Marshal(&trees)
	if err != nil {
		return nil, err
	}

	if err = tool.RedisClient.Set(context.Background(), "menus", string(jsonTress), time.Second*2*360).Err(); err != nil {
		return nil, err
	}

	return trees, nil
}

func tree(menus []*system.Menu, pid int) []*system.MenuTree {
	var Trees []*system.MenuTree
	for _, v := range menus {
		if v.Parent == pid {
			if v.Component[0] != '/' {
				v.Component = "/" + v.Component
			}
			menuTree := &system.MenuTree{
				Id:         v.Id,
				Parent:     v.Parent,
				Path:       v.Path,
				Redirect:   v.Redirect,
				Component:  v.Component,
				Name:       v.Name,
				AlwaysShow: v.AlwaysShow,
				Title:      v.Title,
				Icon:       v.Icon,
				Children:   tree(menus, int(v.Id)),
			}

			Trees = append(Trees, menuTree)
		}
	}

	return Trees
}

func (m ServicesMenu) removeCache() error {
	if err := tool.RedisClient.Del(context.Background(), "menus").Err(); err != nil {
		return err
	}

	return nil
}
