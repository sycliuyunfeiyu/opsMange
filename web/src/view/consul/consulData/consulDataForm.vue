<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="consul注册id:" prop="consulId">
          <el-input v-model="formData.consulId" :clearable="true"  placeholder="请输入consul注册id" />
       </el-form-item>
        <el-form-item label="consul注册名称:" prop="consulName">
          <el-input v-model="formData.consulName" :clearable="true"  placeholder="请输入consul注册名称" />
       </el-form-item>
        <el-form-item label="状态:" prop="consulStatus">
          <el-input v-model.number="formData.consulStatus" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="consul注册地址:" prop="consulUrl">
          <el-input v-model="formData.consulUrl" :clearable="true"  placeholder="请输入consul注册地址" />
       </el-form-item>
        <el-form-item label="consul注册名称:" prop="consulUuid">
        <el-select  v-model="formData.consulUuid" placeholder="请选择consul注册名称" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.consulUuid" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="consul组注册id:" prop="groupUuid">
          <el-input v-model="formData.groupUuid" :clearable="true"  placeholder="请输入consul组注册id" />
       </el-form-item>
        <el-form-item label="备注:" prop="remake">
          <el-input v-model="formData.remake" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="服务地址:" prop="serviceAddress">
          <el-input v-model="formData.serviceAddress" :clearable="true"  placeholder="请输入服务地址" />
       </el-form-item>
        <el-form-item label="服务检查:" prop="serviceChecks">
          <el-input v-model="formData.serviceChecks" :clearable="true"  placeholder="请输入服务检查" />
       </el-form-item>
        <el-form-item label="服务内容:" prop="serviceConnect">
          <el-input v-model="formData.serviceConnect" :clearable="true"  placeholder="请输入服务内容" />
       </el-form-item>
        <el-form-item label="服务标签:" prop="serviceEnableTagOverride">
          <el-input v-model="formData.serviceEnableTagOverride" :clearable="true"  placeholder="请输入服务标签" />
       </el-form-item>
        <el-form-item label="服务kind:" prop="serviceKind">
          <el-input v-model="formData.serviceKind" :clearable="true"  placeholder="请输入服务kind" />
       </el-form-item>
        <el-form-item label="服务元数据:" prop="serviceMeta">
          <el-input v-model="formData.serviceMeta" :clearable="true"  placeholder="请输入服务元数据" />
       </el-form-item>
        <el-form-item label="服务端口:" prop="servicePort">
          <el-input v-model.number="formData.servicePort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="代理对象:" prop="serviceProxy">
          <el-input v-model="formData.serviceProxy" :clearable="true"  placeholder="请输入代理对象" />
       </el-form-item>
        <el-form-item label="服务套接字地址:" prop="serviceSocketPath">
          <el-input v-model="formData.serviceSocketPath" :clearable="true"  placeholder="请输入服务套接字地址" />
       </el-form-item>
        <el-form-item label="服务附加地址:" prop="serviceTaggedAdds">
          <el-input v-model="formData.serviceTaggedAdds" :clearable="true"  placeholder="请输入服务附加地址" />
       </el-form-item>
        <el-form-item label="服务标签:" prop="serviceTags">
          <el-input v-model="formData.serviceTags" :clearable="true"  placeholder="请输入服务标签" />
       </el-form-item>
        <el-form-item label="服务token:" prop="serviceToken">
          <el-input v-model="formData.serviceToken" :clearable="true"  placeholder="请输入服务token" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getConsulDataDataSource,
  createConsulData,
  updateConsulData,
  findConsulData
} from '@/api/consulData/consulData'

defineOptions({
    name: 'ConsulDataForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
// 验证规则
const rule = reactive({
               consulId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               consulName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               consulStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               consulUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serviceAddress : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               servicePort : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getConsulDataDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findConsulData({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
             default:
               res = await createConsulData(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
