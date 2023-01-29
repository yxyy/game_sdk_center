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
	"sort"
	"time"
)

type ServicesPermission struct {
	system.Permission
	conmon.Format
	Permissions []int `json:"permissions"`
	// Routers     []string `json:"routers"`
}

func NewServicesPermission() ServicesPermission {
	return ServicesPermission{}
}

func (s ServicesPermission) Create() error {
	if s.Name == "" {
		return errors.New("标题不能空")
	}

	if len(s.Permissions) == 0 {
		return errors.New("权限不能为空不能空")
	}
	sort.Ints(s.Permissions)

	marshal, err := json.Marshal(s.Permissions)
	if err != nil {
		return err
	}
	s.Permission.Permission = string(marshal)

	router, err := s.Router()
	if err != nil {
		return err
	}
	s.Permission.Router = router

	if err = s.Permission.Create(); err != nil {
		return err
	}

	return s.removeCache()
}

func (s ServicesPermission) Update() error {
	if s.Id <= 0 {
		return errors.New("id 无效")
	}
	sort.Ints(s.Permissions)
	marshal, err := json.Marshal(s.Permissions)
	if err != nil {
		return err
	}
	s.Permission.Permission = string(marshal)

	router, err := s.Router()
	if err != nil {
		return err
	}
	s.Permission.Router = router

	if err = s.Permission.Update(); err != nil {
		return err
	}

	return s.removeCache()
}

func (s ServicesPermission) List(params common.Params) (servicesPermissionList []*ServicesPermission, total int64, err error) {

	params.Check()
	list, total, err := s.Permission.List(params)
	if err != nil {
		return nil, 0, err
	}

	userMap, err := mapping.User()
	if err != nil {
		return
	}
	for _, v := range list {
		tmp := &ServicesPermission{
			Permission: *v,
			Format: conmon.Format{
				UpdateDate:  time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
				OptUserName: userMap[v.OptUser],
			},
		}
		servicesPermissionList = append(servicesPermissionList, tmp)
	}

	return servicesPermissionList, total, nil

}

func (s ServicesPermission) Lists() (servicesPermissionList []*ServicesPermission, err error) {

	permissions, err := s.Permission.GetAll()
	if err != nil {
		return nil, err
	}

	for _, v := range permissions {
		tmp := &ServicesPermission{
			Permission: *v,
			Format: conmon.Format{
				UpdateDate: time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
			},
		}
		servicesPermissionList = append(servicesPermissionList, tmp)
	}

	return servicesPermissionList, nil

}

func (s ServicesPermission) Router() (string, error) {

	var m system.Menu
	menus, err := m.GetByIds(s.Permissions)
	if err != nil {
		return "", err
	}

	var routers []string
	for _, menu := range menus {
		routers = append(routers, menu.Path)
	}

	marshal, err := json.Marshal(&routers)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

func (s ServicesPermission) removeCache() error {
	if err := tool.RedisClient.Del(context.Background(), "menus").Err(); err != nil {
		return err
	}

	return nil
}