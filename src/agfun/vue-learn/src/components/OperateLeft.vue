<template>
  <div>
<!--    <el-row :gutter="0" v-for="item in items">-->
<!--      <el-col :span="24">-->
<!--        <div class="grid-content bg-purple">-->
<!--          <el-button type="primary">-->
<!--            {{ item.name }}-->
<!--          </el-button>-->
<!--        </div>-->
<!--      </el-col>-->
<!--    </el-row>-->

    <el-menu
      default-active="1"
      class="el-menu-vertical-demo"
      @open="handleOpen"
      @close="handleClose"
      background-color="#FFFFFF"
      text-color="#606266"
      active-text-color="#409EFF">
      <el-menu-item v-for="item in items" :index="item.id">
        <i class="el-icon-menu"></i>
        <span slot="title">{{item.name}}</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  props: {
		items: {
			type: Object
		}
	},
	data() {
    return {};
  },
  methods: {
    handleOpen: function () {
      console.log("1111111111111111111111")
    }
  }
}
</script>

<style scoped>
  /*.grid-content{*/
  /*  height: 72px;*/
  /*}*/
  /*.bg-purple{*/
  /*  background: #FFFFFF;*/
  /*}*/
</style>
