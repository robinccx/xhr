<template>
  <div class="archive-nurse-add">
    <!-- 添加或修改护士对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.archive-nurse-add .el-dialog', '.archive-nurse-add .el-dialog__header']">添加护士</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-row :gutter="15">
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="编号" prop="nurseCode">
              <el-input v-model="formData.nurseCode" placeholder="请输入编号" />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="名称" prop="nurseName">
              <el-input v-model="formData.nurseName" placeholder="请输入名称" />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="性别" prop="sex">
              <el-select v-model="formData.sex" placeholder="请选择性别" class="w100">
                <el-option v-for="dict in sexOptions" :key="dict.value" :label="dict.label"
                  :value="dict.value"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="出生年月" prop="birthday">
              <el-date-picker clearable class="w100" v-model="formData.birthday" type="date" placeholder="选择出生年月">
              </el-date-picker>
            </el-form-item></el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="院区" prop="areaId">
              <el-select v-model="formData.areaId" placeholder="请选择院区" clearable class="w100">
                <el-option v-for="dict in areaIdOptions" :key="dict.value" :label="dict.label"
                  :value="dict.value"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="科室" prop="deptId">
              <el-select v-model="formData.deptId" placeholder="请选择科室" @click.native="getDepartmentItemsDeptId" clearable
                class="w100">
                <el-option v-for="item in deptIdOptions" :key="item.key" :label="item.value"
                  :value="item.key"></el-option>
              </el-select>
            </el-form-item></el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="职务" prop="positionId">
              <el-select v-model="formData.positionId" placeholder="请选择职务" @click.native="getPositionItemsPositionId"
                clearable class="w100">
                <el-option v-for="item in positionIdOptions" :key="item.key" :label="item.value"
                  :value="item.key"></el-option>
              </el-select>
            </el-form-item></el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="学历" prop="educationId">
              <el-select v-model="formData.educationId" placeholder="请选择学历" @click.native="getEducationItemsEducationId"
                clearable class="w100">
                <el-option v-for="item in educationIdOptions" :key="item.key" :label="item.value"
                  :value="item.key"></el-option>
              </el-select>
            </el-form-item></el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="职称" prop="titleId">
              <el-select v-model="formData.titleId" placeholder="请选择职称" @click.native="getTitleItemsTitleId" clearable
                class="w100">
                <el-option v-for="item in titleIdOptions" :key="item.key" :label="item.value"
                  :value="item.key"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="人员类别" prop="staffType">
              <el-select v-model="formData.staffType" placeholder="请选择人员类别" clearable class="w100">
                <el-option v-for="dict in staffTypeOptions" :key="dict.value" :label="dict.label"
                  :value="dict.value"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="身份证号" prop="idNo">
              <el-input v-model="formData.idNo" placeholder="身份证号" />
            </el-form-item>
          </el-col>          
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="入职日期" prop="startDate">
              <el-date-picker clearable class="w100" v-model="formData.startDate" type="date" placeholder="选择入职日期">
              </el-date-picker>
            </el-form-item></el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="执业证书号" prop="certNo">
              <el-input v-model="formData.certNo" placeholder="请输入执业证书号" />
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="执业注册有效期" prop="certEnddate">
              <el-date-picker clearable class="w100" v-model="formData.certEnddate" type="date" placeholder="选择执业注册有效期">
              </el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="执业注册时间" prop="regDate">
              <el-date-picker clearable class="w100" v-model="formData.regDate" type="date" placeholder="选择执业注册时间">
              </el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12" class="mb20">
            <el-form-item label="离职日期" prop="endDate">
              <el-date-picker clearable class="w100" v-model="formData.endDate" type="date" placeholder="选择离职日期">
              </el-date-picker>
            </el-form-item></el-col>
        </el-row>

        <el-form-item label="备注" prop="note">
          <el-input v-model="formData.note" type="textarea" placeholder="请输入备注" />
        </el-form-item>

        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio v-for="dict in statusOptions" :key="dict.value" :label="dict.value">{{ dict.label }}</el-radio>
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
import { reactive, toRefs, defineComponent, ref, unref, getCurrentInstance } from 'vue';
import { ElMessageBox, ElMessage, FormInstance, UploadProps } from 'element-plus';
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
  linkedNursePosition,
} from "/@/views/archive/nurse/list/component/model"
export default defineComponent({
  name: "apiV1ArchiveNurseEdit",
  components: {
  },
  props: {
    sexOptions: {
      type: Array,
      default: () => []
    },
    areaIdOptions: {
      type: Array,
      default: () => []
    },
    deptIdOptions: {
      type: Array,
      default: () => []
    },
    positionIdOptions: {
      type: Array,
      default: () => []
    },    
    educationIdOptions: {
      type: Array,
      default: () => []
    },
    titleIdOptions: {
      type: Array,
      default: () => []
    },
    staffTypeOptions: {
      type: Array,
      default: () => []
    },
    statusOptions: {
      type: Array,
      default: () => []
    },
  },
  setup(props, { emit }) {
    const { proxy } = <any>getCurrentInstance()
    const formRef = ref<HTMLElement | null>(null);
    const menuRef = ref();
    const state = reactive<NurseEditState>({
      loading: false,
      isShowDialog: false,
      formData: {
        nurseId: undefined,
        nurseCode: undefined,
        nurseName: undefined,
        sex: undefined,
        birthday: undefined,
        areaId: undefined,
        deptId: undefined,
        positionId: undefined,
        startDate: undefined,
        educationId: undefined,
        titleId: undefined,
        staffType: undefined,
        note: undefined,
        endDate: undefined,
        idNo: undefined,
        certNo: undefined,
        certEnddate: undefined,
        regDate: undefined,
        status: false,
        createdAt: undefined,
        createdBy: undefined,
        updatedAt: undefined,
        updatedBy: undefined,
        deletedAt: undefined,
        deletedBy: undefined,
        linkedNurseDepartment: {
          deptId: undefined,    //          
          deptName: undefined,    //          
        },
        linkedNurseEducation: {
          educationId: undefined,    //          
          educationName: undefined,    //          
        },
        linkedNurseTitle: {
          titleId: undefined,    //          
          titleName: undefined,    //          
        },
        linkedNursePosition: {
          positionId: undefined,    //          
          positionName: undefined,    //          
        },
      },
      // 表单校验
      rules: {
        nurseId: [
          { required: true, message: "ID不能为空", trigger: "blur" }
        ],
        nurseName: [
          { required: true, message: "名称不能为空", trigger: "blur" }
        ],
        status: [
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
          addNurse(state.formData).then(() => {
            ElMessage.success('添加成功');
            closeDialog(); // 关闭弹窗
            emit('nurseList')
          }).finally(() => {
            state.loading = false;
          })
        }
      });
    };
    const resetForm = () => {
      state.formData = {
        nurseId: undefined,
        nurseCode: undefined,
        nurseName: undefined,
        sex: undefined,
        birthday: undefined,
        areaId: undefined,
        deptId: undefined,
        startDate: undefined,
        educationId: undefined,
        positionId: undefined,
        titleId: undefined,
        staffType: undefined,
        note: undefined,
        endDate: undefined,
        idNo: undefined,
        certNo: undefined,
        certEnddate: undefined,
        regDate: undefined,
        status: false,
        createdAt: undefined,
        createdBy: undefined,
        updatedAt: undefined,
        updatedBy: undefined,
        deletedAt: undefined,
        deletedBy: undefined,
        linkedNurseDepartment: {
          deptId: undefined,    //          
          deptName: undefined,    //          
        },
        linkedNurseEducation: {
          educationId: undefined,    //          
          educationName: undefined,    //          
        },
        linkedNurseTitle: {
          titleId: undefined,    //          
          titleName: undefined,    //          
        },
      }
    };
    //关联t_department表选项
    const getDepartmentItemsDeptId = () => {
      emit("getDepartmentItemsDeptId")
    }
    //关联t_position表选项
    const getPositionItemsPositionId = () => {
      emit("getPositionItemsPositionId")
    }
    //关联t_education表选项
    const getEducationItemsEducationId = () => {
      emit("getEducationItemsEducationId")
    }
    //关联t_title表选项
    const getTitleItemsTitleId = () => {
      emit("getTitleItemsTitleId")
    }
    return {
      proxy,
      openDialog,
      closeDialog,
      onCancel,
      onSubmit,
      menuRef,
      formRef,
      getDepartmentItemsDeptId,
      getPositionItemsPositionId,
      getEducationItemsEducationId,
      getTitleItemsTitleId,
      ...toRefs(state),
    };
  }
})
</script>
<style scoped></style>