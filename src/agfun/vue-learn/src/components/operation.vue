<template>
	<div>
    <el-row :gutter="20">
      <el-col :span="4"><div class="grid-content bg-both-sides"></div></el-col>
      <el-col :span="20">
        <div class="grid-content bg-purple">
          <el-row :gutter="20">
            <el-col :span="6" v-for="app in items">
              <div class="grid-content bg-purple">
                <OperateItem :item=app></OperateItem>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-col>
    </el-row>
	</div>
</template>

<script>
import { mapGetters } from 'vuex'
import OperateItem from  'components/OperateItem'

export default {
	data() {
		return {
		  items:[
        {name:'视频上传', path: '/upload-video'}
        // {name:'其他'}
      ]
    }
  },
  components: {
    OperateItem
  },
  // computed: {
	// 	...mapGetters([
	// 		'requesting',
	// 		'error',
	// 		'toplist'
	// 	])
  // },
  // mounted(){
  //   this.$store.dispatch('toplist')
  // }
}

</script>

<style scoped>
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
</style>
