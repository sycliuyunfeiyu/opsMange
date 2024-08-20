package consulGroup

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulGroup"
	consulGroupReq "github.com/flipped-aurora/gin-vue-admin/server/model/consulGroup/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConsulGroupApi struct{}

// CreateConsulGroup 创建监控数据源consul
// @Tags ConsulGroup
// @Summary 创建监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body consulGroup.ConsulGroup true "创建监控数据源consul"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /consulG/createConsulGroup [post]
func (consulGApi *ConsulGroupApi) CreateConsulGroup(c *gin.Context) {
	var consulG consulGroup.ConsulGroup
	consulG.GroupUuid = uuid.New().String()
	err := c.ShouldBindJSON(&consulG)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	consulG.CreatedBy = utils.GetUserID(c)
	err = consulGService.CreateConsulGroup(&consulG)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteConsulGroup 删除监控数据源consul
// @Tags ConsulGroup
// @Summary 删除监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body consulGroup.ConsulGroup true "删除监控数据源consul"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /consulG/deleteConsulGroup [delete]
func (consulGApi *ConsulGroupApi) DeleteConsulGroup(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := consulGService.DeleteConsulGroup(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteConsulGroupByIds 批量删除监控数据源consul
// @Tags ConsulGroup
// @Summary 批量删除监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /consulG/deleteConsulGroupByIds [delete]
func (consulGApi *ConsulGroupApi) DeleteConsulGroupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := consulGService.DeleteConsulGroupByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateConsulGroup 更新监控数据源consul
// @Tags ConsulGroup
// @Summary 更新监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body consulGroup.ConsulGroup true "更新监控数据源consul"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /consulG/updateConsulGroup [put]
func (consulGApi *ConsulGroupApi) UpdateConsulGroup(c *gin.Context) {
	var consulG consulGroup.ConsulGroup
	err := c.ShouldBindJSON(&consulG)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	consulG.UpdatedBy = utils.GetUserID(c)
	err = consulGService.UpdateConsulGroup(consulG)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindConsulGroup 用id查询监控数据源consul
// @Tags ConsulGroup
// @Summary 用id查询监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query consulGroup.ConsulGroup true "用id查询监控数据源consul"
// @Success 200 {object} response.Response{data=consulGroup.ConsulGroup,msg=string} "查询成功"
// @Router /consulG/findConsulGroup [get]

func (consulGApi *ConsulGroupApi) FindConsulGroup(c *gin.Context) {
	ID := c.Query("ID")
	reconsulG, err := consulGService.GetConsulGroup(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reconsulG, c)
}

func (consulGApi *ConsulGroupApi) FindConsulGroupUuid(c *gin.Context) {
	Uuid := c.Query("groupUuid")
	fmt.Println(Uuid)
	reconsulG, err := consulGService.GetConsulGroupUuid(Uuid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reconsulG, c)
}

// GetConsulGroupList 分页获取监控数据源consul列表
// @Tags ConsulGroup
// @Summary 分页获取监控数据源consul列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query consulGroupReq.ConsulGroupSearch true "分页获取监控数据源consul列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /consulG/getConsulGroupList [get]
func (consulGApi *ConsulGroupApi) GetConsulGroupList(c *gin.Context) {
	var pageInfo consulGroupReq.ConsulGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := consulGService.GetConsulGroupInfoList(pageInfo)
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

// GetConsulGroupPublic 不需要鉴权的监控数据源consul接口
// @Tags ConsulGroup
// @Summary 不需要鉴权的监控数据源consul接口
// @accept application/json
// @Produce application/json
// @Param data query consulGroupReq.ConsulGroupSearch true "分页获取监控数据源consul列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /consulG/getConsulGroupPublic [get]
func (consulGApi *ConsulGroupApi) GetConsulGroupPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的监控数据源consul接口信息",
	}, "获取成功", c)
}
