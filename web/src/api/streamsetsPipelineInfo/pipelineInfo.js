import service from '@/utils/request'

// @Tags PipelineInfo
// @Summary 创建流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PipelineInfo true "创建流水线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pipelineInfo/createPipelineInfo [post]
export const createPipelineInfo = (data) => {
  return service({
    url: '/pipelineInfo/createPipelineInfo',
    method: 'post',
    data
  })
}

// @Tags PipelineInfo
// @Summary 删除流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PipelineInfo true "删除流水线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pipelineInfo/deletePipelineInfo [delete]
export const deletePipelineInfo = (params) => {
  return service({
    url: '/pipelineInfo/deletePipelineInfo',
    method: 'delete',
    params
  })
}

// @Tags PipelineInfo
// @Summary 批量删除流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除流水线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pipelineInfo/deletePipelineInfo [delete]
export const deletePipelineInfoByIds = (params) => {
  return service({
    url: '/pipelineInfo/deletePipelineInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags PipelineInfo
// @Summary 更新流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PipelineInfo true "更新流水线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pipelineInfo/updatePipelineInfo [put]
export const updatePipelineInfo = (data) => {
  return service({
    url: '/pipelineInfo/updatePipelineInfo',
    method: 'put',
    data
  })
}

// @Tags PipelineInfo
// @Summary 更新流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PipelineInfo true "更新流水线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pipelineInfo/updatePipelineInfo [put]
export const rsyncPipelineInfo = (data) => {
  return service({
    url: '/pipelineInfo/rsyncPipelineInfo',
    method: 'put',
    data
  })
}

// @Tags PipelineInfo
// @Summary 用id查询流水线信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PipelineInfo true "用id查询流水线信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pipelineInfo/findPipelineInfo [get]
export const findPipelineInfo = (params) => {
  return service({
    url: '/pipelineInfo/findPipelineInfo',
    method: 'get',
    params
  })
}

// @Tags PipelineInfo
// @Summary 分页获取流水线信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取流水线信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pipelineInfo/getPipelineInfoList [get]
export const getPipelineInfoList = (params) => {
  return service({
    url: '/pipelineInfo/getPipelineInfoList',
    method: 'get',
    params
  })
}
