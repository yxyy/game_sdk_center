package system

import (
	"crypto/md5"
	"errors"
	"fmt"
	"game.sdk.center/tool"
	"gorm.io/gorm"
)

type Loginer interface {
	check() error
	login() (User, error)
}

type LoginParams struct {
	*Account
	*Mobile
}

func NewLoginParams() *LoginParams {
	return &LoginParams{}
}

// Login 登录控制
func Login(l Loginer) (user User, err error) {
	if err = l.check(); err != nil {
		return
	}
	user, err = l.login()
	if err != nil {
		return
	}
	return
}

// Account 账号登录
type Account struct {
	Account  string `json:"account" form:"account" gorm:"account"`
	Password string `json:"password" form:"password" gorm:"password"`
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) check() error {
	if a.Account == "" {
		return errors.New("账号不能为空")
	}
	if a.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}

func (a *Account) login() (user User, err error) {
	if err = tool.MysqlDb.Model(&user).Where("account", a.Account).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("账号不存在")
		}
		return user, err
	}

	a.Password = fmt.Sprintf("%x", md5.Sum([]byte(a.Password+user.Salt)))
	if a.Password != user.Password {
		return user, errors.New("密码错误")
	}

	return
}

// Mobile 手机登录
type Mobile struct {
	Phone int `json:"phone" form:"phone" gorm:"phone"`
	Code  int `json:"code" form:"code"`
}

func NewMobile() *Mobile {
	return &Mobile{}
}
func (m *Mobile) check() error {
	if m.Phone <= 0 {
		return errors.New("账号不能为空")
	}
	if m.Code <= 0 {
		return errors.New("账号不能为空")
	}

	return nil
}

func (m *Mobile) login() (user User, err error) {
	if m.Phone <= 0 {
		return user, errors.New("手机号不能为空")
	}

	return
}
