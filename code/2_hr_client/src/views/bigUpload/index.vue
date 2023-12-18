<template>
  <div>
    <el-card shadow="hover">
      <div class="system-user-search mb15">
        <el-form :model="queryParams" ref="queryRef" :inline="true" label-width="68px">
          <el-form-item label="文件名称" prop="name">
            <el-input
                v-model="queryParams.name"
                placeholder="请输入文件名称"
                clearable
                size="default"
                @keyup.enter.native="getBigFileList"
            />
          </el-form-item>

          <el-form-item>
            <el-button size="default" type="primary" class="ml10" @click="getBigFileList">
              <el-icon>
                <ele-Search/>
              </el-icon>
              查询
            </el-button>
            <el-button size="default" @click="resetQuery(queryRef)">
              <el-icon>
                <ele-Refresh/>
              </el-icon>
              重置
            </el-button>
            <el-button size="default" type="success" class="ml10" @click="uploadHandle">
              <el-icon>
                <ele-FolderAdd/>
              </el-icon>
              文件上传
            </el-button>
            <el-button size="default" type="danger" class="ml10" @click="delMult">
              <el-icon>
                <ele-Delete/>
              </el-icon>
              删除文件
            </el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-table :data="tableData" style="width: 100%" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"/>
        <el-table-column label="标题" align="center" prop="name"/>
        <el-table-column label="大小" align="center" prop="size">
          <template #default="scope">
            {{byteText(scope.row.size)}}
          </template>
        </el-table-column>
        <el-table-column label="文件类型" align="center" prop="mimeType"/>
<!--        <el-table-column label="文件描述" align="center" prop="describe"/>-->
        <el-table-column label="创建时间" align="center" prop="createdAt"/>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" text type="primary" @click="edit(scope.row)">修改</el-button>
            <el-button size="small" text type="primary" @click="del(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <pagination
          v-show="total>0"
          :total="total"
          v-model:page="queryParams.pageNum"
          v-model:limit="queryParams.pageSize"
          @pagination="getBigFileList"
      />


      <EditBigUpload ref="editBigUploadRef" @success="getBigFileList"></EditBigUpload>
    </el-card>
  </div>
</template>

<script lang="ts">

import { onMounted, ref, defineComponent, getCurrentInstance} from 'vue';
import getTableData from '/@/views/bigUpload/tableData'
import {addBigFile, deleteBigFile} from "/@/api/system/bigFile";
import {ElMessage, ElMessageBox} from "element-plus";
import EditBigUpload from '/@/views/bigUpload/component/editBigUpload.vue'


export default defineComponent({
  name: "bigUpload",
  components:{
    EditBigUpload
  },
  setup() {

    const {
      total,
      queryParams,
      tableData,
      queryRef,
      resetQuery,
      getBigFileList,
      resetBigFileList
    } = getTableData()


    const editBigUploadRef = ref();


    const {proxy} = <any>getCurrentInstance();

    const selected = ref<number[]>([])

    // 文件上传
    const uploadHandle = function () {
      proxy.mittBus.emit("bigUploader.uploadFile")
    }

    onMounted(()=> {
      proxy.mittBus.on("bigUploader.uploadFileSuccess", (res:any) => {
        //console.log(res)
        if (res.skipUpload === true) {
          // 秒传
        }
        const {filename, totalSize, url, identifier,  fileType} = res
        addBigFile({
          name: filename,
          size: totalSize,
          path: url,
          mimeType: fileType,
          source: 0,
          md5: identifier
        }).then(() => {
          resetBigFileList()
        })
      })
    })


    const handleSelectionChange =  (selection:any[]) => {
      selected.value = selection.map(item => item.id)
    }

    const edit = function (row:any) {
      editBigUploadRef.value.openDialog(row.id)
    }

    const del = function (row:any) {
      ElMessageBox.confirm(`是否删除文件:${row.name}`, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(()=> {
        deleteBigFile([row.id]).then(()=> {
          getBigFileList()
        })
      })
    }


    const delMult = function () {
      if(selected.value.length===0) {
        ElMessage.error('请选择要删除的数据。');
        return
      }
      ElMessageBox.confirm(`确定删除所选数据?`, '提示', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(()=> {
        deleteBigFile([...selected.value]).then(()=> {
          getBigFileList()
        })
      })

    }



    const byteText  = function (value:number) :string {
      if (value > 1048576 ) {
        return (value / 1024 / 1024).toFixed(2) + "MB"
      } else if (value > 0)  {
        return (value / 1024).toFixed(2) + "KB"
      } else {
        return ""
      }
    }


    return {
      total,
      queryParams,
      tableData,
      queryRef,
      resetQuery,
      getBigFileList,
      resetBigFileList,
      uploadHandle,
      handleSelectionChange,
      byteText,
      edit,
      del,
      delMult,
      editBigUploadRef
    }
  }
})
</script>

<style scoped lang="scss">

</style>
