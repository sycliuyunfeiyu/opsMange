package streamsetsData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsData"
	streamsetsDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsData/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type StreamsetsDataApi struct{}

// CreateStreamsetsData 创建streamsets数据源
// @Tags StreamsetsData
// @Summary 创建streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body streamsetsData.StreamsetsData true "创建streamsets数据源"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sdc/createStreamsetsData [post]
func (sdcApi *StreamsetsDataApi) CreateStreamsetsData(c *gin.Context) {
	var sdc streamsetsData.StreamsetsData
	err := c.ShouldBindJSON(&sdc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sdc.StreamsetsUuid = uuid.New().String()
	sdc.CreatedBy = utils.GetUserID(c)
	err = sdcService.CreateStreamsetsData(&sdc)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteStreamsetsData 删除streamsets数据源
// @Tags StreamsetsData
// @Summary 删除streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body streamsetsData.StreamsetsData true "删除streamsets数据源"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sdc/deleteStreamsetsData [delete]
func (sdcApi *StreamsetsDataApi) DeleteStreamsetsData(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := sdcService.DeleteStreamsetsData(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteStreamsetsDataByIds 批量删除streamsets数据源
// @Tags StreamsetsData
// @Summary 批量删除streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sdc/deleteStreamsetsDataByIds [delete]
func (sdcApi *StreamsetsDataApi) DeleteStreamsetsDataByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := sdcService.DeleteStreamsetsDataByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateStreamsetsData 更新streamsets数据源
// @Tags StreamsetsData
// @Summary 更新streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body streamsetsData.StreamsetsData true "更新streamsets数据源"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sdc/updateStreamsetsData [put]
func (sdcApi *StreamsetsDataApi) UpdateStreamsetsData(c *gin.Context) {
	var sdc streamsetsData.StreamsetsData
	err := c.ShouldBindJSON(&sdc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sdc.UpdatedBy = utils.GetUserID(c)
	err = sdcService.UpdateStreamsetsData(sdc)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindStreamsetsData 用id查询streamsets数据源
// @Tags StreamsetsData
// @Summary 用id查询streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query streamsetsData.StreamsetsData true "用id查询streamsets数据源"
// @Success 200 {object} response.Response{data=streamsetsData.StreamsetsData,msg=string} "查询成功"
// @Router /sdc/findStreamsetsData [get]
func (sdcApi *StreamsetsDataApi) FindStreamsetsData(c *gin.Context) {
	ID := c.Query("ID")
	resdc, err := sdcService.GetStreamsetsData(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(resdc, c)
}

// GetStreamsetsDataList 分页获取streamsets数据源列表
// @Tags StreamsetsData
// @Summary 分页获取streamsets数据源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query streamsetsDataReq.StreamsetsDataSearch true "分页获取streamsets数据源列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sdc/getStreamsetsDataList [get]
func (sdcApi *StreamsetsDataApi) GetStreamsetsDataList(c *gin.Context) {
	var pageInfo streamsetsDataReq.StreamsetsDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sdcService.GetStreamsetsDataInfoList(pageInfo)
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

// GetStreamsetsDataPublic 不需要鉴权的streamsets数据源接口
// @Tags StreamsetsData
// @Summary 不需要鉴权的streamsets数据源接口
// @accept application/json
// @Produce application/json
// @Param data query streamsetsDataReq.StreamsetsDataSearch true "分页获取streamsets数据源列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sdc/getStreamsetsDataPublic [get]
func (sdcApi *StreamsetsDataApi) GetStreamsetsDataPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的streamsets数据源接口信息",
	}, "获取成功", c)
}
