package consulData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConsulDataRouter struct {}

// InitConsulDataRouter 初始化 consulData表 路由信息
func (s *ConsulDataRouter) InitConsulDataRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	consulDRouter := Router.Group("consulD").Use(middleware.OperationRecord())
	consulDRouterWithoutRecord := Router.Group("consulD")
	consulDRouterWithoutAuth := PublicRouter.Group("consulD")
	{
		consulDRouter.POST("createConsulData", consulDApi.CreateConsulData)   // 新建consulData表
		consulDRouter.DELETE("deleteConsulData", consulDApi.DeleteConsulData) // 删除consulData表
		consulDRouter.DELETE("deleteConsulDataByIds", consulDApi.DeleteConsulDataByIds) // 批量删除consulData表
		consulDRouter.PUT("updateConsulData", consulDApi.UpdateConsulData)    // 更新consulData表
	}
	{
		consulDRouterWithoutRecord.GET("findConsulData", consulDApi.FindConsulData)        // 根据ID获取consulData表
		consulDRouterWithoutRecord.GET("getConsulDataList", consulDApi.GetConsulDataList)  // 获取consulData表列表
	}
	{
	    consulDRouterWithoutAuth.GET("getConsulDataDataSource", consulDApi.GetConsulDataDataSource)  // 获取consulData表数据源
	    consulDRouterWithoutAuth.GET("getConsulDataPublic", consulDApi.GetConsulDataPublic)  // 获取consulData表列表
	}
}
