package ff_page

type Page struct {
	CurOffset  int64
	PreOffset  int64
	NextOffset int64
	TotalSize  int64
	PageSize   int64
}

func NewPage(offset int64, pageSize int64, totalSize int64) *Page {
	page := &Page{}
	page.TotalSize = totalSize
	page.PageSize = pageSize
	page.CurOffset = offset
	if offset > pageSize {
		page.PreOffset = offset - pageSize
	}
	page.NextOffset = offset
	if offset < totalSize {
		page.NextOffset = offset + pageSize
	}
	return page
}

func (p *Page) SetTotalSize(totalSize int64) {
	p.TotalSize = totalSize
	if p.CurOffset < totalSize {
		p.NextOffset = p.CurOffset + p.PageSize
	}
}
