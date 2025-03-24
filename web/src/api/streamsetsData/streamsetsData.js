import service from '@/utils/request'

// @Tags StreamsetsData
// @Summary 创建streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StreamsetsData true "创建streamsets数据源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sdc/createStreamsetsData [post]
export const createStreamsetsData = (data) => {
  return service({
    url: '/sdc/createStreamsetsData',
    method: 'post',
    data
  })
}

// @Tags StreamsetsData
// @Summary 删除streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StreamsetsData true "删除streamsets数据源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sdc/deleteStreamsetsData [delete]
export const deleteStreamsetsData = (params) => {
  return service({
    url: '/sdc/deleteStreamsetsData',
    method: 'delete',
    params
  })
}

// @Tags StreamsetsData
// @Summary 批量删除streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除streamsets数据源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sdc/deleteStreamsetsData [delete]
export const deleteStreamsetsDataByIds = (params) => {
  return service({
    url: '/sdc/deleteStreamsetsDataByIds',
    method: 'delete',
    params
  })
}

// @Tags StreamsetsData
// @Summary 更新streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StreamsetsData true "更新streamsets数据源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sdc/updateStreamsetsData [put]
export const updateStreamsetsData = (data) => {
  return service({
    url: '/sdc/updateStreamsetsData',
    method: 'put',
    data
  })
}

// @Tags StreamsetsData
// @Summary 用id查询streamsets数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.StreamsetsData true "用id查询streamsets数据源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sdc/findStreamsetsData [get]
export const findStreamsetsData = (params) => {
  return service({
    url: '/sdc/findStreamsetsData',
    method: 'get',
    params
  })
}

// @Tags StreamsetsData
// @Summary 分页获取streamsets数据源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取streamsets数据源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sdc/getStreamsetsDataList [get]
export const getStreamsetsDataList = (params) => {
  return service({
    url: '/sdc/getStreamsetsDataList',
    method: 'get',
    params
  })
}
