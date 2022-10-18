package common

type Model struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}
