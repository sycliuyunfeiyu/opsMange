// 自动生成模板PipelineInfo
package streamsetsPipelineInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 流水线信息 结构体  PipelineInfo
type PipelineInfo struct {
	global.GVA_MODEL
	StreamsetsUuid string     `json:"streamsetsUuid" form:"streamsetsUuid" gorm:"index;column:streamsetsUuid;comment:外键关联数据源;"`                  //关联sdcuuid
	PipelineUuid   string     `json:"pipelineUuid" form:"pipelineUuid" gorm:"index;column:pipelineUuid;comment:唯一uuid;"`                         //流水线uuid
	Title_name     string     `json:"title_name" form:"title_name" gorm:"column:title_name;comment:;"`                                           //流水线名称
	Pipeline_id    string     `json:"pipeline_id" form:"pipeline_id" gorm:"uniqueIndex;column:pipeline_id;comment:流水线sdcId;" binding:"required"` //流水线id
	Last_updated   *time.Time `json:"last_updated" form:"last_updated" gorm:"column:last_updated;comment:;"`                                     //last_updated
	Status         string     `json:"status" form:"status" gorm:"column:status;comment:流水线状态;"`                                                  //流水线状态
	Message        string     `json:"message" form:"message" gorm:"column:message;comment:信息;type:text;"`                                        //信息
	Link           string     `json:"link" form:"link" gorm:"column:link;comment:流水线链接;"`                                                        //流水线链接
	CreatedBy      uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 流水线信息 PipelineInfo自定义表名 pipelineInfo
func (PipelineInfo) TableName() string {
	return "pipelineInfo"
}

type PipelineSdcInfo struct {
	PipelineId   string `json:"pipelineId" form:"pipelineId"`
	Name         string `json:"title" form:"title"`
	LastModified int    `json:"lastModified" form:"lastModified"`
	SdcId        string `json:"sdcId" form:"sdcId"`
}

type PipelineSdcStatusInfo map[string]struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	NextRunTime int    `json:"nextRetryTimeStamp"`
}
