package system

import (
	"crypto/md5"
	"errors"
	"fmt"
	"game.sdk.center/internal/mapping"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/internal/model/system"
	"game.sdk.center/internal/services/conmon"
	"game.sdk.center/tool"
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
	fmt.Printf("%#v\n", u)
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

func (u *ServiceUser) Create() error {
	if u.Account == "" {
		return errors.New("账号不能为空")
	}
	if u.Password == "" {
		return errors.New("密码不能为空")
	}
	if len(u.Password) < 6 {
		return errors.New("密码长度不能小于6位")
	}

	user, err := u.User.Get()
	if err != nil {
		return err
	}
	if user.Account != "" {
		return errors.New("账号已存在")
	}
	if u.GroupId <= 0 {
		return errors.New("账号归属分组错误")
	}

	if u.Nickname == "" {
		u.Nickname = u.Account
	}
	u.Salt = tool.Salt()
	u.Password = fmt.Sprintf("%x", md5.Sum([]byte(u.Password+u.Salt)))
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()

	return u.User.Create()
}

func (u *ServiceUser) Update() error {

	if u.Id <= 0 {
		return errors.New("Id无效")
	}
	if u.Status < 0 || u.Status > 2 {
		return errors.New("状态无效")
	}

	user := system.User{
		Nickname:      u.Nickname,
		Status:        u.Status,
		GroupId:       u.GroupId,
		Avatar:        u.Avatar,
		LastLoginIp:   u.LastLoginIp,
		LastLoginTime: u.LastLoginTime,
		Remarks:       u.Remarks,
	}
	u.User = user

	return u.User.Update()
}

func (u *ServiceUser) Remove() error {
	if u.Id <= 0 {
		return errors.New("Id无效")
	}

	return u.User.Remove()
}
