package utils

import (
	"fmt"
	"math"
	"strconv"

	restful "github.com/emicklei/go-restful"
)

// Paginate 提供与分页相关的帮助方法
type Paginate struct {
	Req *restful.Request

	DefPerPage uint
}

// Page return current page
func (p Paginate) Page() uint {
	page, err := strconv.ParseUint(p.Req.QueryParameter("page"), 10, 0)
	if err != nil || page == 0 {
		page = 1
	}
	return uint(page)
}

// Limit return query limit
func (p Paginate) Limit() uint {
	if p.DefPerPage == 0 {
		p.DefPerPage = 10
	}
	perPage, err := strconv.ParseUint(p.Req.QueryParameter("perpage"), 10, 0)
	if err != nil || perPage == 0 {
		perPage = uint64(p.DefPerPage)
	}
	return uint(perPage)
}

// Offset return current offset
func (p Paginate) Offset() uint {
	return (p.Page() - 1) * p.Limit()
}

// WriterHeader writer response header
func (p Paginate) WriterHeader(total uint, resp *restful.Response) {
	pageMax := math.Ceil(float64(total) / float64(p.Limit()))
	resp.AddHeader("X-Total-Count", fmt.Sprintf("%d", total))
	resp.AddHeader("X-Total-Page", fmt.Sprintf("%d", int(pageMax)))
	resp.AddHeader("X-Current-Page", fmt.Sprintf("%d", p.Page()))
	resp.AddHeader("X-Current-PerPage", fmt.Sprintf("%d", p.Limit()))
}
