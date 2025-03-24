<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="streamsets名称:" prop="streamsetsName">
          <el-input v-model="formData.streamsetsName" :clearable="true"  placeholder="请输入streamsets名称" />
       </el-form-item>
        <el-form-item label="uuid:" prop="streamsetsUuid">
          <el-input v-model="formData.streamsetsUuid" :clearable="true"  placeholder="请输入uuid" />
       </el-form-item>
        <el-form-item label="地址:" prop="url">
          <el-input v-model="formData.url" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="用户名:" prop="user">
          <el-input v-model="formData.user" :clearable="true"  placeholder="请输入用户名" />
       </el-form-item>
        <el-form-item label="密码:" prop="passwd">
          <el-input v-model="formData.passwd" :clearable="true"  placeholder="请输入密码" />
       </el-form-item>
        <el-form-item label="token:" prop="token">
          <el-input v-model="formData.token" :clearable="true"  placeholder="请输入token" />
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
  createStreamsetsData,
  updateStreamsetsData,
  findStreamsetsData
} from '@/api/streamsetsData/streamsetsData'

defineOptions({
    name: 'StreamsetsDataForm'
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
            streamsetsName: '',
            streamsetsUuid: '',
            url: '',
            user: '',
            passwd: '',
            token: '',
        })
// 验证规则
const rule = reactive({
               streamsetsName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               url : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               user : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               passwd : [{
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
      const res = await findStreamsetsData({ ID: route.query.id })
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
               res = await createStreamsetsData(formData.value)
               break
             case 'update':
               res = await updateStreamsetsData(formData.value)
               break
             default:
               res = await createStreamsetsData(formData.value)
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
