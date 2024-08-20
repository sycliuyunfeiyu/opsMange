package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ConsulDataSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    ConsulId  string `json:"consulId" form:"consulId" `
    ConsulName  string `json:"consulName" form:"consulName" `
    ConsulStatus  *int `json:"consulStatus" form:"consulStatus" `
    ConsulUrl  string `json:"consulUrl" form:"consulUrl" `
    ConsulUuid  string `json:"consulUuid" form:"consulUuid" `
    GroupUuid  string `json:"groupUuid" form:"groupUuid" `
    Remake  string `json:"remake" form:"remake" `
    ServiceAddress  string `json:"serviceAddress" form:"serviceAddress" `
    ServiceMeta  string `json:"serviceMeta" form:"serviceMeta" `
    request.PageInfo
}
