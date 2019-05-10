<template>
  <el-row :gutter="20">
    <el-col :span="12">
      <el-form ref="form" :model="movie" label-width="80px" class="form-video">
        <el-form-item label="视频名称">
          <el-input v-model="movie.name"></el-input>
        </el-form-item>
        <el-form-item label="视频描述">
          <el-input v-model="movie.describe"></el-input>
        </el-form-item>
        <el-form-item label="拇指图">
          <el-input v-model="movie.thumb"></el-input>
        </el-form-item>
        <el-form-item label="电影类型">
          <el-checkbox-group v-model="movie.types">
            <el-checkbox label="爱情" name="type"></el-checkbox>
            <el-checkbox label="恐怖" name="type"></el-checkbox>
            <el-checkbox label="喜剧" name="type"></el-checkbox>
            <el-checkbox label="动作" name="type"></el-checkbox>
            <el-checkbox label="科幻" name="type"></el-checkbox>
            <el-checkbox label="悬疑" name="type"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">立即创建</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<script>
  export default {
    name: "MovieStep1"
  }
</script>

<style scoped>

</style>

