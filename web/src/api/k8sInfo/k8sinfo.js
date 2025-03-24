import service from '@/utils/request'

// @Tags K8sInfo
// @Summary 创建k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sInfo true "创建k8s集群信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ki/createK8sInfo [post]
export const createK8sInfo = (data) => {
  return service({
    url: '/ki/createK8sInfo',
    method: 'post',
    data
  })
}

// @Tags K8sInfo
// @Summary 删除k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sInfo true "删除k8s集群信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ki/deleteK8sInfo [delete]
export const deleteK8sInfo = (params) => {
  return service({
    url: '/ki/deleteK8sInfo',
    method: 'delete',
    params
  })
}

// @Tags K8sInfo
// @Summary 批量删除k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除k8s集群信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ki/deleteK8sInfo [delete]
export const deleteK8sInfoByIds = (params) => {
  return service({
    url: '/ki/deleteK8sInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags K8sInfo
// @Summary 更新k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.K8sInfo true "更新k8s集群信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ki/updateK8sInfo [put]
export const updateK8sInfo = (data) => {
  return service({
    url: '/ki/updateK8sInfo',
    method: 'put',
    data
  })
}

export const k8sInspection = (data) => {
  return service({
    url: '/ki/k8sInspection',
    method: 'put',
    data
  })
}
// @Tags K8sInfo
// @Summary 用id查询k8s集群信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.K8sInfo true "用id查询k8s集群信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ki/findK8sInfo [get]
export const findK8sInfo = (params) => {
  return service({
    url: '/ki/findK8sInfo',
    method: 'get',
    params
  })
}




// @Tags K8sInfo
// @Summary 分页获取k8s集群信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取k8s集群信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ki/getK8sInfoList [get]
export const getK8sInfoList = (params) => {
  return service({
    url: '/ki/getK8sInfoList',
    method: 'get',
    params
  })
}

export const k8sInspectionDownFile = (params) => {
  return service({
    // responseType:'blob',
    // responseType:'arraybuffer',
    url: '/ki/k8sInspectionDownFile',
    method: 'get',
    params
  })
}



