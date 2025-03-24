package k8sInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type K8sInfoRouter struct{}

// InitK8sInfoRouter 初始化 k8s集群信息 路由信息
func (s *K8sInfoRouter) InitK8sInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	kiRouter := Router.Group("ki").Use(middleware.OperationRecord())
	kiRouterWithoutRecord := Router.Group("ki")
	kiRouterWithoutAuth := PublicRouter.Group("ki")
	{
		kiRouter.POST("createK8sInfo", kiApi.CreateK8sInfo)             // 新建k8s集群信息
		kiRouter.DELETE("deleteK8sInfo", kiApi.DeleteK8sInfo)           // 删除k8s集群信息
		kiRouter.DELETE("deleteK8sInfoByIds", kiApi.DeleteK8sInfoByIds) // 批量删除k8s集群信息
		kiRouter.PUT("updateK8sInfo", kiApi.UpdateK8sInfo)              // 更新k8s集群信息
		kiRouter.PUT("k8sInspection", kiApi.K8sInspection)              // 巡检k8s集群信息

	}
	{
		kiRouterWithoutRecord.GET("findK8sInfo", kiApi.FindK8sInfo)                     // 根据ID获取k8s集群信息
		kiRouterWithoutRecord.GET("getK8sInfoList", kiApi.GetK8sInfoList)               // 获取k8s集群信息列表
		kiRouterWithoutRecord.GET("k8sInspectionDownFile", kiApi.K8sInspectionDownFile) // 巡检k8s集群信息

	}
	{
		kiRouterWithoutAuth.GET("getK8sInfoPublic", kiApi.GetK8sInfoPublic) // 获取k8s集群信息列表
	}
}
