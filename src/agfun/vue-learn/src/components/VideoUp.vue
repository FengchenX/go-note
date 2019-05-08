<template>
  <el-form ref="form" :model="form" label-width="80px" class="form-video">
    <el-form-item label="视频名称">
      <el-input v-model="form.name"></el-input>
    </el-form-item>
    <el-form-item label="视频链接">
      <el-select v-model="form.url" placeholder="请选择活动区域">
        <el-option label="http://www.videos.com" value="shanghai"></el-option>
<!--        <el-option label="区域二" value="beijing"></el-option>-->
      </el-select>
    </el-form-item>
    <el-form-item label="视频描述">
      <el-input v-model="form.describe"></el-input>
    </el-form-item>
    <el-form-item label="拇指图">
      <el-input v-model="form.thumb"></el-input>
    </el-form-item>

    <el-form-item>
      <el-upload
        class="upload-demo"
        ref="upload"
        action="https://jsonplaceholder.typicode.com/posts/"
        :on-preview="handlePreview"
        :on-remove="handleRemove"
        :file-list="fileList"
        :auto-upload="false">
        <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
<!--        <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>-->
        <div slot="tip" class="el-upload__tip">上传视频可能需要较长时间, 请耐心等待!</div>
      </el-upload>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
	data() {
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      },
      fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'},
        {name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}]
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
    }
  }
}

</script>

<style scoped>
  .form-video{
    text-align: left;
    width: 50%;
  }
</style>
