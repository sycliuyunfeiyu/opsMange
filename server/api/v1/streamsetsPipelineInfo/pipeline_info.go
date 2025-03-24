package streamsetsPipelineInfo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	streamsetsDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsData/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsPipelineInfo"
	streamsetsPipelineInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/streamsetsPipelineInfo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/streamsetsData"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type PipelineInfoApi struct{}

// CreatePipelineInfo 创建流水线信息
// @Tags PipelineInfo
// @Summary 创建流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body streamsetsPipelineInfo.PipelineInfo true "创建流水线信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /pipelineInfo/createPipelineInfo [post]
func (pipelineInfoApi *PipelineInfoApi) CreatePipelineInfo(c *gin.Context) {
	var pipelineInfo streamsetsPipelineInfo.PipelineInfo
	err := c.ShouldBindJSON(&pipelineInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pipelineInfo.CreatedBy = utils.GetUserID(c)
	err = pipelineInfoService.CreatePipelineInfo(&pipelineInfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePipelineInfo 删除流水线信息
// @Tags PipelineInfo
// @Summary 删除流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body streamsetsPipelineInfo.PipelineInfo true "删除流水线信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /pipelineInfo/deletePipelineInfo [delete]
func (pipelineInfoApi *PipelineInfoApi) DeletePipelineInfo(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := pipelineInfoService.DeletePipelineInfo(ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePipelineInfoByIds 批量删除流水线信息
// @Tags PipelineInfo
// @Summary 批量删除流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /pipelineInfo/deletePipelineInfoByIds [delete]
func (pipelineInfoApi *PipelineInfoApi) DeletePipelineInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := pipelineInfoService.DeletePipelineInfoByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePipelineInfo 更新流水线信息
// @Tags PipelineInfo
// @Summary 更新流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body streamsetsPipelineInfo.PipelineInfo true "更新流水线信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /pipelineInfo/updatePipelineInfo [put]
func (pipelineInfoApi *PipelineInfoApi) UpdatePipelineInfo(c *gin.Context) {
	var pipelineInfo streamsetsPipelineInfo.PipelineInfo
	err := c.ShouldBindJSON(&pipelineInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pipelineInfo.UpdatedBy = utils.GetUserID(c)
	err = pipelineInfoService.UpdatePipelineInfo(pipelineInfo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (pipelineInfoApi *PipelineInfoApi) RsyncPipelineInfo(c *gin.Context) {
	var pipelineInfoList = make([]streamsetsPipelineInfo.PipelineInfo, 0, 10)
	//var pipelineInfoList []streamsetsPipelineInfo.PipelineInfo
	//pipelineInfoList = []streamsetsPipelineInfo.PipelineInfo{}

	var sd streamsetsData.StreamsetsDataService
	var sdr streamsetsDataReq.StreamsetsDataSearch
	sdcDataList, _, errgetdata := sd.GetStreamsetsDataInfoList(sdr)
	userId := utils.GetUserID(c)
	for _, sdcData := range sdcDataList {
		pipelineInfoList, _ = getPipelineInfoList(sdcData.Url, sdcData.User, sdcData.Passwd, sdcData.StreamsetsUuid, userId, pipelineInfoList)
	}

	pipelineInfoService.ApplyPipelineInfo(&pipelineInfoList)

	if errgetdata != nil {
		global.GVA_LOG.Error("同步失败!", zap.Error(errgetdata))
		response.FailWithMessage("同步失败", c)
		return
	}
	response.OkWithMessage("同步成功", c)
}

func getPipelineInfoList(url, user, passwd string, sdcuuid string, userId uint, pipelineInfoList []streamsetsPipelineInfo.PipelineInfo) ([]streamsetsPipelineInfo.PipelineInfo, error) {
	var pipelineSdcInfoList []streamsetsPipelineInfo.PipelineSdcInfo
	var pipelineSdcStatusInfoList streamsetsPipelineInfo.PipelineSdcStatusInfo

	sdcUrl := url + "rest/v1/pipelines"
	pipelineStatus := url + "rest/v1/pipelines/status"

	loginUser := fmt.Sprintf("%s:%s", user, passwd)
	authUser := "Basic " + base64.StdEncoding.EncodeToString([]byte(loginUser))
	fmt.Println(authUser)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", sdcUrl, nil)
	reqStatus, _ := http.NewRequest("GET", pipelineStatus, nil)

	req.Header.Add("Host", "127.0.0.1")
	req.Header.Add("Authorization", authUser)
	req.Header.Add("X-Requested-By", "sdc")

	reqStatus.Header.Add("Host", "127.0.0.1")
	reqStatus.Header.Add("Authorization", authUser)
	reqStatus.Header.Add("X-Requested-By", "sdc")
	//req.Header.Add("Cookie", "JSESSIONID_18630=node0890dtth43xut1f174t1xv0c8853.node0")

	response, _ := client.Do(req)
	responseStatus, _ := client.Do(reqStatus)

	sdcList, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(sdcList, &pipelineSdcInfoList)

	sdcStatusList, _ := ioutil.ReadAll(responseStatus.Body)
	fmt.Println(responseStatus.Body)
	json.Unmarshal(sdcStatusList, &pipelineSdcStatusInfoList)

	for i := 0; i < len(pipelineSdcInfoList); i++ {
		var pipelineInfo streamsetsPipelineInfo.PipelineInfo
		pipelineInfo.StreamsetsUuid = sdcuuid
		pipelineInfo.PipelineUuid = uuid.New().String()
		pipelineInfo.Pipeline_id = pipelineSdcInfoList[i].PipelineId
		pipelineInfo.Title_name = pipelineSdcInfoList[i].Name
		pipelineInfo.Message = pipelineSdcStatusInfoList[pipelineSdcInfoList[i].PipelineId].Message
		pipelineInfo.Status = pipelineSdcStatusInfoList[pipelineSdcInfoList[i].PipelineId].Status
		pipelineInfo.Link = url + "collector/pipeline/" + pipelineInfo.Pipeline_id
		pipelineInfo.CreatedBy = userId
		// 加入最终更新的信息
		pipelineInfoList = append(pipelineInfoList, pipelineInfo)

	}
	return pipelineInfoList, nil
}

// FindPipelineInfo 用id查询流水线信息
// @Tags PipelineInfo
// @Summary 用id查询流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query streamsetsPipelineInfo.PipelineInfo true "用id查询流水线信息"
// @Success 200 {object} response.Response{data=streamsetsPipelineInfo.PipelineInfo,msg=string} "查询成功"
// @Router /pipelineInfo/findPipelineInfo [get]
func (pipelineInfoApi *PipelineInfoApi) FindPipelineInfo(c *gin.Context) {
	ID := c.Query("ID")
	repipelineInfo, err := pipelineInfoService.GetPipelineInfo(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(repipelineInfo, c)
}

// GetPipelineInfoList 分页获取流水线信息列表
// @Tags PipelineInfo
// @Summary 分页获取流水线信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query streamsetsPipelineInfoReq.PipelineInfoSearch true "分页获取流水线信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /pipelineInfo/getPipelineInfoList [get]
func (pipelineInfoApi *PipelineInfoApi) GetPipelineInfoList(c *gin.Context) {
	var pageInfo streamsetsPipelineInfoReq.PipelineInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := pipelineInfoService.GetPipelineInfoInfoList(pageInfo)
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

// GetPipelineInfoPublic 不需要鉴权的流水线信息接口
// @Tags PipelineInfo
// @Summary 不需要鉴权的流水线信息接口
// @accept application/json
// @Produce application/json
// @Param data query streamsetsPipelineInfoReq.PipelineInfoSearch true "分页获取流水线信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /pipelineInfo/getPipelineInfoPublic [get]
func (pipelineInfoApi *PipelineInfoApi) GetPipelineInfoPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的流水线信息接口信息",
	}, "获取成功", c)
}
