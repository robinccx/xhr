import {onMounted, reactive, ref} from "vue";
import {getBigFileList as apiGetBigFileList} from "/@/api/system/bigFile";
import {FormInstance} from "element-plus/es";

export default function () {
    const tableData = ref<any[]>([])
    let total = ref<number>(0)
    const queryParams = reactive({
        pageNum: 1,
        pageSize: 10,
        name: '',
        orderBy:'created_at desc'
    })
    const queryRef = ref<FormInstance>();

    const resetQuery = (formEl: FormInstance | undefined) => {
        if (!formEl) return
        formEl.resetFields()
        getBigFileList()
    }


    const getBigFileList = async () =>  {
        const result:any = await apiGetBigFileList(queryParams).then((res:any) => res.code === 0 ? res.data : {})
        tableData.value =  Array.isArray(result.list) ? result.list : []
        total.value = result.total || 0
    }

    const resetBigFileList = async () => {
        queryParams.pageNum = 1
        queryParams.pageSize = 10
        queryParams.name = ''
        await getBigFileList()
    }


    onMounted(getBigFileList)

    return {
        total,
        queryParams,
        tableData,
        queryRef,
        resetQuery,
        getBigFileList,
        resetBigFileList
    }
}
