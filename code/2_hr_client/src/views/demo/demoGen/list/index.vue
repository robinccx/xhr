<template>
  <div class="demo-demoGen-container">
    <el-card shadow="hover">
        <div class="demo-demoGen-search mb15">
            <el-form :model="tableData.param" ref="queryRef" :inline="true" label-width="100px">
            <el-row>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="姓名" prop="demoName">
                    <el-input
                        v-model="tableData.param.demoName"
                        placeholder="请输入姓名"
                        clearable                        
                        @keyup.enter.native="demoGenList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" class="colBlock">
                  <el-form-item label="年龄" prop="demoAge">
                    <el-input
                        v-model="tableData.param.demoAge"
                        placeholder="请输入年龄"
                        clearable                        
                        @keyup.enter.native="demoGenList"
                    />                    
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="!showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="demoGenList"><el-icon><ele-Search /></el-icon>搜索</el-button>
                    <el-button  @click="resetQuery(queryRef)"><el-icon><ele-Refresh /></el-icon>重置</el-button>
                    <el-button type="primary" link  @click="toggleSearch">
                      {{ word }}
                      <el-icon v-show="showAll"><ele-ArrowUp/></el-icon>
                      <el-icon v-show="!showAll"><ele-ArrowDown /></el-icon>
                    </el-button>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="班级" prop="classes">
                    <el-select v-model="tableData.param.classes" placeholder="请选择班级" clearable   @click.native="getDemoGenClassItemsClasses">
                      <el-option                      
                          v-for="item in classesOptions"                      
                          :key="item.key"
                          :label="item.value"
                          :value="item.key"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="班级2" prop="classesTwo">
                    <el-select v-model="tableData.param.classesTwo" placeholder="请选择班级2" clearable   @click.native="getDemoGenClassItemsClassesTwo">
                      <el-option                      
                          v-for="item in classesTwoOptions"                      
                          :key="item.key"
                          :label="item.value"
                          :value="item.key"
                      />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="出生年月" prop="demoBorn">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.demoBorn"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择出生年月"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="性别" prop="demoGender">
                    <el-select v-model="tableData.param.demoGender" placeholder="请选择性别" clearable >
                        <el-option
                            v-for="dict in sys_user_sex"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="创建日期" prop="createdAt">
                    <el-date-picker
                        clearable  style="width: 200px"
                        v-model="tableData.param.createdAt"
                        format="YYYY-MM-DD HH:mm:ss"
                        value-format="YYYY-MM-DD HH:mm:ss"                    
                        type="datetime"
                        placeholder="选择创建日期"                    
                    ></el-date-picker>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="状态" prop="demoStatus">
                    <el-select v-model="tableData.param.demoStatus" placeholder="请选择状态" clearable >
                        <el-option
                            v-for="dict in sys_normal_disable"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>                
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item label="分类" prop="demoCate">
                    <el-select v-model="tableData.param.demoCate" placeholder="请选择分类" clearable >
                        <el-option
                            v-for="dict in sys_oper_log_status"
                            :key="dict.value"
                            :label="dict.label"
                            :value="dict.value"
                        />
                    </el-select>
                  </el-form-item>
                </el-col>            
                <el-col :span="8" :class="showAll ? 'colBlock' : 'colNone'">
                  <el-form-item>
                    <el-button type="primary"  @click="demoGenList"><el-icon><ele-Search /></el-icon>搜索</el-button>
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
                  v-auth="'api/v1/demo/demoGen/add'"
                ><el-icon><ele-Plus /></el-icon>新增</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="success"
                  :disabled="single"
                  @click="handleUpdate(null)"
                  v-auth="'api/v1/demo/demoGen/edit'"
                ><el-icon><ele-Edit /></el-icon>修改</el-button>
              </el-col>
              <el-col :span="1.5">
                <el-button
                  type="danger"
                  :disabled="multiple"
                  @click="handleDelete(null)"
                  v-auth="'api/v1/demo/demoGen/delete'"
                ><el-icon><ele-Delete /></el-icon>删除</el-button>
              </el-col>
            </el-row>
        </div>
        <el-table v-loading="loading" :data="tableData.data" @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />          
          <el-table-column label="ID" align="center" prop="id"
            min-width="100px"            
             />          
          <el-table-column label="姓名" align="center" prop="demoName"
            min-width="100px"            
             />          
          <el-table-column label="年龄" align="center" prop="demoAge"
            min-width="100px"            
             />          
          <el-table-column label="班级" align="center" prop="linkedClasses.className"
            min-width="100px"            
             />          
          <el-table-column label="班级2" align="center" prop="linkedClassesTwo.className"
            min-width="100px"            
             />          
          <el-table-column label="出生年月" align="center" prop="demoBorn"
            min-width="100px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.demoBorn, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>          
          <el-table-column label="性别" align="center" prop="demoGender" :formatter="demoGenderFormat"
            min-width="100px"            
             />          
          <el-table-column label="创建日期" align="center" prop="createdAt"
            min-width="100px"            
            >
            <template #default="scope">
                <span>{{ proxy.parseTime(scope.row.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
            </template>
          </el-table-column>          
          <el-table-column label="创建人" align="center" prop="createdBy"
            min-width="100px"            
             />          
          <el-table-column label="状态" align="center" prop="demoStatus" :formatter="demoStatusFormat"
            min-width="100px"            
             />          
          <el-table-column label="分类" align="center" prop="demoCate" :formatter="demoCateFormat"
            min-width="100px"            
             />          
          <el-table-column align="center" label="头像"
            min-width="100px"            
            >
            <template #default="scope">
              <el-image
                style="width: 100px; height: 50px"
                v-if="!proxy.isEmpty(scope.row.demoThumb)"
                :src="proxy.getUpFileUrl(scope.row.demoThumb)"
                fit="contain"></el-image>
            </template>
          </el-table-column>        
          <el-table-column label="操作" align="center" class-name="small-padding" min-width="200px" fixed="right">
            <template #default="scope">            
              <el-button
                type="primary"
                link
                @click="handleView(scope.row)"
                v-auth="'api/v1/demo/demoGen/view'"
              ><el-icon><ele-View /></el-icon>详情</el-button>              
              <el-button
                type="primary"
                link
                @click="handleUpdate(scope.row)"
                v-auth="'api/v1/demo/demoGen/edit'"
              ><el-icon><ele-EditPen /></el-icon>修改</el-button>
              <el-button
                type="primary"
                link
                @click="handleDelete(scope.row)"
                v-auth="'api/v1/demo/demoGen/delete'"
              ><el-icon><ele-DeleteFilled /></el-icon>删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <pagination
            v-show="tableData.total>0"
            :total="tableData.total"
            v-model:page="tableData.param.pageNum"
            v-model:limit="tableData.param.pageSize"
            @pagination="demoGenList"
        />
    </el-card>
    <apiV1DemoDemoGenAdd
        ref="addRef"        
        :classesOptions="classesOptions"
        @getDemoGenClassItemsClasses="getDemoGenClassItemsClasses"        
        :classesTwoOptions="classesTwoOptions"
        @getDemoGenClassItemsClassesTwo="getDemoGenClassItemsClassesTwo"        
        :demoGenderOptions="sys_user_sex"        
        :demoStatusOptions="sys_normal_disable"        
        :demoCateOptions="sys_oper_log_status"        
        @demoGenList="demoGenList"
    ></apiV1DemoDemoGenAdd>
    <apiV1DemoDemoGenEdit
       ref="editRef"       
       :classesOptions="classesOptions"
       @getDemoGenClassItemsClasses="getDemoGenClassItemsClasses"       
       :classesTwoOptions="classesTwoOptions"
       @getDemoGenClassItemsClassesTwo="getDemoGenClassItemsClassesTwo"       
       :demoGenderOptions="sys_user_sex"       
       :demoStatusOptions="sys_normal_disable"       
       :demoCateOptions="sys_oper_log_status"       
       @demoGenList="demoGenList"
    ></apiV1DemoDemoGenEdit>
    <apiV1DemoDemoGenDetail
      ref="detailRef"      
      :classesOptions="classesOptions"
      @getDemoGenClassItemsClasses="getDemoGenClassItemsClasses"      
      :classesTwoOptions="classesTwoOptions"
      @getDemoGenClassItemsClassesTwo="getDemoGenClassItemsClassesTwo"      
      :demoGenderOptions="sys_user_sex"      
      :demoStatusOptions="sys_normal_disable"      
      :demoCateOptions="sys_oper_log_status"      
      @demoGenList="demoGenList"
    ></apiV1DemoDemoGenDetail>
  </div>
</template>
<script lang="ts">
import {ItemOptions} from "/@/api/items";
import {toRefs, reactive, onMounted, ref, defineComponent, computed,getCurrentInstance,toRaw} from 'vue';
import {ElMessageBox, ElMessage, FormInstance} from 'element-plus';
import {
    listDemoGen,
    getDemoGen,
    delDemoGen,
    addDemoGen,
    updateDemoGen,    
    listDemoGenClass,    
    getUserList,    
} from "/@/api/demo/demoGen";
import {
    DemoGenTableColumns,
    DemoGenInfoData,
    DemoGenTableDataState,    
    LinkedDemoGenDemoGenClass,    
} from "/@/views/demo/demoGen/list/component/model"
import apiV1DemoDemoGenAdd from "/@/views/demo/demoGen/list/component/add.vue"
import apiV1DemoDemoGenEdit from "/@/views/demo/demoGen/list/component/edit.vue"
import apiV1DemoDemoGenDetail from "/@/views/demo/demoGen/list/component/detail.vue"
export default defineComponent({
    name: "apiV1DemoDemoGenList",
    components:{
        apiV1DemoDemoGenAdd,
        apiV1DemoDemoGenEdit,
        apiV1DemoDemoGenDetail
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
            sys_user_sex,            
            sys_normal_disable,            
            sys_oper_log_status,            
        } = proxy.useDict(            
            'sys_user_sex',            
            'sys_normal_disable',            
            'sys_oper_log_status',            
        )        
        // classesOptions关联表数据
        const classesOptions = ref<Array<ItemOptions>>([])        
        // classesTwoOptions关联表数据
        const classesTwoOptions = ref<Array<ItemOptions>>([])        
        const state = reactive<DemoGenTableDataState>({
            ids:[],
            tableData: {
                data: [],
                total: 0,
                loading: false,
                param: {
                    pageNum: 1,
                    pageSize: 10,                    
                    demoName: undefined,                    
                    demoAge: undefined,                    
                    classes: undefined,                    
                    classesTwo: undefined,                    
                    demoBorn: undefined,                    
                    demoGender: undefined,                    
                    createdAt: undefined,                    
                    demoStatus: undefined,                    
                    demoCate: undefined,                    
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
            demoGenList()
        };
        /** 重置按钮操作 */
        const resetQuery = (formEl: FormInstance | undefined) => {
            if (!formEl) return
            formEl.resetFields()
            demoGenList()
        };
        // 获取列表数据
        const demoGenList = ()=>{
          loading.value = true
          listDemoGen(state.tableData.param).then((res:any)=>{
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
        //关联demo_gen_class表选项
        const getDemoGenClassItemsClasses = () => {
          if (classesOptions.value && classesOptions.value.length > 0) {
            return
          }
          proxy.getItems(listDemoGenClass, {pageSize:10000}).then((res:any) => {
            classesOptions.value = proxy.setItems(res, 'id', 'className')
          })
        }        
        //关联demo_gen_class表选项
        const getDemoGenClassItemsClassesTwo = () => {
          if (classesTwoOptions.value && classesTwoOptions.value.length > 0) {
            return
          }
          proxy.getItems(listDemoGenClass, {pageSize:10000}).then((res:any) => {
            classesTwoOptions.value = proxy.setItems(res, 'id', 'className')
          })
        }        
        // 性别字典翻译
        const demoGenderFormat = (row:DemoGenTableColumns) => {
            return proxy.selectDictLabel(sys_user_sex.value, row.demoGender);
        }        
        // 状态字典翻译
        const demoStatusFormat = (row:DemoGenTableColumns) => {
            return proxy.selectDictLabel(sys_normal_disable.value, row.demoStatus);
        }        
        // 分类字典翻译
        const demoCateFormat = (row:DemoGenTableColumns) => {
                let demoCate = row.demoCate.split(",")
                let data:Array<string> = [];
                demoCate.map(item=>{
                data.push(proxy.selectDictLabel(sys_oper_log_status.value, item))
            })
            return data.join(",")
        }        
        // 多选框选中数据
        const handleSelectionChange = (selection:Array<DemoGenInfoData>) => {
            state.ids = selection.map(item => item.id)
            single.value = selection.length!=1
            multiple.value = !selection.length
        }
        const handleAdd =  ()=>{
            addRef.value.openDialog()
        }
        const handleUpdate = (row: DemoGenTableColumns) => {
            if(!row){
                row = state.tableData.data.find((item:DemoGenTableColumns)=>{
                    return item.id ===state.ids[0]
                }) as DemoGenTableColumns
            }
            editRef.value.openDialog(toRaw(row));
        };
        const handleDelete = (row: DemoGenTableColumns) => {
            let msg = '你确定要删除所选数据？';
            let id:number[] = [] ;
            if(row){
            msg = `此操作将永久删除数据，是否继续?`
            id = [row.id]
            }else{
            id = state.ids
            }
            if(id.length===0){
                ElMessage.error('请选择要删除的数据。');
                return
            }
            ElMessageBox.confirm(msg, '提示', {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: 'warning',
            })
                .then(() => {
                    delDemoGen(id).then(()=>{
                        ElMessage.success('删除成功');
                        demoGenList();
                    })
                })
                .catch(() => {});
        }
        const handleView = (row:DemoGenTableColumns)=>{
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
            demoGenList,
            toggleSearch,            
            //关联表数据选项
            classesOptions,
            //关联demo_gen_class表选项获取数据方法
            getDemoGenClassItemsClasses,            
            //关联表数据选项
            classesTwoOptions,
            //关联demo_gen_class表选项获取数据方法
            getDemoGenClassItemsClassesTwo,            
            demoGenderFormat,
            sys_user_sex,            
            demoStatusFormat,
            sys_normal_disable,            
            demoCateFormat,
            sys_oper_log_status,            
            handleSelectionChange,
            handleAdd,
            handleUpdate,
            handleDelete,            
            handleView,            
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