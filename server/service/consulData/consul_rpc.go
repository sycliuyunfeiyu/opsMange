package consulData

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulData"
	"io/ioutil"
	"net/http"
)

type ConsulRpcService struct {
}

func (crs *ConsulRpcService) ConsulRegister(consulDataModel consulData.ConsulData) error {
	var err error
	var consulRpcModel consulData.ConsulRpcModel

	consulRpcModel.Id = consulDataModel.ConsulId
	consulRpcModel.Name = consulDataModel.ConsulName
	consulRpcModel.Address = consulDataModel.ServiceAddress
	consulRpcModel.Port = *consulDataModel.ServicePort
	data_tags := make([]string, 0)
	err = json.Unmarshal([]byte(consulDataModel.ServiceTags), &data_tags)
	consulRpcModel.Tags = data_tags
	data_meta := make(map[string]string)
	err = json.Unmarshal([]byte(consulDataModel.ServiceMeta), &data_meta)
	consulRpcModel.Meta = data_meta

	consulRpcModel.Checks = func(dataS string) []map[string]string {
		var dataEnd []map[string]string

		data1 := make([]map[string]string, 0)
		err := json.Unmarshal([]byte(dataS), &data1)
		fmt.Println(err)
		for _, data2 := range data1 {
			dataEnd = append(dataEnd, data2)

		}

		return dataEnd
	}(consulDataModel.ServiceChecks)

	consulRpcModelByte, _ := json.Marshal(consulRpcModel)

	client := &http.Client{}
	register_url := consulDataModel.ConsulUrl + "/v1/agent/service/register"

	req, err := http.NewRequest("PUT", register_url, bytes.NewReader(consulRpcModelByte))
	req.Header.Add("x-consul-token", consulDataModel.ServiceToken)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	var consulDataService ConsulDataService
	if err == nil && resp.Status != "200 OK" {
		err = errors.New("consul状态码" + resp.Status + ",注册失败...")
	} else if err == nil {
		consulStatus := 1
		consulDataModel.ConsulStatus = &consulStatus
		err = consulDataService.UpdateConsulData(consulDataModel)
		if err != nil {
			//fmt.Println("更新失败。。。", err.Error())
			return err
		}
	}
	return err

}

func (crs *ConsulRpcService) ConsulDeRegister(consulDataModel consulData.ConsulData) error {
	var err error

	client := &http.Client{}
	deregister_url := consulDataModel.ConsulUrl + "/v1/agent/service/deregister/" + consulDataModel.ConsulId
	req, err := http.NewRequest("PUT", deregister_url, nil)
	req.Header.Add("x-consul-token", consulDataModel.ServiceToken)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	var consulDataService ConsulDataService
	if err == nil && resp.Status != "200 OK" {
		err = errors.New("consul状态码" + resp.Status + ",下线失败...")
	} else if err == nil {
		consulStatus := 2
		consulDataModel.ConsulStatus = &consulStatus
		err = consulDataService.UpdateConsulData(consulDataModel)
		if err != nil {
			//fmt.Println("更新失败。。。", err.Error())
			return err
		}
	}

	return err
}
