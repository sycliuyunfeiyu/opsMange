package streamsetsData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsData"
	streamsetsDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsData/request"
	"gorm.io/gorm"
)

type StreamsetsDataService struct{}

// CreateStreamsetsData 创建streamsets数据源记录
// Author [piexlmax](https://github.com/piexlmax)
func (sdcService *StreamsetsDataService) CreateStreamsetsData(sdc *streamsetsData.StreamsetsData) (err error) {
	err = global.GVA_DB.Create(sdc).Error
	return err
}

// DeleteStreamsetsData 删除streamsets数据源记录
// Author [piexlmax](https://github.com/piexlmax)
func (sdcService *StreamsetsDataService) DeleteStreamsetsData(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&streamsetsData.StreamsetsData{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&streamsetsData.StreamsetsData{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteStreamsetsDataByIds 批量删除streamsets数据源记录
// Author [piexlmax](https://github.com/piexlmax)
func (sdcService *StreamsetsDataService) DeleteStreamsetsDataByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&streamsetsData.StreamsetsData{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&streamsetsData.StreamsetsData{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateStreamsetsData 更新streamsets数据源记录
// Author [piexlmax](https://github.com/piexlmax)
func (sdcService *StreamsetsDataService) UpdateStreamsetsData(sdc streamsetsData.StreamsetsData) (err error) {
	err = global.GVA_DB.Model(&streamsetsData.StreamsetsData{}).Where("id = ?", sdc.ID).Updates(&sdc).Error
	return err
}

// GetStreamsetsData 根据ID获取streamsets数据源记录
// Author [piexlmax](https://github.com/piexlmax)
func (sdcService *StreamsetsDataService) GetStreamsetsData(ID string) (sdc streamsetsData.StreamsetsData, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sdc).Error
	return
}

// GetStreamsetsDataInfoList 分页获取streamsets数据源记录
// Author [piexlmax](https://github.com/piexlmax)
func (sdcService *StreamsetsDataService) GetStreamsetsDataInfoList(info streamsetsDataReq.StreamsetsDataSearch) (list []streamsetsData.StreamsetsData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&streamsetsData.StreamsetsData{})
	var sdcs []streamsetsData.StreamsetsData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StreamsetsName != "" {
		db = db.Where("streamsetsName LIKE ?", "%"+info.StreamsetsName+"%")
	}
	if info.StreamsetsUuid != "" {
		db = db.Where("streamsetsUuid = ?", info.StreamsetsUuid)
	}
	if info.Url != "" {
		db = db.Where("url LIKE ?", "%"+info.Url+"%")
	}
	if info.User != "" {
		db = db.Where("user = ?", info.User)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sdcs).Error
	return sdcs, total, err
}
