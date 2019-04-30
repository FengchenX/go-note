// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VueLazyload from 'vue-lazyload'
import iView from 'iview';
import 'iview/dist/styles/iview.css';

Vue.use(VueLazyload)
Vue.use(iView);

import store from './store'
import router from './router'
/* eslint-disable no-new */
new Vue({
	store,
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
