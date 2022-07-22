package model

//Page 结构
type Page struct {
	Books       []*Book //每一页所展示的图书，放在一个切片里
	PageNo      int64   //当前页码
	PageSize    int64   //每页图书的数量
	TotalPage   int64   //总页数
	TotalRecord int64   //总记录数
	MinPrice    string  //价格范围
	MaxPrice    string
	IsLogin     bool   //是否登录
	Username    string //用户名
}

//判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPage
}

//获取上一页
func (p *Page) GetPrevPage() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//获取下一页
func (p *Page) GetNextPage() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPage
	}

}
