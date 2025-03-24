<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联sdcuuid:" prop="streamsetsUuid">
          <el-input v-model="formData.streamsetsUuid" :clearable="true"  placeholder="请输入关联sdcuuid" />
       </el-form-item>
        <el-form-item label="流水线uuid:" prop="pipelineUuid">
          <el-input v-model="formData.pipelineUuid" :clearable="true"  placeholder="请输入流水线uuid" />
       </el-form-item>
        <el-form-item label="流水线名称:" prop="title_name">
          <el-input v-model="formData.title_name" :clearable="true"  placeholder="请输入流水线名称" />
       </el-form-item>
        <el-form-item label="流水线id:" prop="pipeline_id">
          <el-input v-model="formData.pipeline_id" :clearable="true"  placeholder="请输入流水线id" />
       </el-form-item>
        <el-form-item label="last_updated:" prop="last_updated">
          <el-date-picker v-model="formData.last_updated" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="流水线状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入流水线状态" />
       </el-form-item>
        <el-form-item label="信息:" prop="message">
          <RichEdit v-model="formData.message"/>
       </el-form-item>
        <el-form-item label="流水线链接:" prop="link">
          <el-input v-model="formData.link" :clearable="true"  placeholder="请输入流水线链接" />
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
  createPipelineInfo,
  updatePipelineInfo,
  findPipelineInfo
} from '@/api/streamsetsPipelineInfo/pipelineInfo'

defineOptions({
    name: 'PipelineInfoForm'
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
            streamsetsUuid: '',
            pipelineUuid: '',
            title_name: '',
            pipeline_id: '',
            last_updated: new Date(),
            status: '',
            message: '',
            link: '',
        })
// 验证规则
const rule = reactive({
               pipeline_id : [{
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
      const res = await findPipelineInfo({ ID: route.query.id })
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
               res = await createPipelineInfo(formData.value)
               break
             case 'update':
               res = await updatePipelineInfo(formData.value)
               break
             default:
               res = await createPipelineInfo(formData.value)
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
