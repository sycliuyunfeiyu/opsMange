package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StreamsetsDataSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StreamsetsName  string `json:"streamsetsName" form:"streamsetsName" `
    StreamsetsUuid  string `json:"streamsetsUuid" form:"streamsetsUuid" `
    Url  string `json:"url" form:"url" `
    User  string `json:"user" form:"user" `
    request.PageInfo
}
