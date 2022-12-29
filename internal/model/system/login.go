package system

import (
	"errors"
)

type Loginer interface {
	check() error
	login() error
	logout() error
}

type LoginParams struct {
	Account
	Mobile
}

func NewLoginParams() *LoginParams {
	return &LoginParams{}
}

// Login 登录控制
func Login(l Loginer) error {
	if err := l.check(); err != nil {
		return err
	}
	if err := l.login(); err != nil {
		return err
	}
	return nil
}

// Logout 注销登录
func Logout(l Loginer) error {
	if err := l.check(); err != nil {
		return err
	}
	if err := l.login(); err != nil {
		return err
	}
	return nil
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
		return errors.New("账号不能为空")
	}

	return nil
}

func (a *Account) login() error {
	return nil
}

func (a *Account) logout() error {
	return nil
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

func (m *Mobile) login() error {
	if m.Phone <= 0 {
		return errors.New("手机号不能为空")
	}

	return nil
}

func (m *Mobile) logout() error {
	return nil
}
