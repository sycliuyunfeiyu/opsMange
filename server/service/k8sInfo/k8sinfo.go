package k8sInfo

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/k8sInfo"
	k8sInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sInfo/request"
	"github.com/jung-kurt/gofpdf"
	"go.uber.org/zap"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsclientset "k8s.io/metrics/pkg/client/clientset/versioned"
	"strconv"
	"strings"
)

type K8sInfoService struct{}
type K8sClusterInfoService struct{}

// CreateK8sInfo 创建k8s集群信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (kiService *K8sInfoService) CreateK8sInfo(ki *k8sInfo.K8sInfo) (err error) {
	err = global.GVA_DB.Create(ki).Error
	return err
}

// DeleteK8sInfo 删除k8s集群信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (kiService *K8sInfoService) DeleteK8sInfo(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&k8sInfo.K8sInfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&k8sInfo.K8sInfo{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteK8sInfoByIds 批量删除k8s集群信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (kiService *K8sInfoService) DeleteK8sInfoByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&k8sInfo.K8sInfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&k8sInfo.K8sInfo{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateK8sInfo 更新k8s集群信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (kiService *K8sInfoService) UpdateK8sInfo(ki k8sInfo.K8sInfo) (err error) {
	err = global.GVA_DB.Model(&k8sInfo.K8sInfo{}).Where("id = ?", ki.ID).Updates(&ki).Error
	return err
}

// GetK8sInfo 根据ID获取k8s集群信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (kiService *K8sInfoService) GetK8sInfo(ID string) (ki k8sInfo.K8sInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ki).Error
	return
}

// GetK8sInfoInfoList 分页获取k8s集群信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (kiService *K8sInfoService) GetK8sInfoInfoList(info k8sInfoReq.K8sInfoSearch) (list []k8sInfo.K8sInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&k8sInfo.K8sInfo{})
	var kis []k8sInfo.K8sInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.K8sName != "" {
		db = db.Where("k8s_name LIKE ?", "%"+info.K8sName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&kis).Error
	return kis, total, err
}

func (kniService *K8sClusterInfoService) GetK8sHealthInfo(reki k8sInfo.K8sInfo) (k8sInfo.K8sClusterInfo, error) {
	var K8sClusterInfoData k8sInfo.K8sClusterInfo
	var err error
	K8sClusterInfoData.K8sInfoUuid = reki.Uuid
	K8sClusterInfoData.K8sInfoName = reki.K8sName
	K8sClusterInfoData.K8sNodeMetrics, err = kniService.GetNodeInfo(reki.K8sConfig)
	if err != nil {
		return K8sClusterInfoData, err
	}
	k8sComponentsData, err := kniService.Getk8sComponentsHealthInfo(reki.K8sPrometheusUrl)

	K8sClusterInfoData.K8sComponentsData = k8sComponentsData
	if err != nil {
		fmt.Println(err.Error())
		return K8sClusterInfoData, err
	}
	return K8sClusterInfoData, nil

}

func (kniService *K8sClusterInfoService) GetNodeInfo(K8sConfig string) (map[string]k8sInfo.K8sNodeMetrics, error) {

	k8sNodeMetrics := make(map[string]k8sInfo.K8sNodeMetrics)
	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(K8sConfig))
	if err != nil {
		global.GVA_LOG.Error("查询失败!加载kubeconfig失败！！", zap.Error(err))
		return nil, errors.New("查询失败!加载kubeconfig失败！！" + err.Error())

	}
	kubeClient, kcErr := kubernetes.NewForConfig(kubeConfig)
	kubeMetricsClient, kmErr := metricsclientset.NewForConfig(kubeConfig)
	if kcErr != nil && kmErr != nil {
		global.GVA_LOG.Error("查询失败!创建客户端失败！！")
		return nil, errors.New("查询失败!创建客户端失败！！")
	}

	nodeList, kubeNodeErr := kubeClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if kubeNodeErr != nil {
		global.GVA_LOG.Error("k8s查询失败!"+kubeNodeErr.Error(), zap.Error(err))
		return nil, errors.New("k8s查询失败!" + " :" + kubeNodeErr.Error())
	}

	var nodeErrList []string
	nodeErrList = []string{""}

	var isNodeErr bool
	isNodeErr = false
	for _, nds := range nodeList.Items {
		K8sNodeMetricsData := k8sInfo.K8sNodeMetrics{}
		kubeMetricsNodeList, kubeMetricsErr := kubeMetricsClient.MetricsV1beta1().NodeMetricses().Get(context.TODO(), nds.Name, metav1.GetOptions{})
		if kubeMetricsErr != nil {
			isNodeErr = true
			nodeErrList = append(nodeErrList, nds.Name+kubeMetricsErr.Error())
			continue
		}
		K8sNodeMetricsData.NodeIp = nds.Status.Addresses[0].Address
		K8sNodeMetricsData.CpuTotal = nds.Status.Allocatable.Cpu().MilliValue() / 1000
		K8sNodeMetricsData.MemTotal = nds.Status.Allocatable.Memory().MilliValue() / 1024 / 1024 / 1024 / 1024
		K8sNodeMetricsData.PodTotal, _ = nds.Status.Allocatable.Pods().AsInt64()
		K8sNodeMetricsData.PodTotal, _ = nds.Status.Allocatable.Pods().AsInt64()

		K8sNodeMetricsData.CpuUsed = kubeMetricsNodeList.Usage.Cpu().MilliValue()
		K8sNodeMetricsData.MemUsed = kubeMetricsNodeList.Usage.Memory().MilliValue() / 1024 / 1024 / 1024 / 1024

		for _, condition := range nds.Status.Conditions {
			K8sNodeMetricsData.NodeInfo = append(K8sNodeMetricsData.NodeInfo, condition.Message)
		}

		//K8sNodeMetricsData.PodUsed, _ = kubeMetricsNodeList.Usage.Pods().MilliValue()

		podList, kubePodErr := kubeClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{FieldSelector: "spec.nodeName=" + nds.Name})
		//K8sNodeMetricsData.PodList = podList.Items
		for _, pod := range podList.Items {
			K8sNodeMetricsData.PodList = append(K8sNodeMetricsData.PodList, k8sInfo.PodInfo{PodName: pod.Name, Namespace: pod.Namespace, PodIp: pod.Status.PodIP, PodStatus: string(pod.Status.Phase)})

		}

		if kubePodErr != nil {
			global.GVA_LOG.Error("k8s查询失败!"+kubePodErr.Error(), zap.Error(err))
			return nil, errors.New("k8s查询失败!" + " :" + kubePodErr.Error())
		}
		k8sNodeMetrics[nds.Name] = K8sNodeMetricsData

	}

	if isNodeErr == true {
		global.GVA_LOG.Error("部份节点查询失败!" + strings.Join(nodeErrList, ","))
		return nil, errors.New("部份节点查询失败!" + strings.Join(nodeErrList, ","))

	}

	if err != nil {
		global.GVA_LOG.Error("查询失败!"+err.Error(), zap.Error(err))
		return nil, errors.New("查询失败!" + strings.Join(nodeErrList, ",") + " :" + err.Error())

	}
	return k8sNodeMetrics, nil

}

func (kniService *K8sClusterInfoService) Getk8sComponentsHealthInfo(prometheusUrl string) (k8sComponentsData k8sInfo.K8sComponents, err error) {

	// etcd leader 切换
	const etcdLeaderIpSql = "sum(etcd_server_is_leader) by (instance)"
	// etcd 使用内存
	const etcdUseMemSql = "process_resident_memory_bytes{pod=~\".*etcd-client.*\"}/1024/1024"
	// etcd 磁盘提交时间
	const etcdDiskBackendCommitSql = "histogram_quantile(0.99, sum(rate(etcd_disk_wal_fsync_duration_seconds_bucket[5m])) by (instance, le))"
	// etcd wal 磁盘提交时间
	const etcdWalDiskBackendCommitSql = "histogram_quantile(0.99, sum(rate(etcd_disk_backend_commit_duration_seconds_bucket[5m])) by (instance, le))"
	// etcd key 数量
	const etcdKeyNumSql = "sum(etcd_debugging_mvcc_keys_total) by (instance)"
	// etcd db 大小
	const etcdDbSizeSql = "etcd_mvcc_db_total_size_in_bytes"
	// etcd 切换次数
	const etcdLeaderSwitchedNumSql = "increase(etcd_server_leader_changes_seen_total[12h])"
	// apiserve 请求延迟
	const apiserveReqDelaySql = "histogram_quantile(0.99, sum(rate(apiserver_request_duration_seconds_bucket[3h])) by (le, verb))"
	// apiserve 与 etcd 之间延迟
	const apiserveReqEtcdDelaySql = "histogram_quantile(0.99, sum(rate(etcd_request_duration_seconds_bucket[3h])) by (le))"
	// coredns 缓存命中率
	const corednsCacheMissSql = "(sum(rate(coredns_cache_hits_total[5m])) / (sum(rate(coredns_cache_hits_total[5m])) + sum(rate(coredns_cache_misses_total[5m]))))*100"
	// coredns 请求持续时间
	const corednsRequestDurationSql = "histogram_quantile(0.99, sum(rate(coredns_dns_request_duration_seconds_bucket[3h])) by (le, job))"
	// coredns 请求类型统计
	const corednsRequestTypeCountSql = "sum(rate(coredns_dns_request_type_count_total[3h])) by (type) or sum(rate(coredns_dns_requests_total[3h])) by (type)"
	// 调度队列
	const schedulerQueueSql = "rate(scheduler_queue_incoming_pods_total{event=\"PodAdd\"}[5m])"
	// kubelet 与 apiserve 延迟
	const kubeletApiserveDelaySql = "sum(rate(rest_client_requests_total{job=\"kubelet\", code=~\"5..\"}[12h]))/ sum(rate(rest_client_requests_total{job=\"kubelet\"}[12h]))"
	// kubelet pleg 耗时
	const kubeletPlegRelistDurationSql = "histogram_quantile(0.99, sum(rate(kubelet_pleg_relist_duration_seconds_bucket{job=\"kubelet\", metrics_path=\"/metrics\"}[5m])) by (instance, le))"
	// 服务发现数量
	const svcNumVectorSql = "sum(kube_service_info)"
	// endpoint 数量
	const endpointNumSql = "sum(kube_endpoint_address_available)"
	// iptables 同步延迟
	const iptablesSyncDelaySql = "histogram_quantile(0.95, rate(kubeproxy_sync_proxy_rules_duration_seconds_bucket[24h]))"

	var PromClient k8sInfo.PrometheusClient
	PromClient = k8sInfo.PrometheusClient{
		Url: prometheusUrl,
	}

	err = PromClient.NewClient()
	if err != nil {
		return k8sInfo.K8sComponents{}, err
	}

	etcdInfoDataMap := make(k8sInfo.EtcdInfoMap, 5)
	etcdLeaderIpVectorList, err := PromClient.Query(etcdLeaderIpSql)
	for _, vector := range etcdLeaderIpVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":2379")[0]
		if int(vector.Value) == 1 {
			etcdInfo := etcdInfoDataMap["cluster"]
			etcdInfo.LeaderIp = instance
			etcdInfoDataMap["cluster"] = etcdInfo
		}

		fmt.Printf("%s etcd leader ip：%f \n", vector.Metric["instance"], vector.Value)
	}

	etcdUseMemVectorList, err := PromClient.Query(etcdUseMemSql)
	for _, vector := range etcdUseMemVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":2379")[0]
		etcdInfo := etcdInfoDataMap[instance]
		etcdInfo.MemUsed = int(vector.Value)
		etcdInfoDataMap[instance] = etcdInfo
		fmt.Printf("%s etcd 使用内存：%f \n", vector.Metric["instance"], vector.Value)
	}

	etcdDiskBackendCommitVectorList, err := PromClient.Query(etcdDiskBackendCommitSql)
	for _, vector := range etcdDiskBackendCommitVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":2379")[0]
		etcdInfo := etcdInfoDataMap[instance]
		etcdInfo.SyncTime, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		etcdInfoDataMap[instance] = etcdInfo
		fmt.Printf("%s etcd 5分钟内DB磁盘同步持续时间：%.4f \n", vector.Metric["instance"], vector.Value)
	}

	etcdWalDiskBackendCommitVectorList, err := PromClient.Query(etcdWalDiskBackendCommitSql)
	for _, vector := range etcdWalDiskBackendCommitVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":2379")[0]
		etcdInfo := etcdInfoDataMap[instance]
		etcdInfo.WalSyncTime, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		etcdInfoDataMap[instance] = etcdInfo
		fmt.Printf("%s etcd 5分钟内wal磁盘同步持续时间：%.4f \n", vector.Metric["instance"], vector.Value)
	}

	etcdKeyNumVectorList, err := PromClient.Query(etcdKeyNumSql)
	for _, vector := range etcdKeyNumVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":2379")[0]
		etcdInfo := etcdInfoDataMap[instance]
		etcdInfo.EtcdKeyNum = int(vector.Value)
		etcdInfoDataMap[instance] = etcdInfo
		fmt.Printf("%s etcd key数量：%.4f \n", vector.Metric["instance"], vector.Value)
	}

	etcdDbSizeVectorList, err := PromClient.Query(etcdDbSizeSql)
	for _, vector := range etcdDbSizeVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":2379")[0]
		etcdInfo := etcdInfoDataMap[instance]
		etcdInfo.DataSize, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value/1024/1024/1024)), 64)
		etcdInfoDataMap[instance] = etcdInfo
		fmt.Printf("%s  etcd db 文件大小：%.3f G \n", vector.Metric["instance"], vector.Value/1024/1024/1024)
	}

	etcdLeaderSwitchedNumVectorList, err := PromClient.Query(etcdLeaderSwitchedNumSql)
	for _, vector := range etcdLeaderSwitchedNumVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":2379")[0]
		etcdInfo := etcdInfoDataMap[instance]
		etcdInfo.LeaderSwitch = int(vector.Value)
		etcdInfoDataMap[instance] = etcdInfo

		fmt.Printf("%s  etcd 12H 内 leader切换数量：%v \n", vector.Metric["instance"], vector.Value)
	}

	apiserveInfoDataMap := make(k8sInfo.ApiserveInfo, 5)
	apiserveReqDelayVectorList, err := PromClient.Query(apiserveReqDelaySql)
	for _, vector := range apiserveReqDelayVectorList {
		apiserveDelay := apiserveInfoDataMap["cluster"]
		delay, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		apiserveDelay.ReqDelay = apiserveDelay.ReqDelay + string(vector.Metric["verb"]) + ":" + strconv.FormatFloat(delay, 'f', -1, 32) + "/"
		apiserveInfoDataMap["cluster"] = apiserveDelay
		fmt.Printf("apisierve 处理 %s 的延迟（99 分位数,3h）：%v \n", vector.Metric["verb"], vector.Value)
	}

	apiserveReqEtcdDelayVectorList, err := PromClient.Query(apiserveReqEtcdDelaySql)
	for _, vector := range apiserveReqEtcdDelayVectorList {
		apiserveEtcdDelay := apiserveInfoDataMap["cluster"]
		apiserveEtcdDelay.EtcdDelay, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)

		apiserveInfoDataMap["cluster"] = apiserveEtcdDelay
		fmt.Printf("apisierve与etcd 的延迟（99 分位数,3小时内）：%v \n", vector.Value)
	}

	CoreDnsInfoDataMap := make(k8sInfo.CoreDnsInfo, 5)
	corednsCacheMissVectorList, err := PromClient.Query(corednsCacheMissSql)
	for _, vector := range corednsCacheMissVectorList {
		CoreDnsInfo := CoreDnsInfoDataMap["cluster"]
		CoreDnsInfo.CacheMiss, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		CoreDnsInfoDataMap["cluster"] = CoreDnsInfo

		fmt.Printf("coredns缓存命中率：%v % \n", vector.Value)
	}

	corednsRequestDurationList, err := PromClient.Query(corednsRequestDurationSql)
	for _, vector := range corednsRequestDurationList {
		CoreDnsInfo := CoreDnsInfoDataMap["cluster"]
		CoreDnsInfo.ReqDuration, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		CoreDnsInfoDataMap["cluster"] = CoreDnsInfo

		fmt.Printf("coredns请求延迟：%v \n", vector.Value)
	}

	corednsRequestTypeCount, err := PromClient.Query(corednsRequestTypeCountSql)
	for _, vector := range corednsRequestTypeCount {
		CoreDnsInfo := CoreDnsInfoDataMap["cluster"]
		requestTypeCountStr, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		CoreDnsInfo.RequestTypeCount = CoreDnsInfo.RequestTypeCount + string(vector.Metric["type"]) + ":" + strconv.FormatFloat(requestTypeCountStr, 'f', -1, 32) + "/"
		CoreDnsInfoDataMap["cluster"] = CoreDnsInfo

		fmt.Printf("coredns请求统计：%v \n", CoreDnsInfo.RequestTypeCount)
	}

	SchedulerInfoDataMap := make(k8sInfo.SchedulerInfo, 5)
	schedulerQueueVectorList, err := PromClient.Query(schedulerQueueSql)
	for _, vector := range schedulerQueueVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":10251")[0]
		schedulerInfo := SchedulerInfoDataMap[instance]
		schedulerInfo.QueueNum = int(vector.Value)
		SchedulerInfoDataMap[instance] = schedulerInfo
		fmt.Printf("scheduler 调度队列剩余多少pod：%v \n", vector.Value)
	}
	KubeletInfoDataMap := make(k8sInfo.KubeletInfo, 5)

	kubeletApiserveDelayVectorList, err := PromClient.Query(kubeletApiserveDelaySql)
	for _, vector := range kubeletApiserveDelayVectorList {
		kubeletInfo := KubeletInfoDataMap["cluster"]
		kubeletInfo.ApiserveDelay, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		KubeletInfoDataMap["cluster"] = kubeletInfo
		fmt.Printf("12h 内kubelet 请求apisiver 失败率(状态为5xx)：%v \n", vector.Value)
	}

	kubeletPlegRelistDurationList, err := PromClient.Query(kubeletPlegRelistDurationSql)
	for _, vector := range kubeletPlegRelistDurationList {
		instance := strings.Split(string(vector.Metric["instance"]), ":10250")[0]
		plegRelistDurationInfo := KubeletInfoDataMap[instance]
		plegRelistDurationInfo.PlegRelistDuration, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		KubeletInfoDataMap[instance] = plegRelistDurationInfo
		fmt.Printf("pleg 检测pod状态时间：%v \n", vector.Value)
	}

	ProxyInfoDataMap := make(k8sInfo.ProxyInfo, 5)
	svcNumVectorList, err := PromClient.Query(svcNumVectorSql)
	for _, vector := range svcNumVectorList {
		proxyInfo := ProxyInfoDataMap["cluster"]
		proxyInfo.SvcNum = int(vector.Value)
		ProxyInfoDataMap["cluster"] = proxyInfo
		fmt.Printf("svc服务数量: %v \n", vector.Value)
	}

	endpointNumVectorList, err := PromClient.Query(endpointNumSql)
	for _, vector := range endpointNumVectorList {
		proxyInfo := ProxyInfoDataMap["cluster"]
		proxyInfo.EndpointNum = int(vector.Value)
		ProxyInfoDataMap["cluster"] = proxyInfo
		fmt.Printf("Endpoint 服务数量: %v \n", vector.Value)
	}

	iptablesSyncDelayVectorList, err := PromClient.Query(iptablesSyncDelaySql)
	for _, vector := range iptablesSyncDelayVectorList {
		instance := strings.Split(string(vector.Metric["instance"]), ":10249")[0]
		proxyInfo := ProxyInfoDataMap[instance]
		proxyInfo.IptablesSyncDelay, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(vector.Value)), 64)
		ProxyInfoDataMap[instance] = proxyInfo
		fmt.Printf("iptables 规则同步延迟: %v \n", vector.Value)
	}

	k8sComponentsData.ProxyInfo = ProxyInfoDataMap
	k8sComponentsData.KubeletInfo = KubeletInfoDataMap
	k8sComponentsData.ApiserveInfo = apiserveInfoDataMap
	k8sComponentsData.EtcdInfoData = etcdInfoDataMap
	k8sComponentsData.SchedulerInfo = SchedulerInfoDataMap
	k8sComponentsData.CoreDnsInfos = CoreDnsInfoDataMap

	return k8sComponentsData, nil

}

func (kniService *K8sClusterInfoService) GetPdfFile(pdfData k8sInfo.K8sClusterInfo) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.AddUTF8Font("fs", "", "fs.ttf")
	pdf.SetFont("fs", "", 20)

	pdf.CellFormat(0, 10, pdfData.K8sInfoName+"集群节点数量为"+strconv.Itoa(len(pdfData.K8sNodeMetrics)), "", 1, "", false, 0, "")

	pdf.CellFormat(0, 10, "cluster: apiserve 请求延迟: ", "", 1, "", false, 0, "")
	for _, strinfo := range strings.Split(pdfData.K8sComponentsData.ApiserveInfo["cluster"].ReqDelay, "/") {
		pdf.CellFormat(0, 10, strinfo, "", 1, "", false, 0, "")
	}

	pdf.CellFormat(0, 10, "cluster: apiserve 请求etcd 延迟 info: "+strconv.FormatFloat(pdfData.K8sComponentsData.ApiserveInfo["cluster"].EtcdDelay, 'f', -1, 64), "", 1, "", false, 0, "")

	// kubelet 组件
	pdf.AddPage()
	pdf.CellFormat(0, 10, "kubelet 组件巡检：", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 10, "cluster: kubelet 与apiserve 延迟: "+strconv.FormatFloat(pdfData.K8sComponentsData.KubeletInfo["cluster"].ApiserveDelay, 'f', -1, 64)+" s", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 10, "", "", 1, "", false, 0, "")
	for nodeip, plegRelistDurationData := range pdfData.K8sComponentsData.KubeletInfo {
		if nodeip == "cluster" {
			continue
		}
		pdf.CellFormat(0, 10, nodeip+" :Pleg 检测pod状态所需时间: "+strconv.FormatFloat(plegRelistDurationData.PlegRelistDuration, 'f', -1, 64), "", 1, "", false, 0, "")
	}

	//etcd 组件
	pdf.AddPage()
	pdf.CellFormat(0, 10, "Etcd 组件巡检：", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 10, "Etcd leader ip："+pdfData.K8sComponentsData.EtcdInfoData["cluster"].LeaderIp, "", 1, "", false, 0, "")
	for nodeip, etcdNode := range pdfData.K8sComponentsData.EtcdInfoData {
		if nodeip == "cluster" {
			continue
		}
		pdf.CellFormat(0, 10, nodeip+" 使用内存: "+strconv.Itoa(etcdNode.MemUsed)+" M", "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, nodeip+" 5分钟内wal磁盘同步持续时间: "+strconv.FormatFloat(etcdNode.WalSyncTime, 'f', -1, 64)+" 秒", "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, nodeip+" 5分钟内DB磁盘同步持续时间: "+strconv.FormatFloat(etcdNode.SyncTime, 'f', -1, 64)+" 秒", "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, nodeip+" etcd key数量：: "+strconv.Itoa(etcdNode.EtcdKeyNum), "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, nodeip+" 数据存储大小: "+strconv.FormatFloat(etcdNode.DataSize, 'f', -1, 64)+" G", "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, nodeip+" 切换leader次数: "+strconv.Itoa(etcdNode.LeaderSwitch), "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, " ", "", 1, "", false, 0, "")

	}

	// scheduler 组件
	pdf.AddPage()
	pdf.CellFormat(0, 10, "Scheduler 组件巡检：", "", 1, "", false, 0, "")
	for nodeip, schedulerNode := range pdfData.K8sComponentsData.SchedulerInfo {
		pdf.CellFormat(0, 10, nodeip+" :Scheduler 队列调度数量: "+strconv.Itoa(schedulerNode.QueueNum), "", 1, "", false, 0, "")
	}

	// coredns 组件
	pdf.AddPage()
	pdf.CellFormat(0, 10, "CoreDns 组件巡检：", "", 1, "", false, 0, "")
	for _, strinfo := range strings.Split(pdfData.K8sComponentsData.CoreDnsInfos["cluster"].RequestTypeCount, "/") {
		pdf.CellFormat(0, 10, strinfo, "", 1, "", false, 0, "")
	}
	pdf.CellFormat(0, 10, "cluster: coredns 缓存命中率: "+strconv.FormatFloat(pdfData.K8sComponentsData.CoreDnsInfos["cluster"].CacheMiss, 'f', -1, 64)+"%", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 10, "cluster: coredns 请求延迟: "+strconv.FormatFloat(pdfData.K8sComponentsData.CoreDnsInfos["cluster"].ReqDuration, 'f', -1, 64), "", 1, "", false, 0, "")

	// proxy 组件
	pdf.AddPage()
	pdf.CellFormat(0, 10, "proxy 组件巡检：", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 10, "cluster: svc 规则数量: "+strconv.Itoa(pdfData.K8sComponentsData.ProxyInfo["cluster"].SvcNum), "", 1, "", false, 0, "")
	pdf.CellFormat(0, 10, "cluster: Endpoint 数量: "+strconv.Itoa(pdfData.K8sComponentsData.ProxyInfo["cluster"].EndpointNum), "", 1, "", false, 0, "")
	for nodeip, proxyNodeData := range pdfData.K8sComponentsData.ProxyInfo {
		pdf.CellFormat(0, 10, nodeip+": iptables同步延迟："+strconv.FormatFloat(proxyNodeData.IptablesSyncDelay, 'f', -1, 64), "", 1, "", false, 0, "")
	}

	for _, k8sInfoMetrics := range pdfData.K8sNodeMetrics {
		pdf.AddPage()
		pdf.CellFormat(0, 10, "节点ip："+k8sInfoMetrics.NodeIp, "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, "总量内存："+strconv.FormatInt(k8sInfoMetrics.MemTotal, 10)+" G", "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, "使用内存："+strconv.FormatInt(k8sInfoMetrics.MemUsed, 10)+" G", "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, "总量cpu："+strconv.FormatInt(k8sInfoMetrics.CpuTotal, 10), "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, "使用cpu："+strconv.FormatInt(k8sInfoMetrics.CpuUsed, 10), "", 1, "", false, 0, "")
		pdf.CellFormat(0, 10, "运行的pod数量："+strconv.Itoa(len(k8sInfoMetrics.PodList))+"/"+strconv.Itoa(int(k8sInfoMetrics.PodTotal)), "", 1, "", false, 0, "")

		pdf.CellFormat(0, 10, "节点最近信息：", "", 1, "", false, 0, "")
		for _, nodeI := range k8sInfoMetrics.NodeInfo {
			pdf.CellFormat(0, 10, nodeI, "", 1, "", false, 0, "")
		}
		pdf.CellFormat(0, 10, " ", "", 1, "", false, 0, "")

	}

	pdf.OutputFileAndClose("/tmp/" + pdfData.K8sInfoName + ".pdf")
	return nil
}
