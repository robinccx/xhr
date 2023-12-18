<template>
  <!-- 护士详情抽屉 -->  
  <div class="archive-nurse-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>护士详情</h4>
      </template>
      <el-form ref="formRef" :model="formData" label-width="100px">          
        <el-row>        
          <el-col :span="12">          
            <el-form-item label="ID">{{ formData.nurseId }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="编号">{{ formData.nurseCode }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="名称">{{ formData.nurseName }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="性别">{{ proxy.getOptionValue(formData.sex, sexOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="出生年月">{{ proxy.parseTime(formData.birthday, '{y}-{m}-{d}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">            
            <el-form-item label="科室">{{ formData.linkedDeptId?formData.linkedDeptId.deptName:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="入职日期">{{ proxy.parseTime(formData.startDate, '{y}-{m}-{d}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">            
            <el-form-item label="学历">{{ formData.linkedEducationId?formData.linkedEducationId.educationName:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">            
            <el-form-item label="职称">{{ formData.linkedTitleId?formData.linkedTitleId.titleName:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="人员类别">{{ proxy.getOptionValue(formData.staffType, staffTypeOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="备注">{{ formData.note }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="离职日期">{{ proxy.parseTime(formData.endDate, '{y}-{m}-{d}') }}</el-form-item>
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
          <el-col :span="12">          
            <el-form-item label="院区">{{ proxy.getOptionValue(formData.areaId, areaIdOptions,'value','label') }}</el-form-item>          
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
    listNurse,
    getNurse,
    delNurse,
    addNurse,
    updateNurse,    
    listDepartment,    
    listEducation,    
    listTitle,    
    getUserList,    
  } from "/@/api/archive/nurse";  
  import {
    NurseTableColumns,
    NurseInfoData,
    NurseTableDataState,
    NurseEditState,    
    LinkedNurseDepartment,    
    LinkedNurseEducation,    
    LinkedNurseTitle,    
  } from "/@/views/archive/nurse/list/component/model"
  export default defineComponent({
    name:"apiV1ArchiveNurseDetail",
    components:{      
    },
    props:{      
      sexOptions:{
        type:Array,
        default:()=>[]
      },      
      areaIdOptions:{
        type:Array,
        default:()=>[]
      },      
      deptIdOptions:{
        type:Array,
        default:()=>[]
      },      
      educationIdOptions:{
        type:Array,
        default:()=>[]
      },      
      titleIdOptions:{
        type:Array,
        default:()=>[]
      },      
      staffTypeOptions:{
        type:Array,
        default:()=>[]
      },      
      statusOptions:{
        type:Array,
        default:()=>[]
      },      
    },
    setup(props,{emit}) {      
      const {proxy} = <any>getCurrentInstance()
      const formRef = ref<HTMLElement | null>(null);
      const menuRef = ref();      
      const state = reactive<NurseEditState>({
        loading:false,
        isShowDialog: false,
        formData: {          
          nurseId: undefined,          
          nurseCode: undefined,          
          nurseName: undefined,          
          sex: undefined,          
          birthday: undefined,          
          areaId: undefined,          
          deptId: undefined,          
          linkedDeptId:{deptId:undefined,deptName:undefined },          
          startDate: undefined,          
          educationId: undefined,          
          linkedEducationId:{educationId:undefined,educationName:undefined },          
          titleId: undefined,          
          linkedTitleId:{titleId:undefined,titleName:undefined },          
          staffType: undefined,          
          note: undefined,          
          endDate: undefined,          
          status: false ,          
          createdAt: undefined,          
          createdBy: undefined,          
          updatedAt: undefined,          
          updatedBy: undefined,          
          deletedAt: undefined,          
          deletedBy: undefined,          
          linkedNurseDepartment: {            
            deptId:undefined,    //            
            deptName:undefined,    //            
          },          
          linkedNurseEducation: {            
            educationId:undefined,    //            
            educationName:undefined,    //            
          },          
          linkedNurseTitle: {            
            titleId:undefined,    //            
            titleName:undefined,    //            
          },          
        },
        // 表单校验
        rules: {          
          nurseId : [
              { required: true, message: "ID不能为空", trigger: "blur" }
          ],          
          nurseName : [
              { required: true, message: "名称不能为空", trigger: "blur" }
          ],          
          status : [
              { required: true, message: "状态不能为空", trigger: "blur" }
          ],          
        }
      });
        // 打开弹窗
        const openDialog = (row?: NurseInfoData) => {
          resetForm();
          if(row) {
            getNurse(row.nurseId!).then((res:any)=>{
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
            nurseId: undefined,            
            nurseCode: undefined,            
            nurseName: undefined,            
            sex: undefined,            
            birthday: undefined,            
            areaId: undefined,            
            deptId: undefined,            
            linkedDeptId:{deptId:undefined,deptName:undefined },            
            startDate: undefined,            
            educationId: undefined,            
            linkedEducationId:{educationId:undefined,educationName:undefined },            
            titleId: undefined,            
            linkedTitleId:{titleId:undefined,titleName:undefined },            
            staffType: undefined,            
            note: undefined,            
            endDate: undefined,            
            status: false ,            
            createdAt: undefined,            
            createdBy: undefined,            
            updatedAt: undefined,            
            updatedBy: undefined,            
            deletedAt: undefined,            
            deletedBy: undefined,            
            linkedNurseDepartment: {              
              deptId:undefined,    //              
              deptName:undefined,    //              
            },            
            linkedNurseEducation: {              
              educationId:undefined,    //              
              educationName:undefined,    //              
            },            
            linkedNurseTitle: {              
              titleId:undefined,    //              
              titleName:undefined,    //              
            },            
          }
        };        
        //关联t_department表选项
        const getDepartmentItemsDeptId = () => {
          emit("getDepartmentItemsDeptId")
        }
        const getDeptIdOp = computed(()=>{
          getDepartmentItemsDeptId()
          return props.deptIdOptions
        })        
        //关联t_education表选项
        const getEducationItemsEducationId = () => {
          emit("getEducationItemsEducationId")
        }
        const getEducationIdOp = computed(()=>{
          getEducationItemsEducationId()
          return props.educationIdOptions
        })        
        //关联t_title表选项
        const getTitleItemsTitleId = () => {
          emit("getTitleItemsTitleId")
        }
        const getTitleIdOp = computed(()=>{
          getTitleItemsTitleId()
          return props.titleIdOptions
        })        
        return {
          proxy,
          openDialog,
          closeDialog,
          onCancel,
          menuRef,
          formRef,          
          getDepartmentItemsDeptId,
          getDeptIdOp,          
          getEducationItemsEducationId,
          getEducationIdOp,          
          getTitleItemsTitleId,
          getTitleIdOp,          
          ...toRefs(state),
        };
      }
  })
</script>
<style scoped>  
  .archive-nurse-detail :deep(.el-form-item--large .el-form-item__label){
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