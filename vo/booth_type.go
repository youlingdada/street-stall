package vo

type BoothTypeVo struct {
	TId         uint32 `json:"tId"`         // 类型id
	Name        string `json:"name"`        // 类型名称
	Description string `json:"description"` // 类型描述
	IconUrl     string `json:"iconUrl"`     // 图标文件id
	Status      uint8  `json:"status"`      // 类型状态 1 正常 2 下架
}
