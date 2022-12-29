package system

import (
	"crypto/md5"
	"errors"
	"fmt"
	"game.sdk.center/internal/model/common"
	"game.sdk.center/tool"
	"gorm.io/gorm"
	"time"
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
	GroupId       int64  `json:"group_id" form:"group_id" gorm:"group_id"`
	Avatar        string `json:"avatar" form:"avatar" gorm:"avatar"`
	Status        int64  `json:"status" form:"status" gorm:"status"`
	LastLoginIp   string `json:"last_login_ip" form:"last_login_ip" gorm:"last_login_ip"`
	LastLoginTime string `json:"last_login_time" form:"last_login_time" gorm:"last_login_time"`
	Remarks       string `json:"remarks" form:"remarks" gorm:"remarks"`
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

func (u *User) List(params common.Params) (user []*User, err error) {
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
		tx = tx.Where("nickname", "like", "%"+u.Nickname+"%")
	}

	if err = tx.Offset(params.Offse).Limit(params.Limit).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
	}

	return
}

func (u *User) Create() error {
	if u.Account == "" {
		return errors.New("账号不能为空")
	}
	if u.Password == "" {
		return errors.New("密码不能为空")
	}
	if len(u.Password) < 6 {
		return errors.New("密码长度不能小于6位")
	}
	user := NewUser()
	if err := tool.MysqlDb.Model(&user).Where("account", u.Account).Where("status", 0).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
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

	if err := tool.MysqlDb.Model(&u).Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Update() error {

	if u.Id <= 0 {
		return errors.New("Id无效")
	}
	if u.Status < 0 || u.Status > 2 {
		return errors.New("状态无效")
	}
	err := tool.MysqlDb.Model(&u).Where("id", u.Id).Updates(User{
		Nickname: u.Nickname,
		Status:   u.Status,
		GroupId:  u.GroupId,
		Avatar:   u.Avatar,
		Remarks:  u.Remarks,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Remove() error {
	if u.Id <= 0 {
		return errors.New("Id无效")
	}
	if err := tool.MysqlDb.Model(&u).Where("id", u.Id).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
