package consulData

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulData"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ConsulRpcApi struct {
}

var consulRpcService = service.ServiceGroupApp.ConsulDataServiceGroup.ConsulRpcService

// UpdateConsul
// @Description: 更新consul 数据
// @receiver CRA
// @param c
func (CRA *ConsulRpcApi) RegisterConsul(c *gin.Context) {
	var consulDataModel consulData.ConsulData
	_ = c.ShouldBindJSON(&consulDataModel)

	if err := consulRpcService.ConsulRegister(consulDataModel); err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}

// UpdateConsul
// @Description: 更新consul 数据
// @receiver CRA
// @param c
func (CRA *ConsulRpcApi) DeRegisterConsul(c *gin.Context) {
	var consulDataModel consulData.ConsulData
	_ = c.ShouldBindJSON(&consulDataModel)

	if err := consulRpcService.ConsulDeRegister(consulDataModel); err != nil {
		global.GVA_LOG.Error("下线失败!", zap.Error(err))
		response.FailWithMessage("下线失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("下线成功", c)
	}

}

func (CRA *ConsulRpcApi) BatchRegisterConsul(c *gin.Context) {
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string("-----------"))
	//
	//fmt.Println(string(data))

	var consulDataModelList []consulData.ConsulData
	var consulDataModelErr []map[string]string
	//consulDataModelIds := map[string][]consulData.ConsulData{"ids": consulDataModelList}

	err := c.ShouldBindJSON(&consulDataModelList)

	fmt.Println(consulDataModelList)
	if err != nil {
		errStr := "下线失败!解析数据出现问题" + string(err.Error())
		global.GVA_LOG.Error(errStr)
		response.FailWithMessage(errStr, c)
	}
	//consulDataModelList = consulDataModelIds["ids"]
	for _, consulDataModel := range consulDataModelList {
		err := consulRpcService.ConsulRegister(consulDataModel)
		if err != nil {
			consulDataModelErr = append(consulDataModelErr, map[string]string{consulDataModel.ConsulId: err.Error()})
		}
	}

	if errListLen := len(consulDataModelErr); errListLen > 0 {
		errS, _ := json.Marshal(consulDataModelErr)
		errStr := "注册失败!" + string(errS)
		global.GVA_LOG.Error(errStr)
		response.FailWithMessage(errStr, c)
	} else {
		response.OkWithMessage("注册成功", c)
	}

}

func (CRA *ConsulRpcApi) BatchDeRegisterConsul(c *gin.Context) {
	var consulDataModelList []consulData.ConsulData
	var consulDataModelErr []map[string]string
	consulDataModelIds := map[string][]consulData.ConsulData{"ids": consulDataModelList}
	err := c.ShouldBindJSON(&consulDataModelIds)
	if err != nil {
		errStr := "下线失败!解析数据出现问题" + string(err.Error())
		global.GVA_LOG.Error(errStr)
		response.FailWithMessage(errStr, c)
	}
	consulDataModelList = consulDataModelIds["ids"]
	for _, consulDataModel := range consulDataModelList {
		err := consulRpcService.ConsulDeRegister(consulDataModel)
		if err != nil {
			consulDataModelErr = append(consulDataModelErr, map[string]string{consulDataModel.ConsulId: err.Error()})
		}
	}
	if errListLen := len(consulDataModelErr); errListLen > 0 {
		errS, _ := json.Marshal(consulDataModelErr)
		errStr := "下线失败!" + string(errS)
		global.GVA_LOG.Error(errStr)
		response.FailWithMessage(errStr, c)
	} else {
		response.OkWithMessage("下线成功", c)
	}

}
