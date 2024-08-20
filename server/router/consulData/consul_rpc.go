package consulData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConsulRpcRouter struct {
}

// InitConsulDataRouter 初始化 ConsulData 路由信息
func (s *ConsulDataRouter) InitConsulRpcRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	consulRpcRouter := Router.Group("consulRpc").Use(middleware.OperationRecord())
	//consulRpcRouterWithoutRecord := Router.Group("consulRpc")
	var consulRpcApi = v1.ApiGroupApp.ConsulDataApiGroup.ConsulRpcApi
	{
		consulRpcRouter.POST("registerConsulData", consulRpcApi.RegisterConsul)               // 注册consul rpc
		consulRpcRouter.POST("deRegisterConsulData", consulRpcApi.DeRegisterConsul)           // 下线consul rpc
		consulRpcRouter.POST("batchRegisterConsulData", consulRpcApi.BatchRegisterConsul)     // 批量上线consul rpc
		consulRpcRouter.POST("batchDeRegisterConsulData", consulRpcApi.BatchDeRegisterConsul) // 批量下线consul rpc

	}

}
