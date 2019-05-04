import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

import home from 'components/home'
import about from 'components/about'
import document from 'components/document'
import application from 'components/application'
import VideoUp from  'components/VideoUp'

let router = new VueRouter({
    mode: 'history',
    routes: [
      {
        path: '/home',
        component: home
      },
      {
        path: '/about',
        component: about
      },
      {
        path: '/document',
        component: document
      },
      {
        path: '/application',
        component: application
        // children:[{
        //   path: 'upload-video',
        //   component: VideoUp
        // }]
      },
      {
        path: '/upload-video',
        component: VideoUp
      },
      {
        path: '*',
        component: home
      }
    ]
})

export default router;
