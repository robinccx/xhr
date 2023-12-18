<template>
    <div class="select-dialog">
        <!-- 列表选择对话框 -->
        <el-dialog v-model="isShowDialog" width="512px" :close-on-click-modal="false" :destroy-on-close="true" draggable>
            <template #header>
                <div>{{ caption }}</div>
            </template>
            <el-form ref="formRef" :model="formData" :rules="rules" label-width="60px">
                <el-form-item prop="selectValue">
                    <el-radio-group v-model="formData.selectValue" class="vertical-radio-group" :style="styleDirection">
                        <el-radio v-for="dict in selectOptions" :key="dict.key" :label="dict.key">{{
                            dict.value }}</el-radio>
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
import { defineComponent, reactive, ref, toRefs, onMounted, unref } from 'vue'

export default defineComponent({
    props: {
        selectOptions: {
            type: Array,
            default: () => []
        },
        styleDirection:{
            type:Object 
        }
    },
    setup(props, { emit }) {
        const formRef = ref<HTMLElement | null>(null);
        const caption = ref()

        const state = reactive({
            loading: false,
            isShowDialog: false,
            formData: {
                selectValue: undefined,
            },
            // 表单校验
            rules: {
                selectValue: [
                    { required: true, message: "请选择内容", trigger: "blur" }
                ],
            }
        });

        // 打开弹窗
        const openDialog = (aCaption: string, aDefalutValue: any) => {
            resetForm()
            if (aCaption) {
                caption.value = aCaption
            }

            if (aDefalutValue) {
                state.formData.selectValue = aDefalutValue;
            } 
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
                    closeDialog(); // 关闭弹窗
                    emit('onDialogConfirm', state.formData.selectValue)

                }
            });

        };

        const resetForm = () => {
            state.formData = {
                selectValue: undefined,
            }
        };

        // 页面加载时
        onMounted(() => {
        });
        return {
            caption,
            openDialog,
            closeDialog,
            onCancel,
            onSubmit,
            resetForm,
            formRef,
            ...toRefs(state),
        }
    }
})
</script>

<style scoped>
.vertical-radio-group {
    display: flex;
    /*flex-direction: column;*/
    align-items: flex-start;
    /* 左对齐 */
}
</style>