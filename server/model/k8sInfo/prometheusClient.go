package k8sInfo

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	promApi "github.com/prometheus/client_golang/api"
	proV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"strings"
	"time"
)

type PrometheusClient struct {
	Url    string         `json:"url"`
	Client promApi.Client `json:"-"`
}

func (PromCS *PrometheusClient) NewClient() (err error) {
	PromCS.Client, err = promApi.NewClient(promApi.Config{
		Address: PromCS.Url,
	})
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		fmt.Printf("Error creating client: %v\n", err)
		return err
	}
	return nil
}

func (PromCS *PrometheusClient) Query(promSql string) (vector model.Vector, err error) {
	v1api := proV1.NewAPI(PromCS.Client)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, promSql, time.Now(), proV1.WithTimeout(5*time.Second))
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
		global.GVA_LOG.Warn("Warnings: %v\n" + strings.Join(warnings, "\n"))
	}
	vector, _ = result.(model.Vector)

	return vector, nil
}
