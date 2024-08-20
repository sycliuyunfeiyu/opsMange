package consulData

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ConsulDataApi
	ConsulRpcApi
}

var consulDService = service.ServiceGroupApp.ConsulDataServiceGroup.ConsulDataService
