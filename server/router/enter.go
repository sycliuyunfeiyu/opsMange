package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/consulData"
	"github.com/flipped-aurora/gin-vue-admin/server/router/consulGroup"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/k8sInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/streamsetsData"
	"github.com/flipped-aurora/gin-vue-admin/server/router/streamsetsPipelineInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System                 system.RouterGroup
	Example                example.RouterGroup
	ConsulGroup            consulGroup.RouterGroup
	ConsulData             consulData.RouterGroup
	StreamsetsPipelineInfo streamsetsPipelineInfo.RouterGroup
	StreamsetsData         streamsetsData.RouterGroup
	K8sInfo                k8sInfo.RouterGroup
}
