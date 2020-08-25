/*
@Time : 2020/7/16 23:44
@Author : xuyiqing
@File : common.go
*/

package serializers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pager struct {
	Page int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
	OffSet int `json:"-"`
	Total int `json:"total"`
	MaxPage int `json:"maxPage"`
}

func (p *Pager) InitPager(ctx *gin.Context) {
	p.Page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	p.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	p.OffSet = (p.Page - 1) * p.PageSize
}

func (p *Pager) GetPager() {
	p.MaxPage = int(p.Total / p.PageSize) + 1
}
