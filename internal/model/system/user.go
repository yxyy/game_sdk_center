package system

import (
	"errors"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
	"gorm.io/gorm"
)

type User struct {
	common.Model
	Account       string `json:"account" form:"account" gorm:""`
	Password      string `json:"-" form:"password" gorm:"password"`
	Salt          string `json:"-" gorm:"salt"`
	Nickname      string `json:"nickname" form:"nickname" gorm:"nickname"`
	Phone         int64  `json:"phone" form:"phone" gorm:"phone"`
	Wechat        string `json:"wechat" form:"wechat" gorm:"wechat"`
	Email         string `json:"email" form:"email" gorm:"email"`
	GroupId       int    `json:"group_id" form:"group_id" gorm:"group_id"`
	Avatar        string `json:"avatar" form:"avatar" gorm:"avatar"`
	Status        int64  `json:"status" form:"status" gorm:"status"`
	LastLoginIp   string `json:"last_login_ip" form:"last_login_ip" gorm:"last_login_ip"`
	LastLoginTime int64  `json:"last_login_time" form:"last_login_time" gorm:"last_login_time"`
	Remarks       string `json:"remarks" form:"remarks" gorm:"remarks"`
}

type TokenInfo struct {
	User
	GroupName    string `json:"group_name"`
	PermissionId int    `json:"permission_id"`
}

func NewTokenInfo() TokenInfo {
	return TokenInfo{}
}

func NewUser() *User {
	return &User{}
}

func (u *User) UserInfo() (user *User, err error) {
	if u.Id <= 0 && u.Account == "" {
		return user, errors.New("参数错误")
	}
	tx := tool.MysqlDb.Model(&u)
	if u.Id > 0 {
		tx = tx.Where("id", u.Id)
	}
	if u.Account != "" {
		tx = tx.Where("account", u.Account)
	}
	if err = tx.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return
}

func (u *User) List(params common.Params) (user []*User, total int64, err error) {
	tx := tool.MysqlDb.Model(&u)
	if u.Id > 0 {
		tx = tx.Where("id", u.Id)
	}
	if u.GroupId > 0 {
		tx = tx.Where("group_id", u.GroupId)
	}
	if u.Status > 0 {
		tx = tx.Where("status", u.Status)
	}
	if u.Phone > 0 {
		tx = tx.Where("phone", u.Phone)
	}
	if u.Account != "" {
		tx = tx.Where("account", u.Account)
	}
	if u.Email != "" {
		tx = tx.Where("email", u.Email)
	}
	if u.Wechat != "" {
		tx = tx.Where("wechat", u.Wechat)
	}
	if u.Nickname != "" {
		tx = tx.Where("nickname like ?", "%"+u.Nickname+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = tx.Offset(params.Offse).Limit(params.Limit).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}

	return
}

func (u *User) All() (users []*User, err error) {

	if err = tool.MysqlDb.Model(&u).Find(&users).Error; err != nil {
		return
	}
	return
}

func (u *User) Create() error {

	return tool.MysqlDb.Model(&u).Create(&u).Error
}

func (u *User) Update() error {

	return tool.MysqlDb.Model(&u).Where("id", u.Id).Updates(&u).Error
}

func (u *User) Remove() error {
	return tool.MysqlDb.Model(&u).Where("id", u.Id).Delete(&u).Error
}

func (u *User) Get() (user User, err error) {

	tx := tool.MysqlDb.Model(&u)
	if u.Id <= 0 {
		tx = tx.Where("id", u.Id)
	}
	if u.Account != "" {
		tx = tx.Where("account", u.Account)
	}
	if err = tx.Find(&user).Error; err != nil {
		return
	}
	return
}
