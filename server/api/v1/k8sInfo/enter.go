package k8sInfo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ K8sInfoApi }

var kiService = service.ServiceGroupApp.K8sInfoServiceGroup.K8sInfoService
