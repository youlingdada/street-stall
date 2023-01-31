package query

// BoothAddQuery 摊位增加请求模型
// @Description 摊位增加请求模型
type BoothAddQuery struct {
	Type      uint32  `json:"type"`      // 摊位类型
	Address   string  `json:"address"`   // 摊位地址
	Status    uint8   `json:"status"`    // 摊位状态
	Name      string  `json:"name"`      // 摊位名称
	UserId    uint64  `json:"userId"`    // 用户id
	Longitude float64 `json:"longitude"` // 经度
	Latitude  float64 `json:"latitude"`  // 维度
}

// EditBoothQuery 编辑摊位请求模型
// @Description 编辑摊位请求模型
type EditBoothQuery struct {
	BId       uint64  `json:"bId"`       //摊位id
	Type      uint32  `json:"type"`      // 摊位类型
	Address   string  `json:"address"`   // 摊位地址
	Name      string  `json:"name"`      // 摊位名称
	Longitude float64 `json:"longitude"` // 经度
	Latitude  float64 `json:"latitude"`  // 维度
}

// PageBoothQuery 分页获取摊位请求模型
//
//	@Description 分页获取摊位请求模型
type PageBoothQuery struct {
	PageNo   uint64 `json:"pageNo"`   // 页号
	PageSize uint64 `json:"pageSize"` // 页数
	Status   uint8  `json:"status"`   // 状态
	Type     uint8  `json:"type"`     // 类型
	UserId   uint64 `json:"userId"`   // 摊主
}
