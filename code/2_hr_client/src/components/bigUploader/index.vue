<template>

  <div class="uploader-min" v-show="!panelShow" @click="openHandle">
    <el-icon style="vertical-align: middle"><ele-DArrowLeft/></el-icon>
  </div>

  <div class="uploader-container" v-show="panelShow">
    <div class="header">
      <span>文件列表</span>
      <span class="shrink" @click="closeHandle">
        <el-icon style="vertical-align: middle"><ele-DArrowRight /></el-icon>
      </span>
    </div>

    <div class="box">
      <uploader :options="options"
                :file-status-text="handleFileStatusText"
                :autoStart="false"
                @file-added="onFileAdded"
                @file-success="onFileSuccess"
                @file-progress="onFileProgress"
                @file-error="onFileError"
                class="uploader-example">

        <uploader-unsupport></uploader-unsupport>

        <div v-show="false">
          <uploader-btn ref="uploadBtn" :attrs="all">选择文件</uploader-btn>
          <uploader-btn :attrs="image">选择图片</uploader-btn>
          <uploader-btn :directory="true">选择文件夹</uploader-btn>
        </div>

        <uploader-list>
          <template v-slot="props">
            <ul class="file-list">
              <li v-for="file in props.fileList" :key="file.id">
                <uploader-file :ref="'file_' + file.id" :class="'file_' + file.id" :file="file" :list="true"></uploader-file>
              </li>
              <div class="no-file" v-if="!props.fileList.length">暂无待上传文件</div>
            </ul>
          </template>
        </uploader-list>
      </uploader>

    </div>


  </div>

</template>

<script>
import SparkMD5 from 'spark-md5';
import axios from "axios"
import {bigUpload} from "/@/stores/bigUpload";
import {getToken} from "/@/utils/gfast"
const stores = bigUpload();

const baseURL = import.meta.env.VITE_API_URL
const instance = axios.create({
  baseURL: baseURL,
  timeout: 1000 * 30,
});
const acceptConfig  = {
  image: ['gif', 'jpg', 'jpeg', 'png', 'bmp', 'webp'],
  video: ['mp4', 'm3u8', 'rmvb', 'avi', 'swf', '3gp', 'mkv', 'flv'],
  audio: ['mp3', 'wav', 'wma', 'ogg', 'aac', 'flac'],
  document: ['doc', 'txt', 'docx', 'pages', 'epub', 'pdf', 'numbers', 'csv', 'xls', 'xlsx', 'keynote', 'ppt', 'pptx'],
  all() {
    return [...this.image, ...this.video, ...this.audio, ...this.document]
  }
}
export default {
  name: "bigUploader",
  data () {
    return {
      options: {
        target: baseURL + 'api/v1/system/bigUpload/upload',   // 目标上传 URL
        chunkSize: '2048000',         // 分片大小
        fileParameterName: 'upfile',      //上传文件时文件的参数名，默认 file
        maxChunkRetries: 3,             //最大自动失败重试上传次数
        testChunks: true,   //是否开启服务器分片校验
        // 服务器分片校验函数，秒传及断点续传基础
        // 可选的函数用于根据 XHR 响应内容检测每个块是否上传成功了，传入的参数是：Uploader.Chunk 实例以及请求响应信息。
        // 这样就没必要上传（测试）所有的块了
        checkChunkUploadedByResponse: function (chunk, message) {
          try {
            let objMessage = JSON.parse(message)
            const {code, data} = objMessage
            if (code === 0) {
              if (data.skipUpload) {
                return true
              }
              return (data.uploaded || []).indexOf(chunk.offset + 1) >= 0
            }
          } catch (e) {
            console.error(e)
          }

          return false
        },
        headers: {
          Authorization:null
        },
        query: (file, chunk) => {
          return {
            ...file.params,
          }
        },
      },
      image: {
        accept: 'image/*'
      },
      all: {
        accept: acceptConfig.all()
      },
      statusText: {
        success: '成功了',
        error: '出错了',
        uploading: '上传中',
        paused: '暂停中',
        waiting: '等待中'
      }
    }
  },
  mounted() {
     this.$nextTick(()=> {
       this.mittBus.on("bigUploader.uploadFile", this.uploadFile)
     })
  },
  computed:{
    panelShow() {
      return stores.panelShow
    },
    uploadBtn () {
      return this.$refs.uploadBtn && this.$refs.uploadBtn.btn
    }
  },
  methods: {
    openHandle () {
      stores.setPanelShow(true)
    },
    closeHandle() {
      stores.setPanelShow(false)
    },
    uploadFile () {
      // console.log("uploadFile")
      if(!this.options.headers.Authorization){
        this.options.headers.Authorization = "Bearer "+getToken()
      }
      if (this.uploadBtn) {
        this.uploadBtn.click()
      }

    },

    handleFileStatusText (status, response) {
      //console.log(status, response)
      if (status === "success") {
        const {code, data} = response
        if (code === 0 && data.needMerge === true) {
          return "文件合并中..."
        } else {
          return this.statusText[status]
        }
      } else {
        return this.statusText[status]
      }
    },

    //上传过程中出错了
    onFileError(rootFile, file, response, chunk) {
      console.error(rootFile, file, response, chunk)
    },

    // 设置自定义状态
    setStatus (id, text) {
      this.$nextTick(()=> {
        const el = this.$refs['file_'+ id][0].$el.getElementsByClassName("uploader-file-status")[0].getElementsByTagName("span")[0]
        const parent = el.parentNode
        const  para = document.createElement("span")
        para.appendChild(document.createTextNode(text));
        para.className = "para"
        el.style.display = 'none'
        parent.appendChild(para)
      })
    },
    // 删除自定义状态
    removeStatus (id) {
      try{
        const els = this.$refs['file_'+ id][0].$el.getElementsByClassName("uploader-file-status")[0].getElementsByClassName("para")
        if (els && els.length > 0) {
          const parent = els[0].parentNode
          for (let i = 0; i < els.length; i++) {
            parent.removeChild(els[i])
          }
        }
        const firstSpan = this.$refs['file_'+ id][0].$el.getElementsByClassName("uploader-file-status")[0].getElementsByTagName("span")[0]
        firstSpan.style.display = ''
      } catch (e) {
        console.error(e)
      }

    },

    /**
     文件上传成功事件
     第一个参数 rootFile 就是成功上传的文件所属的根 Uploader.File 对象，它应该包含或者等于成功上传文件；
     第二个参数 file 就是当前成功的 Uploader.File 对象本身；
     第三个参数 message 就是服务端响应内容，永远都是字符串；
     第四个参数 chunk 就是 Uploader.Chunk 实例， 它就是该文件的最后一个块实例，如果你想得到请求响应码的话，chunk.xhr.status 就是。
     */
    onFileSuccess(rootFile, file, response, chunk) {
      try{
        // console.log(rootFile, file, response, chunk)
        let res = JSON.parse(response);
        // console.log(res)
        const {code, data} = res
        if (code === 0) {
          // 如果服务端返回需要合并
          if (data.needMerge) {
            // 文件合并操作
            instance.post(baseURL+'api/v1/system/bigUpload/uploadMerge', {
              identifier: data.identifier,
              totalChunks: data.totalChunks,
              totalSize: data.totalSize,
              filename: data.filename,
              token:getToken()
            }).then(res=> {
              const {status, data}  = res
              return data
            }).then(res => {
              let {code, data} = res
              if(code === 0) {
                this.setStatus(file.id, this.statusText["success"])
                this.mittBus.emit("bigUploader.uploadFileSuccess", {...data, fileType: file.fileType})
              }
            })
          } else {
            // 不需要合并
            // 秒传或普通上传
            this.mittBus.emit("bigUploader.uploadFileSuccess", {...data, fileType: file.fileType})
          }
        }
      } catch (e) {
        console.error(e)
      }
    },
    // 一个文件在上传中
    onFileProgress(rootFile, file, chunk) {
      console.log(`上传中 ${file.name}，chunk：${chunk.startByte / 1024 / 1024} ~ ${chunk.endByte / 1024 / 1024}`)
    },
    // 这个事件一般用作文件校验，如果说返回了 false，那么这个文件就会被忽略，不会添加到文件上传列表中。
    onFileAdded(file) {
      this.computeMD5(file);
      // 2022/1/10 将额外的参数赋值到每个文件上，解决了不同文件使用不同params的需求
      file.params = this.params

      this.openHandle()
    },
    computeMD5(file) {
      let fileReader = new FileReader();
      let time = new Date().getTime();
      let blobSlice = File.prototype.slice || File.prototype.mozSlice || File.prototype.webkitSlice;
      // 分片序号
      let currentChunk = 0;
      // 分片大小
      const chunkSize = 10 * 1024 * 1000;
      // 分片总数量
      let chunks = Math.ceil(file.size / chunkSize);
      let spark = new SparkMD5.ArrayBuffer();

      this.setStatus(file.id, "md5计算中")
      file.pause();

      loadNext();

      fileReader.onload = (e => {
        spark.append(e.target.result);

        if (currentChunk < chunks) {
          currentChunk++;
          loadNext();
          // 实时展示MD5的计算进度
          this.$nextTick(() => {
            console.log('校验MD5 '+ ((currentChunk/chunks)*100).toFixed(0)+'%')
          })
        } else {
          let md5 = spark.end();
          this.computeMD5Success(md5, file);
          this.removeStatus(file.id)
          console.log(`MD5计算完毕：${file.name} \nMD5：${md5} \n分片：${chunks} 大小:${file.size} 用时：${new Date().getTime() - time} ms`);
        }
      });

      fileReader.onerror = function () {
        console.error(`文件${file.name}读取出错，请检查该文件`)
        file.cancel();
      };

      function loadNext() {
        let start = currentChunk * chunkSize;
        let end = ((start + chunkSize) >= file.size) ? file.size : start + chunkSize;
        fileReader.readAsArrayBuffer(blobSlice.call(file.file, start, end));
      }
    },
    // md5 计算完毕
    computeMD5Success(md5, file) {
      file.uniqueIdentifier = md5;
      file.resume();
    }
  }
}
</script>

<style scoped  lang="scss">
.uploader-min {
  background-color:#fff;
  z-index: 9999;
  position: absolute;
  right: 0;
  bottom:10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, .4);
  padding:15px;
  cursor: pointer;
}
.uploader-container {
  background-color:#fff;
  z-index: 9999;
  position: absolute;
  right: 0;
  bottom:10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, .4);
}

.header {
  background-color:#fff;
  height:35px;
  line-height: 35px;
  display: flex;
  padding: 0 10px;
  justify-content:space-between;
  .shrink {
    cursor: pointer;
  }
}
.box {
  max-height: 200px;
  overflow-y: scroll;
}



.uploader-example {
  width: 550px;
  padding: 15px;
  font-size: 12px;


}
.uploader-example .uploader-btn {
  margin-right: 4px;
}
.uploader-example .uploader-list {
  overflow-x: hidden;
  overflow-y: hidden;
}



</style>
