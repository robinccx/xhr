<template>
  <div class="archive-title-container">
    <el-card shadow="hover">
      <div class="archive-title-search mb15">
        <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
          <el-row>
            <el-col :span="4" class="colBlock">
              <el-form-item label="年份" prop="year">
                <el-input v-model="tableData.param.year" placeholder="请输入年份" clearable @keyup.enter.native="doQuery" />
              </el-form-item>
            </el-col>
            <el-col :span="4" class="colBlock">
              <el-form-item label="院区" prop="areaId">
                <el-select v-model="tableData.param.areaId" placeholder="请选择院区" clearable>
                  <el-option v-for="dict in area_id" :key="dict.value" :label="dict.label" :value="dict.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="4" class="colBlock">
              <el-form-item label="注册与变更" prop="regStatus">
                <el-select v-model="tableData.param.regStatus" placeholder="请选择状态" clearable>
                  <el-option v-for="dict in sys_yes_no" :key="dict.value" :label="dict.label" :value="dict.value" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="4" class="colBlock">
              <el-form-item>
                <el-button type="primary" @click="doQuery"><el-icon><ele-Search /></el-icon>搜索</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>
      <el-table v-loading="loading" :data="tableData.data">
        <el-table-column label="院区" align="center" prop="area_id" :formatter="areaIdFormat" min-width="100px" />
        <el-table-column label="月份" align="center" prop="month_str" min-width="100px" />        
        <el-table-column v-for="column in titleData" :key="column.columnKey" :prop="column.columnKey"
          :label="column.columnName" />
        <el-table-column label="合计" align="center" prop="total_qty" min-width="100px" />
      </el-table>
    </el-card>
  </div>
</template>
<script lang="ts">

import { toRefs, reactive, onMounted, ref, defineComponent, computed, getCurrentInstance } from 'vue';

import {
  ListMonthlyTitle,
} from "/@/api/report/nurse/monthlyReport";

import {
  listTitle,
} from "/@/api/archive/title";
import {
  MonthlyTableColumns,
} from "/@/views/report/nurse/model";


export default defineComponent({
  name: "apiV1ReportMonthlyTotal",
  setup() {
    const { proxy } = <any>getCurrentInstance()
    const loading = ref(false)
    const queryRef = ref()
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
      area_id,
      sys_yes_no,
    } = proxy.useDict(
      'area_id',
      'sys_yes_no',
    )
    // 院区字典翻译
    const areaIdFormat = (row: any) => {
      return proxy.selectDictLabel(area_id.value, row.area_id+"");
    }


    const state = reactive({
      titleData: Array<MonthlyTableColumns>,
      tableData: {
        data: [],
        total: 0,
        loading: false,
        param: {
          pageNum: 1,
          pageSize: 10,
          year: undefined,
          status: undefined,
          dateRange: []
        },
      },
    });

    // 页面加载时
    onMounted(() => {
      initTableData();
    });

    // 初始化数据
    const initTableData = () => {
      listTitle(state.tableData.param).then((res: any) => {
        state.titleData = res.data.list.map((item) => {
          const newItem: MonthlyTableColumns = {
            columnKey: "qty_" + item.titleId,
            columnName: item.titleName,
          }
          return newItem
        })

      })
    };


    // 获取列表数据
    const doQuery = () => {
      loading.value = true
      ListMonthlyTitle(state.tableData.param).then((res: any) => {
        let list = res.data.list ?? [];
        state.tableData.data = list;
        state.tableData.total = res.data.total;
        loading.value = false
      })
    };

    return {
      proxy,
      showAll,
      loading,
      single,
      multiple,
      word,
      area_id,
      sys_yes_no,
      queryRef,
      doQuery,
      areaIdFormat,
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