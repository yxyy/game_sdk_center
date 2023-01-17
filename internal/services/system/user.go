package system

import (
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/internal/services/conmon"
	"time"
)

type ServiceUser struct {
	system.User
	conmon.Format
}

func NewServiceUser() ServiceUser {
	return ServiceUser{}
}

func (u ServiceUser) List(params common.Params) (users []*ServiceUser, total int64, err error) {

	list, total, err := u.User.List(params)
	if err != nil {
		return nil, 0, err
	}

	userMap, err := mapping.User()
	if err != nil {
		return nil, 0, err
	}
	group, err := mapping.Group()
	if err != nil {
		return nil, 0, err
	}
	for _, v := range list {
		tmp := &ServiceUser{
			User: *v,
			Format: conmon.Format{
				CreateDate:    time.Unix(v.CreatedAt, 0).Format("2006-01-02 15:04:05"),
				UpdateDate:    time.Unix(v.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
				LastLoginDate: time.Unix(v.LastLoginTime, 0).Format("2006-01-02 15:04:05"),
				OptUserName:   userMap[v.OptUser],
				GroupName:     group[v.GroupId],
			},
		}

		users = append(users, tmp)
	}

	return
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
