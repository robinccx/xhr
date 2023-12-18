<template>
  <div class="archive-nurseChangeLog-container">
    <el-card shadow="hover">
      <div class="archive-nurseChangeLog-search mb15">
        <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
          <el-row>
            <el-col :span="4" class="colBlock">
              <el-form-item label="护士名称" prop="nurseName">
                <el-input v-model="tableData.param.nurseName" placeholder="请输入" clearable
                  @keyup.enter.native="nurseChangeLogList" />
              </el-form-item>
            </el-col>
            <el-col :span="4" class="colBlock">
              <el-form-item label="变更类型" prop="changeType">
                <el-select v-model="tableData.param.changeType" placeholder="请选择变更类型" clearable>
                  <el-option v-for="dict in nurse_change_type" :key="dict.value" :label="dict.label"
                    :value="dict.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="4" :class="!showAll ? 'colBlock' : 'colNone'">
              <el-form-item>
                <el-button type="primary" @click="nurseChangeLogList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                <el-button @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                <el-button type="primary" link @click="toggleSearch">
                  {{ word }}
                  <el-icon v-show="showAll"><ele-ArrowUp /></el-icon>
                  <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                </el-button>
              </el-form-item>
            </el-col>
            <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="业务时间" prop="transDate">
                <el-date-picker clearable style="width: 200px" v-model="tableData.param.transDate"
                  format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" type="datetime"
                  placeholder="选择变更时间"></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item label="创建时间" prop="createdAt">
                <el-date-picker clearable style="width: 200px" v-model="tableData.param.createdAt"
                  format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" type="datetime"
                  placeholder="选择"></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
              <el-form-item>
                <el-button type="primary" @click="nurseChangeLogList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
            <el-button type="success" :disabled="single" @click="handleUpdate(null)"
              v-auth="'api/v1/archive/nurseChangeLog/edit'"><el-icon><ele-Edit /></el-icon>修改</el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button type="danger" :disabled="multiple" @click="handleDelete(null)"
              v-auth="'api/v1/archive/nurseChangeLog/delete'"><el-icon><ele-Delete /></el-icon>删除</el-button>
          </el-col>
        </el-row>
      </div>
      <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="护士名称" align="center" prop="linkedNurseId.nurseName" min-width="100px" />
        <el-table-column label="变更类型" align="center" prop="changeType" :formatter="changeTypeFormat" min-width="100px" />
        <el-table-column label="原内容" align="center" prop="fromEntityName" min-width="120px" />
        <el-table-column label="新内容" align="center" prop="toEntityName" min-width="120px" />
        <el-table-column label="业务时间" align="center" prop="transDate" min-width="100px">
          <template #default="scope">
            <span>{{ proxy.parseTime(scope.row.transDate, '{y}-{m}-{d}') }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作时间" align="center" prop="createdBy" min-width="100px" />
        <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
          <template #default="scope">
            <el-button type="primary" link @click="handleDelete(scope.row)"
              v-auth="'api/v1/archive/nurseChangeLog/delete'"><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination v-show="tableData.total > 0" :total="tableData.total" v-model:page="tableData.param.pageNum"
        v-model:limit="tableData.param.pageSize" @pagination="nurseChangeLogList" />
    </el-card>
    <apiV1ArchiveNurseChangeLogEdit ref="editRef" :nurseIdOptions="nurseIdOptions"
      @getNurseItemsNurseId="getNurseItemsNurseId" :changeTypeOptions="nurse_change_type" :fromAreaIdOptions="area_id"
      :fromDeptIdOptions="fromDeptIdOptions" @getDepartmentItemsFromDeptId="getDepartmentItemsFromDeptId"
      :toAreaIdOptions="area_id" :toDeptIdOptions="toDeptIdOptions"
      @getDepartmentItemsToDeptId="getDepartmentItemsToDeptId" @nurseChangeLogList="nurseChangeLogList">
    </apiV1ArchiveNurseChangeLogEdit>
  </div>
</template>
<script lang="ts">
import { ItemOptions } from "/@/api/items";
import { toRefs, reactive, onMounted, ref, defineComponent, computed, getCurrentInstance, toRaw } from 'vue';
import { ElMessageBox, ElMessage, FormInstance } from 'element-plus';
import {
  listNurseChangeLog,
  getNurseChangeLog,
  delNurseChangeLog,
  addNurseChangeLog,
  updateNurseChangeLog,  
  listDepartment,
  getUserList,
} from "/@/api/archive/nurseChangeLog";
import {
  NurseChangeLogTableColumns,
  NurseChangeLogInfoData,
  NurseChangeLogTableDataState,
} from "/@/views/archive/nurseChangeLog/list/component/model"
import apiV1ArchiveNurseChangeLogAdd from "/@/views/archive/nurseChangeLog/list/component/add.vue"
import apiV1ArchiveNurseChangeLogEdit from "/@/views/archive/nurseChangeLog/list/component/edit.vue"
import apiV1ArchiveNurseChangeLogDetail from "/@/views/archive/nurseChangeLog/list/component/detail.vue"
export default defineComponent({
  name: "apiV1ArchiveNurseChangeLogList",
  components: {
    apiV1ArchiveNurseChangeLogAdd,
    apiV1ArchiveNurseChangeLogEdit,
    apiV1ArchiveNurseChangeLogDetail
  },
  setup() {
    const { proxy } = <any>getCurrentInstance()
    const loading = ref(false)
    const queryRef = ref()
    const addRef = ref();
    const editRef = ref();
    const detailRef = ref();
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
      nurse_change_type,      
    } = proxy.useDict(
      'nurse_change_type',
    )
    
    const state = reactive<NurseChangeLogTableDataState>({
      ids: [],
      tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
          pageNum: 1,
          pageSize: 10,
          nurseId: undefined,
          changeType: undefined,
          fromAreaId: undefined,
          fromDeptId: undefined,
          toAreaId: undefined,
          toDeptId: undefined,
          transDate: undefined,
          status: undefined,
          createdAt: undefined,
          createdBy: undefined,
          deletedBy: undefined,
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
      nurseChangeLogList()
    };
    /** 重置按钮操作 */
    const resetQuery = (formEl: FormInstance | undefined) => {
      if (!formEl) return
      formEl.resetFields()
      nurseChangeLogList()
    };
    // 获取列表数据
    const nurseChangeLogList = () => {
      loading.value = true
      listNurseChangeLog(state.tableData.param).then((res: any) => {
        let list = res.data.list ?? [];
        state.tableData.data = list;
        state.tableData.total = res.data.total;
        loading.value = false
      })
    };
    const toggleSearch = () => {
      showAll.value = !showAll.value;
    }
    
    // 变更类型字典翻译
    const changeTypeFormat = (row: NurseChangeLogTableColumns) => {
      return proxy.selectDictLabel(nurse_change_type.value, row.changeType);
    }
    
    // 多选框选中数据
    const handleSelectionChange = (selection: Array<NurseChangeLogInfoData>) => {
      state.ids = selection.map(item => item.id)
      single.value = selection.length != 1
      multiple.value = !selection.length
    }
    const handleAdd = () => {
      addRef.value.openDialog()
    }
    const handleUpdate = (row: NurseChangeLogTableColumns) => {
      if (!row) {
        row = state.tableData.data.find((item: NurseChangeLogTableColumns) => {
          return item.id === state.ids[0]
        }) as NurseChangeLogTableColumns
      }
      editRef.value.openDialog(toRaw(row));
    };
    const handleDelete = (row: NurseChangeLogTableColumns) => {
      let msg = '你确定要删除所选数据？';
      let id: number[] = [];
      if (row) {
        msg = `此操作将永久删除数据，是否继续?`
        id = [row.id]
      } else {
        id = state.ids
      }
      if (id.length === 0) {
        ElMessage.error('请选择要删除的数据。');
        return
      }
      ElMessageBox.confirm(msg, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(() => {
          delNurseChangeLog(id).then(() => {
            ElMessage.success('删除成功');
            nurseChangeLogList();
          })
        })
        .catch(() => { });
    }
    const handleView = (row: NurseChangeLogTableColumns) => {
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
      nurseChangeLogList,
      toggleSearch,
      changeTypeFormat,
      nurse_change_type,
      handleSelectionChange,
      handleAdd,
      handleUpdate,
      handleDelete,
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
</style>