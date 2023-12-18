<template>
  <div class="archive-position-container">
    <el-card shadow="hover">
        <div class="archive-position-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="80px">
            <el-row>                
                <el-col :span="5" class="colBlock">
                  <el-form-item label="职务" prop="positionName">
                    <el-input
                        v-model="tableData.param.positionName"
                        placeholder="请输入职务"
                        clearable                        
                        @keyup.enter.native="positionList"
                    />                    
                  </el-form-item>
                </el-col>            
                <el-col :span="4" class="colBlock">
                  <el-form-item>
                    <el-button type="primary"   @click="positionList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                  </el-form-item>
                </el-col>            
              </el-row>
            </el-form>
            <el-row :gutter="10" class="mb8">
              <el-col :span="1.5">
                <el-button
                  type="primary"
                  @click="handleAdd"
                  v-auth="'api/v1/archive/position/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/archive/position/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/archive/position/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="职务" align="center" prop="positionName"
            min-width="100px"            
             />          
          <el-table-column label="次序" align="center" prop="seq"
            min-width="100px"            
             />          
          <el-table-column label="状态" align="center" prop="status" :formatter="statusFormat"
            min-width="100px"            
             />        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/archive/position/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/archive/position/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="positionList"
        />
    </el-card>
    <apiV1ArchivePositionAdd
        ref="addRef"        
        :statusOptions="sys_normal_disable"        
        @positionList="positionList"
    ></apiV1ArchivePositionAdd>
    <apiV1ArchivePositionEdit
       ref="editRef"       
       :statusOptions="sys_normal_disable"       
       @positionList="positionList"
    ></apiV1ArchivePositionEdit>
    <apiV1ArchivePositionDetail
      ref="detailRef"      
      :statusOptions="sys_normal_disable"      
      @positionList="positionList"
    ></apiV1ArchivePositionDetail>
  </div>
</template>
<script lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listPosition,
    getPosition,
    delPosition,
    addPosition,
    updatePosition,    
    getUserList,    
} from "/@/api/archive/position";
import {
    PositionTableColumns,
    PositionInfoData,
    PositionTableDataState,    
} from "/@/views/archive/position/list/component/model"
import apiV1ArchivePositionAdd from "/@/views/archive/position/list/component/add.vue"
import apiV1ArchivePositionEdit from "/@/views/archive/position/list/component/edit.vue"
import apiV1ArchivePositionDetail from "/@/views/archive/position/list/component/detail.vue"
export default defineComponent({
    name: "apiV1ArchivePositionList",
    components:{
        apiV1ArchivePositionAdd,
        apiV1ArchivePositionEdit,
        apiV1ArchivePositionDetail
    },
    setup() {
        const {proxy} = <any>getCurrentInstance()
        const loading = ref(false)
        const queryRef = ref()
        const addRef = ref();
        const editRef = ref();
        const detailRef = ref();
        // 是否显示所有搜索选项
        const showAll =  ref(false)
        // 非单个禁用
        const single = ref(true)
        // 非多个禁用
        const multiple =ref(true)
        const word = computed(()=>{
            if(showAll.value === false) {
                //对文字进行处理
                return "展开搜索";
            } else {
                return "收起搜索";
            }
        })
        // 字典选项数据        
        const {            
            sys_normal_disable,            
        } = proxy.useDict(            
            'sys_normal_disable',            
        )        
        const state = reactive<PositionTableDataState>({
            positionIds:[],
            tableData: {
                data: [],
                total: 0,
                loading: false,
                param: {
                    pageNum: 1,
                    pageSize: 10,                    
                    positionName: undefined,                    
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
            positionList()
        };
        /** 重置按钮操作 */
        const resetQuery = (formEl: FormInstance | undefined) => {
            if (!formEl) return
            formEl.resetFields()
            positionList()
        };
        // 获取列表数据
        const positionList = ()=>{
          loading.value = true
          listPosition(state.tableData.param).then((res:any)=>{
            let list = res.data.list??[];            
            state.tableData.data = list;            
            state.tableData.total = res.data.total;
            loading.value = false
          })
        };
        const toggleSearch = () => {
            showAll.value = !showAll.value;
        }        
        // 状态字典翻译
        const statusFormat = (row:PositionTableColumns) => {
            return proxy.selectDictLabel(sys_normal_disable.value, row.status);
        }        
        // 多选框选中数据
        const handleSelectionChange = (selection:Array<PositionInfoData>) => {
            state.positionIds = selection.map(item => item.positionId)
            single.value = selection.length!=1
            multiple.value = !selection.length
        }
        const handleAdd =  ()=>{
            addRef.value.openDialog()
        }
        const handleUpdate = (row: PositionTableColumns) => {
            if(!row){
                row = state.tableData.data.find((item:PositionTableColumns)=>{
                    return item.positionId ===state.positionIds[0]
                }) as PositionTableColumns
            }
            editRef.value.openDialog(toRaw(row));
        };
        const handleDelete = (row: PositionTableColumns) => {
            let msg = '你确定要删除所选数据？';
            let positionId:number[] = [] ;
            if(row){
            msg = `此操作将永久删除数据，是否继续?`
            positionId = [row.positionId]
            }else{
            positionId = state.positionIds
            }
            if(positionId.length===0){
                ElMessage.error('请选择要删除的数据。');
                return
            }
            ElMessageBox.confirm(msg, '提示', {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: 'warning',
            })
                .then(() => {
                    delPosition(positionId).then(()=>{
                        ElMessage.success('删除成功');
                        positionList();
                    })
                })
                .catch(() => {});
        }
        const handleView = (row:PositionTableColumns)=>{
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
            positionList,
            toggleSearch,            
            statusFormat,
            sys_normal_disable,            
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