package consulData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulData"
	consulDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/consulData/request"
	"gorm.io/gorm"
)

type ConsulDataService struct{}

// CreateConsulData 创建consulData表记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulDService *ConsulDataService) CreateConsulData(consulD *consulData.ConsulData) (err error) {
	err = global.GVA_DB.Create(consulD).Error
	return err
}
func (consulDService *ConsulDataService) CreateBatchConsulData(consulD *[]consulData.ConsulData) (err error) {
	err = global.GVA_DB.Create(consulD).Error
	return err
}

// DeleteConsulData 删除consulData表记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulDService *ConsulDataService) DeleteConsulData(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&consulData.ConsulData{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&consulData.ConsulData{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteConsulDataByIds 批量删除consulData表记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulDService *ConsulDataService) DeleteConsulDataByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&consulData.ConsulData{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&consulData.ConsulData{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateConsulData 更新consulData表记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulDService *ConsulDataService) UpdateConsulData(consulD consulData.ConsulData) (err error) {
	err = global.GVA_DB.Model(&consulData.ConsulData{}).Where("id = ?", consulD.ID).Save(&consulD).Error
	return err
}

// GetConsulData 根据ID获取consulData表记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulDService *ConsulDataService) GetConsulData(ID string) (consulD consulData.ConsulData, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&consulD).Error
	return
}

// GetConsulDataInfoList 分页获取consulData表记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulDService *ConsulDataService) GetConsulDataInfoList(info consulDataReq.ConsulDataSearch) (list []consulData.ConsulData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&consulData.ConsulData{})
	var consulDs []consulData.ConsulData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ConsulId != "" {
		db = db.Where("consul_id LIKE ?", "%"+info.ConsulId+"%")
	}
	if info.ConsulName != "" {
		db = db.Where("consul_name LIKE ?", "%"+info.ConsulName+"%")
	}
	if info.ConsulStatus != nil {
		db = db.Where("consul_status = ?", info.ConsulStatus)
	}
	if info.ConsulUrl != "" {
		db = db.Where("consul_url LIKE ?", "%"+info.ConsulUrl+"%")
	}
	if info.ConsulUuid != "" {
		db = db.Where("consul_uuid LIKE ?", "%"+info.ConsulUuid+"%")
	}
	if info.GroupUuid != "" {
		db = db.Where("group_uuid LIKE ?", "%"+info.GroupUuid+"%")
	}
	if info.Remake != "" {
		db = db.Where("remake LIKE ?", "%"+info.Remake+"%")
	}
	if info.ServiceAddress != "" {
		db = db.Where("service_address LIKE ?", "%"+info.ServiceAddress+"%")
	}
	if info.ServiceMeta != "" {
		db = db.Where("service_meta LIKE ?", "%"+info.ServiceMeta+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&consulDs).Error
	return consulDs, total, err
}
func (consulDService *ConsulDataService) GetConsulDataDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	groupUuid := make([]map[string]any, 0)
	global.GVA_DB.Table("consul_group").Select("group_uuid as label,group_uuid as value").Scan(&groupUuid)
	res["groupUuid"] = groupUuid
	return
}
