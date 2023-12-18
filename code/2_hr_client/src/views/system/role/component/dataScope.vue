<template>
  <div class="system-sysJob-edit">
    <!-- 分配角色数据权限对话框 -->
    <el-dialog title="数据权限设置" v-model="openDataScope" width="769px" append-to-body :close-on-click-modal="false">
      <el-form ref="formRef" :model="form" label-width="80px">
        <el-form-item label="角色名称">
          <el-input v-model="form.roleName" :disabled="true" />
        </el-form-item>

        <el-form-item label="权限范围">
          <el-select v-model="form.dataScope">
            <el-option
                v-for="item in dataScopeOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="数据权限" v-show="form.dataScope == 2">
          <el-row :gutter="35">
            <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
              <el-checkbox v-model="deptExpand" @change="handleCheckedTreeExpand($event)">展开/折叠</el-checkbox>
              <el-checkbox v-model="deptNodeAll" @change="handleCheckedTreeNodeAll($event)">全选/全不选</el-checkbox>
              <el-checkbox v-model="form.deptCheckStrictly" @change="handleCheckedTreeConnect($event)">父子联动</el-checkbox>
            </el-col>
            <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="mb20">
              <el-tree
                  class="tree-border"
                  :data="deptOptions"
                  show-checkbox
                  default-expand-all
                  ref="deptRef"
                  node-key="deptId"
                  :check-strictly="!form.deptCheckStrictly"
                  :props="deptProps"
              ></el-tree>
            </el-col>
          </el-row>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitDataScope">确 定</el-button>
          <el-button @click="cancelDataScope">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import {defineComponent, reactive, ref} from "vue";
import {dataScope, getRole, roleDeptTreeSelect} from "/@/api/system/role";
import {ElMessage} from "element-plus/es";

export default defineComponent({
  name: "dataScope",
  setup(props,{emit}){
    const openDataScope = ref(false)
    const deptExpand =  ref(true)
    const deptNodeAll = ref(false)
    const deptOptions = ref<Array<any>>([])
    const menuOptions = ref([])
    const deptRef = ref();
    const formRef = ref();
    const deptProps = ref({
      children: "children",
          label: "deptName"
    })
    // 数据范围选项
    const dataScopeOptions = ref([
      {
        value: "1",
        label: "全部数据权限"
      },
      {
        value: "2",
        label: "自定数据权限"
      },
      {
        value: "3",
        label: "本部门数据权限"
      },
      {
        value: "4",
        label: "本部门及以下数据权限"
      },
      {
        value: "5",
        label: "仅本人数据权限"
      }
    ])
    const form = reactive({
      roleId:undefined,
      roleName:'',
      dataScope:'3',
      deptCheckStrictly: false,
      deptIds:[]
    })
    // 树权限（展开/折叠）
    const handleCheckedTreeExpand = (value:any) => {
      let treeList = deptOptions.value;
      for (let i = 0; i < treeList.length; i++) {
        deptRef.value.store.nodesMap[treeList[i].deptId].expanded = value;
      }
    }
    // 树权限（全选/全不选）
    const handleCheckedTreeNodeAll = (value:any) => {
      deptRef.value.setCheckedNodes(value ? deptOptions.value: []);
    }
    /** 树权限（父子联动） */
    const handleCheckedTreeConnect = (value:any) => {
      form.deptCheckStrictly = value ? true : false;
    }
    /** 提交按钮（数据权限） */
    const submitDataScope = () => {
      if (form.roleId != undefined) {
        form.deptIds = getDeptAllCheckedKeys();
        dataScope(form).then((response:any) => {
          if (response.code === 0) {
            ElMessage.success("设置成功");
            openDataScope.value = false;
            emit('getRoleList')
          } else {
            ElMessage.error("设置失败")
          }
        });
      }
    }
    const cancelDataScope = () => {
      closeDialog();
    };
    // 打开弹窗
    const openDialog = (row: any) => {
      resetForm();
      if(row) {
        getRole(row.id).then((res:any)=>{
          if(res.data.role){
            form.roleName = res.data.role.name;
            form.roleId = res.data.role.id;
            form.dataScope=res.data.role.dataScope.toString();
          }
        })
        getRoleDeptTreeselect(row.id)
      }
      openDataScope.value = true;
    };
    /** 根据角色ID查询部门树结构 */
    const getRoleDeptTreeselect = (roleId:number) =>{
      roleDeptTreeSelect(roleId).then(response => {
        deptOptions.value = response.data.depts||[];
        deptRef.value.setCheckedKeys(response.data.checkedKeys, true);
      });
    }
    // 关闭弹窗
    const closeDialog = () => {
      openDataScope.value = false;
    };
    // 所有部门节点数据
    const getDeptAllCheckedKeys = () => {
      // 目前被选中的部门节点
      let checkedKeys = deptRef.value.getCheckedKeys();
      return checkedKeys;
    }
    const resetForm = ()=>{
      form.roleId=undefined
      form.roleName=''
      form.dataScope='3'
      form.deptCheckStrictly= false
      form.deptIds=[]
    };
    return {
      openDialog,
      dataScopeOptions,
      deptExpand,
      openDataScope,
      deptNodeAll,
      deptOptions,
      deptProps,
      menuOptions,
      deptRef,
      formRef,
      cancelDataScope,
      submitDataScope,
      handleCheckedTreeExpand,
      handleCheckedTreeNodeAll,
      handleCheckedTreeConnect,
      form
    }
  }
})
</script>

<style scoped lang="scss">
.tree-border {
  margin-top: 5px;
  border: 1px solid #e5e6e7!important;
  border-radius: 4px;
}
.system-edit-role-container {
  .menu-data-tree {
    border: var(--el-input-border, var(--el-border-base));
    border-radius: var(--el-input-border-radius, var(--el-border-radius-base));
    padding: 5px;
  }
}
</style>
