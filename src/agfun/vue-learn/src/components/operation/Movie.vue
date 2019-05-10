<template>
  <el-tabs v-model="activeName" @tab-click="handleClick" class="movie-title">
    <el-tab-pane label="添加电影" name="first">
      <span slot="label">
        <img src="../../assets/images/operation/upload.svg" style="height: 20px;">
      </span>
      <AddMovie></AddMovie>
    </el-tab-pane>
    <el-tab-pane label="查看电影" name="second">
      <span slot="label">
        <img src="../../assets/images/operation/refresh.svg" style="height: 20px;">
      </span>
      查看电影
    </el-tab-pane>
<!--    <el-tab-pane label="角色管理" name="third">角色管理</el-tab-pane>-->
<!--    <el-tab-pane label="定时任务补偿" name="fourth">定时任务补偿</el-tab-pane>-->
  </el-tabs>
</template>
<script>
  import AddMovie from 'components/operation/AddMovie'

  export default {
    data() {
      return {
        activeName: 'first'
      };
    },
    methods: {
      handleClick(tab, event) {
        console.log(tab, event);
      }
    },
    components: {
      AddMovie
    }
  };
</script>
<style scoped>
  .movie-title{
    line-height: 40px;
  }
  .el-icon-upload {
    content: '';
  }
</style>
