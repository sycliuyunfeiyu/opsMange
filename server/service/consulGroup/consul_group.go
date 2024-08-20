package consulGroup

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulGroup"
	consulGroupReq "github.com/flipped-aurora/gin-vue-admin/server/model/consulGroup/request"
	"gorm.io/gorm"
)

type ConsulGroupService struct{}

// CreateConsulGroup 创建监控数据源consul记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulGService *ConsulGroupService) CreateConsulGroup(consulG *consulGroup.ConsulGroup) (err error) {
	err = global.GVA_DB.Create(consulG).Error
	return err
}

// DeleteConsulGroup 删除监控数据源consul记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulGService *ConsulGroupService) DeleteConsulGroup(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&consulGroup.ConsulGroup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&consulGroup.ConsulGroup{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteConsulGroupByIds 批量删除监控数据源consul记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulGService *ConsulGroupService) DeleteConsulGroupByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&consulGroup.ConsulGroup{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&consulGroup.ConsulGroup{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateConsulGroup 更新监控数据源consul记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulGService *ConsulGroupService) UpdateConsulGroup(consulG consulGroup.ConsulGroup) (err error) {
	err = global.GVA_DB.Model(&consulGroup.ConsulGroup{}).Where("id = ?", consulG.ID).Updates(&consulG).Error
	return err
}

// GetConsulGroup 根据ID获取监控数据源consul记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulGService *ConsulGroupService) GetConsulGroup(ID string) (consulG consulGroup.ConsulGroup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&consulG).Error
	return
}
func (consulGService *ConsulGroupService) GetConsulGroupUuid(Uuid string) (consulG consulGroup.ConsulGroup, err error) {
	err = global.GVA_DB.Where("group_uuid = ?", Uuid).First(&consulG).Error
	return
}

// GetConsulGroupInfoList 分页获取监控数据源consul记录
// Author [piexlmax](https://github.com/piexlmax)
func (consulGService *ConsulGroupService) GetConsulGroupInfoList(info consulGroupReq.ConsulGroupSearch) (list []consulGroup.ConsulGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&consulGroup.ConsulGroup{})
	var consulGs []consulGroup.ConsulGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.GroupConsulUrl != "" {
		db = db.Where("group_consul_url LIKE ?", "%"+info.GroupConsulUrl+"%")
	}
	if info.GroupInfo != "" {
		db = db.Where("group_info LIKE ?", "%"+info.GroupInfo+"%")
	}
	if info.GroupName != "" {
		db = db.Where("group_name LIKE ?", "%"+info.GroupName+"%")
	}
	if info.GroupToken != "" {
		db = db.Where("group_token LIKE ?", "%"+info.GroupToken+"%")
	}
	if info.GroupUuid != "" {
		db = db.Where("group_uuid LIKE ?", "%"+info.GroupUuid+"%")
	}
	if info.Remake != "" {
		db = db.Where("remake LIKE ?", "%"+info.Remake+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&consulGs).Error
	return consulGs, total, err
}
