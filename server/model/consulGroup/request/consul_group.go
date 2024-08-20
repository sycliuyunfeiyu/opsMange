package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ConsulGroupSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    GroupConsulUrl  string `json:"groupConsulUrl" form:"groupConsulUrl" `
    GroupInfo  string `json:"groupInfo" form:"groupInfo" `
    GroupName  string `json:"groupName" form:"groupName" `
    GroupToken  string `json:"groupToken" form:"groupToken" `
    GroupUuid  string `json:"groupUuid" form:"groupUuid" `
    Remake  string `json:"remake" form:"remake" `
    request.PageInfo
}
