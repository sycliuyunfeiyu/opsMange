<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标识符:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入标识符" />
       </el-form-item>
        <el-form-item label="k8s集群名称:" prop="k8sName">
          <el-input v-model="formData.k8sName" :clearable="true"  placeholder="请输入k8s集群名称" />
       </el-form-item>
        <el-form-item label="kubeconfig:" prop="k8sConfig">
          <RichEdit v-model="formData.k8sConfig"/>
       </el-form-item>
        <el-form-item label="巡检信息:" prop="inspectionStatus">
          <el-input v-model="formData.inspectionStatus" :clearable="true"  placeholder="请输入巡检信息" />
       </el-form-item>
        <el-form-item label="上次巡检时间:" prop="lastInspectionDate">
          <el-date-picker v-model="formData.lastInspectionDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createK8sInfo,
  updateK8sInfo,
  findK8sInfo
} from '@/api/k8sInfo/k8sinfo'

defineOptions({
    name: 'K8sInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            uuid: '',
            k8sName: '',
            k8sConfig: '',
            inspectionStatus: '',
            lastInspectionDate: new Date(),
        })
// 验证规则
const rule = reactive({
               k8sName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               k8sConfig : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findK8sInfo({ ID: route.query.id })
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
               res = await createK8sInfo(formData.value)
               break
             case 'update':
               res = await updateK8sInfo(formData.value)
               break
             default:
               res = await createK8sInfo(formData.value)
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
