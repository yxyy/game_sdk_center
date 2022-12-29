package common

type Params struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
	Offse int
}

func NewParams() *Params {
	return &Params{
		Page:  1,
		Limit: 10,
	}
}

func (p *Params) Check() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
	p.Offse = (p.Page - 1) * p.Limit
}
