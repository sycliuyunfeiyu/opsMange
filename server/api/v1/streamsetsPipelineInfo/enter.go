package streamsetsPipelineInfo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PipelineInfoApi }

var pipelineInfoService = service.ServiceGroupApp.StreamsetsPipelineInfoServiceGroup.PipelineInfoService
