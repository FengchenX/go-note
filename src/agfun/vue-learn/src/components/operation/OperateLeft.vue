<template>
  <div class="left-pane">
    <div class="left-title">
      <i></i>
      运维管理
    </div>
    <el-menu
      class="el-menu-vertical-demo"
      @select="handleSelect"
      :router="true"
      background-color="#FFFFFF"
      text-color="#606266"
      active-text-color="#409EFF">
      <el-menu-item v-for="item in items" :index="item.path">
        <i class="el-icon-menu"></i>
        <span slot="title">{{item.name}}</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: 'OperateLeft',
  props: {
		items: {
			type: Array
		}
	},
	data() {
    return {};
  },
  methods: {
    handleSelect: function () {
    }
  }
}
</script>

<style scoped>
  .left-title{
    background: #409EFF;
    color: white;
    height: 40px;
    line-height: 40px;
  }
  .el-menu-vertical-demo{
    text-align: left;
  }
</style>
