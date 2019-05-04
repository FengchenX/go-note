<template>
	<div>
    <el-row :gutter="20">
      <el-col :span="6" v-for="app in apps">
        <div class="grid-content bg-purple">
          <AppItem :application=app></AppItem>
        </div>
      </el-col>
    </el-row>
	</div>
</template>

<script>
import { mapGetters } from 'vuex'
import AppItem from  'components/AppItem'

export default {
	data() {
		return {
		  apps:[
        {name:'视频上传', path: '/upload-video'}
        // {name:'其他'}
      ]
    }
  },
  components: {
    AppItem
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

</style>
