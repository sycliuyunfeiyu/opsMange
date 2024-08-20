import service from '@/utils/request'

// @Tags ConsulGroup
// @Summary 创建监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConsulGroup true "创建监控数据源consul"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /consulG/createConsulGroup [post]
export const createConsulGroup = (data) => {
  return service({
    url: '/consulG/createConsulGroup',
    method: 'post',
    data
  })
}

// @Tags ConsulGroup
// @Summary 删除监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConsulGroup true "删除监控数据源consul"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /consulG/deleteConsulGroup [delete]
export const deleteConsulGroup = (params) => {
  return service({
    url: '/consulG/deleteConsulGroup',
    method: 'delete',
    params
  })
}

// @Tags ConsulGroup
// @Summary 批量删除监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除监控数据源consul"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /consulG/deleteConsulGroup [delete]
export const deleteConsulGroupByIds = (params) => {
  return service({
    url: '/consulG/deleteConsulGroupByIds',
    method: 'delete',
    params
  })
}

// @Tags ConsulGroup
// @Summary 更新监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ConsulGroup true "更新监控数据源consul"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /consulG/updateConsulGroup [put]
export const updateConsulGroup = (data) => {
  return service({
    url: '/consulG/updateConsulGroup',
    method: 'put',
    data
  })
}

// @Tags ConsulGroup
// @Summary 用id查询监控数据源consul
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ConsulGroup true "用id查询监控数据源consul"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /consulG/findConsulGroup [get]
export const findConsulGroup = (params) => {
  return service({
    url: '/consulG/findConsulGroup',
    method: 'get',
    params
  })
}

// @Tags ConsulGroup
// @Summary 分页获取监控数据源consul列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取监控数据源consul列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /consulG/getConsulGroupList [get]
export const getConsulGroupList = (params) => {
  return service({
    url: '/consulG/getConsulGroupList',
    method: 'get',
    params
  })
}
export const findConsulGroupUuid = (params) => {
  return service({
    url: '/consulG/findConsulGroupUuid',
    method: 'get',
    params
  })
}

