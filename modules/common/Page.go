package common

import (
	"biligo/util"

	"github.com/gin-gonic/gin"
)

const (
	// 页数参数名
	pageNoName = "pageNo"
	// 每页条数参数名
	pageSizeName = "pageSize"

	// 默认页数
	pageNoDefault = "1"
	// 默认每页条数
	pageSizeDefault = "10"
)

// Pagination - 分页组件对象
type Pagination struct {
	PageNo    int           `json:"pageNo"`
	PageSize  int           `json:"pageSize"`
	TotalSize int           `json:"totalSize"`
	TotalPage int           `json:"totalPage"`
	List      []interface{} `json:"list"`
}

// NewPage - 初始化一个分页对象
func NewPage(pageNo int, pageSize int) *Pagination {
	pagination := Pagination{PageNo: pageNo, PageSize: pageSize}
	return &pagination
}

// NewPageFromGin - 从 gin.Context 初始化一个分页对象
func NewPageFromGin(c *gin.Context) *Pagination {
	pageNo, _ := c.Params.Get(pageNoName)
	pageSize, _ := c.Params.Get(pageSizeName)
	if !util.IsInt(pageNo) {
		pageNo = pageNoDefault
	}
	if !util.IsInt(pageSize) {
		pageSize = pageSizeDefault
	}
	return NewPage(util.ToInt(pageNo), util.ToInt(pageSize))
}

// SetTotalSize - 设置总数，并计算出总页数
func (pagination *Pagination) SetTotalSize(totalSize int) {
	pagination.TotalSize = totalSize
	pagination.TotalPage = (totalSize + pagination.PageSize - 1) / totalSize
}
