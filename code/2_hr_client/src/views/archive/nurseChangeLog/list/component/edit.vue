<template>  
  <div class="archive-nurseChangeLog-edit">
    <!-- 添加或修改变更记录查询对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.archive-nurseChangeLog-edit .el-dialog', '.archive-nurseChangeLog-edit .el-dialog__header']">{{(!formData.id || formData.id==0?'添加':'修改')+'变更记录查询'}}</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="90px">        
        <el-form-item label="" prop="nurseId">
          <el-input v-model="formData.nurseId" placeholder="请输入" />
        </el-form-item>          
        <el-form-item label="变更类型" prop="changeType">
          <el-select v-model="formData.changeType" placeholder="请选择变更类型" >
            <el-option
              v-for="dict in changeTypeOptions"
              :key="dict.value"
              :label="dict.label"              
                  :value="dict.value"              
            ></el-option>
          </el-select>
        </el-form-item>          
        <el-form-item label="" prop="fromAreaId">
          <el-select v-model="formData.fromAreaId" placeholder="请选择" >
            <el-option
              v-for="dict in fromAreaIdOptions"
              :key="dict.value"
              :label="dict.label"              
                  :value="dict.value"              
            ></el-option>
          </el-select>
        </el-form-item>          
        <el-form-item label="" prop="fromDeptId">
          <el-select v-model="formData.fromDeptId" placeholder="请选择"   @click.native="getDepartmentItemsFromDeptId">
              <el-option              
                  v-for="item in fromDeptIdOptions"              
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="" prop="toAreaId">
          <el-input v-model="formData.toAreaId" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="toDeptId">
          <el-input v-model="formData.toDeptId" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="变更时间" prop="transDate">
          <el-date-picker clearable  style="width: 200px"
            v-model="formData.transDate"
            type="datetime"
            placeholder="选择变更时间">
          </el-date-picker>
        </el-form-item>          
        <el-form-item label="" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio>请选择字典生成</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="" prop="deletedBy">
          <el-input v-model="formData.deletedBy" placeholder="请输入" />
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
  name:"apiV1ArchiveNurseChangeLogEdit",
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
        changeType: undefined,        
        fromAreaId: undefined,        
        fromDeptId: undefined,        
        fromEntityId: undefined,        
        fromEntityName: undefined,        
        toAreaId: undefined,        
        toDeptId: undefined,        
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
        getNurseItemsNurseId()        
        getDepartmentItemsFromDeptId()        
        getDepartmentItemsToDeptId()        
        getNurseChangeLog(row.id!).then((res:any)=>{
          const data = res.data;          
          data.changeType = ''+data.changeType          
          data.fromAreaId = ''+data.fromAreaId          
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
          if(!state.formData.id || state.formData.id===0){
            //添加
          addNurseChangeLog(state.formData).then(()=>{
              ElMessage.success('添加成功');
              closeDialog(); // 关闭弹窗
              emit('nurseChangeLogList')
            }).finally(()=>{
              state.loading = false;
            })
          }else{
            //修改
          updateNurseChangeLog(state.formData).then(()=>{
              ElMessage.success('修改成功');
              closeDialog(); // 关闭弹窗
              emit('nurseChangeLogList')
            }).finally(()=>{
              state.loading = false;
            })
          }
        }
      });
    };
    const resetForm = ()=>{
      state.formData = {        
        id: undefined,        
        nurseId: undefined,        
        changeType: undefined,        
        fromAreaId: undefined,        
        fromDeptId: undefined,        
        fromEntityId: undefined,        
        fromEntityName: undefined,        
        toAreaId: undefined,        
        toDeptId: undefined,        
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
    //关联t_department表选项
    const getDepartmentItemsFromDeptId = () => {
      emit("getDepartmentItemsFromDeptId")
    }    
    //关联t_department表选项
    const getDepartmentItemsToDeptId = () => {
      emit("getDepartmentItemsToDeptId")
    }    
    return {
      proxy,
      openDialog,
      closeDialog,
      onCancel,
      onSubmit,
      menuRef,
      formRef,      
      getNurseItemsNurseId,      
      getDepartmentItemsFromDeptId,      
      getDepartmentItemsToDeptId,      
      ...toRefs(state),
    };
  }
})
</script>
<style scoped>  
</style>