package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		consulGroupRouter := router.RouterGroupApp.ConsulGroup
		consulGroupRouter.InitConsulGroupRouter(privateGroup, publicGroup)
	}
	{
		consulDataRouter := router.RouterGroupApp.ConsulData
		consulDataRouter.InitConsulDataRouter(privateGroup, publicGroup)
		consulDataRouter.InitConsulRpcRouter(privateGroup, publicGroup)
	}
	{
		streamsetsDataRouter := router.RouterGroupApp.StreamsetsData
		streamsetsDataRouter.InitStreamsetsDataRouter(privateGroup, publicGroup)
	}
	{
		streamsetsPipelineInfoRouter := router.RouterGroupApp.StreamsetsPipelineInfo
		streamsetsPipelineInfoRouter.InitPipelineInfoRouter(privateGroup, publicGroup)
	}
	{
		k8sInfoRouter := router.RouterGroupApp.K8sInfo
		k8sInfoRouter.InitK8sInfoRouter(privateGroup, publicGroup)
	}
}
