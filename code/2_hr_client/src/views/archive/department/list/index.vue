<template>
  <div class="archive-department-container">
    <el-card shadow="hover">
        <div class="archive-department-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="5" class="colBlock">
                  <el-form-item label="科室代码" prop="deptCode">
                    <el-input
                        v-model="tableData.param.deptCode"
                        placeholder="请输入科室代码"
                        clearable                        
                        @keyup.enter.native="departmentList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="5" class="colBlock">
                  <el-form-item label="科室名称" prop="deptName">
                    <el-input
                        v-model="tableData.param.deptName"
                        placeholder="请输入科室名称"
                        clearable                        
                        @keyup.enter.native="departmentList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="5" class="colBlock">
                  <el-form-item label="片区" prop="deptType">
                    <el-input
                        v-model="tableData.param.deptType"
                        placeholder="请输入片区"
                        clearable                        
                        @keyup.enter.native="departmentList"
                    />                    
                  </el-form-item>
                </el-col> 
              
                <el-col :span="5" class="colBlock">
                  <el-form-item label="状态" prop="status">
                    <el-select v-model="tableData.param.status" placeholder="请选择状态" clearable >
                        <el-option
                            v-for="dict in sys_normal_disable"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col> 

                <el-col :span="4" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="departmentList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>   

                <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="属性1" prop="att1">
                    <el-input
                        v-model="tableData.param.att1"
                        placeholder="请输入属性1"
                        clearable                        
                        @keyup.enter.native="departmentList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="属性2" prop="att2">
                    <el-input
                        v-model="tableData.param.att2"
                        placeholder="请输入属性2"
                        clearable                        
                        @keyup.enter.native="departmentList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="5" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="属性3" prop="att3">
                    <el-input
                        v-model="tableData.param.att3"
                        placeholder="请输入属性3"
                        clearable                        
                        @keyup.enter.native="departmentList"
                    />                    
                  </el-form-item>
                </el-col>            
                <el-col :span="4" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="departmentList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                  v-auth="'api/v1/archive/department/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/archive/department/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/archive/department/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="科室代码" align="center" prop="deptCode"
            min-width="100px"            
             />          
          <el-table-column label="科室名称" align="center" prop="deptName"
            min-width="100px"            
             />          
          <el-table-column label="片区" align="center" prop="deptType"
            min-width="100px"            
             />          
          <el-table-column label="负责人" align="center" prop="leader"
            min-width="100px"            
             />          
          <el-table-column label="电话" align="center" prop="phone"
            min-width="100px"            
             />          
          <el-table-column label="邮件" align="center" prop="email"
            min-width="100px"            
             />          
          <el-table-column label="属性1" align="center" prop="att1"
            min-width="100px"            
             />          
          <el-table-column label="属性2" align="center" prop="att2"
            min-width="100px"            
             />          
          <el-table-column label="属性3" align="center" prop="att3"
            min-width="100px"            
             />          
          <el-table-column label="状态" align="center" prop="status" :formatter="statusFormat"
            min-width="100px"            
             />          
          <el-table-column label="创建时间" align="center" prop="createdAt"
            min-width="100px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>          
          <el-table-column label="创建用户" align="center" prop="createdBy"
            min-width="100px"            
             />        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="160px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/archive/department/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/archive/department/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="departmentList"
        />
    </el-card>
    <apiV1ArchiveDepartmentAdd
        ref="addRef"        
        :statusOptions="sys_normal_disable"        
        @departmentList="departmentList"
    ></apiV1ArchiveDepartmentAdd>
    <apiV1ArchiveDepartmentEdit
       ref="editRef"       
       :statusOptions="sys_normal_disable"       
       @departmentList="departmentList"
    ></apiV1ArchiveDepartmentEdit>
    <apiV1ArchiveDepartmentDetail
      ref="detailRef"      
      :statusOptions="sys_normal_disable"      
      @departmentList="departmentList"
    ></apiV1ArchiveDepartmentDetail>
  </div>
</template>
<script lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listDepartment,
    getDepartment,
    delDepartment,
    addDepartment,
    updateDepartment,    
    getUserList,    
} from "/@/api/archive/department";
import {
    DepartmentTableColumns,
    DepartmentInfoData,
    DepartmentTableDataState,    
} from "/@/views/archive/department/list/component/model"
import apiV1ArchiveDepartmentAdd from "/@/views/archive/department/list/component/add.vue"
import apiV1ArchiveDepartmentEdit from "/@/views/archive/department/list/component/edit.vue"
import apiV1ArchiveDepartmentDetail from "/@/views/archive/department/list/component/detail.vue"
export default defineComponent({
    name: "apiV1ArchiveDepartmentList",
    components:{
        apiV1ArchiveDepartmentAdd,
        apiV1ArchiveDepartmentEdit,
        apiV1ArchiveDepartmentDetail
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
        const state = reactive<DepartmentTableDataState>({
            deptIds:[],
            tableData: {
                data: [],
                total: 0,
                loading: false,
                param: {
                    pageNum: 1,
                    pageSize: 10,                    
                    deptCode: undefined,                    
                    deptName: undefined,                    
                    status: undefined,                    
                    deptType: undefined,                    
                    att1: undefined,                    
                    att2: undefined,                    
                    att3: undefined,                    
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
            departmentList()
        };
        /** 重置按钮操作 */
        const resetQuery = (formEl: FormInstance | undefined) => {
            if (!formEl) return
            formEl.resetFields()
            departmentList()
        };
        // 获取列表数据
        const departmentList = ()=>{
          loading.value = true
          listDepartment(state.tableData.param).then((res:any)=>{
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
        // 状态字典翻译
        const statusFormat = (row:DepartmentTableColumns) => {
            return proxy.selectDictLabel(sys_normal_disable.value, row.status);
        }        
        // 多选框选中数据
        const handleSelectionChange = (selection:Array<DepartmentInfoData>) => {
            state.deptIds = selection.map(item => item.deptId)
            single.value = selection.length!=1
            multiple.value = !selection.length
        }
        const handleAdd =  ()=>{
            addRef.value.openDialog()
        }
        const handleUpdate = (row: DepartmentTableColumns) => {
            if(!row){
                row = state.tableData.data.find((item:DepartmentTableColumns)=>{
                    return item.deptId ===state.deptIds[0]
                }) as DepartmentTableColumns
            }
            editRef.value.openDialog(toRaw(row));
        };
        const handleDelete = (row: DepartmentTableColumns) => {
            let msg = '你确定要删除所选数据？';
            let deptId:number[] = [] ;
            if(row){
            msg = `此操作将永久删除数据，是否继续?`
            deptId = [row.deptId]
            }else{
            deptId = state.deptIds
            }
            if(deptId.length===0){
                ElMessage.error('请选择要删除的数据。');
                return
            }
            ElMessageBox.confirm(msg, '提示', {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: 'warning',
            })
                .then(() => {
                    delDepartment(deptId).then(()=>{
                        ElMessage.success('删除成功');
                        departmentList();
                    })
                })
                .catch(() => {});
        }
        const handleView = (row:DepartmentTableColumns)=>{
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
            departmentList,
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