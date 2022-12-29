package common

type Model struct {
	Id        int64 `json:"id" form:"id" gorm:"primarykey"`
	OptUser   int64 `json:"opt_user" form:"opt_user" gorm:"opt_user"`
	CreatedAt int64 `json:"created_at" form:"created_at" gorm:"created_at"`
	UpdatedAt int64 `json:"updated_at" form:"updated_at" gorm:"updated_at"`
	DeletedAt int64 `json:"deleted_at" form:"deleted_at" gorm:"deleted_at"`
}
