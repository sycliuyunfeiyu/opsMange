package consulGroup

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ConsulGroupApi }

var consulGService = service.ServiceGroupApp.ConsulGroupServiceGroup.ConsulGroupService
