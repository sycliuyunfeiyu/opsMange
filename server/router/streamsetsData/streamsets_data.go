package streamsetsData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StreamsetsDataRouter struct {}

// InitStreamsetsDataRouter 初始化 streamsets数据源 路由信息
func (s *StreamsetsDataRouter) InitStreamsetsDataRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sdcRouter := Router.Group("sdc").Use(middleware.OperationRecord())
	sdcRouterWithoutRecord := Router.Group("sdc")
	sdcRouterWithoutAuth := PublicRouter.Group("sdc")
	{
		sdcRouter.POST("createStreamsetsData", sdcApi.CreateStreamsetsData)   // 新建streamsets数据源
		sdcRouter.DELETE("deleteStreamsetsData", sdcApi.DeleteStreamsetsData) // 删除streamsets数据源
		sdcRouter.DELETE("deleteStreamsetsDataByIds", sdcApi.DeleteStreamsetsDataByIds) // 批量删除streamsets数据源
		sdcRouter.PUT("updateStreamsetsData", sdcApi.UpdateStreamsetsData)    // 更新streamsets数据源
	}
	{
		sdcRouterWithoutRecord.GET("findStreamsetsData", sdcApi.FindStreamsetsData)        // 根据ID获取streamsets数据源
		sdcRouterWithoutRecord.GET("getStreamsetsDataList", sdcApi.GetStreamsetsDataList)  // 获取streamsets数据源列表
	}
	{
	    sdcRouterWithoutAuth.GET("getStreamsetsDataPublic", sdcApi.GetStreamsetsDataPublic)  // 获取streamsets数据源列表
	}
}
