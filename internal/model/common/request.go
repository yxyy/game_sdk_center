package common

type Params struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func NewParams() *Params {
	return &Params{
		Page:  0,
		Limit: 10,
	}
}
