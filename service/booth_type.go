package service

import (
	"errors"
	"github.com/youlingdada/street-stall/model"
	"github.com/youlingdada/street-stall/query"
	"github.com/youlingdada/street-stall/utils"
	"github.com/youlingdada/street-stall/vo"
	"log"
	"time"
)

type boothTypeService struct{}

var BoothTypeService boothTypeService

// Add 添加摊位类型信息
func (bts boothTypeService) Add(bType query.BoothTypeAddQuery) error {
	var boothType model.BoothType
	boothType.CreateAt = time.Now()
	boothType.UpdateAt = time.Now()
	boothType.Status = 1
	boothType.Description = bType.Description
	boothType.Description = bType.Name
	if err := model.BoothTypeModel.Add(boothType); err != nil {
		log.Printf("添加摊位类型失败,err->%s\n", err)
		return errors.New("添加摊位类型失败")
	}
	return nil
}

// Edit 编辑摊位类型信息
func (bts boothTypeService) Edit(bType query.BoothTypeEditQuery) error {
	var boothType model.BoothType
	boothType.UpdateAt = time.Now()
	boothType.Description = bType.Description
	boothType.Description = bType.Name
	if err := model.BoothTypeModel.Edit(boothType); err != nil {
		log.Printf("更新摊位类型失败,err->%s\n", err)
		return errors.New("更新摊位类型失败")
	}
	return nil
}

// Detail 获取摊位列表详情
func (bts boothTypeService) Detail(id uint32) (model.BoothType, error) {
	var bType model.BoothType
	if bType, err := model.BoothTypeModel.Detail(id); err != nil {
		log.Printf("获取摊位类型详情失败,err->%s\n", err)
		return bType, errors.New("获取摊位类型详情失败")
	}
	return bType, nil
}

// Page 获取摊位列表
func (bts boothTypeService) Page(pageNo, pageSize uint64) (utils.PageData, error) {
	pageData, err := model.BoothTypeModel.Page((pageNo-1)*pageSize, pageSize)
	if err != nil {
		log.Printf("获取摊位列表失败,err->%s\n", err)
		return pageData, errors.New("获取摊位列表失败")
	}
	// 处理类型的图标id
	var ids []uint64
	boothTypes := pageData.Rows.([]model.BoothType)
	for _, v := range boothTypes {
		ids = append(ids, v.IconId)
	}
	listByIdsMap, err := FileService.ListByIdsMap(ids)
	if err != nil {
		return pageData, err
	}
	var rows []vo.BoothTypeVo
	for _, t := range boothTypes {
		var item vo.BoothTypeVo
		item.Name = t.Name
		item.TId = t.TId
		item.Status = t.Status
		item.Description = t.Description
		fileVo, ok := listByIdsMap[t.IconId]
		if ok {
			item.IconUrl = fileVo.FileUrl
		}
		rows = append(rows, item)
	}
	pageData.Rows = rows
	return pageData, nil
}
