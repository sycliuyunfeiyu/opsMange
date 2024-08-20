package consulGroup

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConsulGroupRouter struct{}

// InitConsulGroupRouter 初始化 监控数据源consul 路由信息
func (s *ConsulGroupRouter) InitConsulGroupRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	consulGRouter := Router.Group("consulG").Use(middleware.OperationRecord())
	consulGRouterWithoutRecord := Router.Group("consulG")
	consulGRouterWithoutAuth := PublicRouter.Group("consulG")
	{
		consulGRouter.POST("createConsulGroup", consulGApi.CreateConsulGroup)             // 新建监控数据源consul
		consulGRouter.DELETE("deleteConsulGroup", consulGApi.DeleteConsulGroup)           // 删除监控数据源consul
		consulGRouter.DELETE("deleteConsulGroupByIds", consulGApi.DeleteConsulGroupByIds) // 批量删除监控数据源consul
		consulGRouter.PUT("updateConsulGroup", consulGApi.UpdateConsulGroup)              // 更新监控数据源consul
	}
	{
		consulGRouterWithoutRecord.GET("findConsulGroup", consulGApi.FindConsulGroup)         // 根据ID获取监控数据源consul
		consulGRouterWithoutRecord.GET("findConsulGroupUuid", consulGApi.FindConsulGroupUuid) // 根据ID获取监控数据源consul
		consulGRouterWithoutRecord.GET("getConsulGroupList", consulGApi.GetConsulGroupList)   // 获取监控数据源consul列表
	}
	{
		consulGRouterWithoutAuth.GET("getConsulGroupPublic", consulGApi.GetConsulGroupPublic) // 获取监控数据源consul列表
	}
}
