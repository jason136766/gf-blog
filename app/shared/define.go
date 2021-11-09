package shared

import "github.com/gogf/gf/database/gdb"

type PageReq struct {
	Page     int `p:"page" d:"1" v:"integer|min:1"`
	PageSize int `p:"page_size" d:"5" v:"integer|in:5,15,50"`
}

type PageRes struct {
	Page     int        `json:"page"`
	PageSize int        `json:"page_size"`
	Count    uint64     `json:"count"`
	Result   gdb.Result `json:"result"`
}
