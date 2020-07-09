package ff_page

type Page struct {
	CurOffset  int `json:"cur_offset"`  //当前页从第几条开始查询
	PreOffset  int `json:"pre_offset"`  //上一页从第几条开始查询
	NextOffset int `json:"next_offset"` //下一页从第几条开始查询
	TotalSize  int `json:"total_size"`  //数据表总共有多少条记录
	PageSize   int `json:"page_size"`   //一页多少条记录
}

func NewPage(offset int, pageSize int, totalSize int) *Page {
	page := &Page{}
	page.TotalSize = totalSize
	page.PageSize = pageSize
	page.CurOffset = offset
	if offset > pageSize { //满足这个条件应该是第二页数据了
		page.PreOffset = offset - pageSize
	}
	page.NextOffset = offset
	if offset < totalSize { //满足这个条件说明后面还有数据
		page.NextOffset = offset + pageSize
	}
	return page
}

func (p *Page) SetTotalSize(totalSize int) {
	p.TotalSize = totalSize
	if p.CurOffset < totalSize {
		p.NextOffset = p.CurOffset + p.PageSize
	}
}
