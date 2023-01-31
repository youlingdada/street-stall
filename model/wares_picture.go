// Package model /*
package model

type WaresPicture struct {
	PId     uint64 `json:"pId" gorm:"column:p_id;primaryKey"` // 图片id
	FileId  string `json:"fileId" gorm:"column:file_id"`      // 图片文件id
	Status  uint8  `json:"status" gorm:"column:status"`       // 图片状态 1 正常 2 禁用
	WaresId uint64 `json:"waresId" gorm:"column:wares_id"`    // 商品id
	Model
}

type waresPictureModel struct{}

var WaresPictureModel waresPictureModel

// Add 添加商品图
func (wpm waresPictureModel) Add(picture WaresPicture) error {
	return DB.Model(&WaresPicture{}).Create(&picture).Error
}

// Detail 商品图详情
func (wpm waresPictureModel) Detail(pId uint64) (WaresPicture, error) {
	var wp WaresPicture
	tx := DB.Model(&WaresPicture{}).Where("p_id", pId).Find(&wp)
	return wp, tx.Error
}

// ListByWares 查询指定商品的所有商品图
func (wpm waresPictureModel) ListByWares(wId uint64) ([]WaresPicture, error) {
	var wps []WaresPicture
	tx := DB.Model(&WaresPicture{}).Where("wares_id", wId)
	tx = tx.Where("status", 1).Order("create_at aes").Find(&wps)
	return wps, tx.Error
}

// Del 删除指定的商品图
func (wpm waresPictureModel) Del(pId uint64) error {
	var wp WaresPicture
	wp.PId = pId
	return DB.Model(&WaresPicture{}).Delete(&wp).Error
}
