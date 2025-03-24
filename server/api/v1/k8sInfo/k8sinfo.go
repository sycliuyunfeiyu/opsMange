package k8sInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	k8sInfoModel "github.com/flipped-aurora/gin-vue-admin/server/model/k8sInfo"
	k8sInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/k8sInfo/request"
	k8sInfoService "github.com/flipped-aurora/gin-vue-admin/server/service/k8sInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io/ioutil"
)

type K8sInfoApi struct{}

// CreateK8sInfo 创建k8s集群信息
// @Tags K8sInfo
// @Summary 创建k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sInfo.K8sInfo true "创建k8s集群信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ki/createK8sInfo [post]un
func (kiApi *K8sInfoApi) CreateK8sInfo(c *gin.Context) {
	var ki k8sInfoModel.K8sInfo
	err := c.ShouldBindJSON(&ki)
	ki.Uuid = uuid.New().String()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ki.CreatedBy = utils.GetUserID(c)
	err = kiService.CreateK8sInfo(&ki)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteK8sInfo 删除k8s集群信息
// @Tags K8sInfo
// @Summary 删除k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sInfo.K8sInfo true "删除k8s集群信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ki/deleteK8sInfo [delete]
func (kiApi *K8sInfoApi) DeleteK8sInfo(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := kiService.DeleteK8sInfo(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteK8sInfoByIds 批量删除k8s集群信息
// @Tags K8sInfo
// @Summary 批量删除k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ki/deleteK8sInfoByIds [delete]
func (kiApi *K8sInfoApi) DeleteK8sInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := kiService.DeleteK8sInfoByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateK8sInfo 更新k8s集群信息
// @Tags K8sInfo
// @Summary 更新k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body k8sInfo.K8sInfo true "更新k8s集群信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ki/updateK8sInfo [put]
func (kiApi *K8sInfoApi) UpdateK8sInfo(c *gin.Context) {
	var ki k8sInfoModel.K8sInfo
	err := c.ShouldBindJSON(&ki)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ki.UpdatedBy = utils.GetUserID(c)
	err = kiService.UpdateK8sInfo(ki)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindK8sInfo 用id查询k8s集群信息
// @Tags K8sInfo
// @Summary 用id查询k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query k8sInfo.K8sInfo true "用id查询k8s集群信息"
// @Success 200 {object} response.Response{data=k8sInfo.K8sInfo,msg=string} "查询成功"
// @Router /ki/findK8sInfo [get]
func (kiApi *K8sInfoApi) FindK8sInfo(c *gin.Context) {
	ID := c.Query("ID")
	reki, err := kiService.GetK8sInfo(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reki, c)
}

// GetK8sInfoList 分页获取k8s集群信息列表
// @Tags K8sInfo
// @Summary 分页获取k8s集群信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query k8sInfoReq.K8sInfoSearch true "分页获取k8s集群信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ki/getK8sInfoList [get]
func (kiApi *K8sInfoApi) GetK8sInfoList(c *gin.Context) {
	var pageInfo k8sInfoReq.K8sInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := kiService.GetK8sInfoInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetK8sInfoPublic 不需要鉴权的k8s集群信息接口
// @Tags K8sInfo
// @Summary 不需要鉴权的k8s集群信息接口
// @accept application/json
// @Produce application/json
// @Param data query k8sInfoReq.K8sInfoSearch true "分页获取k8s集群信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ki/getK8sInfoPublic [get]
func (kiApi *K8sInfoApi) GetK8sInfoPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的k8s集群信息接口信息",
	}, "获取成功", c)
}

func (kiApi *K8sInfoApi) K8sInspection(c *gin.Context) {
	var reki k8sInfoModel.K8sInfo
	err := c.ShouldBindJSON(&reki)

	if err != nil {
		global.GVA_LOG.Error("获取k8s集群失败 "+err.Error(), zap.Error(err))
		response.FailWithMessage("获取k8s集群失败 "+err.Error(), c)
		return

	}

	var K8sNodeInfoS k8sInfoService.K8sClusterInfoService

	// 获取节点信息（node Metrics）
	k8sInfoMetricsData, err := K8sNodeInfoS.GetK8sHealthInfo(reki)
	if err != nil {
		response.FailWithMessage("获取k8s集群数据失败 "+err.Error(), c)
	}
	K8sNodeInfoS.GetPdfFile(k8sInfoMetricsData)

	//data1, err := json.Marshal(k8sInfoMetricsData)
	//fmt.Print(string(data1))
	response.OkWithData(reki, c)
}

func (kiApi *K8sInfoApi) K8sInspectionDownFile(c *gin.Context) {

	var k8sname struct {
		K8sName string `json:"K8sName"`
	}
	err := c.ShouldBindQuery(&k8sname)
	if err != nil {
		global.GVA_LOG.Error("获取k8s集群 名称失败 "+err.Error(), zap.Error(err))
		response.FailWithMessage("获取k8s 名称失败 "+err.Error(), c)
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline;filename="+k8sname.K8sName+".pdf")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	content, err := ioutil.ReadFile("/tmp/" + k8sname.K8sName + ".pdf")
	c.Data(200, "application/pdf;chartset=UTF-8", content)

	response.OkWithMessage("aaa", c)

}
