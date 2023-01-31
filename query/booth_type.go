package query

// BoothTypeAddQuery 添加摊位类型模型
type BoothTypeAddQuery struct {
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
	IconId      uint64 `json:"iconId"`      // 图标文件id
}

// BoothTypeEditQuery 编辑摊位类型模型
type BoothTypeEditQuery struct {
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
	IconId      uint64 `json:"iconId"`      // 图标文件id
}
