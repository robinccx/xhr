<template>
  <!-- 导入详情抽屉 -->  
  <div class="archive-nurseImport-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>导入详情</h4>
      </template>
      <el-form ref="formRef" :model="formData" label-width="100px">          
        <el-row>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.nurseId }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.sessionId }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.nurseCode }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.nurseName }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.startDate }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.endDate }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.sex }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.birthday }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.areaId }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.areaCode }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.deptId }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.deptCode }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.educationId }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.educationCode }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.titleId }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.titleCode }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.staffType }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.staffTypeCode }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.note }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.status }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="">{{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.createdBy }}</el-form-item>          
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
    listNurseImport,
    getNurseImport,
    delNurseImport,
    addNurseImport,
    updateNurseImport,    
    getUserList,    
  } from "/@/api/archive/nurseImport";  
  import {
    NurseImportTableColumns,
    NurseImportInfoData,
    NurseImportTableDataState,
    NurseImportEditState,    
  } from "/@/views/archive/nurseImport/list/component/model"
  export default defineComponent({
    name:"apiV1ArchiveNurseImportDetail",
    components:{      
    },
    props:{      
    },
    setup(props,{emit}) {      
      const {proxy} = <any>getCurrentInstance()
      const formRef = ref<HTMLElement | null>(null);
      const menuRef = ref();      
      const state = reactive<NurseImportEditState>({
        loading:false,
        isShowDialog: false,
        formData: {          
          nurseId: undefined,          
          sessionId: undefined,          
          nurseCode: undefined,          
          nurseName: undefined,          
          startDate: undefined,          
          endDate: undefined,          
          sex: undefined,          
          birthday: undefined,          
          areaId: undefined,          
          areaCode: undefined,          
          deptId: undefined,          
          deptCode: undefined,          
          educationId: undefined,          
          educationCode: undefined,          
          titleId: undefined,          
          titleCode: undefined,          
          staffType: undefined,          
          staffTypeCode: undefined,          
          note: undefined,          
          status: false ,          
          createdAt: undefined,          
          createdBy: undefined,          
        },
        // 表单校验
        rules: {          
          nurseId : [
              { required: true, message: "不能为空", trigger: "blur" }
          ],          
          nurseName : [
              { required: true, message: "不能为空", trigger: "blur" }
          ],          
          status : [
              { required: true, message: "不能为空", trigger: "blur" }
          ],          
        }
      });
        // 打开弹窗
        const openDialog = (row?: NurseImportInfoData) => {
          resetForm();
          if(row) {
            getNurseImport(row.nurseId!).then((res:any)=>{
              const data = res.data;              
              let listUid = [];              
              listUid.push(data.createdBy)              
              getUserList(listUid).then((response:any) =>{
                let users = response.data.list||[]
                users.forEach((user:any)=>{                  
                  if(data.createdBy==user.id){
                    data.createdBy = user.userNickname
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
            nurseId: undefined,            
            sessionId: undefined,            
            nurseCode: undefined,            
            nurseName: undefined,            
            startDate: undefined,            
            endDate: undefined,            
            sex: undefined,            
            birthday: undefined,            
            areaId: undefined,            
            areaCode: undefined,            
            deptId: undefined,            
            deptCode: undefined,            
            educationId: undefined,            
            educationCode: undefined,            
            titleId: undefined,            
            titleCode: undefined,            
            staffType: undefined,            
            staffTypeCode: undefined,            
            note: undefined,            
            status: false ,            
            createdAt: undefined,            
            createdBy: undefined,            
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
  .archive-nurseImport-detail :deep(.el-form-item--large .el-form-item__label){
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