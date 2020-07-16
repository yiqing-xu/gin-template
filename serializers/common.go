/*
@Time : 2020/7/16 23:44
@Author : xuyiqing
@File : common.py
*/

package serializers

type Pagination struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
}

func (p *Pagination) Paginate() {
	offset := p.Page * p.PageSize
}