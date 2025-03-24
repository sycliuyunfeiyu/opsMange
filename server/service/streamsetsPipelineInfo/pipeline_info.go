package streamsetsPipelineInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsPipelineInfo"
	streamsetsPipelineInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsPipelineInfo/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PipelineInfoService struct{}

// CreatePipelineInfo 创建流水线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pipelineInfoService *PipelineInfoService) CreatePipelineInfo(pipelineInfo *streamsetsPipelineInfo.PipelineInfo) (err error) {
	err = global.GVA_DB.Create(pipelineInfo).Error
	return err
}

func (pipelineInfoService *PipelineInfoService) ApplyPipelineInfo(pipelineInfoList *[]streamsetsPipelineInfo.PipelineInfo) (err error) {

	err = global.GVA_DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pipeline_id"}, {Name: "streamsetsUuid"}},
		DoUpdates: clause.AssignmentColumns([]string{"title_name", "status", "message", "link", "last_updated"}),
		UpdateAll: false,
	}).Create(&pipelineInfoList).Error
	return err
}

// DeletePipelineInfo 删除流水线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pipelineInfoService *PipelineInfoService) DeletePipelineInfo(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&streamsetsPipelineInfo.PipelineInfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&streamsetsPipelineInfo.PipelineInfo{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePipelineInfoByIds 批量删除流水线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pipelineInfoService *PipelineInfoService) DeletePipelineInfoByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&streamsetsPipelineInfo.PipelineInfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&streamsetsPipelineInfo.PipelineInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePipelineInfo 更新流水线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pipelineInfoService *PipelineInfoService) UpdatePipelineInfo(pipelineInfo streamsetsPipelineInfo.PipelineInfo) (err error) {
	err = global.GVA_DB.Model(&streamsetsPipelineInfo.PipelineInfo{}).Where("id = ?", pipelineInfo.ID).Updates(&pipelineInfo).Error
	return err
}

// GetPipelineInfo 根据ID获取流水线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pipelineInfoService *PipelineInfoService) GetPipelineInfo(ID string) (pipelineInfo streamsetsPipelineInfo.PipelineInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pipelineInfo).Error
	return
}

// GetPipelineInfoInfoList 分页获取流水线信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (pipelineInfoService *PipelineInfoService) GetPipelineInfoInfoList(info streamsetsPipelineInfoReq.PipelineInfoSearch) (list []streamsetsPipelineInfo.PipelineInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&streamsetsPipelineInfo.PipelineInfo{})
	var pipelineInfos []streamsetsPipelineInfo.PipelineInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StreamsetsUuid != "" {
		db = db.Where("streamsetsUuid = ?", info.StreamsetsUuid)
	}
	if info.PipelineUuid != "" {
		db = db.Where("pipelineUuid = ?", info.PipelineUuid)
	}
	if info.Title_name != "" {
		db = db.Where("title_name LIKE ?", "%"+info.Title_name+"%")
	}
	if info.Pipeline_id != "" {
		db = db.Where("pipeline_id LIKE ?", "%"+info.Pipeline_id+"%")
	}
	if info.StartLast_updated != nil && info.EndLast_updated != nil {
		db = db.Where("last_updated BETWEEN ? AND ? ", info.StartLast_updated, info.EndLast_updated)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&pipelineInfos).Error
	return pipelineInfos, total, err
}
