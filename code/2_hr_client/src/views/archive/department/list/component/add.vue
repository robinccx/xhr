<template>  
  <div class="archive-department-add">
    <!-- 添加或修改科室对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.archive-department-add .el-dialog', '.archive-department-add .el-dialog__header']">添加科室</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="90px">        
        <el-form-item label="科室代码" prop="deptCode">
          <el-input v-model="formData.deptCode" placeholder="请输入科室代码" />
        </el-form-item>        
        <el-form-item label="科室名称" prop="deptName">
          <el-input v-model="formData.deptName" placeholder="请输入科室名称" />
        </el-form-item>          
        <el-form-item label="片区" prop="deptType">
          <el-select v-model="formData.deptType" placeholder="请选择片区" >
            <el-option label="请选择字典生成" value="" />
          </el-select>
        </el-form-item>        
        <el-form-item label="负责人" prop="leader">
          <el-input v-model="formData.leader" placeholder="请输入负责人" />
        </el-form-item>        
        <el-form-item label="电话" prop="phone">
          <el-input v-model="formData.phone" placeholder="请输入电话" />
        </el-form-item>        
        <el-form-item label="邮件" prop="email">
          <el-input v-model="formData.email" placeholder="请输入邮件" />
        </el-form-item>        
        <el-form-item label="属性1" prop="att1">
          <el-input v-model="formData.att1" placeholder="请输入属性1" />
        </el-form-item>        
        <el-form-item label="属性2" prop="att2">
          <el-input v-model="formData.att2" placeholder="请输入属性2" />
        </el-form-item>        
        <el-form-item label="属性3" prop="att3">
          <el-input v-model="formData.att3" placeholder="请输入属性3" />
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
  listDepartment,
  getDepartment,
  delDepartment,
  addDepartment,
  updateDepartment,  
  getUserList,  
} from "/@/api/archive/department";
import {
  DepartmentTableColumns,
  DepartmentInfoData,
  DepartmentTableDataState,
  DepartmentEditState,  
} from "/@/views/archive/department/list/component/model"
export default defineComponent({
  name:"apiV1ArchiveDepartmentEdit",
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
    const state = reactive<DepartmentEditState>({
      loading:false,
      isShowDialog: false,
      formData: {        
        deptId: undefined,        
        parentId: undefined,        
        deptCode: undefined,        
        deptName: undefined,        
        deptType: undefined,        
        leader: undefined,        
        phone: undefined,        
        email: undefined,        
        att1: undefined,        
        att2: undefined,        
        att3: undefined,        
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
        deptId : [
            { required: true, message: "科室ID不能为空", trigger: "blur" }
        ],        
        deptCode : [
            { required: true, message: "科室代码不能为空", trigger: "blur" }
        ],        
        deptName : [
            { required: true, message: "科室名称不能为空", trigger: "blur" }
        ],        
        status : [
            { required: true, message: "状态不能为空", trigger: "blur" }
        ],        
      }
    });
    // 打开弹窗
    const openDialog = () => {
      resetForm();
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
          addDepartment(state.formData).then(()=>{
              ElMessage.success('添加成功');
              closeDialog(); // 关闭弹窗
              emit('departmentList')
            }).finally(()=>{
              state.loading = false;
            })
        }
      });
    };
    const resetForm = ()=>{
      state.formData = {        
        deptId: undefined,        
        parentId: undefined,        
        deptCode: undefined,        
        deptName: undefined,        
        deptType: undefined,        
        leader: undefined,        
        phone: undefined,        
        email: undefined,        
        att1: undefined,        
        att2: undefined,        
        att3: undefined,        
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