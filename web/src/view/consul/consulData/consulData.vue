<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">

        <el-form-item label="注册名称" prop="consulName">
         <el-input v-model="searchInfo.consulName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态" prop="consulStatus">
             <el-input v-model.number="searchInfo.consulStatus" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="服务地址" prop="serviceAddress">
          <el-input v-model="searchInfo.serviceAddress" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="备注" prop="remake">
         <el-input v-model="searchInfo.remake" placeholder="搜索条件" />
        </el-form-item>
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->

          <el-form-item label="注册id" prop="consulId">
            <el-input v-model="searchInfo.consulId" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="consul地址" prop="consulUrl">
            <el-input v-model="searchInfo.consulUrl" placeholder="搜索条件" />

          </el-form-item>
          <el-form-item label="Uuid" prop="consulUuid">
            <el-input v-model="searchInfo.consulUuid" placeholder="搜索条件" />

          </el-form-item>
          <el-form-item label="组Uuid" prop="groupUuid">
            <el-input v-model="searchInfo.groupUuid" placeholder="搜索条件" />

          </el-form-item>
          <el-form-item label="服务元数据" prop="serviceMeta">
            <el-input v-model="searchInfo.serviceMeta" placeholder="搜索条件" />

          </el-form-item>
          <el-form-item label="创建日期" prop="createdAt">
            <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
            </template>
            <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
          </el-form-item>
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>

          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>

        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" label="注册id" prop="consulId" width="120" />
        <el-table-column align="left" label="注册名称" prop="consulName" width="120" />
        <el-table-column align="left" label="状态" prop="consulStatus" width="120" />
        <el-table-column align="left" label="consul地址" prop="consulUrl" width="120" />
        <el-table-column align="left" label="注册名称" prop="consulUuid" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.consulUuid,scope.row.consulUuid) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column align="left" label="组Uuid" prop="groupUuid" width="120" />
        <el-table-column align="left" label="备注" prop="remake" width="120" />
        <el-table-column align="left" label="服务地址" prop="serviceAddress" width="120" />
        <el-table-column align="left" label="服务检查" prop="serviceChecks" width="120" />
        <el-table-column align="left" label="服务内容" prop="serviceConnect" width="120" />
        <el-table-column align="left" label="服务标签" prop="serviceEnableTagOverride" width="120" />
        <el-table-column align="left" label="服务kind" prop="serviceKind" width="120" />
        <el-table-column align="left" label="服务元数据" prop="serviceMeta" width="120" />
        <el-table-column align="left" label="服务端口" prop="servicePort" width="120" />
        <el-table-column align="left" label="代理对象" prop="serviceProxy" width="120" />
        <el-table-column align="left" label="服务套接字地址" prop="serviceSocketPath" width="120" />
        <el-table-column align="left" label="服务附加地址" prop="serviceTaggedAdds" width="120" />
        <el-table-column align="left" label="服务标签" prop="serviceTags" width="120" />
        <el-table-column align="left" label="服务token" prop="serviceToken" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateConsulDataFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="registerConsulDataFunc(scope.row)">注册</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="deregisterConsulDataFunc(scope.row)">下线</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
<!--                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>-->
                <span class="text-lg" v-if="type==='create'">
                  添加
                </span>
                <span class="text-lg" v-if="type==='update'">
                  修改
                </span>
                <span class="text-lg" v-if="type==='register'">
                  注册
                </span>
                <span class="text-lg" v-if="type==='deregister'">
                  下线
                </span>

                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">

<!--            <el-form-item label="consul UUID:"  prop="consulUuid" >-->
<!--&lt;!&ndash;              <el-select  v-model="formData.consulUuid" placeholder="请选择consul UUID" style="width:100%" :clearable="true" @change="updateFromData(formData)">&ndash;&gt;-->
<!--&lt;!&ndash;                <el-option v-for="(item,key) in dataSource.consulUuid" :key="key" :label="item.label" :value="item.value" />&ndash;&gt;-->
<!--&lt;!&ndash;              </el-select>&ndash;&gt;-->
<!--            </el-form-item>-->

              <el-form-item  label="组Uuid:"  prop="groupUuid" >
<!--                <el-input v-model="formData.groupUuid" :clearable="true"  placeholder="请输入组Uuid" />-->
                <el-select  v-model="formData.groupUuid" placeholder="请选择组Uuid" style="width:100%" :clearable="true" @change="updateFromData(formData)">
                  <el-option v-for="(item,key) in dataSource.groupUuid" :key="key" :label="item.label" :value="item.value" />
                </el-select>
              </el-form-item>

            <el-form-item label="注册名称:"  prop="consulName" >
              <el-input v-model="formData.consulName" :clearable="true"  placeholder="请输入注册名称" />
            </el-form-item>
            <el-form-item label="consul地址:"  prop="consulUrl" >
              <el-input v-model="formData.consulUrl" :clearable="true"  placeholder="请输入consul地址" />
              <!--              <el-option v-for="(item,key) in dataSource.consulUrl" :key="key" :label="item.label" :value="item.value" />-->
            </el-form-item>
            <el-form-item label="状态:"  prop="consulStatus" >
<!--              <el-input v-model.number="formData.consulStatus" :clearable="true" placeholder="请输入状态" />-->
              <el-select v-model.number="formData.consulStatus" placeholder="请选择状态">
                <el-option
                    v-for="item in service_status"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>

            <el-form-item  label="注册id:"  prop="consulId" >
              <el-input v-model="formData.consulId" :clearable="true"  placeholder="请输入注册id" />
            </el-form-item>
            <el-form-item label="服务地址:"  prop="serviceAddress" >
              <el-input v-model="formData.serviceAddress" :clearable="true"  placeholder="请输入服务地址" />
            </el-form-item>
            <el-form-item label="服务token:"  prop="serviceToken" >
              <el-input v-model="formData.serviceToken" :clearable="true"  placeholder="请输入服务token" />
            </el-form-item>
            <el-form-item label="服务端口:"  prop="servicePort" >
              <el-input v-model.number="formData.servicePort" :clearable="true" placeholder="请输入服务端口" />
            </el-form-item>
            <el-form-item label="服务元数据:"  prop="serviceMeta" >
              <el-input v-model="formData.serviceMeta" :clearable="true"  placeholder="请输入服务元数据" />
            </el-form-item>
            <el-form-item label="服务标签:"  prop="serviceTags" >
              <el-input v-model="formData.serviceTags" :clearable="true"  placeholder="请输入服务标签" />
            </el-form-item>

            <el-form-item label="服务检查:"  prop="serviceChecks" >
              <el-input v-model="formData.serviceChecks" :clearable="true"  placeholder="请输入服务检查" />
            </el-form-item>
            <el-form-item label="服务内容:"  prop="serviceConnect" >
              <el-input v-model="formData.serviceConnect" :clearable="true"  placeholder="请输入服务内容" />
            </el-form-item>
<!--            <el-form-item label="服务标签:"  prop="serviceEnableTagOverride" >-->
<!--              <el-input v-model="formData.serviceEnableTagOverride" :clearable="true"  placeholder="请输入服务标签" />-->
<!--            </el-form-item>-->
            <el-form-item label="服务kind:"  prop="serviceKind" >
              <el-input v-model="formData.serviceKind" :clearable="true"  placeholder="请输入服务kind" />
            </el-form-item>

            <el-form-item label="代理对象:"  prop="serviceProxy" >
              <el-input v-model="formData.serviceProxy" :clearable="true"  placeholder="请输入代理对象" />
            </el-form-item>
            <el-form-item label="服务套接字地址:"  prop="serviceSocketPath" >
              <el-input v-model="formData.serviceSocketPath" :clearable="true"  placeholder="请输入服务套接字地址" />
            </el-form-item>
            <el-form-item label="服务附加地址:"  prop="serviceTaggedAdds" >
              <el-input v-model="formData.serviceTaggedAdds" :clearable="true"  placeholder="请输入服务附加地址" />
            </el-form-item>

            <el-form-item label="备注:"  prop="remake" >
              <el-input v-model="formData.remake" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>

          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getConsulDataDataSource,
  createConsulData,
  deleteConsulData,
  deleteConsulDataByIds,
  updateConsulData,
  findConsulData,
  getConsulDataList, registerConsulData, deregisterConsulData
} from '@/api/consulData/consulData'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import {findConsulGroupUuid, getConsulGroupList} from "@/api/consulGroup/consulGroup";

defineOptions({
    name: 'ConsulData'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        consulId: '',
        consulName: '',
        consulStatus: undefined,
        consulUrl: '',
        consulUuid: '',
        groupUuid: '',
        remake: '',
        serviceAddress: '',
        serviceChecks: '',
        serviceConnect: '',
        serviceEnableTagOverride: '',
        serviceKind: '',
        serviceMeta: '',
        servicePort: undefined,
        serviceProxy: '',
        serviceSocketPath: '',
        serviceTaggedAdds: '',
        serviceTags: '',
        serviceToken: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getConsulDataDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()
const service_status = [{
  value: '0',
  label: '未注册'
},{
  value: '1',
  label: '已注册'
},{
  value: '0',
  label: '已下线'
},]

// 验证规则
const rule = reactive({
               consulId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               consulName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               consulStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               consulUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               serviceAddress : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               servicePort : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()


var groupUuid = {}

const updateFromData  = async (val) => {
  const groupUuidData = await findConsulGroupUuid({groupUuid :formData.value.groupUuid})
  console.log(groupUuidData)
  // groupUuid = groupUuidData.data


  val.consulUrl = groupUuidData.data.groupConsulUrl
  val.consulName = groupUuidData.data.groupName
  val.serviceToken = groupUuidData.data.groupToken

}

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}
const getGroupData = async() => {
  const groupData = await getConsulGroupList()
  // updateFromData(groupData)
  if (groupData.code === 0) {
    for (var i = 0; i < groupData.data.list.length; i++) {
      groupUuid[groupData.data.list[i].groupUuid] = groupData.data.list[i]
    }
  }

  console.log('======groupUuid=======')
}


// 查询
const getTableData = async() => {
  const table = await getConsulDataList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
// getGroupData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteConsulDataFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteConsulDataByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateConsulDataFunc = async(row) => {
    // const guuid = await getGroupData({groupUuid: row.groupUuid})

    const res = await findConsulData({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}
// 注册行
const registerConsulDataFunc = async(row) => {
  const res = await findConsulData({ ID: row.ID })
  type.value = 'register'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 下线行
const deregisterConsulDataFunc = async(row) => {
  const res = await findConsulData({ ID: row.ID })
  type.value = 'deregister'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteConsulDataFunc = async (row) => {
    const res = await deleteConsulData({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        consulId: '',
        consulName: '',
        consulStatus: undefined,
        consulUrl: '',
        consulUuid: '',
        groupUuid: '',
        remake: '',
        serviceAddress: '',
        serviceChecks: '',
        serviceConnect: '',
        serviceEnableTagOverride: '',
        serviceKind: '',
        serviceMeta: '',
        servicePort: undefined,
        serviceProxy: '',
        serviceSocketPath: '',
        serviceTaggedAdds: '',
        serviceTags: '',
        serviceToken: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createConsulData(formData.value)
                  break
                case 'update':
                  res = await updateConsulData(formData.value)
                  break
                case 'register':
                  res = await registerConsulData(formData.value)
                  break
                case 'deregister':
                  res = await deregisterConsulData(formData.value)
                  break
                default:
                  res = await createConsulData(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
