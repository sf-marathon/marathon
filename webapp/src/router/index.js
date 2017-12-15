import Vue from 'vue'
import Router from 'vue-router'

import Home from '../pages/Home.vue'
const Join = () => import('../pages/Join.vue')
const Share = () => import('../pages/Share.vue')

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/join',
      name: 'Join',
      component: Join
    },
    {
      path: '/share',
      name: 'Share',
      component: Share
    }
  ]
})
