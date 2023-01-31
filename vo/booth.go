package vo

import "time"

type BoothVo struct {
	BId       uint64    `json:"bId"`       //摊位id
	Type      uint32    `json:"type"`      // 摊位类型
	Address   string    `json:"address"`   // 摊位地址
	Status    uint8     `json:"status"`    // 摊位状态
	Name      string    `json:"name"`      // 摊位名称
	UserId    uint64    `json:"userId"`    // 用户id
	PosterUrl string    `json:"posterUrl"` // 摊位海报图片id
	Longitude float64   `json:"longitude"` // 经度
	Latitude  float64   `json:"latitude"`  // 维度
	CreateAt  time.Time `json:"createAt"`  // 创建时间
	UpdateAt  time.Time `json:"updateAt"`  // 更新时间
}
