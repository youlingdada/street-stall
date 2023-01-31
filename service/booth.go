package service

import (
	"errors"
	"github.com/youlingdada/street-stall/vo"
	"log"
	"time"

	"github.com/youlingdada/street-stall/model"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/utils"
)

type boothService struct{}

var BoothService boothService

// Add 添加摊位
func (bs boothService) Add(query query.BoothAddQuery) error {
	var booth model.Booth
	booth.CreateAt = time.Now()
	booth.UpdateAt = time.Now()
	booth.UserId = query.UserId
	booth.Address = query.Address
	booth.Latitude = query.Latitude
	booth.Longitude = query.Longitude
	booth.Type = query.Type
	booth.Status = query.Status
	booth.Name = query.Name
	tx := model.BoothModel.BoothDB().Create(&booth)
	err := tx.Error
	if err != nil {
		log.Printf("添加摊位失败,err->%s", err)
		return errors.New("添加摊位失败")
	}
	return nil
}

// Edit 编辑摊位信息
func (bs boothService) Edit(query query.EditBoothQuery) error {
	var booth model.Booth
	booth.BId = query.BId
	booth.Address = query.Address
	booth.Latitude = query.Latitude
	booth.Longitude = query.Longitude
	booth.Type = query.Type
	booth.Name = query.Name

	tx := model.BoothModel.BoothDB().Where("b_id", query.BId).Updates(booth)
	if tx.Error != nil {
		log.Printf("编辑摊位信息失败,err->%s\n", tx.Error)
		return tx.Error
	}
	return nil
}

// Detail 摊位详细信息
func (bs boothService) Detail(bId uint64) (model.Booth, error) {
	var booth model.Booth
	tx := model.BoothModel.BoothDB().Where("b_id = ?", bId).First(&booth)
	if tx.Error != nil {
		log.Printf("查询摊位详细信息失败,err->%s\n", tx.Error)
		return booth, tx.Error
	}
	return booth, nil
}

func (bs boothService) List() {

}

// Page 分页查询摊位信息
func (bs boothService) Page(query query.PageBoothQuery) (utils.PageData, error) {
	var data utils.PageData
	var rows []model.Booth
	tx := model.BoothModel.BoothDB()
	if query.UserId != 0 {
		tx.Where("user_id", query.UserId)
	}
	if query.Status != 0 {
		tx.Where("status", 1)
	} else {
		tx.Where("status", 0)
	}
	tx.Offset(int((query.PageNo - 1) * query.PageSize)).Limit(int(query.PageSize))
	var count int64
	tx.Count(&count)

	tx.Find(&rows)

	err := tx.Error
	if err != nil {
		log.Printf("分页查询摊位信息失败,err->%s\n", err)
		return data, errors.New("分页查询摊位信息失败")
	}
	data.TotalRow = uint64(count)
	data.TotalPage = uint64(count)/query.PageSize + func() uint64 {
		if uint64(count)%query.PageSize == 0 {
			return 0
		}
		return 1
	}()
	data.PageNo = query.PageNo
	data.PageSize = query.PageSize
	var ids []uint64
	for _, v := range rows {
		ids = append(ids, v.PosterId)
	}
	idsMap, err := FileService.ListByIdsMap(ids)
	if err != nil {
		return data, err
	}
	var ret []vo.BoothVo
	for _, v := range rows {
		var item vo.BoothVo
		item.BId = v.BId
		item.Name = v.Name
		item.Status = v.Status
		item.Type = v.Type
		item.UserId = v.UserId
		item.UpdateAt = v.UpdateAt
		item.CreateAt = v.CreateAt
		item.Latitude = v.Latitude
		item.Longitude = v.Longitude
		item.Address = v.Address
		fileVo, ok := idsMap[v.PosterId]
		if ok {
			item.PosterUrl = fileVo.FileUrl
		}
		ret = append(ret, item)
	}
	data.Rows = ret
	return data, nil
}
