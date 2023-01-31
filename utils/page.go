package utils

type PageData struct {
	PageNo    uint64      `json:"pageNo"`    // 页号
	PageSize  uint64      `json:"pageSize"`  // 页数
	TotalRow  uint64      `json:"totalRow"`  // 总行数
	TotalPage uint64      `json:"totalPage"` // 总页数
	Rows      interface{} `json:"rows"`      // 当前数据
}
