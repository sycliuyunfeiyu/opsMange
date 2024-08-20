<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="组consul地址:" prop="groupConsulUrl">
          <el-input v-model="formData.groupConsulUrl" :clearable="true"  placeholder="请输入组consul地址" />
       </el-form-item>
        <el-form-item label="groupInfo字段:" prop="groupInfo">
          <el-input v-model="formData.groupInfo" :clearable="true"  placeholder="请输入groupInfo字段" />
       </el-form-item>
        <el-form-item label="组名称:" prop="groupName">
          <el-input v-model="formData.groupName" :clearable="true"  placeholder="请输入组名称" />
       </el-form-item>
        <el-form-item label="组token:" prop="groupToken">
          <el-input v-model="formData.groupToken" :clearable="true"  placeholder="请输入组token" />
       </el-form-item>
        <el-form-item label="唯一id:" prop="groupUuid">
          <el-input v-model="formData.groupUuid" :clearable="true"  placeholder="请输入唯一id" />
       </el-form-item>
        <el-form-item label="备注:" prop="remake">
          <el-input v-model="formData.remake" :clearable="true"  placeholder="请输入备注" />
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
  createConsulGroup,
  updateConsulGroup,
  findConsulGroup
} from '@/api/consulGroup/consulGroup'

defineOptions({
    name: 'ConsulGroupForm'
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
            groupConsulUrl: '',
            groupInfo: '',
            groupName: '',
            groupToken: '',
            groupUuid: '',
            remake: '',
        })
// 验证规则
const rule = reactive({
               groupConsulUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               groupInfo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               groupName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               groupToken : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               groupUuid : [{
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
      const res = await findConsulGroup({ ID: route.query.id })
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
               res = await createConsulGroup(formData.value)
               break
             case 'update':
               res = await updateConsulGroup(formData.value)
               break
             default:
               res = await createConsulGroup(formData.value)
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
