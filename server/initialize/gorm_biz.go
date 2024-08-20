package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulData"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulGroup"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(consulGroup.ConsulGroup{}, consulData.ConsulData{})
	if err != nil {
		return err
	}
	return nil
}
