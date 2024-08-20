package consulData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/consulData"
	consulDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/consulData/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConsulDataApi struct{}

// CreateConsulData 创建consulData表
// @Tags ConsulData
// @Summary 创建consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body consulData.ConsulData true "创建consulData表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /consulD/createConsulData [post]
func (consulDApi *ConsulDataApi) CreateConsulData(c *gin.Context) {
	var consulD consulData.ConsulData
	consulD.ConsulUuid = uuid.New().String()

	err := c.ShouldBindJSON(&consulD)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	consulD.CreatedBy = utils.GetUserID(c)
	err = consulDService.CreateConsulData(&consulD)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteConsulData 删除consulData表
// @Tags ConsulData
// @Summary 删除consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body consulData.ConsulData true "删除consulData表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /consulD/deleteConsulData [delete]
func (consulDApi *ConsulDataApi) DeleteConsulData(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := consulDService.DeleteConsulData(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteConsulDataByIds 批量删除consulData表
// @Tags ConsulData
// @Summary 批量删除consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /consulD/deleteConsulDataByIds [delete]
func (consulDApi *ConsulDataApi) DeleteConsulDataByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := consulDService.DeleteConsulDataByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateConsulData 更新consulData表
// @Tags ConsulData
// @Summary 更新consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body consulData.ConsulData true "更新consulData表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /consulD/updateConsulData [put]
func (consulDApi *ConsulDataApi) UpdateConsulData(c *gin.Context) {
	var consulD consulData.ConsulData
	err := c.ShouldBindJSON(&consulD)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	consulD.UpdatedBy = utils.GetUserID(c)
	err = consulDService.UpdateConsulData(consulD)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindConsulData 用id查询consulData表
// @Tags ConsulData
// @Summary 用id查询consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query consulData.ConsulData true "用id查询consulData表"
// @Success 200 {object} response.Response{data=consulData.ConsulData,msg=string} "查询成功"
// @Router /consulD/findConsulData [get]
func (consulDApi *ConsulDataApi) FindConsulData(c *gin.Context) {
	ID := c.Query("ID")
	reconsulD, err := consulDService.GetConsulData(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reconsulD, c)
}

// GetConsulDataList 分页获取consulData表列表
// @Tags ConsulData
// @Summary 分页获取consulData表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query consulDataReq.ConsulDataSearch true "分页获取consulData表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /consulD/getConsulDataList [get]
func (consulDApi *ConsulDataApi) GetConsulDataList(c *gin.Context) {
	var pageInfo consulDataReq.ConsulDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := consulDService.GetConsulDataInfoList(pageInfo)
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

// GetConsulDataDataSource 获取ConsulData的数据源
// @Tags ConsulData
// @Summary 获取ConsulData的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /consulD/getConsulDataDataSource [get]
func (consulDApi *ConsulDataApi) GetConsulDataDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	dataSource, err := consulDService.GetConsulDataDataSource()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetConsulDataPublic 不需要鉴权的consulData表接口
// @Tags ConsulData
// @Summary 不需要鉴权的consulData表接口
// @accept application/json
// @Produce application/json
// @Param data query consulDataReq.ConsulDataSearch true "分页获取consulData表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /consulD/getConsulDataPublic [get]
func (consulDApi *ConsulDataApi) GetConsulDataPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的consulData表接口信息",
	}, "获取成功", c)
}
