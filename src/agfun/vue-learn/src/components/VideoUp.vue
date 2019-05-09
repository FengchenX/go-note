<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-upload
          class="upload-demo"
          ref="upload"
          action="https://jsonplaceholder.typicode.com/posts/"
          :on-preview="handlePreview"
          :on-remove="handleRemove"
          :file-list="fileList"
          :before-upload="beforeUpload"
          :auto-upload="false">
          <el-button slot="trigger" size="small" type="primary">
            选取文件
          </el-button>
          <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">
            上传到服务器
          </el-button>
          <div slot="tip" class="el-upload__tip">
            上传视频可能需要较长时间, 请耐心等待!
          </div>
        </el-upload>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-form ref="form" :model="form" label-width="80px" class="form-video">
          <el-form-item label="视频名称">
            <el-input v-model="form.name"></el-input>
          </el-form-item>
          <el-form-item label="视频链接">
            <el-select v-model="form.url" placeholder="请选择视频链接">
              <el-option label="http://www.videos.com" value="shanghai"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="视频描述">
            <el-input v-model="form.describe"></el-input>
          </el-form-item>
          <el-form-item label="拇指图">
            <el-input v-model="form.thumb"></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即创建</el-button>
            <el-button>取消</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
	data() {
    return {
      form: {
        name: '',
        url: '',
        describe: '',
        thumb: ''
      },
      fileList: []
    }
  },
  // components: {
  // },
  // computed: {
	// 	...mapGetters([
	// 		'requesting',
	// 		'error',
	// 		'toplist'
	// 	])
  // },
  // mounted(){
  //   this.$store.dispatch('toplist')
  // },
  methods: {
    onSubmit: function() {
      console.log('submit!');

    },

    submitUpload: function() {
      this.$refs.upload.submit();
    },
    handleRemove: function(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview: function(file) {
      console.log(file);
    },
    beforeUpload: function (file) {

    }
  }
}

</script>

<style scoped>
  .form-video{
    text-align: left;
  }
  .upload-demo{
    line-height: 50px;
  }
</style>
