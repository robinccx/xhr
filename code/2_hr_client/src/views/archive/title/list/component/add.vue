<template>  
  <div class="archive-title-add">
    <!-- 添加或修改职称对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.archive-title-add .el-dialog', '.archive-title-add .el-dialog__header']">添加职称</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="90px">        
        <el-form-item label="职称" prop="titleName">
          <el-input v-model="formData.titleName" placeholder="请输入职称" />
        </el-form-item>        
        <el-form-item label="次序" prop="seq">
          <el-input v-model="formData.seq" placeholder="请输入次序" />
        </el-form-item>          
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio
              v-for="dict in statusOptions"
              :key="dict.value"
              :label="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>       
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="onSubmit">确 定</el-button>
          <el-button @click="onCancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>
<script lang="ts">
import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance } from 'vue';
import {ElMessageBox, ElMessage, FormInstance,UploadProps} from 'element-plus';
import {
  listTitle,
  getTitle,
  delTitle,
  addTitle,
  updateTitle,  
  getUserList,  
} from "/@/api/archive/title";
import {
  TitleTableColumns,
  TitleInfoData,
  TitleTableDataState,
  TitleEditState,  
} from "/@/views/archive/title/list/component/model"
export default defineComponent({
  name:"apiV1ArchiveTitleEdit",
  components:{    
  },
  props:{    
    statusOptions:{
      type:Array,
      default:()=>[]
    },    
  },
  setup(props,{emit}) {    
    const {proxy} = <any>getCurrentInstance()
    const formRef = ref<HTMLElement | null>(null);
    const menuRef = ref();    
    const state = reactive<TitleEditState>({
      loading:false,
      isShowDialog: false,
      formData: {        
        titleId: undefined,        
        titleName: undefined,        
        seq: undefined,        
        status: false ,        
        createdAt: undefined,        
        createdBy: undefined,        
        updatedAt: undefined,        
        updatedBy: undefined,        
        deletedAt: undefined,        
        deletedBy: undefined,        
      },
      // 表单校验
      rules: {        
        titleName : [
            { required: true, message: "职称不能为空", trigger: "blur" }
        ],        
        status : [
            { required: true, message: "状态不能为空", trigger: "blur" }
        ],        
      }
    });
    // 打开弹窗
    const openDialog = () => {
      resetForm();
      state.formData.status = "1";
      state.isShowDialog = true;
    };
    // 关闭弹窗
    const closeDialog = () => {
      state.isShowDialog = false;
    };
    // 取消
    const onCancel = () => {
      closeDialog();
    };
    // 提交
    const onSubmit = () => {
      const formWrap = unref(formRef) as any;
      if (!formWrap) return;
      formWrap.validate((valid: boolean) => {
        if (valid) {
          state.loading = true;
          //添加
          addTitle(state.formData).then(()=>{
              ElMessage.success('添加成功');
              closeDialog(); // 关闭弹窗
              emit('titleList')
            }).finally(()=>{
              state.loading = false;
            })
        }
      });
    };
    const resetForm = ()=>{
      state.formData = {        
        titleId: undefined,        
        titleName: undefined,        
        seq: undefined,        
        status: false ,        
        createdAt: undefined,        
        createdBy: undefined,        
        updatedAt: undefined,        
        updatedBy: undefined,        
        deletedAt: undefined,        
        deletedBy: undefined,        
      }
    };    
    return {
      proxy,
      openDialog,
      closeDialog,
      onCancel,
      onSubmit,
      menuRef,
      formRef,      
      ...toRefs(state),
    };
  }
})
</script>
<style scoped>  
</style>