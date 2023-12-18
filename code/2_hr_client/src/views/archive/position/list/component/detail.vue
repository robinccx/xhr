<template>
  <!-- 职务详情抽屉 -->  
  <div class="archive-position-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>职务详情</h4>
      </template>
      <el-form ref="formRef" :model="formData" label-width="100px">      
      </el-form>
    </el-drawer>
  </div>
</template>
<script lang="ts">
  import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance,computed } from 'vue';
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
    name:"apiV1ArchivePositionDetail",
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
          menuRef,
          formRef,          
          ...toRefs(state),
        };
      }
  })
</script>
<style scoped>  
  .archive-position-detail :deep(.el-form-item--large .el-form-item__label){
    font-weight: bolder;
  }
  .pic-block{
    margin-right: 8px;
  }
  .file-block{
    width: 100%;
    border: 1px solid var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
    margin-bottom: 5px;
    padding: 3px 6px;
  }
  .ml-2{margin-right: 5px;}
</style>