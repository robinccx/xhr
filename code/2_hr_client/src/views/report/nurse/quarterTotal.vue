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
              <el-form-item>
                <el-button type="primary" @click="doQuery"><el-icon><ele-Search /></el-icon>搜索</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button type="primary" @click="handleExport">
              <el-icon><ele-Plus /></el-icon>导出</el-button>
          </el-col>
        </el-row>
      </div>
      <el-table v-loading="loading" :data="tableData.data">

        <el-table-column label="分类名称" align="center" prop="itemClass" min-width="150px" />
        <el-table-column label="变量名称" align="center" prop="itemName" min-width="150px" />
        <el-table-column label="第一季度" align="center">
          <el-table-column label="杨浦" align="center" prop="quarter1Qty1" min-width="65px" />
          <el-table-column label="安亭" align="center" prop="quarter1Qty2" min-width="65px" />
          <el-table-column label="全院" align="center" prop="quarter1QtyTotal" min-width="65px" />
        </el-table-column>

        <el-table-column label="第二季度" align="center">
          <el-table-column label="杨浦" align="center" prop="quarter2Qty1" min-width="65px" />
          <el-table-column label="安亭" align="center" prop="quarter2Qty2" min-width="65px" />
          <el-table-column label="全院" align="center" prop="quarter2QtyTotal" min-width="65px" />
        </el-table-column>

        <el-table-column label="第三季度" align="center">
          <el-table-column label="杨浦" align="center" prop="quarter3Qty1" min-width="65px" />
          <el-table-column label="安亭" align="center" prop="quarter3Qty2" min-width="65px" />
          <el-table-column label="全院" align="center" prop="quarter3QtyTotal" min-width="65px" />
        </el-table-column>

        <el-table-column label="第四季度" align="center">
          <el-table-column label="杨浦" align="center" prop="quarter4Qty1" min-width="65px" />
          <el-table-column label="安亭" align="center" prop="quarter4Qty2" min-width="65px" />
          <el-table-column label="全院" align="center" prop="quarter4QtyTotal" min-width="65px" />
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>
<script lang="ts">

import { toRefs, reactive, onMounted, ref, defineComponent, computed, getCurrentInstance } from 'vue';
import { ElMessageBox, ElMessage, FormInstance } from 'element-plus';
import {
  ListQuarterTotal,
} from "/@/api/report/nurse/monthlyReport";
import * as XLSX from 'xlsx';

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


    const state = reactive({
      titleIds: [],
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

    });

    // 导出
    const handleExport = () => {
      if (!state.tableData.data) {
        return
      }

      if (state.tableData.data.length < 1) {
        return
      }
      let msg = '导出数据,是否继续？';
      ElMessageBox.confirm(msg, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(() => {
          console.log("export.begin")
          const exportData: string[][] = [];

          exportData.push(["全院填报数据采集"]);
          exportData.push(["分类名称", "变量名称", "第一季度", "", "", "第二季度", "", "", "第三季度", "", "", "第四季度", "", ""]);

          exportData.push(["分类名称", "变量名称", "杨浦", "安亭", "全院",
            "杨浦", "安亭", "全院",
            "杨浦", "安亭", "全院",
            "杨浦", "安亭", "全院"
          ]);
          state.tableData.data.forEach((d: any) => {
            exportData.push([
              d.itemClass, d.itemName,
              d.quarter1Qty1,
              d.quarter1Qty2,
              d.quarter1QtyTotal,

              d.quarter2Qty1,
              d.quarter2Qty2,
              d.quarter2QtyTotal,

              d.quarter3Qty1,
              d.quarter3Qty2,
              d.quarter3QtyTotal,

              d.quarter4Qty1,
              d.quarter4Qty2,
              d.quarter4QtyTotal
            ]);
          })
          console.log("export.exportData:", exportData)



          const ws = XLSX.utils.aoa_to_sheet(exportData);
          // 设置列宽
          const columnWidths = [{ wch: 34 }, { wch: 37 }];
          ws['!cols'] = columnWidths;

          // 设置合并单元格，从第 1 行、第 1 列到第 1 行、第 14 列
          const mergeRange = { s: { r: 0, c: 0 }, e: { r: 0, c: 13 } };
          if (!ws['!merges']) ws['!merges'] = [];
          ws['!merges'].push(mergeRange);

          // 设置合并单元格，从第 2 行、第 1 列到第 3 行、第 1 列
          const mergeRange2 = { s: { r: 1, c: 0 }, e: { r: 2, c: 0 } };
          ws['!merges'].push(mergeRange2);

          // 设置合并单元格，从第 2 行、第 2 列到第 3 行、第 2 列
          const mergeRange3 = { s: { r: 1, c: 1 }, e: { r: 2, c: 1 } };
          ws['!merges'].push(mergeRange3);

          // 设置合并单元格，从第 2 行、第 3 列到第 2行、第 5 列
          const mergeRange4 = { s: { r: 1, c: 2 }, e: { r: 1, c: 4 } };
          ws['!merges'].push(mergeRange4);


          // 设置合并单元格，从第 2 行、第 6 列到第 2行、第 8 列
          const mergeRange5 = { s: { r: 1, c: 5 }, e: { r: 1, c: 7 } };
          ws['!merges'].push(mergeRange5);


          // 设置合并单元格，从第 2 行、第 9 列到第 2行、第 11 列
          const mergeRange6 = { s: { r: 1, c: 8 }, e: { r: 1, c: 10 } };
          ws['!merges'].push(mergeRange6);

          // 设置合并单元格，从第 2 行、第 12 列到第 2行、第 14 列
          const mergeRange7 = { s: { r: 1, c: 11 }, e: { r: 1, c: 13 } };
          ws['!merges'].push(mergeRange7);

          // 设置合并单元格，从第 4 行、第 1 列到第 7行、第 1 列
          const mergeRange8 = { s: { r: 3, c: 0 }, e: { r: 6, c: 0 } };
          ws['!merges'].push(mergeRange8);

          // 设置合并单元格，从第 8 行、第 1 列到第 17行、第 1 列
          const mergeRange9 = { s: { r: 7, c: 0 }, e: { r: 16, c: 0 } };
          ws['!merges'].push(mergeRange9);

          
          // 设置合并单元格，从第 18 行、第 1 列到第29行、第 1 列
          const mergeRange10 = { s: { r: 17, c: 0 }, e: { r: 28, c: 0 } };
          ws['!merges'].push(mergeRange10);

          
          // 设置合并单元格，从第 30 行、第 1 列到第 41 行、第 1 列
          const mergeRange11 = { s: { r: 29, c: 0 }, e: { r: 40, c: 0 } };
          ws['!merges'].push(mergeRange11);

          // 设置合并单元格，从第 42 行、第 1 列到第 51 行、第 1 列
          const mergeRange12 = { s: { r: 41, c: 0 }, e: { r: 50, c: 0 } };
          ws['!merges'].push(mergeRange12);
          

          // 设置单元格样式
          const style: XLSX.Style = {
            font: { bold: true }, // 字体样式
            alignment: { vertical: 'center', horizontal: 'center' }, // 对齐方式
            border: {
              top: { style: 'thin', color: { rgb: '000000' } }, // 边框样式
              bottom: { style: 'thin', color: { rgb: '000000' } },
              left: { style: 'thin', color: { rgb: '000000' } },
              right: { style: 'thin', color: { rgb: '000000' } },
            },
          };
          // 应用样式到单元格 A1:N4
          const range = XLSX.utils.decode_range('A1:N4');
          for (let R = range.s.r; R <= range.e.r; ++R) {
            for (let C = range.s.c; C <= range.e.c; ++C) {
              const cell_address = { c: C, r: R };
              if (!ws[XLSX.utils.encode_cell(cell_address)]) {
                ws[XLSX.utils.encode_cell(cell_address)] = {};
              }
              ws[XLSX.utils.encode_cell(cell_address)].s = style;
            }
          }

          // 设置字体样式
          const font = { bold: true };
          const alignment = { vertical: 'center', horizontal: 'center' };
          // 应用样式到单元格 A1
          if (!ws['A1']) ws['A1'] = {};
          if (!ws['A2']) ws['A2'] = {};
          ws['A1'].s = { font }; // s 表示样式
          ws['A2'].s = { "font": font, "alignment": alignment }; // s 表示样式

          const wb = XLSX.utils.book_new();
          XLSX.utils.book_append_sheet(wb, ws, '全院数据采集表单');


          XLSX.writeFile(wb, '全院填报数据汇总.xlsx');
          console.log("export.end")
        })
        .catch(() => { });
    };



    // 获取列表数据
    const doQuery = () => {
      loading.value = true
      ListQuarterTotal(state.tableData.param).then((res: any) => {
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
      queryRef,
      doQuery,
      handleExport,
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