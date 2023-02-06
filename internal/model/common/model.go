package common

type Model struct {
	Id        int   `json:"id" form:"id" gorm:"primarykey"`
	OptUser   int   `json:"opt_user,omitempty" form:"opt_user" gorm:"opt_user"`
	CreatedAt int64 `json:"created_at,omitempty" form:"created_at" gorm:"created_at"`
	UpdatedAt int64 `json:"updated_at,omitempty" form:"updated_at" gorm:"updated_at"`
	DeletedAt int64 `json:"deleted_at,omitempty" form:"deleted_at" gorm:"deleted_at"`
}
