package query

// WaresAddQuery 添加商品模型
// @Description 添加商品模型
type WaresAddQuery struct {
	WId         uint64  `json:"wId"`         // 商品id
	Name        string  `json:"name"`        // 名称
	Description string  `json:"description"` // 描述
	Type        uint8   `json:"type"`        //类型
	Price       float64 `json:"price"`       // 价格
	Stock       uint32  `json:"stock"`       // 库存
	Status      uint8   `json:"status"`      // 状态
	BoothId     uint64  `json:"boothId"`     // 摊位id
}

// WaresEditQuery 更新商品模型
// @Description 更新商品模型
type WaresEditQuery struct {
	WId    uint64  `json:"wId"`    // 商品id
	Name   string  `json:"name"`   // 名称
	Type   uint8   `json:"type"`   //类型
	Price  float64 `json:"price"`  // 价格
	Stock  uint32  `json:"stock"`  // 库存
	Status uint8   `json:"status"` // 状态
}

// WaresPageQuery 分页获取商品请求模型
// @Description 分页获取商品请求模型
type WaresPageQuery struct {
	PageNo   uint64 `json:"pageNo"`   // 页号
	PageSize uint64 `json:"pageSize"` // 页数
	Status   uint8  `json:"status"`   // 状态
	Type     uint8  `json:"type"`     // 类型
	UserId   uint64 `json:"userId"`   // 摊主
	BoothId  uint64 `json:"boothId"`  // 摊位id
}
