<template>  
  <div class="archive-nurseImport-add">
    <!-- 添加或修改导入对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.archive-nurseImport-add .el-dialog', '.archive-nurseImport-add .el-dialog__header']">添加导入</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="90px">        
        <el-form-item label="" prop="sessionId">
          <el-input v-model="formData.sessionId" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="nurseCode">
          <el-input v-model="formData.nurseCode" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="nurseName">
          <el-input v-model="formData.nurseName" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="startDate">
          <el-input v-model="formData.startDate" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="endDate">
          <el-input v-model="formData.endDate" placeholder="请输入" />
        </el-form-item>          
        <el-form-item label="" prop="sex">
          <el-select v-model="formData.sex" placeholder="请选择" >
            <el-option label="请选择字典生成" value="" />
          </el-select>
        </el-form-item>        
        <el-form-item label="" prop="birthday">
          <el-input v-model="formData.birthday" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="areaId">
          <el-input v-model="formData.areaId" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="areaCode">
          <el-input v-model="formData.areaCode" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="deptId">
          <el-input v-model="formData.deptId" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="deptCode">
          <el-input v-model="formData.deptCode" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="educationId">
          <el-input v-model="formData.educationId" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="educationCode">
          <el-input v-model="formData.educationCode" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="titleId">
          <el-input v-model="formData.titleId" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="titleCode">
          <el-input v-model="formData.titleCode" placeholder="请输入" />
        </el-form-item>          
        <el-form-item label="" prop="staffType">
          <el-select v-model="formData.staffType" placeholder="请选择" >
            <el-option label="请选择字典生成" value="" />
          </el-select>
        </el-form-item>        
        <el-form-item label="" prop="staffTypeCode">
          <el-input v-model="formData.staffTypeCode" placeholder="请输入" />
        </el-form-item>        
        <el-form-item label="" prop="note">
          <el-input v-model="formData.note" placeholder="请输入" />
        </el-form-item>          
        <el-form-item label="" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio>请选择字典生成</el-radio>
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
  name:"apiV1ArchiveNurseImportEdit",
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
          addNurseImport(state.formData).then(()=>{
              ElMessage.success('添加成功');
              closeDialog(); // 关闭弹窗
              emit('nurseImportList')
            }).finally(()=>{
              state.loading = false;
            })
        }
      });
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