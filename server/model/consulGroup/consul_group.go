// 自动生成模板ConsulGroup
package consulGroup

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 监控数据源consul 结构体  ConsulGroup
type ConsulGroup struct {
	global.GVA_MODEL
	GroupConsulUrl string `json:"groupConsulUrl" form:"groupConsulUrl" gorm:"column:group_consul_url;comment:组consul地址;size:100;" binding:"required"` //组consul地址
	GroupInfo      string `json:"groupInfo" form:"groupInfo" gorm:"column:group_info;comment:;" binding:"required"`                                   //groupInfo字段
	GroupName      string `json:"groupName" form:"groupName" gorm:"column:group_name;comment:组名称;size:100;" binding:"required"`                       //组名称
	GroupToken     string `json:"groupToken" form:"groupToken" gorm:"column:group_token;comment:组token;size:100;" binding:"required"`                 //组token
	GroupUuid      string `json:"groupUuid" form:"groupUuid" gorm:"uniqueIndex;column:group_uuid;comment:唯一id;size:100;" binding:"required"`          //唯一id
	Remake         string `json:"remake" form:"remake" gorm:"column:remake;comment:备注;size:255;"`                                                     //备注
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`
	//Data           []consulData.ConsulData
}

// TableName 监控数据源consul ConsulGroup自定义表名 consul_group
func (ConsulGroup) TableName() string {
	return "consul_group"
}
