import service from '@/utils/request'

// @Tags ConsulData
// @Summary 创建consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConsulData true "创建consulData表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /consulD/createConsulData [post]
export const createConsulData = (data) => {
  return service({
    url: '/consulD/createConsulData',
    method: 'post',
    data
  })
}

export const createBatchConsulData = (data) => {
  return service({
    url: '/consulD/createBatchConsulData',
    method: 'post',
    data
  })
}
// @Tags ConsulData
// @Summary 删除consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConsulData true "删除consulData表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /consulD/deleteConsulData [delete]
export const deleteConsulData = (params) => {
  return service({
    url: '/consulD/deleteConsulData',
    method: 'delete',
    params
  })
}

// @Tags ConsulData
// @Summary 批量删除consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除consulData表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /consulD/deleteConsulData [delete]
export const deleteConsulDataByIds = (params) => {
  return service({
    url: '/consulD/deleteConsulDataByIds',
    method: 'delete',
    params
  })
}

// @Tags ConsulData
// @Summary 更新consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConsulData true "更新consulData表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /consulD/updateConsulData [put]
export const updateConsulData = (data) => {
  return service({
    url: '/consulD/updateConsulData',
    method: 'put',
    data
  })
}

// @Tags ConsulData
// @Summary 用id查询consulData表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ConsulData true "用id查询consulData表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /consulD/findConsulData [get]
export const findConsulData = (params) => {
  return service({
    url: '/consulD/findConsulData',
    method: 'get',
    params
  })
}

// @Tags ConsulData
// @Summary 分页获取consulData表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取consulData表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /consulD/getConsulDataList [get]
export const getConsulDataList = (params) => {
  return service({
    url: '/consulD/getConsulDataList',
    method: 'get',
    params
  })
}
// @Tags ConsulData
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /consulD/findConsulDataDataSource [get]
export const getConsulDataDataSource = () => {
  return service({
    url: '/consulD/getConsulDataDataSource',
    method: 'get',
  })
}

export const registerConsulData = (data) => {
  return service({
    url: '/consulRpc/registerConsulData',
    method: 'post',
    data
  })
}

export const registerConsulDataByItems = (data) => {
  return service({
    url: '/consulRpc/batchRegisterConsulData',
    method: 'post',
    data
  })
}
export const deregisterConsulDataByItems = (data) => {
  return service({
    url: '/consulRpc/batchDeRegisterConsulData',
    method: 'post',
    data
  })
}

export const deregisterConsulData = (data) => {
  return service({
    url: '/consulRpc/deRegisterConsulData',
    method: 'post',
    data
  })
}

// export const deregisterBatchConsulData = (params) => {
//   return service({
//     url: '/consulRpc/batchRegisterConsulData',
//     method: 'post',
//     params
//   })
// }
