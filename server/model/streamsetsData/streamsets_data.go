// 自动生成模板StreamsetsData
package streamsetsData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// streamsets数据源 结构体  StreamsetsData
type StreamsetsData struct {
    global.GVA_MODEL
    StreamsetsName  string `json:"streamsetsName" form:"streamsetsName" gorm:"uniqueIndex;column:streamsetsName;comment:streamsets 名称;" binding:"required"`  //streamsets名称 
    StreamsetsUuid  string `json:"streamsetsUuid" form:"streamsetsUuid" gorm:"uniqueIndex;column:streamsetsUuid;comment:唯一uuid;"`  //uuid 
    Url  string `json:"url" form:"url" gorm:"column:url;comment:;" binding:"required"`  //地址 
    User  string `json:"user" form:"user" gorm:"column:user;comment:用户名;" binding:"required"`  //用户名 
    Passwd  string `json:"passwd" form:"passwd" gorm:"column:passwd;comment:;" binding:"required"`  //密码 
    Token  string `json:"token" form:"token" gorm:"column:token;comment:;"`  //token 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName streamsets数据源 StreamsetsData自定义表名 streamsetsData
func (StreamsetsData) TableName() string {
    return "streamsetsData"
}

