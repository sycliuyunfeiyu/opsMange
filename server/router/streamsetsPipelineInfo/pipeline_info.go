package streamsetsPipelineInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PipelineInfoRouter struct{}

// InitPipelineInfoRouter 初始化 流水线信息 路由信息
func (s *PipelineInfoRouter) InitPipelineInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	pipelineInfoRouter := Router.Group("pipelineInfo").Use(middleware.OperationRecord())
	pipelineInfoRouterWithoutRecord := Router.Group("pipelineInfo")
	pipelineInfoRouterWithoutAuth := PublicRouter.Group("pipelineInfo")
	{
		pipelineInfoRouter.POST("createPipelineInfo", pipelineInfoApi.CreatePipelineInfo)             // 新建流水线信息
		pipelineInfoRouter.DELETE("deletePipelineInfo", pipelineInfoApi.DeletePipelineInfo)           // 删除流水线信息
		pipelineInfoRouter.DELETE("deletePipelineInfoByIds", pipelineInfoApi.DeletePipelineInfoByIds) // 批量删除流水线信息
		pipelineInfoRouter.PUT("updatePipelineInfo", pipelineInfoApi.UpdatePipelineInfo)              // 更新流水线信息
		pipelineInfoRouter.PUT("rsyncPipelineInfo", pipelineInfoApi.RsyncPipelineInfo)                // 更新流水线信息

	}
	{
		pipelineInfoRouterWithoutRecord.GET("findPipelineInfo", pipelineInfoApi.FindPipelineInfo)       // 根据ID获取流水线信息
		pipelineInfoRouterWithoutRecord.GET("getPipelineInfoList", pipelineInfoApi.GetPipelineInfoList) // 获取流水线信息列表
	}
	{
		pipelineInfoRouterWithoutAuth.GET("getPipelineInfoPublic", pipelineInfoApi.GetPipelineInfoPublic) // 获取流水线信息列表
	}
}
