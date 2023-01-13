package system

import (
	"game.sdk.center/internal/model/system"
)

type ServiceUser struct {
	system.User
}

func NewServiceUser() ServiceUser {
	return ServiceUser{}
}

func (u ServiceUser) UserInfos() (map[string]interface{}, error) {
	user, err := u.UserInfo()
	if err != nil {
		return nil, err
	}

	servicesMenu := NewServicesMenu()
	getTree, err := servicesMenu.GetTree()
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	data["userInfo"] = user
	data["menus"] = getTree

	return data, nil
}
