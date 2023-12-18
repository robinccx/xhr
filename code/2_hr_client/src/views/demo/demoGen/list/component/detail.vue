<template>
  <!-- 代码生成测试详情抽屉 -->  
  <div class="demo-demoGen-detail">
    <el-drawer v-model="isShowDialog" size="80%" direction="ltr">
      <template #header>
        <h4>代码生成测试详情</h4>
      </template>
      <el-form ref="formRef" :model="formData" label-width="100px">          
        <el-row>        
          <el-col :span="12">          
            <el-form-item label="ID">{{ formData.id }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="姓名">{{ formData.demoName }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="年龄">{{ formData.demoAge }}</el-form-item>          
          </el-col>        
          <el-col :span="12">            
            <el-form-item label="班级">{{ formData.linkedClasses?formData.linkedClasses.className:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">            
            <el-form-item label="班级2">{{ formData.linkedClassesTwo?formData.linkedClassesTwo.className:'' }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="出生年月">{{ proxy.parseTime(formData.demoBorn, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="性别">{{ proxy.getOptionValue(formData.demoGender, demoGenderOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="创建日期">{{ proxy.parseTime(formData.createdAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="创建人">{{ formData.createdBy }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="修改人">{{ formData.updatedBy }}</el-form-item>          
          </el-col>        
          <el-col :span="12">          
            <el-form-item label="状态">{{ proxy.getOptionValue(formData.demoStatus, demoStatusOptions,'value','label') }}</el-form-item>          
          </el-col>        
          <el-col :span="12">
            <el-form-item label="分类">
              <el-tag class="ml-2" type="success" v-for="(item,key) in formData.demoCate" :key="'demoCate-'+key">
                {{ proxy.getOptionValue(item, demoCateOptions,'value','label') }}
              </el-tag>
            </el-form-item>
          </el-col>        
          <el-col :span="12">
            <el-form-item label="头像">
              <el-image
                style="width: 150px; height: 150px"
                v-if="!proxy.isEmpty(formData.demoThumb)"
                :src="proxy.getUpFileUrl(formData.demoThumb)"
                fit="contain"></el-image>
            </el-form-item>
          </el-col>        
          <el-col :span="12">
            <el-form-item label="相册">
              <div class="pic-block" v-for="(img,key) in formData.demoPhoto" :key="'demoPhoto-'+key">
                <el-image
                  style="width: 150px; height: 150px"
                  v-if="!proxy.isEmpty(img.url)"
                  :src="proxy.getUpFileUrl(img.url)"
                  fit="contain"></el-image>
              </div>
            </el-form-item>
          </el-col>        
          <el-col :span="12">
            <el-form-item label="个人描述">
              <p v-html="formData.demoInfo"></p>
            </el-form-item>
          </el-col>        
          <el-col :span="12">
            <el-form-item label="相关附件">
              <div class="file-block" v-for="(item,key) in formData.demoFile" :key="'demoFile-'+key">
                <el-link type="primary" :underline="false"
                         :href="proxy.getUpFileUrl(item.url)" target="_blank">{{item.name}}</el-link>
              </div>
            </el-form-item>
          </el-col>      
        </el-row>      
      </el-form>
    </el-drawer>
  </div>
</template>
<script lang="ts">
  import { reactive, toRefs, defineComponent,ref,unref,getCurrentInstance,computed } from 'vue';
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
    name:"apiV1DemoDemoGenDetail",
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
          linkedClasses:{id:undefined,className:undefined },          
          classesTwo: undefined,          
          linkedClassesTwo:{id:undefined,className:undefined },          
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
        const openDialog = (row?: DemoGenInfoData) => {
          resetForm();
          if(row) {
            getDemoGen(row.id!).then((res:any)=>{
              const data = res.data;              
              data.demoCate = data.demoCate.split(",")              
              //单图地址赋值
              imageUrlDemoThumb.value = data.demoThumb ? proxy.getUpFileUrl(data.demoThumb) : ''              
              data.demoPhoto =data.demoPhoto?JSON.parse(data.demoPhoto) : []              
              data.demoFile =data.demoFile?JSON.parse(data.demoFile) : []              
              let listUid = [];              
              listUid.push(data.createdBy,data.updatedBy)              
              getUserList(listUid).then((response:any) =>{
                let users = response.data.list||[]
                users.forEach((user:any)=>{                  
                  if(data.createdBy==user.id){
                    data.createdBy = user.userNickname
                  }                  
                  if(data.updatedBy==user.id){
                    data.updatedBy = user.userNickname
                  }                  
                })
                state.formData = data;
              })              
            })
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
        const resetForm = ()=>{
          state.formData = {            
            id: undefined,            
            demoName: undefined,            
            demoAge: undefined,            
            classes: undefined,            
            linkedClasses:{id:undefined,className:undefined },            
            classesTwo: undefined,            
            linkedClassesTwo:{id:undefined,className:undefined },            
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
        const getClassesOp = computed(()=>{
          getDemoGenClassItemsClasses()
          return props.classesOptions
        })        
        //关联demo_gen_class表选项
        const getDemoGenClassItemsClassesTwo = () => {
          emit("getDemoGenClassItemsClassesTwo")
        }
        const getClassesTwoOp = computed(()=>{
          getDemoGenClassItemsClassesTwo()
          return props.classesTwoOptions
        })        
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
          menuRef,
          formRef,          
          getDemoGenClassItemsClasses,
          getClassesOp,          
          getDemoGenClassItemsClassesTwo,
          getClassesTwoOp,          
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
  .demo-demoGen-detail :deep(.avatar-uploader .avatar) {
    width: 178px;
    height: 178px;
    display: block;
  }
  .demo-demoGen-detail :deep(.avatar-uploader .el-upload) {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
  }
  .demo-demoGen-detail :deep(.avatar-uploader .el-upload:hover) {
    border-color: var(--el-color-primary);
  }
  .demo-demoGen-detail :deep(.el-icon.avatar-uploader-icon) {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
  }  
  .demo-demoGen-detail :deep(.el-form-item--large .el-form-item__label){
    font-weight: bolder;
  }
  .pic-block{
    margin-right: 8px;
  }
  .file-block{
    width: 100%;
    border: 1px solid var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
    margin-bottom: 5px;
    padding: 3px 6px;
  }
  .ml-2{margin-right: 5px;}
</style>