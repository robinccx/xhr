<template>  
  <div class="demo-demoGen-add">
    <!-- 添加或修改代码生成测试对话框 -->
    <el-dialog v-model="isShowDialog" width="769px" :close-on-click-modal="false" :destroy-on-close="true">
      <template #header>
        <div v-drag="['.demo-demoGen-add .el-dialog', '.demo-demoGen-add .el-dialog__header']">添加代码生成测试</div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="90px">        
        <el-form-item label="姓名" prop="demoName">
          <el-input v-model="formData.demoName" placeholder="请输入姓名" />
        </el-form-item>        
        <el-form-item label="年龄" prop="demoAge">
          <el-input v-model="formData.demoAge" placeholder="请输入年龄" />
        </el-form-item>          
        <el-form-item label="班级" prop="classes">
          <el-select v-model="formData.classes" placeholder="请选择班级"   @click.native="getDemoGenClassItemsClasses">
              <el-option              
                  v-for="item in classesOptions"              
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
        </el-form-item>          
        <el-form-item label="班级2" prop="classesTwo">
          <el-select v-model="formData.classesTwo" placeholder="请选择班级2"   @click.native="getDemoGenClassItemsClassesTwo">
              <el-option              
                  v-for="item in classesTwoOptions"              
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="出生年月" prop="demoBorn">
          <el-date-picker clearable  style="width: 200px"
            v-model="formData.demoBorn"
            type="datetime"
            placeholder="选择出生年月">
          </el-date-picker>
        </el-form-item>          
        <el-form-item label="性别" prop="demoGender">
          <el-radio-group v-model="formData.demoGender">
            <el-radio
              v-for="dict in demoGenderOptions"
              :key="dict.value"
              :label="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>          
        <el-form-item label="状态" prop="demoStatus">
          <el-radio-group v-model="formData.demoStatus">
            <el-radio
              v-for="dict in demoStatusOptions"
              :key="dict.value"
              :label="dict.value"
            >{{dict.label }}</el-radio>
          </el-radio-group>
        </el-form-item>        
        <el-form-item label="分类" prop="demoCate">
          <el-checkbox-group v-model="formData.demoCate">
            <el-checkbox
              v-for="dict in demoCateOptions"
              :key="dict.value"
              :label="dict.value"
            >{{dict.label }}</el-checkbox>
          </el-checkbox-group>
        </el-form-item>        
        <el-form-item label="头像" prop="demoThumb">
          <el-upload
            v-loading="upLoadingDemoThumb"
            :action="baseURL+'api/v1/system/upload/singleImg'"
            :before-upload="beforeAvatarUploadDemoThumb"
            :data="setUpData()"
            :on-success="handleAvatarSuccessDemoThumb"
            :show-file-list="false"
            class="avatar-uploader"
            name="file"
          >
            <img v-if="!proxy.isEmpty(imageUrlDemoThumb)" :src="imageUrlDemoThumb" class="avatar">
            <el-icon v-else class="avatar-uploader-icon"><ele-Plus /></el-icon>
          </el-upload>
        </el-form-item>        
        <el-form-item label="相册" prop="demoPhoto" >
          <upload-img :action="baseURL+'api/v1/system/upload/singleImg'" v-model="formData.demoPhoto" @uploadData="setUpImgListDemoPhoto" :limit="10"></upload-img>
        </el-form-item>        
        <el-form-item label="个人描述">
          <gf-ueditor editorId="ueDemoGenDemoInfo" v-model="formData.demoInfo" @setEditContent="setDemoInfoEditContent"></gf-ueditor>
        </el-form-item>        
        <el-form-item label="相关附件" prop="demoFile" >
          <upload-file :action="baseURL+'api/v1/system/upload/singleFile'" v-model="formData.demoFile" @upFileData="setUpFileListDemoFile" :limit="10"></upload-file>
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
import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance } from 'vue';
import {ElMessageBox, ElMessage, FormInstance,UploadProps} from 'element-plus';
import {
  listDemoGen,
  getDemoGen,
  delDemoGen,
  addDemoGen,
  updateDemoGen,  
  listDemoGenClass,  
  getUserList,  
} from "/@/api/demo/demoGen";
import GfUeditor from "/@/components/ueditor/index.vue"
import {getToken} from "/@/utils/gfast"
import uploadImg from "/@/components/uploadImg/index.vue"
import uploadFile from "/@/components/uploadFile/index.vue"
import {
  DemoGenTableColumns,
  DemoGenInfoData,
  DemoGenTableDataState,
  DemoGenEditState,  
  LinkedDemoGenDemoGenClass,  
} from "/@/views/demo/demoGen/list/component/model"
export default defineComponent({
  name:"apiV1DemoDemoGenEdit",
  components:{    
    GfUeditor,    
    uploadImg,    
    uploadFile,    
  },
  props:{    
    classesOptions:{
      type:Array,
      default:()=>[]
    },    
    classesTwoOptions:{
      type:Array,
      default:()=>[]
    },    
    demoGenderOptions:{
      type:Array,
      default:()=>[]
    },    
    demoStatusOptions:{
      type:Array,
      default:()=>[]
    },    
    demoCateOptions:{
      type:Array,
      default:()=>[]
    },    
  },
  setup(props,{emit}) {    
    const baseURL:string|undefined|boolean = import.meta.env.VITE_API_URL    
    const {proxy} = <any>getCurrentInstance()
    const formRef = ref<HTMLElement | null>(null);
    const menuRef = ref();    
    //图片上传地址
    const imageUrlDemoThumb = ref('')
    //上传加载
    const upLoadingDemoThumb = ref(false)    
    const state = reactive<DemoGenEditState>({
      loading:false,
      isShowDialog: false,
      formData: {        
        id: undefined,        
        demoName: undefined,        
        demoAge: undefined,        
        classes: undefined,        
        classesTwo: undefined,        
        demoBorn: undefined,        
        demoGender: false ,        
        createdAt: undefined,        
        updatedAt: undefined,        
        deletedAt: undefined,        
        createdBy: undefined,        
        updatedBy: undefined,        
        demoStatus: false ,        
        demoCate: [] ,        
        demoThumb: undefined,        
        demoPhoto: [] ,        
        demoInfo: undefined,        
        demoFile: [] ,        
        linkedDemoGenDemoGenClass: {          
          id:undefined,    // 分类id          
          className:undefined,    // 分类名          
        },        
      },
      // 表单校验
      rules: {        
        id : [
            { required: true, message: "ID不能为空", trigger: "blur" }
        ],        
        demoName : [
            { required: true, message: "姓名不能为空", trigger: "blur" }
        ],        
        demoStatus : [
            { required: true, message: "状态不能为空", trigger: "blur" }
        ],        
      }
    });
    // 打开弹窗
    const openDialog = () => {
      resetForm();
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
          //添加
          addDemoGen(state.formData).then(()=>{
              ElMessage.success('添加成功');
              closeDialog(); // 关闭弹窗
              emit('demoGenList')
            }).finally(()=>{
              state.loading = false;
            })
        }
      });
    };
    const resetForm = ()=>{
      state.formData = {        
        id: undefined,        
        demoName: undefined,        
        demoAge: undefined,        
        classes: undefined,        
        classesTwo: undefined,        
        demoBorn: undefined,        
        demoGender: false ,        
        createdAt: undefined,        
        updatedAt: undefined,        
        deletedAt: undefined,        
        createdBy: undefined,        
        updatedBy: undefined,        
        demoStatus: false ,        
        demoCate: [] ,        
        demoThumb: undefined,        
        demoPhoto: [] ,        
        demoInfo: undefined,        
        demoFile: [] ,        
        linkedDemoGenDemoGenClass: {          
          id:undefined,    // 分类id          
          className:undefined,    // 分类名          
        },        
      }
    };    
    //关联demo_gen_class表选项
    const getDemoGenClassItemsClasses = () => {
      emit("getDemoGenClassItemsClasses")
    }    
    //关联demo_gen_class表选项
    const getDemoGenClassItemsClassesTwo = () => {
      emit("getDemoGenClassItemsClassesTwo")
    }    
    //单图上传头像
    const handleAvatarSuccessDemoThumb:UploadProps['onSuccess'] = (res, file) => {
      if (res.code === 0) {
        imageUrlDemoThumb.value = URL.createObjectURL(file.raw!)
        state.formData.demoThumb = res.data.path
      } else {
        ElMessage.error(res.msg)
      }
      upLoadingDemoThumb.value = false
    }
    const beforeAvatarUploadDemoThumb = () => {
      upLoadingDemoThumb.value = true
      return true
    }    
    const setUpData = () => {
      return { token: getToken() }
    }    
    const setUpImgListDemoPhoto = (data:any)=>{
      state.formData.demoPhoto = data
    }    
    //富文本编辑器个人描述
    const setDemoInfoEditContent = (data:string) => {
      state.formData.demoInfo = data
    }    
    const setUpFileListDemoFile = (data:any)=>{
      state.formData.demoFile = data
    }    
    return {
      proxy,
      openDialog,
      closeDialog,
      onCancel,
      onSubmit,
      menuRef,
      formRef,      
      getDemoGenClassItemsClasses,      
      getDemoGenClassItemsClassesTwo,      
      //图片上传地址
      imageUrlDemoThumb,
      //上传加载
      upLoadingDemoThumb,
      handleAvatarSuccessDemoThumb,
      beforeAvatarUploadDemoThumb,      
      setUpData,      
      setUpImgListDemoPhoto,      
      //富文本编辑器个人描述
      setDemoInfoEditContent,      
      setUpFileListDemoFile,      
      baseURL,      
      ...toRefs(state),
    };
  }
})
</script>
<style scoped>  
  .demo-demoGen-add :deep(.avatar-uploader .avatar) {
    width: 178px;
    height: 178px;
    display: block;
  }
  .demo-demoGen-add :deep(.avatar-uploader .el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }
  .demo-demoGen-add :deep(.avatar-uploader .el-upload:hover) {
    border-color: var(--el-color-primary);
  }
  .demo-demoGen-add :deep(.el-icon.avatar-uploader-icon) {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
  }  
</style>