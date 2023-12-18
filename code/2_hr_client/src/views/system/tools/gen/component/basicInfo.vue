<template>
	<div class="form-adapt-container">
		<el-card shadow="never" header="基本信息配置">
			<el-form ref="basicInfoFormRef" :model="info" :rules="rules" label-width="150px" size="default" class="mt35 mb35">
				<el-row :gutter="35">
					<el-col class="mb20" :span="12">
						<el-form-item label="表名称" prop="tableName">
							<el-input placeholder="请输入仓库名称" v-model="info.tableName" />
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="表描述" prop="tableComment">
							<el-input placeholder="请输入" v-model="info.tableComment" />
						</el-form-item>
					</el-col>

					<el-col class="mb20" :span="12">
						<el-form-item label="实体类名称" prop="className">
							<el-input placeholder="请输入" v-model="info.className" />
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="作者" prop="functionAuthor">
							<el-input placeholder="请输入" v-model="info.functionAuthor" />
						</el-form-item>
					</el-col>

					<el-col class="mb20" :span="12">
						<el-form-item label="排序字段" prop="sortColumn">
							<el-select v-model="info.sortColumn" placeholder="请选择">
								<el-option v-for="(column, index) in info.columns" :key="index" :label="column.columnComment" :value="column.columnName"> </el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="排序方式" prop="sortType">
							<el-select placeholder="请选择" v-model="info.sortType">
								<el-option label="顺序" value="asc" />
								<el-option label="倒序" value="desc" />
							</el-select>
						</el-form-item>
					</el-col>

					<el-col class="mb20" :span="24">
						<el-form-item label="备注" prop="remark">
							<el-input type="textarea" :rows="3" v-model="info.remark"></el-input>
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="是否覆盖原有文件" prop="overwrite">
							<el-checkbox v-model="info.overwrite" />
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="12">
						<el-form-item label="是否有显示详情功能" prop="showDetail">
							<el-checkbox v-model="info.showDetail" />
						</el-form-item>
					</el-col>
				</el-row>
			</el-form>
			<el-row>
				<el-col :span="10"></el-col>
				<el-col :span="4" class="base-footer">
					<span>
						<el-button size="default" @click="close">关 闭</el-button>
						<el-button type="primary" @click="nextTip" size="default">下一步</el-button>
					</span>
				</el-col>
				<el-col :span="10"></el-col>
			</el-row>
		</el-card>
	</div>
</template>

<script lang="ts">
import { defineComponent, inject, reactive, ref, unref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';
import type { TableDataInfo } from '/@/views/system/tools/gen/component/model';
import { ElMessage } from 'element-plus';
export default defineComponent({
	name: 'toolsBasicInfoForm',
	emits: ['goNext','close'],
	setup(props, { emit }) {
		const basicInfoFormRef = ref<FormInstance>();
		const info = inject<TableDataInfo>('tableData') as TableDataInfo;
		const nextTip = (): boolean => {
			//表单验证
			let tag = false;
			const formWrap = unref(basicInfoFormRef) as any;
			if (!formWrap) return false;
			formWrap.validate((valid: boolean) => {
				if (valid) {
					emit('goNext');
					tag = true;
				} else {
					ElMessage.error('请将表单填写完整');
				}
			});
			return tag;
		};
    const close = ()=>{
      console.log('xxxxxxxxxxxxxxxxxxxxxxx')
      emit('close')
    }
		const rules = reactive<FormRules>({
			tableName: [{ required: true, message: '请输入表名称', trigger: 'blur' }],
			tableComment: [{ required: true, message: '请输入表描述', trigger: 'blur' }],
			className: [{ required: true, message: '请输入实体类名称', trigger: 'blur' }],
			functionAuthor: [{ required: true, message: '请输入作者', trigger: 'blur' }],
			sortColumn: [{ required: true, message: '请选择排序字段', trigger: 'blur' }],
			sortType: [{ required: true, message: '请选择排序类型', trigger: 'blur' }],
		});
		return {
			basicInfoFormRef,
			info,
			rules,
			nextTip,
      close
		};
	},
});
</script>

<style scoped>
.base-footer {
	text-align: center;
}
</style>
