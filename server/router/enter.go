package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/consulData"
	"github.com/flipped-aurora/gin-vue-admin/server/router/consulGroup"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	ConsulGroup consulGroup.RouterGroup
	ConsulData  consulData.RouterGroup
}
