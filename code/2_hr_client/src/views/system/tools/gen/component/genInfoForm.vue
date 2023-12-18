<template>
	<el-card shadow="never" header="生成信息配置">
		<el-form ref="genInfoFormRef" :model="info" :rules="rules" label-width="150px">
			<el-row :gutter="35">
				<el-col class="mb20" :span="12">
					<el-form-item prop="tplCategory">
						<template #label>
							<span>生成模板</span>
						</template>
						<el-select v-model="info.tplCategory">
							<el-option label="单表（增删改查）" value="crud" />
							<el-option label="树表（增删改查）" value="tree" />
						</el-select>
					</el-form-item>
				</el-col>

				<el-col class="mb20" :span="12">
					<el-form-item prop="packageName">
						<template #label>
							<span>
								生成包路径
								<el-tooltip content="生成在哪个包下，例如 gfast/app/system" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
						</template>
						<el-input v-model="info.packageName" />
					</el-form-item>
				</el-col>

				<el-col class="mb20" :span="12">
					<el-form-item prop="moduleName">
						<template #label>
							<span>
								生成模块名
								<el-tooltip content="可理解为子系统名，例如 system" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
						</template>
						<el-input v-model="info.moduleName" />
					</el-form-item>
				</el-col>

				<el-col class="mb20" :span="12">
					<el-form-item prop="businessName">
						<template #label>
							<span>
								生成业务名
								<el-tooltip content="可理解为功能英文名，例如 user" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
						</template>
						<el-input v-model="info.businessName" />
					</el-form-item>
				</el-col>

				<el-col class="mb20" :span="12">
					<el-form-item prop="functionName">
						<template #label>
							<span>
								生成功能名
								<el-tooltip content="用作类描述，例如 用户" placement="top">
									<el-icon><ele-QuestionFilled /></el-icon>
								</el-tooltip>
							</span>
						</template>
						<el-input v-model="info.functionName" />
					</el-form-item>
				</el-col>
			</el-row>

			<div v-show="info.tplCategory == 'tree'">
				<h4 class="form-header">其他信息</h4>
				<el-divider />
				<el-row :gutter="35">
					<el-col class="mb20" :span="8">
						<el-form-item>
							<template #label>
								<span>
									树编码字段
									<el-tooltip content="树显示的编码字段名， 如：dept_id" placement="top">
										<el-icon><ele-QuestionFilled /></el-icon>
									</el-tooltip>
								</span>
							</template>
							<el-select v-model="info.treeCode" placeholder="请选择">
								<el-option
									v-for="column in info.columns"
									:key="column.htmlField"
									:label="column.columnName + '：' + column.columnComment"
									:value="column.htmlField"
								></el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="8">
						<el-form-item>
							<template #label>
								<span>
									树父编码字段
									<el-tooltip content="树显示的父编码字段名， 如：parent_Id" placement="top">
										<el-icon><ele-QuestionFilled /></el-icon>
									</el-tooltip>
								</span>
							</template>
							<el-select v-model="info.treeParentCode" placeholder="请选择">
								<el-option
									v-for="column in info.columns"
									:key="column.htmlField"
									:label="column.columnName + '：' + column.columnComment"
									:value="column.htmlField"
								></el-option>
							</el-select>
						</el-form-item>
					</el-col>
					<el-col class="mb20" :span="8">
						<el-form-item>
							<template #label>
								<span>
									树名称字段
									<el-tooltip content="树节点的显示名称字段名， 如：dept_name" placement="top">
										<el-icon><ele-QuestionFilled /></el-icon>
									</el-tooltip>
								</span>
							</template>
							<el-select v-model="info.treeName" placeholder="请选择">
								<el-option
									v-for="column in info.columns"
									:key="column.htmlField"
									:label="column.columnName + '：' + column.columnComment"
									:value="column.htmlField"
								></el-option>
							</el-select>
						</el-form-item>
					</el-col>
				</el-row>
			</div>
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
</template>

<script lang="ts">
import {defineComponent, inject, reactive, ref, unref} from 'vue';
import type { FormInstance, FormRules } from 'element-plus';
import { TableDataInfo } from '/@/views/system/tools/gen/component/model';
import {ElMessage} from "element-plus";
export default defineComponent({
	name: 'genInfoForm',
  emits:['goNext','close'],
	setup(props,{emit}) {
		const genInfoFormRef = ref<FormInstance>();
		const info = inject<TableDataInfo>('tableData') as TableDataInfo;
		const rules = reactive<FormRules>({
			tplCategory: [{ required: true, message: '请选择生成模板', trigger: 'blur' }],
			packageName: [{ required: true, message: '请输入生成包路径', trigger: 'blur' }],
			moduleName: [{ required: true, message: '请输入生成模块名', trigger: 'blur' }],
			businessName: [{ required: true, message: '请输入生成业务名', trigger: 'blur' }],
			functionName: [{ required: true, message: '请输入生成功能名', trigger: 'blur' }],
		});
    const nextTip = ()=>{
      //表单验证
      let tag = false;
      const formWrap = unref(genInfoFormRef) as any;
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
    }
    const close = ()=>{
      emit('close')
    }
		return {
      genInfoFormRef,
			info,
			rules,
      nextTip,
      close
		};
	},
});
</script>

<style scoped>
.form-header {
}
</style>
