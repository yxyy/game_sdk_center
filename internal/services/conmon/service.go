package conmon

import (
	"errors"
	"game.sdk.center/internal/model/common"
	"net/url"
	"time"
)

type Format struct {
	CreateDate    string `json:"create_date,omitempty"`
	UpdateDate    string `json:"update_date,omitempty"`
	DeleteDate    string `json:"delete_date,omitempty"`
	LastLoginDate string `json:"last_login_date,omitempty"`
	OptUserName   string `json:"opt_user_name,omitempty"`
	GroupName     string `json:"group_name,omitempty"`
}

func FormatTime(sex int64) string {
	return time.Unix(sex, 0).Format("2006-01-02 15:04:05")
}

// func FormatUser(userId int) string {
// 	userMap, _ := mapping.User()
//
// 	return userMap[userId]
// }

func Formats(model common.Model) (format Format) {
	if model.CreatedAt > 0 {
		format.CreateDate = FormatTime(model.CreatedAt)
	}
	if model.UpdatedAt > 0 {
		format.UpdateDate = FormatTime(model.UpdatedAt)
	}
	if model.DeletedAt > 0 {
		format.DeleteDate = FormatTime(model.DeletedAt)
	}

	// if model.OptUser > 0 {
	// 	format.OptUserName = FormatUser(model.OptUser)
	// }

	return format
}

func ParseUrl(urls string) error {

	uri, err := url.ParseRequestURI(urls)
	if err != nil {
		return err
	}

	if uri.Scheme != "http" && uri.Scheme != "https" {
		return errors.New("协议错误")
	}

	if uri.Host == "" {
		return errors.New("域名错误")
	}

	return nil
}
