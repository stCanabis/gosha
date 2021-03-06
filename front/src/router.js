import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/entity',
      name: 'entity',
      component: () => import(/* webpackChunkName: "about" */ './views/Entity.vue')
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import(/* webpackChunkName: "about" */ './views/Entity.vue')
    },
    {
      path: '/bpm',
      name: 'bpm',
      component: () => import(/* webpackChunkName: "about" */ './views/Bpm.vue')
    },
  ]
})
