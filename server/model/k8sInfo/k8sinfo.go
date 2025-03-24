// 自动生成模板K8sInfo
package k8sInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// k8s集群信息 结构体  K8sInfo
type K8sInfo struct {
	global.GVA_MODEL
	Uuid             string `json:"uuid" form:"uuid" gorm:"primarykey;column:uuid;comment:标识符;"`                                            //标识符
	K8sName          string `json:"k8sName" form:"k8sName" gorm:"not null;uniqueIndex;column:k8s_name;comment:;" binding:"required"`        //k8s集群名称
	K8sPrometheusUrl string `json:"k8sPrometheusUrl" form:"k8sPrometheusUrl" gorm:"column:k8s_prometheus_url;comment:;" binding:"required"` //k8s集群名称

	K8sConfig          string     `json:"k8sConfig" form:"k8sConfig" gorm:"not null;column:k8s_config;comment:kubeconfig;type:text;" binding:"required"` //kubeconfig
	InspectionStatus   string     `json:"inspectionStatus" form:"inspectionStatus" gorm:"column:inspection_status;comment:巡检信息;"`                        //巡检信息
	LastInspectionDate *time.Time `json:"lastInspectionDate" form:"lastInspectionDate" gorm:"column:last_inspection_date;comment:;"`                     //上次巡检时间
	CreatedBy          uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy          uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy          uint       `gorm:"column:deleted_by;comment:删除者"`
}

type K8sClusterInfo struct {
	K8sInfoUuid       string                    `json:"k8sInfoUuid"`
	K8sInfoName       string                    `json:"k8sInfoName"`
	K8sNodeMetrics    map[string]K8sNodeMetrics `json:"k8sNodeMetrics"`
	K8sComponentsData K8sComponents             `json:"k8sInspectionData"`
}
type K8sNodeMetrics struct {
	NodeIp   string `json:"nodeIp"`
	CpuTotal int64  `json:"cpuTotal"`
	MemTotal int64  `json:"memTotal"`
	PodTotal int64  `json:"podTotal"`
	CpuUsed  int64  `json:"cpuUsed"`
	MemUsed  int64  `json:"memUsed"`
	PodUsed  int64  `json:"podUsed"`

	PodList  []PodInfo `json:"podList"`
	NodeInfo []string  `json:"nodeInfo"`
}
type PodInfo struct {
	PodName   string `json:"podname"`
	Namespace string `json:"namespace"`
	PodIp     string `json:"podIp"`
	PodStatus string `json:"podStatus"`
}
type K8sComponents struct {
	CpuData     int `json:"cpuData"`
	MemData     int `json:"memData"`
	PodNum      int `json:"podNum"`
	StorageData int `json:"storageData"`

	NodeNum int    `json:"nodeNum"`
	NodeErr string `json:"nodeErr"`

	EtcdInfoData EtcdInfoMap `json:"etcdInfoData"`

	ApiserveInfo  ApiserveInfo  `json:"apiserveInfo"`
	CoreDnsInfos  CoreDnsInfo   `json:"coreDnsInfos"`
	SchedulerInfo SchedulerInfo `json:"schedulerInfo"`
	KubeletInfo   KubeletInfo   `json:"kubeletInfo"`
	ProxyInfo     ProxyInfo     `json:"proxyInfo"`
}

type EtcdInfoMap map[string]struct {
	LeaderIp    string  `json:"leaderIp"`
	MemUsed     int     `json:"memUsed"`
	EtcdKeyNum  int     `json:"etcdKeyNum"`
	DataSize    float64 `json:"dataSize"`
	SyncTime    float64 `json:"syncTime"`
	WalSyncTime float64 `json:"walSyncTime"`

	LeaderSwitch int `json:"leaderSwitch"`
}
type ApiserveInfo map[string]struct {
	ReqDelay  string  `json:"reqDelay"`
	EtcdDelay float64 `json:"etcdDelay"`
}
type CoreDnsInfo map[string]struct {
	CacheMiss        float64 `json:"cacheMiss"`
	ReqDuration      float64 `json:"reqDuration"`
	RequestTypeCount string  `json:"requestTypeCount"`
}
type SchedulerInfo map[string]struct {
	QueueNum int `json:"queueNum"`
}
type KubeletInfo map[string]struct {
	ApiserveDelay      float64 `json:"apiServeDelay"`
	PlegRelistDuration float64 `json:"legRelistDuration"`
}
type ProxyInfo map[string]struct {
	SvcNum            int     `json:"svcNum"`
	EndpointNum       int     `json:"endpointNum"`
	IptablesSyncDelay float64 `json:"iptablesSyncDelay"`
}

// TableName k8s集群信息 K8sInfo自定义表名 k8sInfo
func (K8sInfo) TableName() string {
	return "k8sInfo"
}

type K8sNodeInfo struct {
	Ip string `json:"ip"`
}
