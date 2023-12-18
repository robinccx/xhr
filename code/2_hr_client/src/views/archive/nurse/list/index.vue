<template>
  <div class="archive-nurse-container">
    <el-card shadow="hover">
      <div class="archive-nurse-search mb15">
        <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="80px">
          <el-row>
            <el-col :span="5" class="colBlock">
              <el-form-item label="院区" prop="areaId">
                <el-select v-model="tableData.param.areaId" placeholder="请选择院区" clearable>
                  <el-option v-for="dict in area_id" :key="dict.value" :label="dict.label" :value="dict.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" class="colBlock">
              <el-form-item label="名称" prop="nurseName">
                <el-input v-model="tableData.param.nurseName" placeholder="请输入名称" clearable
                  @keyup.enter.native="nurseList" />
              </el-form-item>
            </el-col>
            <el-col :span="5" class="colBlock">
              <el-form-item label="科室" prop="deptId">
                <el-select v-model="tableData.param.deptId" placeholder="请选择科室" clearable
                  @click.native="getDepartmentItemsDeptId">
                  <el-option v-for="item in deptIdOptions" :key="item.key" :label="item.value" :value="item.key" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" class="colBlock">
              <el-form-item label="状态" prop="status">
                <el-select v-model="tableData.param.status" placeholder="请选择状态" clearable>
                  <el-option v-for="dict in staff_status" :key="dict.value" :label="dict.label" :value="dict.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="4" :class="!showAll ? 'colBlock' : 'colNone'">
              <el-form-item>
                <el-button type="primary" @click="nurseList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                <el-button @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                <el-button type="primary" link @click="toggleSearch">
                  {{ word }}
                  <el-icon v-show="showAll"><ele-ArrowUp /></el-icon>
                  <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                </el-button>
              </el-form-item>
            </el-col>
            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="职务" prop="deptId">
                <el-select v-model="tableData.param.positionId" placeholder="请选择职务" clearable
                  @click.native="getPositionItemsPositionId">
                  <el-option v-for="item in positionIdOptions" :key="item.key" :label="item.value" :value="item.key" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="学历" prop="educationId">
                <el-select v-model="tableData.param.educationId" placeholder="请选择学历" clearable
                  @click.native="getEducationItemsEducationId">
                  <el-option v-for="item in educationIdOptions" :key="item.key" :label="item.value" :value="item.key" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="职称" prop="titleId">
                <el-select v-model="tableData.param.titleId" placeholder="请选择职称" clearable
                  @click.native="getTitleItemsTitleId">
                  <el-option v-for="item in titleIdOptions" :key="item.key" :label="item.value" :value="item.key" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="人员类别" prop="staffType">
                <el-select v-model="tableData.param.staffType" placeholder="请选择人员类别" clearable>
                  <el-option v-for="dict in staff_type" :key="dict.value" :label="dict.label" :value="dict.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="编号" prop="nurseCode">
                <el-input v-model="tableData.param.nurseCode" placeholder="请输入编号" clearable
                  @keyup.enter.native="nurseList" />
              </el-form-item>
            </el-col>
            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="注册" prop="regStatus">
                <el-select v-model="tableData.param.regStatus" placeholder="请选择状态" clearable>
                  <el-option v-for="dict in sys_yes_no" :key="dict.value" :label="dict.label" :value="dict.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="入职日期" prop="startDate">
                <el-date-picker clearable style="width: 200px" v-model="tableData.param.startDate"
                  value-format="YYYY-MM-DD" type="daterange" range-separator="至" start-placeholder="开始日期"
                  end-placeholder="结束日期"></el-date-picker>
              </el-form-item>
            </el-col>

            <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="离职日期" prop="endDate">
                <el-date-picker clearable style="width: 200px" v-model="tableData.param.endDate" value-format="YYYY-MM-DD"
                  type="daterange" range-separator="至" start-placeholder="开始日期" end-placeholder="结束日期"></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item>
                <el-button type="primary" @click="nurseList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                <el-button @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                <el-button type="primary" link @click="toggleSearch">
                  {{ word }}
                  <el-icon v-show="showAll"><ele-ArrowUp /></el-icon>
                  <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                </el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button type="primary" @click="handleAdd"
              v-auth="'api/v1/archive/nurse/add'"><el-icon><ele-Plus /></el-icon>新增</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button type="success" :disabled="single" @click="handleUpdate(null)"
              v-auth="'api/v1/archive/nurse/edit'"><el-icon><ele-Edit /></el-icon>修改</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button type="info" @click="handleImport"><el-icon><ele-UploadFilled /></el-icon>导入</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button type="info" @click="handleRegInfoImport"><el-icon><ele-UploadFilled /></el-icon>信息导入</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button type="danger" :disabled="multiple" @click="handleDelete(null)"
              v-auth="'api/v1/archive/nurse/delete'"><el-icon><ele-Delete /></el-icon>删除</el-button>
          </el-col>
        </el-row>
      </div>
      <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
        <el-table-column type="expand">
          <template #default="props">
            <el-form class="table-expand">
              <el-form-item label="身份证号:"><span>{{ props.row.idNo }}</span></el-form-item>
              <el-form-item label="执业证书号:"><span>{{ props.row.certNo }}</span></el-form-item>
              <el-form-item label="执业注册有效期:"><span>{{ proxy.parseTime(props.row.certEnddate, '{y}-{m}-{d}')
              }}</span></el-form-item>
              <el-form-item label="执业注册时间:"><span>{{ proxy.parseTime(props.row.regDate, '{y}-{m}-{d}')
              }}</span></el-form-item>
              <el-form-item label="离职日期:">
                <span>{{ proxy.parseTime(props.row.endDate, '{y}-{m}-{d}') }}</span></el-form-item>
                
              <el-form-item label="备注:"><span>{{ props.row.note }}</span></el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="状态" align="center" prop="status" :formatter="statusFormat" min-width="50px" />
        <el-table-column label="编号" align="center" prop="nurseCode" min-width="85px" />
        <el-table-column label="名称" align="center" prop="nurseName" min-width="85px" />
        <el-table-column label="性别" align="center" prop="sex" :formatter="sexFormat" min-width="50px" />
        <el-table-column label="出生年月" align="center" prop="birthday" min-width="90px">
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.birthday, '{y}-{m}-{d}') }}</span>
          </template>
        </el-table-column>
        <el-table-column label="院区" align="center" prop="areaId" :formatter="areaIdFormat" min-width="80px" />
        <el-table-column label="科室" align="center" prop="linkedDeptId.deptName" min-width="120px" />
        <el-table-column label="职务" align="center" prop="linkedPositionId.positionName" min-width="80px" />
        <el-table-column label="入职日期" align="center" prop="startDate" min-width="90px">
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.startDate, '{y}-{m}-{d}') }}</span>
          </template>
        </el-table-column>
        <el-table-column label="学历" align="center" prop="linkedEducationId.educationName" min-width="55px" />
        <el-table-column label="职称" align="center" prop="linkedTitleId.titleName" min-width="55px" />
        <el-table-column label="人员类别" align="center" prop="staffType" :formatter="staffTypeFormat" min-width="70px" />
        <!-- <el-table-column label="职业证书" align="center" prop="certNo" min-width="100px" /> -->
        <!-- <el-table-column label="职业证书有效期" align="center" prop="certEndDate" min-width="100px">
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.certEndDate, '{y}-{m}-{d}') }}</span>
          </template>
        </el-table-column> -->
        <!-- <el-table-column label="注册与变更日期" align="center" prop="regDate" min-width="100px">
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.regDate, '{y}-{m}-{d}') }}</span>
          </template>
        </el-table-column> -->
        <!-- <el-table-column label="离职日期" align="center" prop="endDate" min-width="100px">
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.endDate, '{y}-{m}-{d}') }}</span>
          </template>
        </el-table-column> -->
        <!-- <el-table-column label="备注" align="center" prop="note" min-width="0px" /> -->
        <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleUpdate(scope.row)"
              v-auth="'api/v1/archive/nurse/edit'"><el-icon><ele-EditPen /></el-icon>修改</el-button>
            <el-button type="primary" link @click="handleDelete(scope.row)"
              v-auth="'api/v1/archive/nurse/delete'"><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="tableData.total > 0" :total="tableData.total" v-model:page="tableData.param.pageNum"
        v-model:limit="tableData.param.pageSize" @pagination="nurseList" />
    </el-card>
    <apiV1ArchiveNurseAdd ref="addRef" :sexOptions="sys_user_sex" :areaIdOptions="area_id" :deptIdOptions="deptIdOptions"
      @getDepartmentItemsDeptId="getDepartmentItemsDeptId" :educationIdOptions="educationIdOptions"
      @getEducationItemsEducationId="getEducationItemsEducationId" :titleIdOptions="titleIdOptions"
      @getTitleItemsTitleId="getTitleItemsTitleId" :staffTypeOptions="staff_type" :statusOptions="staff_status"
      @getPositionItemsPositionId="getPositionItemsPositionId" :positionIdOptions="positionIdOptions"
      @nurseList="nurseList"></apiV1ArchiveNurseAdd>
    <apiV1ArchiveNurseEdit ref="editRef" :sexOptions="sys_user_sex" :areaIdOptions="area_id"
      :deptIdOptions="deptIdOptions" @getDepartmentItemsDeptId="getDepartmentItemsDeptId"
      :educationIdOptions="educationIdOptions" @getEducationItemsEducationId="getEducationItemsEducationId"
      :titleIdOptions="titleIdOptions" @getTitleItemsTitleId="getTitleItemsTitleId" :staffTypeOptions="staff_type"
      :statusOptions="staff_status" @getPositionItemsPositionId="getPositionItemsPositionId"
      :positionIdOptions="positionIdOptions" @nurseList="nurseList"></apiV1ArchiveNurseEdit>
    <apiV1ArchiveNurseDetail ref="detailRef" :sexOptions="sys_user_sex" :areaIdOptions="area_id"
      :deptIdOptions="deptIdOptions" @getDepartmentItemsDeptId="getDepartmentItemsDeptId"
      :educationIdOptions="educationIdOptions" @getEducationItemsEducationId="getEducationItemsEducationId"
      :titleIdOptions="titleIdOptions" @getTitleItemsTitleId="getTitleItemsTitleId" :staffTypeOptions="staff_type"
      :statusOptions="staff_status" @nurseList="nurseList"></apiV1ArchiveNurseDetail>
    <ExcelReader ref="xlsRef" @onSheetReader="onExcelReader"></ExcelReader>
    <ExcelReader ref="xls2Ref" @onSheetReader="onExcel2Reader"></ExcelReader>
    <listDialog ref="selectRef" :selectOptions="selectAreaList" :isShowDialog="true"
      @onDialogConfirm="onImportAreaConfirm">
    </listDialog>
  </div>
</template>
<script lang="ts">
import { ItemOptions } from "/@/api/items";
import { toRefs, reactive, onMounted, ref, defineComponent, computed, getCurrentInstance, toRaw } from 'vue';
import { ElMessageBox, ElMessage, FormInstance } from 'element-plus';
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
  importNurse,
  updateByImport,
  listPosition,
} from "/@/api/archive/nurse";
import {
  NurseTableColumns,
  NurseInfoData,
  NurseTableDataState,
  LinkedNurseDepartment,
  LinkedNurseEducation,
  LinkedNurseTitle,
} from "/@/views/archive/nurse/list/component/model"
import apiV1ArchiveNurseAdd from "/@/views/archive/nurse/list/component/add.vue"
import apiV1ArchiveNurseEdit from "/@/views/archive/nurse/list/component/edit.vue"
import apiV1ArchiveNurseDetail from "/@/views/archive/nurse/list/component/detail.vue"

import ExcelReader from '/@/views/common/excelReader.vue';
import listDialog from "/@/views/common/listDialog.vue"
export default defineComponent({
  name: "apiV1ArchiveNurseList",
  components: {
    apiV1ArchiveNurseAdd,
    apiV1ArchiveNurseEdit,
    apiV1ArchiveNurseDetail,
    ExcelReader,
    listDialog
  },
  setup() {
    const { proxy } = <any>getCurrentInstance()
    const loading = ref(false)
    const queryRef = ref()
    const addRef = ref();
    const editRef = ref();
    const detailRef = ref();
    const xlsRef = ref()
    const xls2Ref = ref()
    const selectRef = ref();
    const curImportAreaId = ref();
    const curImportType = ref();

    // 是否显示所有搜索选项
    const showAll = ref(false)
    // 非单个禁用
    const single = ref(true)
    // 非多个禁用
    const multiple = ref(true)
    const word = computed(() => {
      if (showAll.value === false) {
        //对文字进行处理
        return "展开搜索";
      } else {
        return "收起搜索";
      }
    })
    // 字典选项数据        
    const {
      sys_user_sex,
      area_id,
      staff_type,
      staff_status,
      sys_yes_no,
    } = proxy.useDict(
      'sys_user_sex',
      'area_id',
      'staff_type',
      'staff_status',
      'sys_yes_no'
    )
    const selectAreaList = ref<Array<ItemOptions>>([])

    // deptIdOptions关联表数据
    const deptIdOptions = ref<Array<ItemOptions>>([])

    // positionIdOptions关联表数据
    const positionIdOptions = ref<Array<ItemOptions>>([])

    // educationIdOptions关联表数据
    const educationIdOptions = ref<Array<ItemOptions>>([])
    // titleIdOptions关联表数据
    const titleIdOptions = ref<Array<ItemOptions>>([])
    const state = reactive<NurseTableDataState>({
      nurseIds: [],
      tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
          pageNum: 1,
          pageSize: 10,
          areaId: undefined,
          nurseName: undefined,
          deptId: undefined,
          status: undefined,
          regStatus: undefined,
          startDate: [],
          endDate: [],
          educationId: undefined,
          titleId: undefined,
          staffType: undefined,
          nurseCode: undefined,
          dateRange: []
        },
      },
    });
    // 页面加载时
    onMounted(() => {
      initTableData();
    });
    // 初始化表格数据
    const initTableData = () => {
      nurseList()
    };

    // 导入
    const handleImport = () => {
      // console.log('handleImport.begin');
      curImportType.value = 1
      if (selectAreaList.value.length <= 0) {
        selectAreaList.value = area_id.value.map((p: any) => ({ key: p.value, value: p.label }))
      }
      // console.log('selectAreaList:', selectAreaList);
      selectRef.value.openDialog("院区选择")
      /*
      导入步骤：
      1，选择院区
      2，选择excel文件
      3，选择sheet
      4，读取sheet数据，调用api
      */
    };


    // 导入
    const handleRegInfoImport = () => {
      curImportType.value = 2

      if (selectAreaList.value.length <= 0) {
        selectAreaList.value = area_id.value.map((p: any) => ({ key: p.value, value: p.label }))
      }

      selectRef.value.openDialog("院区选择")
      /*
      导入步骤：
      1，选择院区
      2，选择excel文件
      3，选择sheet
      4，读取sheet数据，调用api
      */

    };

    // 导入
    const onImportAreaConfirm = (data: any) => {
      curImportAreaId.value = data;

      if (curImportType.value == 1) {
        xlsRef.value.openDialog();
      } else {
        xls2Ref.value.openDialog();
      }
      // xlsRef.value.openDialog();
    }

    // 导入
    const onExcelReader = (data: any) => {
      // console.log('import.jsonData:', data);
      const importData = reactive({ "data": data, "areaId": curImportAreaId.value })

      // console.log('importData:', data);

      importNurse(importData).then((res: any) => {
        if (res.code != "0") {
          // ElMessage.success('导入失败:' + res.message);
          ElMessageBox.alert(
            res.message, '导入失败',
            {
              dangerouslyUseHTMLString: true,
              confirmButtonText: '确认',
            }
          )

          return;
        }

        if (res.data.errorCount < 1) {
          let errorMsg = "成功记录数:" + res.data.successCount
          if (res.data.existsCount > 0) {
            errorMsg = errorMsg + ";已存在记录数:" + res.data.existsCount
          }
          ElMessageBox.alert(
            errorMsg, '导入成功',
            {
              dangerouslyUseHTMLString: true,
              confirmButtonText: '确认',
            }
          )
          return;
        }
        let errorList = res.data.errorData ?? [];
        // console.log("errorList:",errorList)

        let errorMsg: string = "<table><tr><th>行数</th><th>错误</th></tr>"

        for (var i = 0; i < errorList.length; i++) {
          if (i >= 10) {
            errorMsg = errorMsg + "<tr><td>" + (errorList.length - 10).toString() + "行错误......</td><td></td></tr>"
            break;
          }
          errorMsg = errorMsg + "<tr><td>" + errorList[i].index + "</td><td>" + errorList[i].errorMsg + "</td></tr>"
        }

        errorMsg = errorMsg + "</table>"
        ElMessageBox.alert(
          errorMsg, '导入失败',
          {
            dangerouslyUseHTMLString: true,
            confirmButtonText: '确认',
          }
        )

      })
    };

    // 注册信秘导入
    const onExcel2Reader = (data: any) => {
      // console.log('import.jsonData:', data);
      const importData = reactive({ "data": data, "areaId": curImportAreaId.value })

      // console.log('importData:', data);

      updateByImport(importData).then((res: any) => {
        if (res.code != "0") {
          // ElMessage.success('导入失败:' + res.message);
          ElMessageBox.alert(
            res.message, '导入失败',
            {
              dangerouslyUseHTMLString: true,
              confirmButtonText: '确认',
            }
          )

          return;
        }

        if (res.data.errorCount < 1) {
          let errorMsg = "成功记录数:" + res.data.successCount
          if (res.data.existsCount > 0) {
            errorMsg = errorMsg + ";已存在记录数:" + res.data.existsCount
          }
          ElMessageBox.alert(
            errorMsg, '导入成功',
            {
              dangerouslyUseHTMLString: true,
              confirmButtonText: '确认',
            }
          )
          return;
        }
        let errorList = res.data.errorData ?? [];

        let errorMsg: string = "<table><tr><th>行数</th><th>错误</th></tr>"

        for (var i = 0; i < errorList.length; i++) {
          if (i >= 10) {
            errorMsg = errorMsg + "<tr><td>" + (errorList.length - 10).toString() + "行错误......</td><td></td></tr>"
            break;
          }
          errorMsg = errorMsg + "<tr><td>" + errorList[i].index + "</td><td>" + errorList[i].errorMsg + "</td></tr>"
        }

        errorMsg = errorMsg + "</table>"
        ElMessageBox.alert(
          errorMsg, '导入失败',
          {
            dangerouslyUseHTMLString: true,
            confirmButtonText: '确认',
          }
        )

      })
    };


    /** 重置按钮操作 */
    const resetQuery = (formEl: FormInstance | undefined) => {
      if (!formEl) return
      formEl.resetFields()
      nurseList()
    };
    // 获取列表数据
    const nurseList = () => {
      loading.value = true
      listNurse(state.tableData.param).then((res: any) => {
        let list = res.data.list ?? [];
        state.tableData.data = list;
        state.tableData.total = res.data.total;
        loading.value = false
      })
    };
    const toggleSearch = () => {
      showAll.value = !showAll.value;
    }
    // 性别字典翻译
    const sexFormat = (row: NurseTableColumns) => {
      return proxy.selectDictLabel(sys_user_sex.value, row.sex);
    }
    // 院区字典翻译
    const areaIdFormat = (row: NurseTableColumns) => {
      return proxy.selectDictLabel(area_id.value, row.areaId);
    }
    //关联t_department表选项
    const getDepartmentItemsDeptId = () => {
      if (deptIdOptions.value && deptIdOptions.value.length > 0) {
        return
      }
      proxy.getItems(listDepartment, { pageSize: 10000 }).then((res: any) => {
        deptIdOptions.value = proxy.setItems(res, 'deptId', 'deptName')
      })
    }

    //关联t_position表选项
    const getPositionItemsPositionId = () => {
      if (positionIdOptions.value && deptIdOptions.value.length > 0) {
        return
      }
      proxy.getItems(listPosition, { pageSize: 10000 }).then((res: any) => {
        positionIdOptions.value = proxy.setItems(res, 'positionId', 'positionName')
      })
    }


    //关联t_education表选项
    const getEducationItemsEducationId = () => {
      if (educationIdOptions.value && educationIdOptions.value.length > 0) {
        return
      }
      proxy.getItems(listEducation, { pageSize: 10000 }).then((res: any) => {
        educationIdOptions.value = proxy.setItems(res, 'educationId', 'educationName')
      })
    }
    //关联t_title表选项
    const getTitleItemsTitleId = () => {
      if (titleIdOptions.value && titleIdOptions.value.length > 0) {
        return
      }
      proxy.getItems(listTitle, { pageSize: 10000 }).then((res: any) => {
        titleIdOptions.value = proxy.setItems(res, 'titleId', 'titleName')
      })
    }
    // 人员类别字典翻译
    const staffTypeFormat = (row: NurseTableColumns) => {
      return proxy.selectDictLabel(staff_type.value, row.staffType);
    }
    // 状态字典翻译
    const statusFormat = (row: NurseTableColumns) => {
      return proxy.selectDictLabel(staff_status.value, row.status);
    }
    // 多选框选中数据
    const handleSelectionChange = (selection: Array<NurseInfoData>) => {
      state.nurseIds = selection.map(item => item.nurseId)
      single.value = selection.length != 1
      multiple.value = !selection.length
    }
    const handleAdd = () => {
      addRef.value.openDialog()
    }
    const handleUpdate = (row: NurseTableColumns) => {
      if (!row) {
        row = state.tableData.data.find((item: NurseTableColumns) => {
          return item.nurseId === state.nurseIds[0]
        }) as NurseTableColumns
      }
      editRef.value.openDialog(toRaw(row));
    };
    const handleDelete = (row: NurseTableColumns) => {
      let msg = '你确定要删除所选数据？';
      let nurseId: number[] = [];
      if (row) {
        msg = `此操作将永久删除数据，是否继续?`
        nurseId = [row.nurseId]
      } else {
        nurseId = state.nurseIds
      }
      if (nurseId.length === 0) {
        ElMessage.error('请选择要删除的数据。');
        return
      }
      ElMessageBox.confirm(msg, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(() => {
          delNurse(nurseId).then(() => {
            ElMessage.success('删除成功');
            nurseList();
          })
        })
        .catch(() => { });
    }
    const handleView = (row: NurseTableColumns) => {
      detailRef.value.openDialog(toRaw(row));
    }
    return {
      proxy,
      addRef,
      editRef,
      detailRef,
      showAll,
      loading,
      single,
      multiple,
      word,
      queryRef,
      resetQuery,
      nurseList,
      toggleSearch,
      sexFormat,
      sys_user_sex,
      areaIdFormat,
      area_id,
      selectAreaList,
      //关联表数据选项
      deptIdOptions,
      //关联t_department表选项获取数据方法
      getDepartmentItemsDeptId,
      //关联表数据选项
      educationIdOptions,
      //关联t_education表选项获取数据方法
      getEducationItemsEducationId,
      //关联表数据选项
      titleIdOptions,
      //关联t_title表选项获取数据方法
      getTitleItemsTitleId,
      //关联t_Position表选项获取数据方法
      getPositionItemsPositionId,
      positionIdOptions,
      staffTypeFormat,
      staff_type,
      statusFormat,
      staff_status,
      sys_yes_no,
      handleSelectionChange,
      handleAdd,
      handleUpdate,
      handleDelete,
      handleImport,
      handleRegInfoImport,
      xlsRef,
      xls2Ref,
      onExcelReader,
      onExcel2Reader,
      selectRef,
      onImportAreaConfirm,
      curImportType,
      curImportAreaId,
      ...toRefs(state),
    }
  }
})
</script>
<style lang="scss" scoped>
.colBlock {
  display: block;
}

.colNone {
  display: none;
}
.table-expand {
  font-size: 0;
  display: flex;
  flex-wrap: wrap;
}

.table-expand label {
  width: 90px;
  color: #99a9bf;
}

.table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 15%;
}
</style>