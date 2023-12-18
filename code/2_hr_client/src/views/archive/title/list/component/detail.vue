<template>
  <!-- 职称详情抽屉 -->  
  <div class="archive-title-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>职称详情</h4>
      </template>
      <el-form ref="formRef" :model="formData" label-width="100px">          
        <el-row>        
          <el-col :span="12">          
            <el-form-item label="职称">{{ formData.titleName }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="次序">{{ formData.seq }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="状态">{{ proxy.getOptionValue(formData.status, statusOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="创建时间">{{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="创建用户">{{ formData.createdBy }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="修改用户">{{ formData.updatedBy }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="删除用户">{{ formData.deletedBy }}</el-form-item>          
          </el-col>      
        </el-row>      
      </el-form>
    </el-drawer>
  </div>
</template>
<script lang="ts">
  import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance,computed } from 'vue';
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
    name:"apiV1ArchiveTitleDetail",
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
        const openDialog = (row?: TitleInfoData) => {
          resetForm();
          if(row) {
            getTitle(row.titleId!).then((res:any)=>{
              const data = res.data;              
              let listUid = [];              
              listUid.push(data.createdBy,data.updatedBy)              
              getUserList(listUid).then((response:any) =>{
                let users = response.data.list||[]
                users.forEach((user:any)=>{                  
                  if(data.createdBy==user.id){
                    data.createdBy = user.userNickname
                  }                  
                  if(data.updatedBy==user.id){
                    data.updatedBy = user.userNickname
                  }                  
                })
                state.formData = data;
              })              
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
          menuRef,
          formRef,          
          ...toRefs(state),
        };
      }
  })
</script>
<style scoped>  
  .archive-title-detail :deep(.el-form-item--large .el-form-item__label){
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