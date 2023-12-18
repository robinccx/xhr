<template>
  <!-- 变更记录查询详情抽屉 -->  
  <div class="archive-nurseChangeLog-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>变更记录查询详情</h4>
      </template>
      <el-form ref="formRef" :model="formData" label-width="100px">          
        <el-row>        
          <el-col :span="12">            
            <el-form-item label="">{{ formData.linkedNurseId?formData.linkedNurseId.nurseId:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="变更类型">{{ proxy.getOptionValue(formData.changeType, changeTypeOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ proxy.getOptionValue(formData.fromAreaId, fromAreaIdOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">            
            <el-form-item label="">{{ formData.linkedFromDeptId?formData.linkedFromDeptId.deptId:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ proxy.getOptionValue(formData.toAreaId, toAreaIdOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">            
            <el-form-item label="">{{ formData.linkedToDeptId?formData.linkedToDeptId.deptId:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="变更时间">{{ proxy.parseTime(formData.transDate, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.status }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="">{{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="操作时间">{{ formData.createdBy }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.updatedBy }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="">{{ formData.deletedBy }}</el-form-item>          
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
    listNurseChangeLog,
    getNurseChangeLog,
    delNurseChangeLog,
    addNurseChangeLog,
    updateNurseChangeLog,    
    listNurse,    
    listDepartment,    
    getUserList,    
  } from "/@/api/archive/nurseChangeLog";  
  import {
    NurseChangeLogTableColumns,
    NurseChangeLogInfoData,
    NurseChangeLogTableDataState,
    NurseChangeLogEditState,    
    LinkedNurseChangeLogNurse,    
    LinkedNurseChangeLogDepartment,    
  } from "/@/views/archive/nurseChangeLog/list/component/model"
  export default defineComponent({
    name:"apiV1ArchiveNurseChangeLogDetail",
    components:{      
    },
    props:{      
      nurseIdOptions:{
        type:Array,
        default:()=>[]
      },      
      changeTypeOptions:{
        type:Array,
        default:()=>[]
      },      
      fromAreaIdOptions:{
        type:Array,
        default:()=>[]
      },      
      fromDeptIdOptions:{
        type:Array,
        default:()=>[]
      },      
      toAreaIdOptions:{
        type:Array,
        default:()=>[]
      },      
      toDeptIdOptions:{
        type:Array,
        default:()=>[]
      },      
    },
    setup(props,{emit}) {      
      const {proxy} = <any>getCurrentInstance()
      const formRef = ref<HTMLElement | null>(null);
      const menuRef = ref();      
      const state = reactive<NurseChangeLogEditState>({
        loading:false,
        isShowDialog: false,
        formData: {          
          id: undefined,          
          nurseId: undefined,          
          linkedNurseId:{nurseId:undefined,nurseId:undefined },          
          changeType: undefined,          
          fromAreaId: undefined,          
          fromDeptId: undefined,          
          linkedFromDeptId:{deptId:undefined,deptId:undefined },          
          fromEntityId: undefined,          
          fromEntityName: undefined,          
          toAreaId: undefined,          
          toDeptId: undefined,          
          linkedToDeptId:{deptId:undefined,deptId:undefined },          
          toEntityId: undefined,          
          toEntityName: undefined,          
          transDate: undefined,          
          status: false ,          
          createdAt: undefined,          
          createdBy: undefined,          
          updatedAt: undefined,          
          updatedBy: undefined,          
          deletedAt: undefined,          
          deletedBy: undefined,          
          linkedNurseChangeLogNurse: {            
            nurseId:undefined,    //            
          },          
          linkedNurseChangeLogDepartment: {            
            deptId:undefined,    //            
          },          
        },
        // 表单校验
        rules: {          
          fromEntityName : [
              { required: true, message: "原内容不能为空", trigger: "blur" }
          ],          
          toEntityName : [
              { required: true, message: "新内容不能为空", trigger: "blur" }
          ],          
          status : [
              { required: true, message: "不能为空", trigger: "blur" }
          ],          
        }
      });
        // 打开弹窗
        const openDialog = (row?: NurseChangeLogInfoData) => {
          resetForm();
          if(row) {
            getNurseChangeLog(row.id!).then((res:any)=>{
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
            id: undefined,            
            nurseId: undefined,            
            linkedNurseId:{nurseId:undefined,nurseId:undefined },            
            changeType: undefined,            
            fromAreaId: undefined,            
            fromDeptId: undefined,            
            linkedFromDeptId:{deptId:undefined,deptId:undefined },            
            fromEntityId: undefined,            
            fromEntityName: undefined,            
            toAreaId: undefined,            
            toDeptId: undefined,            
            linkedToDeptId:{deptId:undefined,deptId:undefined },            
            toEntityId: undefined,            
            toEntityName: undefined,            
            transDate: undefined,            
            status: false ,            
            createdAt: undefined,            
            createdBy: undefined,            
            updatedAt: undefined,            
            updatedBy: undefined,            
            deletedAt: undefined,            
            deletedBy: undefined,            
            linkedNurseChangeLogNurse: {              
              nurseId:undefined,    //              
            },            
            linkedNurseChangeLogDepartment: {              
              deptId:undefined,    //              
            },            
          }
        };        
        //关联t_nurse表选项
        const getNurseItemsNurseId = () => {
          emit("getNurseItemsNurseId")
        }
        const getNurseIdOp = computed(()=>{
          getNurseItemsNurseId()
          return props.nurseIdOptions
        })        
        //关联t_department表选项
        const getDepartmentItemsFromDeptId = () => {
          emit("getDepartmentItemsFromDeptId")
        }
        const getFromDeptIdOp = computed(()=>{
          getDepartmentItemsFromDeptId()
          return props.fromDeptIdOptions
        })        
        //关联t_department表选项
        const getDepartmentItemsToDeptId = () => {
          emit("getDepartmentItemsToDeptId")
        }
        const getToDeptIdOp = computed(()=>{
          getDepartmentItemsToDeptId()
          return props.toDeptIdOptions
        })        
        return {
          proxy,
          openDialog,
          closeDialog,
          onCancel,
          menuRef,
          formRef,          
          getNurseItemsNurseId,
          getNurseIdOp,          
          getDepartmentItemsFromDeptId,
          getFromDeptIdOp,          
          getDepartmentItemsToDeptId,
          getToDeptIdOp,          
          ...toRefs(state),
        };
      }
  })
</script>
<style scoped>  
  .archive-nurseChangeLog-detail :deep(.el-form-item--large .el-form-item__label){
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