package streamsetsPipelineInfo

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PipelineInfoRouter }

var pipelineInfoApi = api.ApiGroupApp.StreamsetsPipelineInfoApiGroup.PipelineInfoApi
