<template>
  <div class="archive-nurseImport-container">
    <el-card shadow="hover">
        <div class="archive-nurseImport-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="4" class="colBlock">
                  <el-form-item label="" prop="nurseId">
                    <el-input
                        v-model="tableData.param.nurseId"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" class="colBlock">
                  <el-form-item label="" prop="sessionId">
                    <el-input
                        v-model="tableData.param.sessionId"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" class="colBlock">
                  <el-form-item label="" prop="nurseCode">
                    <el-input
                        v-model="tableData.param.nurseCode"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" class="colBlock">
                  <el-form-item label="" prop="nurseName">
                    <el-input
                        v-model="tableData.param.nurseName"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="nurseImportList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="startDate">
                    <el-input
                        v-model="tableData.param.startDate"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="endDate">
                    <el-input
                        v-model="tableData.param.endDate"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="sex">
                    <el-select v-model="tableData.param.sex" placeholder="请选择" clearable >
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="birthday">
                    <el-input
                        v-model="tableData.param.birthday"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="areaId">
                    <el-input
                        v-model="tableData.param.areaId"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="areaCode">
                    <el-input
                        v-model="tableData.param.areaCode"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="deptId">
                    <el-input
                        v-model="tableData.param.deptId"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="deptCode">
                    <el-input
                        v-model="tableData.param.deptCode"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="educationId">
                    <el-input
                        v-model="tableData.param.educationId"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="educationCode">
                    <el-input
                        v-model="tableData.param.educationCode"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="titleId">
                    <el-input
                        v-model="tableData.param.titleId"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="titleCode">
                    <el-input
                        v-model="tableData.param.titleCode"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="staffType">
                    <el-select v-model="tableData.param.staffType" placeholder="请选择" clearable >
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="staffTypeCode">
                    <el-input
                        v-model="tableData.param.staffTypeCode"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="note">
                    <el-input
                        v-model="tableData.param.note"
                        placeholder="请输入"
                        clearable                        
                        @keyup.enter.native="nurseImportList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="status">
                    <el-select v-model="tableData.param.status" placeholder="请选择" clearable >
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>            
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="nurseImportList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                        {{ word }}
                        <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                        <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>            
              </el-row>
            </el-form>
            <el-row :gutter="10" class="mb8">
              <el-col :span="1.5">
                <el-button
                  type="primary"
                  @click="handleAdd"
                  v-auth="'api/v1/archive/nurseImport/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/archive/nurseImport/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/archive/nurseImport/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="" align="center" prop="nurseId"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="sessionId"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="nurseCode"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="nurseName"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="startDate"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="endDate"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="sex"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="birthday"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="areaId"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="areaCode"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="deptId"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="deptCode"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="educationId"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="educationCode"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="titleId"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="titleCode"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="staffType"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="staffTypeCode"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="note"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="status"
            min-width="100px"            
             />          
          <el-table-column label="" align="center" prop="createdAt"
            min-width="100px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>          
          <el-table-column label="" align="center" prop="createdBy"
            min-width="100px"            
             />        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/archive/nurseImport/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/archive/nurseImport/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="nurseImportList"
        />
    </el-card>
    <apiV1ArchiveNurseImportAdd
        ref="addRef"        
        @nurseImportList="nurseImportList"
    ></apiV1ArchiveNurseImportAdd>
    <apiV1ArchiveNurseImportEdit
       ref="editRef"       
       @nurseImportList="nurseImportList"
    ></apiV1ArchiveNurseImportEdit>
    <apiV1ArchiveNurseImportDetail
      ref="detailRef"      
      @nurseImportList="nurseImportList"
    ></apiV1ArchiveNurseImportDetail>
  </div>
</template>
<script lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listNurseImport,
    getNurseImport,
    delNurseImport,
    addNurseImport,
    updateNurseImport,    
    getUserList,    
} from "/@/api/archive/nurseImport";
import {
    NurseImportTableColumns,
    NurseImportInfoData,
    NurseImportTableDataState,    
} from "/@/views/archive/nurseImport/list/component/model"
import apiV1ArchiveNurseImportAdd from "/@/views/archive/nurseImport/list/component/add.vue"
import apiV1ArchiveNurseImportEdit from "/@/views/archive/nurseImport/list/component/edit.vue"
import apiV1ArchiveNurseImportDetail from "/@/views/archive/nurseImport/list/component/detail.vue"
export default defineComponent({
    name: "apiV1ArchiveNurseImportList",
    components:{
        apiV1ArchiveNurseImportAdd,
        apiV1ArchiveNurseImportEdit,
        apiV1ArchiveNurseImportDetail
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
        } = proxy.useDict(            
        )        
        const state = reactive<NurseImportTableDataState>({
            nurseIds:[],
            tableData: {
                data: [],
                total: 0,
                loading: false,
                param: {
                    pageNum: 1,
                    pageSize: 10,                    
                    nurseId: undefined,                    
                    sessionId: undefined,                    
                    nurseCode: undefined,                    
                    nurseName: undefined,                    
                    startDate: undefined,                    
                    endDate: undefined,                    
                    sex: undefined,                    
                    birthday: undefined,                    
                    areaId: undefined,                    
                    areaCode: undefined,                    
                    deptId: undefined,                    
                    deptCode: undefined,                    
                    educationId: undefined,                    
                    educationCode: undefined,                    
                    titleId: undefined,                    
                    titleCode: undefined,                    
                    staffType: undefined,                    
                    staffTypeCode: undefined,                    
                    note: undefined,                    
                    status: undefined,                    
                    createdAt: undefined,                    
                    createdBy: undefined,                    
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
            nurseImportList()
        };
        /** 重置按钮操作 */
        const resetQuery = (formEl: FormInstance | undefined) => {
            if (!formEl) return
            formEl.resetFields()
            nurseImportList()
        };
        // 获取列表数据
        const nurseImportList = ()=>{
          loading.value = true
          listNurseImport(state.tableData.param).then((res:any)=>{
            let list = res.data.list??[];            
            let listUid = [];            
            listUid = list.map((item:any)=>{
                return item.createdBy
            });            
            if(listUid.length>0){
                getUserList(listUid).then((response:any) =>{
                    let users = response.data.list||[]
                    list.map((item:any)=>{
                        users.forEach((user:any)=>{                            
                            if(item.createdBy==user.id){
                                item.createdBy = user.userNickname
                            }                            
                        })
                    })
                    state.tableData.data = list;
                })
            }else{
                state.tableData.data = list;
            }            
            state.tableData.total = res.data.total;
            loading.value = false
          })
        };
        const toggleSearch = () => {
            showAll.value = !showAll.value;
        }        
        // 多选框选中数据
        const handleSelectionChange = (selection:Array<NurseImportInfoData>) => {
            state.nurseIds = selection.map(item => item.nurseId)
            single.value = selection.length!=1
            multiple.value = !selection.length
        }
        const handleAdd =  ()=>{
            addRef.value.openDialog()
        }
        const handleUpdate = (row: NurseImportTableColumns) => {
            if(!row){
                row = state.tableData.data.find((item:NurseImportTableColumns)=>{
                    return item.nurseId ===state.nurseIds[0]
                }) as NurseImportTableColumns
            }
            editRef.value.openDialog(toRaw(row));
        };
        const handleDelete = (row: NurseImportTableColumns) => {
            let msg = '你确定要删除所选数据？';
            let nurseId:number[] = [] ;
            if(row){
            msg = `此操作将永久删除数据，是否继续?`
            nurseId = [row.nurseId]
            }else{
            nurseId = state.nurseIds
            }
            if(nurseId.length===0){
                ElMessage.error('请选择要删除的数据。');
                return
            }
            ElMessageBox.confirm(msg, '提示', {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: 'warning',
            })
                .then(() => {
                    delNurseImport(nurseId).then(()=>{
                        ElMessage.success('删除成功');
                        nurseImportList();
                    })
                })
                .catch(() => {});
        }
        const handleView = (row:NurseImportTableColumns)=>{
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
            nurseImportList,
            toggleSearch,            
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