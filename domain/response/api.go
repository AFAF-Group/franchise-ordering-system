package response

import (
	"afaf-group.com/pkg/common"
	"github.com/biezhi/gorm-paginator/pagination"
)

type APIResponse struct {
	Code     int         `json:"code,omitempty"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	PageInfo interface{} `json:"page_info,omitempty"`
	Errors   interface{} `json:"errors,omitempty"`
} // @name APIResponse

type PageInfo struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int   `json:"total_page"`
	Offset      int   `json:"offset"`
	Limit       *int  `json:"limit"`
	Page        int   `json:"page"`
	PrevPage    *int  `json:"prev_page"`
	NextPage    *int  `json:"next_page"`
} // @name PageInfo

func NewPageInfo() *PageInfo {
	return &PageInfo{}
}

func (PageInfo) SetPageInfo(paginator *pagination.Paginator) *PageInfo {
	pageInfo := new(PageInfo)
	if paginator.Limit < 0 {
		pageInfo.Limit = nil
		pageInfo.NextPage = nil
	} else {
		pageInfo.Limit = &paginator.Limit
	}

	if paginator.TotalPage < 0 {
		pageInfo.TotalPage = 1
	} else {
		pageInfo.TotalPage = paginator.TotalPage
	}

	if pageInfo.TotalPage == 1 {
		pageInfo.PrevPage = nil
		pageInfo.NextPage = nil
	} else {
		pageInfo.PrevPage = &paginator.PrevPage
		pageInfo.NextPage = &paginator.NextPage
	}
	pageInfo.TotalRecord = int64(paginator.TotalRecord)
	pageInfo.Offset = paginator.Offset
	pageInfo.Page = paginator.Page
	return pageInfo
}

func (PageInfo) ToPageInfo(paginator *common.Pagination) *PageInfo {
	pageInfo := new(PageInfo)
	if paginator.Limit < 0 {
		pageInfo.Limit = nil
		pageInfo.NextPage = nil
	} else {
		pageInfo.Limit = &paginator.Limit
	}

	if paginator.TotalPage < 0 {
		pageInfo.TotalPage = 1
	} else {
		pageInfo.TotalPage = paginator.TotalPage
	}

	if pageInfo.TotalPage == 1 {
		pageInfo.PrevPage = nil
		pageInfo.NextPage = nil
	} else {
		pageInfo.PrevPage = &paginator.PrevPage
		pageInfo.NextPage = &paginator.NextPage
	}
	pageInfo.TotalRecord = paginator.TotalRecord
	pageInfo.Offset = paginator.Offset
	pageInfo.Page = paginator.Page
	return pageInfo
}
