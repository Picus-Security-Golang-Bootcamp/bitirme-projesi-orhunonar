package pagination

// Pagination when listing data
type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

func (p *Pagination) SetPagination(page, limit, totalData int) {
	p.Page = page
	p.Limit = limit
	p.TotalData = totalData
	p.TotalPage = p.TotalData / p.Limit
	if p.TotalData%p.Limit != 0 {
		p.TotalPage++
	}

}
