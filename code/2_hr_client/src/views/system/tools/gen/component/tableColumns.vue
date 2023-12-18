<template>
	<el-card shadow="never" header="字段信息配置">
		<el-alert title="⚠️表字段中的updated_at、deleted_at的字段在此列表中已经隐藏" type="warning" show-icon />
		<el-table ref="tableColumnsRef" :data="info.columns" row-key="columnId" :max-height="tableHeight" width="100%">
			<el-table-column label="序号" type="index" width="50" fixed />
			<el-table-column label="字段列名" prop="columnName" width="150" fixed :show-overflow-tooltip="true" />
			<el-table-column label="字段描述" width="150" fixed>
				<template #default="scope">
					<textarea class="el-textarea__inner" v-model.lazy="scope.row.columnComment"></textarea>
				</template>
			</el-table-column>
			<el-table-column class-name="allowDrag" label="物理类型" prop="columnType" width="120" :show-overflow-tooltip="true" />
			<el-table-column label="go类型" width="120">
				<template #default="scope">
					<el-select v-model="scope.row.goType">
						<el-option label="int" value="int" />
						<el-option label="uint" value="uint" />
						<el-option label="int64" value="int64" />
						<el-option label="uint64" value="uint64" />
						<el-option label="float64" value="float64" />
						<el-option label="string" value="string" />
						<el-option label="Time" value="Time" />
						<el-option label="byte" value="byte" />
						<el-option label="bool" value="bool" />
					</el-select>
				</template>
			</el-table-column>
      <el-table-column label="TS类型" width="120">
        <template #default="scope">
          <el-select v-model="scope.row.tsType">
            <el-option label="string" value="string" />
            <el-option label="number" value="number" />
            <el-option label="boolean" value="boolean" />
            <el-option label="string[]" value="string[]" />
            <el-option label="number[]" value="number[]" />
            <el-option label="any" value="any" />
            <el-option label="tuple" value="tuple" />
            <el-option label="enum" value="enum" />
          </el-select>
        </template>
      </el-table-column>
			<el-table-column label="go属性" width="150">
				<template #default="scope">
					<textarea class="el-textarea__inner" v-model.lazy="scope.row.goField"></textarea>
				</template>
			</el-table-column>

			<el-table-column label="json属性" width="150">
				<template #default="scope">
          <textarea class="el-textarea__inner" v-model.lazy="scope.row.htmlField" ></textarea>
				</template>
			</el-table-column>

			<el-table-column label="插入" width="50">
				<template #default="scope">
					<el-checkbox
						v-model="scope.row.isInsert"
						:disabled="
							scope.row.isPk == '1' ||
							scope.row.columnName == 'created_at' ||
							scope.row.columnName == 'created_by' ||
							scope.row.columnName == 'updated_by'
						"
					></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="编辑" width="50">
				<template #default="scope">
					<el-checkbox
						v-model="scope.row.isEdit"
						:disabled="
							scope.row.isPk == '1' ||
							scope.row.columnName == 'created_at' ||
							scope.row.columnName == 'created_by' ||
							scope.row.columnName == 'updated_by'
						"
					></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="列表" width="50">
				<template #default="scope">
					<el-checkbox v-model="scope.row.isList"></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="列表固定" width="50">
				<template #default="scope">
					<el-checkbox v-model="scope.row.isFixed"></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="过长隐藏" width="50">
				<template #default="scope">
					<el-checkbox v-model="scope.row.isOverflowTooltip" />
				</template>
			</el-table-column>
			<el-table-column label="详情" width="50">
				<template #default="scope">
					<el-checkbox v-model="scope.row.isDetail"></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="查询" width="50">
				<template #default="scope">
					<el-checkbox v-model="scope.row.isQuery" :disabled="scope.row.htmlField == info.treeParentCode"></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="查询方式" width="120">
				<template #default="scope">
					<el-select v-model="scope.row.queryType">
						<el-option label="=" value="EQ" />
						<el-option label="!=" value="NE" />
						<el-option label=">" value="GT" />
						<el-option label=">=" value="GTE" />
						<el-option label="<" value="LT" />
						<el-option label="<=" value="LTE" />
						<el-option label="LIKE" value="LIKE" />
						<el-option label="BETWEEN" value="BETWEEN" />
					</el-select>
				</template>
			</el-table-column>
			<el-table-column label="必填" width="50">
				<template #default="scope">
					<el-checkbox v-model="scope.row.isRequired"></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="显示类型" width="140">
				<template #default="scope">
					<el-select v-model="scope.row.htmlType" :disabled="scope.row.htmlField == info.treeParentCode">
						<el-option label="文本框" value="input" />
						<el-option label="文本域" value="textarea" />
						<el-option label="下拉框" value="select" />
						<el-option label="单选框" value="radio" />
						<el-option label="复选框" value="checkbox" />
						<el-option label="日期控件" value="date" />
						<el-option label="日期时间控件" value="datetime" />
						<el-option label="富文本" value="richtext" />
						<el-option label="单图上传" value="imagefile" />
						<el-option label="多图上传" value="images" />
						<el-option label="单文件上传" value="file" />
						<el-option label="多文件上传" value="files" />
					</el-select>
				</template>
			</el-table-column>
			<el-table-column label="列表排序" width="100">
				<template #default="scope">
          <textarea class="el-textarea__inner" v-model.lazy="scope.row.sortOrderList" ></textarea>
				</template>
			</el-table-column>
			<el-table-column label="详情排序" width="100">
				<template #default="scope">
          <textarea class="el-textarea__inner" v-model.lazy="scope.row.sortOrderDetail" ></textarea>
				</template>
			</el-table-column>
			<el-table-column label="编辑排序" width="100">
				<template #default="scope">
          <textarea class="el-textarea__inner" v-model.lazy="scope.row.sortOrderEdit" ></textarea>
				</template>
			</el-table-column>
			<el-table-column label="查询排序" width="100">
				<template #default="scope">
          <textarea class="el-textarea__inner" v-model.lazy="scope.row.sortOrderQuery" ></textarea>
				</template>
			</el-table-column>
			<el-table-column label="列表宽度" width="100">
				<template #default="scope">
          <textarea class="el-textarea__inner" v-model.lazy="scope.row.minWidth" ></textarea>
				</template>
			</el-table-column>
			<el-table-column label="详情占列" width="70">
				<template #default="scope">
          <textarea class="el-textarea__inner" v-model.lazy="scope.row.colSpan" ></textarea>
				</template>
			</el-table-column>
			<el-table-column label="详情起新行" width="50">
				<template #default="scope">
					<el-checkbox true-label="1" v-model="scope.row.isRowStart"></el-checkbox>
				</template>
			</el-table-column>
			<el-table-column label="字典类型" width="160">
				<template #default="scope">
					<el-select v-model="scope.row.dictType" clearable filterable placeholder="请选择">
						<el-option v-for="dict in dictOptions" :key="dict.dictType" :label="dict.dictName" :value="dict.dictType">
							<span style="float: left">{{ dict.dictName }}</span>
							<span style="float: right; color: #8492a6; font-size: 13px">{{ dict.dictType }}</span>
						</el-option>
					</el-select>
				</template>
			</el-table-column>
			<el-table-column label="关联表" width="160">
				<template #default="scope">
					<el-select
						v-model="scope.row.linkTableName"
						clearable
						filterable
						placeholder="请选择"
						@change="handleChangeConfig(scope.row)"
					>
						<el-option v-for="table in relationTable" :key="table.tableName" :label="table.tableName" :value="table.tableName">
							<span style="float: left">{{ table.tableName }}</span>
							<span style="float: right; color: #8492a6; font-size: 13px">{{ table.tableComment }}</span>
						</el-option>
					</el-select>
				</template>
			</el-table-column>
			<el-table-column label="关联表key" width="150">
				<template #default="scope">
					<el-select v-model="scope.row.linkLabelId" clearable filterable placeholder="请选择">
						<el-option v-for="column in scope.row.columns" :key="column.columnName" :label="column.columnName" :value="column.columnName">
							<span style="float: left">{{ column.htmlField }}</span>
							<span style="float: right; color: #8492a6; font-size: 13px">{{ column.columnComment }}</span>
						</el-option>
					</el-select>
				</template>
			</el-table-column>
			<el-table-column label="关联表value" width="150">
				<template #default="scope">
					<el-select v-model="scope.row.linkLabelName" clearable filterable placeholder="请选择">
						<el-option v-for="column in scope.row.columns" :key="column.columnName" :label="column.columnName" :value="column.columnName">
							<span style="float: left">{{ column.htmlField }}</span>
							<span style="float: right; color: #8492a6; font-size: 13px">{{ column.columnComment }}</span>
						</el-option>
					</el-select>
				</template>
			</el-table-column>
		</el-table>
	</el-card>
</template>

<script lang="ts">
import {defineComponent, inject, onBeforeMount, ref} from 'vue';
import type { FormInstance } from 'element-plus';
import {TableDataInfo} from '/@/views/system/tools/gen/component/model';
import {getRelationTable} from "/@/api/system/tools/gen";
import {optionselect} from "/@/api/system/dict/type";
export default defineComponent({
	name: 'genTableColumns',
	setup() {
		const tableColumnsRef = ref<FormInstance>();
		const info = inject<TableDataInfo>('tableData') as TableDataInfo;
		// 表格的高度
		const tableHeight = ref(document.documentElement.scrollHeight - 300 + 'px');
    const dictOptions = ref([])
    const relationTable = ref<TableDataInfo[]>([])
    onBeforeMount(()=>{
      //获取字典选项
      optionselect().then((res:any)=>{
        dictOptions.value = res.data.dictType??[]
      })
      //获取关联表数据
      getRelationTable().then((res:any)=>{
        relationTable.value = <TableDataInfo[]>(res.data?.data)??[]
      })
    })
    const handleChangeConfig = (row:TableDataInfo)=>{
      relationTable.value.filter(function(item) {
        if (item.tableName === row.linkTableName) {
          row.columns = item.columns
        }
      })
    }
		return {
			tableColumnsRef,
      info,
			tableHeight,
      dictOptions,
      relationTable,
      handleChangeConfig,
		};
	},
});
</script>

<style scoped>
.el-textarea__inner{
 height: 36px;
}
</style>
