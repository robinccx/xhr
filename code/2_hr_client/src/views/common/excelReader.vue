<template>
    <input type="file" ref="fileInput" style="display: none;" @change="handleFileChange" />
    <listDialog ref="selectRef" :selectOptions="sheetNames" :isShowDialog="true" @onDialogConfirm="onDialogConfirm" :styleDirection="{'flex-direction': 'column'}">
    </listDialog>
</template>

<script lang="ts">
// 选项类型接口
export interface ItemOptions {
    key: string,
    value: string
}

import { defineComponent, Ref, ref } from 'vue'
import listDialog from "./listDialog.vue"
import * as XLSX from 'xlsx';

export default defineComponent({
    components: {
        listDialog
    },
    emits: ['onSheetReader'],
    setup(props, { emit }) {

        const fileInput: Ref<HTMLInputElement | null> = ref(null);
        const selectRef = ref();
        const sheetNames = ref<Array<ItemOptions>>([])
        const workbook = ref();
        // 导入
        const openDialog = () => {            
            const files = fileInput.value?.files;
            if (files && files.length > 0) {                 
                fileInput.value.value=""
            }
            fileInput.value?.click();
        };
        // 关
        const onDialogConfirm = (data: any) => {
            // console.log("Dialog选择数据:", data);
            const worksheet1 = workbook.value.Sheets[data]
            const jsonData = XLSX.utils.sheet_to_json(worksheet1, {
                header: 1,
                cellDates: true,
                cellText: false,
                dateNF: 'yyyy-mm-dd'
            });

            // console.log('XLSX.jsonData:', jsonData);
            emit('onSheetReader', jsonData)
        }

        const handleFileChange = () => {
            console.log('handleFileChange.begin');
            const files = fileInput.value?.files;
            
            if (files && files.length > 0) {
                const selectedFile = files[0];
                console.log('selectedFile:',selectedFile);
                // 读取上传的 Excel 文件
                const reader = new FileReader();
                // 当文件读取完成时触发的事件
                reader.onload = (e: ProgressEvent<FileReader>) => {
                    const content = e.target?.result; // 读取到的文件内容
                    if (content) {
                        // 将文件内容赋值给变量
                        // console.log('XLSX.read.begin');
                        workbook.value = XLSX.read(content, { type: 'array' });
                        // console.log('XLSX.SheetNames:', workbook.value.SheetNames);
                        sheetNames.value.splice(0)
                        workbook.value.SheetNames.forEach((e: any) => {
                            sheetNames.value.push({
                                key: e.toString(),
                                value: e.toString(),
                            })
                        });

                        selectRef.value.openDialog("Sheet选择")
                    }
                };
                console.log('reader.readAsArrayBuffer.begin');
                reader.readAsArrayBuffer(selectedFile)

            }
        };

        return {
            fileInput,
            selectRef,
            sheetNames,
            openDialog,
            handleFileChange,
            onDialogConfirm
        }
    }
})
</script>

<style scoped></style>