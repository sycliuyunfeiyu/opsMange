// 自动生成模板ConsulData
package consulData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// consulData表 结构体  ConsulData
type ConsulData struct {
	global.GVA_MODEL
	ConsulId     string `json:"consulId" form:"consulId" gorm:"column:consul_id;comment:consul注册id;size:100;" binding:"required"`                     //consul注册id
	ConsulName   string `json:"consulName" form:"consulName" gorm:"column:consul_name;comment:consul注册名称;size:100;" binding:"required"`               //consul注册名称
	ConsulStatus *int   `json:"consulStatus" form:"consulStatus" gorm:"column:consul_status;comment:状态 0 未注册 1 已注册 2已删除;size:10;" binding:"required"` //状态
	ConsulUrl    string `json:"consulUrl" form:"consulUrl" gorm:"column:consul_url;comment:consul注册地址;size:100;" binding:"required"`                  //consul注册地址
	ConsulUuid   string `json:"consulUuid" form:"consulUuid" gorm:"column:consul_uuid;comment:consul注册名称;size:100;"`                                  //consul注册名称
	GroupUuid    string `json:"groupUuid" form:"groupUuid" gorm:"column:group_uuid;comment:consul组注册id;size:100;"`                                    //consul组注册id
	//GroupUuid consulGroup.ConsulGroup `json:"groupUuid" form:"groupUuid" gorm:"column:group_uuid;comment:consul组注册id;size:100;"` //consul组注册id
	//ConsulGroup consulGroup.ConsulGroup `gorm:"foreignKey:GroupUuid;references:GroupUuid"`

	Remake                   string `json:"remake" form:"remake" gorm:"column:remake;comment:备注;size:255;"`                                                                          //备注
	ServiceAddress           string `json:"serviceAddress" form:"serviceAddress" gorm:"column:service_address;comment:服务地址;size:100;" binding:"required"`                            //服务地址
	ServiceChecks            string `json:"serviceChecks" form:"serviceChecks" gorm:"column:service_checks;comment:健康检查后段地址{\"checks\":[{\"http\": \"\", \"interval\": \"3s\"}]};\"` //服务检查
	ServiceConnect           string `json:"serviceConnect" form:"serviceConnect" gorm:"column:service_connect;comment:;size:255;"`                                                   //服务内容
	ServiceEnableTagOverride string `json:"serviceEnableTagOverride" form:"serviceEnableTagOverride" gorm:"column:service_enable_tag_override;comment:;size:255;"`                   //服务标签
	ServiceKind              string `json:"serviceKind" form:"serviceKind" gorm:"column:service_kind;comment:;size:255;"`                                                            //服务kind
	ServiceMeta              string `json:"serviceMeta" form:"serviceMeta" gorm:"column:service_meta;comment:服务元数据 {\"xx\":\"xxxx\"};"`                                              //服务元数据
	ServicePort              *int   `json:"servicePort" form:"servicePort" gorm:"column:service_port;comment:服务端口;size:10;" binding:"required"`                                      //服务端口
	ServiceProxy             string `json:"serviceProxy" form:"serviceProxy" gorm:"column:service_proxy;comment:代理对象;size:255;"`                                                     //代理对象
	ServiceSocketPath        string `json:"serviceSocketPath" form:"serviceSocketPath" gorm:"column:service_socket_path;comment:指定服务套接字路径的字符串值;size:255;"`                           //服务套接字地址
	ServiceTaggedAdds        string `json:"serviceTaggedAdds" form:"serviceTaggedAdds" gorm:"column:service_tagged_adds;comment:标记地址是可以为节点或服务定义的附加地址;size:255;"`                     //服务附加地址
	ServiceTags              string `json:"serviceTags" form:"serviceTags" gorm:"column:service_tags;comment:服务标签 [\"n9e\",\"loca\"];\"`                                             //服务标签
	ServiceToken             string `json:"serviceToken" form:"serviceToken" gorm:"column:service_token;comment:没有则使用默认token;size:255;"`                                             //服务token
	CreatedBy                uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy                uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy                uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName consulData表 ConsulData自定义表名 consul_data
func (ConsulData) TableName() string {
	return "consul_data"
}
