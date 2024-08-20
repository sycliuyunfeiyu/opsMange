package consulData

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ConsulDataRouter
	ConsulRpcRouter
}

var consulDApi = api.ApiGroupApp.ConsulDataApiGroup.ConsulDataApi
