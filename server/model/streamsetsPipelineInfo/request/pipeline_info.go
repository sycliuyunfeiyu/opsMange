package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PipelineInfoSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StreamsetsUuid  string `json:"streamsetsUuid" form:"streamsetsUuid" `
    PipelineUuid  string `json:"pipelineUuid" form:"pipelineUuid" `
    Title_name  string `json:"title_name" form:"title_name" `
    Pipeline_id  string `json:"pipeline_id" form:"pipeline_id" `
    StartLast_updated  *time.Time  `json:"startLast_updated" form:"startLast_updated"`
    EndLast_updated  *time.Time  `json:"endLast_updated" form:"endLast_updated"`
    Status  string `json:"status" form:"status" `
    request.PageInfo
}
