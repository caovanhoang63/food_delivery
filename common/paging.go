package common

type Paging struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total" form:"total"`
}

func (p *Paging) FullFill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 50
	}
}

func (p *Paging) GetOffSet() int {
	return (p.Page - 1) * p.Limit
}
