<template>
	<div class="app-container">
		<el-card shadow="hover">
			<div class="mb15">
				<el-row>
					<el-col :span="4"></el-col>
					<el-col :span="16">
						<el-steps :active="active" align-center simple>
							<el-step title="1.基本信息" @click="changeSteps(1)"
								><template #icon
									><el-icon><ele-Comment /></el-icon></template
							></el-step>
							<el-step title="2.生成信息" @click="changeSteps(2)"
								><template #icon
									><el-icon><ele-Document /></el-icon></template
							></el-step>
							<el-step title="3.字段信息" @click="changeSteps(3)"
								><template #icon
									><el-icon><ele-Expand /></el-icon></template
							></el-step>
						</el-steps>
					</el-col>
					<el-col :span="4"></el-col>
				</el-row>
			</div>
			<tools-basic-info-form ref="toolsBasicInfoForm" v-show="active === 1" @goNext="goNext" @close="close"></tools-basic-info-form>
			<gen-info-form ref="genInfoForm" v-show="active === 2" @goNext="goNext" @close="close"></gen-info-form>
      <gen-table-columns ref="tableColumns" v-show="active === 3"></gen-table-columns>
			<div class="tools-footer" v-show="active === 3">
				<el-button size="default" @click="close">关 闭</el-button>
				<el-button type="primary" @click="onSubmit" size="default">保 存</el-button>
			</div>
		</el-card>
	</div>
</template>

<script lang="ts">
import {defineComponent, onBeforeMount, ref, provide, getCurrentInstance} from 'vue';
import { useRoute } from 'vue-router';
import toolsBasicInfoForm from '/@/views/system/tools/gen/component/basicInfo.vue';
import { TableColumns, TableDataInfo } from '/@/views/system/tools/gen/component/model';
import { ElMessage } from 'element-plus/es';
import {getGenTable, saveEdit} from '/@/api/system/tools/gen';
import genInfoForm from '/@/views/system/tools/gen/component/genInfoForm.vue';
import genTableColumns from '/@/views/system/tools/gen/component/tableColumns.vue';
export default defineComponent({
	name: 'apiV1SystemToolsGenColumnList',
	components: { toolsBasicInfoForm, genInfoForm,genTableColumns },
	setup() {
    const { proxy } = getCurrentInstance() as any;
		const route = useRoute();
		const toolsBasicInfoForm = ref();
		const genInfoForm = ref();
    const tableColumns = ref()
		const active = ref(1);

		const changeSteps = (args: number) => {
			if (args > active.value) {
				switch (args) {
					case 2:
						//验证基本信息表单
						if (toolsBasicInfoForm.value.nextTip()) {
							active.value = args;
						}
						break;
					case 3:
						//验证生成信息表单
            if (active.value==1){
              //验证基本信息表单
              if (toolsBasicInfoForm.value.nextTip()) {
                active.value = args;
              }
            }else{
              //验证生成信息表单
              if (genInfoForm.value.nextTip()) {
                active.value = args;
              }
            }
						break;
				}
			} else {
				active.value = args;
			}
		};
		// 表详细信息
		const info = ref(<TableDataInfo>{});
		provide('tableData', info);
		onBeforeMount(() => {
			const tableId = route.query?.tableId;
			if (tableId) {
				// 获取表详细信息
				getGenTable(tableId).then((res) => {
					let columnsTmp:TableColumns[]= [];
					columnsTmp = res.data.list.filter((item: TableColumns) => {
						return !['updated_at', 'deleted_at'].includes(item.columnName);
					});
					const data: TableDataInfo = res.data.info as TableDataInfo;
					data.columns = columnsTmp;
					info.value = data;
				});
			} else {
				ElMessage.error('参数错误');
			}
		});
    // 表单组件验证
    const formRulesValidate = (pageRef: string, sonRef: string) => {
      return new Promise((resolve,reject) => {
        proxy.$refs[pageRef].$refs[sonRef].validate((valid: boolean) => {
          if (valid) {
            resolve(valid)
          }else{
            reject('请将表单填写完整')
          }
        });
      });
    };
		const onSubmit = () => {
      Promise.all([
        formRulesValidate('toolsBasicInfoForm', 'basicInfoFormRef'),
        formRulesValidate('genInfoForm', 'genInfoFormRef'),
      ]).then(() => {
        saveEdit(info.value).then(()=>{
          ElMessage.success('保存成功');
          close();
        })
      }).catch(e=>{
        ElMessage.error(e.toString());
      });
		};
		const goNext = () => {
			if (active.value > 2) {
				active.value = 1;
			} else {
				active.value++;
			}
		};
    //取消返回
    const close = ()=>{
      proxy.mittBus.emit('onCurrentContextmenuClick', Object.assign({}, { contextMenuClickId: 1, ...route }));
    }
		return {
			toolsBasicInfoForm,
			genInfoForm,
      tableColumns,
			active,
			changeSteps,
			info,
			onSubmit,
			goNext,
      close
		};
	},
});
</script>

<style scoped>
.tools-footer {
	margin: 30px auto;
	text-align: center;
}
</style>
