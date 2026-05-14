package types

type Pagination struct {
	Page     uint
	PageSize uint
}

func (p *Pagination) Normalize() {
	if p == nil {
		return
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
