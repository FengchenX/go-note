<template>
  <div>
    <el-row class="movie-status">
      <el-col :span="4">
        <div style="height: 36px"></div>
      </el-col>
      <el-col :span="16">
        <el-steps :active="active" finish-status="success">
          <el-step title="步骤 1" icon="el-icon-edit"></el-step>
          <el-step title="步骤 2" icon="el-icon-upload"></el-step>
          <el-step title="步骤 3" icon="el-icon-picture"></el-step>
        </el-steps>
      </el-col>
      <el-col :span="4">
        <div style="height: 36px"></div>
      </el-col>
    </el-row>

    <MovieStep1></MovieStep1>
    <MovieStep2></MovieStep2>

    <el-button style="margin-top: 12px;" @click="next">下一步</el-button>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'

  import MovieStep1 from 'components/operation/MovieStep1'
  import MovieStep2 from 'components/operation/MovieStep2'

  var conf = require('config/conf.js');

  export default {
    name: 'AddMovie',
    data() {
      return {
        active: 0
      };
    },
    components: {
      MovieStep1,
      MovieStep2
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
    // },
    methods: {
      next: function() {
        if (this.active++ > 2) this.active = 0;
      }
    }
  }

</script>

<style scoped>
  .movie-status{
    line-height: 30px;
    text-align: left;
  }
</style>
