package conmon

type Format struct {
	CreateDate    string `json:"create_date,omitempty"`
	UpdateDate    string `json:"update_date,omitempty"`
	LastLoginDate string `json:"last_login_date,omitempty"`
	OptUserName   string `json:"opt_user_name,omitempty"`
	GroupName     string `json:"group_name,omitempty"`
}
