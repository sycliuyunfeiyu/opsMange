package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulData"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulGroup"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsData"
	"github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsPipelineInfo"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(consulGroup.ConsulGroup{}, consulData.ConsulData{}, streamsetsData.StreamsetsData{}, streamsetsPipelineInfo.PipelineInfo{}, k8sInfo.K8sInfo{})
	if err != nil {
		return err
	}
	return nil
}
