<template>
  <el-dialog v-model="isShowDialog" width="769px">
    <template #header>修改</template>
    <el-form :model="ruleForm" ref="formRef" :rules="rules" size="default" label-width="90px">
      <el-form-item label="标题" prop="name">
        <el-input v-model="ruleForm.name" placeholder="请输入标题" />
      </el-form-item>
      <el-form-item label="备注" prop="describe">
        <el-input v-model="ruleForm.describe" type="textarea" placeholder="请输入描述" />
      </el-form-item>
    </el-form>
    <template #footer>
				<span class="dialog-footer">
					<el-button @click="closeDialog" size="default">取 消</el-button>
					<el-button type="primary" @click="onSubmit" size="default">修改</el-button>
				</span>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import {ref, reactive, defineComponent} from 'vue';
import {getBigFile, editBigFile} from "/@/api/system/bigFile";

import {ElMessage} from "element-plus";
export default defineComponent({
  name: "editBigUpload",
  props:{},
  emits: ["success"],
  setup(prop, {emit}) {
    const formRef = ref<HTMLElement | null>(null);
    const isShowDialog = ref<boolean>(false)
    const ruleForm = reactive<any>({id:0, name:"", describe:""})
    const rules = reactive<any>({
      name: [
        { required: true, message: "标题不能为空", trigger: "blur" }
      ],
    })
    const openDialog = async (id:number) => {
      resetForm()
      const result = await getBigFile(id).then((res:any) => res.code === 0? res.data || {} : {})
      const {name, describe} = result
      ruleForm.id = result.id
      ruleForm.name = name
      ruleForm.describe = describe
      isShowDialog.value = true
    }
    const closeDialog =  () => {
      isShowDialog.value = false
    }


    const resetForm = () => {
      ruleForm.id = 0
      ruleForm.name = ""
      ruleForm.describe = ""
    }

    const  onSubmit = async () => {
      const formWrap = formRef.value as any
      if (!formWrap) return;

      formWrap.validate(async (valid: boolean) => {
        if (valid) {
          const result:any = await editBigFile(ruleForm)
          if (result.code === 0) {
            ElMessage.success('修改成功');
            closeDialog()
            emit('success')
          } else {
            ElMessage.error("修改失败")
          }
        }
      });
    }

    return {
      openDialog,
      closeDialog,
      isShowDialog,
      ruleForm,
      rules,
      onSubmit,
      formRef,
      resetForm
    }
  }
})
</script>

<style scoped>

</style>
