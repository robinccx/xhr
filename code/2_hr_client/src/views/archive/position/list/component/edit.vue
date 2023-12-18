<template>  
  <div class="archive-position-edit">
    <!-- 添加或修改职务对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.archive-position-edit .el-dialog', '.archive-position-edit .el-dialog__header']">{{(!formData.positionId || formData.positionId==0?'添加':'修改')+'职务'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="90px">        
        <el-form-item label="职务" prop="positionName">
          <el-input v-model="formData.positionName" placeholder="请输入职务" />
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
  listPosition,
  getPosition,
  delPosition,
  addPosition,
  updatePosition,  
  getUserList,  
} from "/@/api/archive/position";
import {
  PositionTableColumns,
  PositionInfoData,
  PositionTableDataState,
  PositionEditState,  
} from "/@/views/archive/position/list/component/model"
export default defineComponent({
  name:"apiV1ArchivePositionEdit",
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
    const state = reactive<PositionEditState>({
      loading:false,
      isShowDialog: false,
      formData: {        
        positionId: undefined,        
        positionCode: undefined,        
        positionName: undefined,        
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
        positionId : [
            { required: true, message: "职务ID不能为空", trigger: "blur" }
        ],        
        positionName : [
            { required: true, message: "职务不能为空", trigger: "blur" }
        ],        
        status : [
            { required: true, message: "状态不能为空", trigger: "blur" }
        ],        
      }
    });
    // 打开弹窗
    const openDialog = (row?: PositionInfoData) => {
      resetForm();
      if(row) {        
        getPosition(row.positionId!).then((res:any)=>{
          const data = res.data;          
          data.status = ''+data.status          
          state.formData = data;
      })
    }
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
          if(!state.formData.positionId || state.formData.positionId===0){
            //添加
          addPosition(state.formData).then(()=>{
              ElMessage.success('添加成功');
              closeDialog(); // 关闭弹窗
              emit('positionList')
            }).finally(()=>{
              state.loading = false;
            })
          }else{
            //修改
          updatePosition(state.formData).then(()=>{
              ElMessage.success('修改成功');
              closeDialog(); // 关闭弹窗
              emit('positionList')
            }).finally(()=>{
              state.loading = false;
            })
          }
        }
      });
    };
    const resetForm = ()=>{
      state.formData = {        
        positionId: undefined,        
        positionCode: undefined,        
        positionName: undefined,        
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