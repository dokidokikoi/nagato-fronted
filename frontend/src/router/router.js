// import Vue from 'vue'
import {createRouter, createWebHashHistory} from 'vue-router'

const floderTree = () => import(/* webpackChunkName: 'wel' */'@/views/foldertree/index.vue')
const blank = () => import(/* webpackChunkName: 'wel' */'@/views/blank/index.vue')
const blankedit = () => import(/* webpackChunkName: 'wel' */'@/views/blank/edit/index.vue')
const blankview = () => import(/* webpackChunkName: 'wel' */'@/views/blank/view/index.vue')

// Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'floderTree',
    component: floderTree,
  },
  {
    path: '/blank',
    name: 'blank',
    component: blank
  },
  {
    path: '/blank/edit',
    name: 'blankedit',
    component: blankedit,
  },
  {
    path: '/blank/view',
    name: 'blankview',
    component: blankview,
  },
]

const router = createRouter({
  // base: process.env.NODE_ENV === 'production' ? '/newProject/index/': 'index', 
	history: createWebHashHistory(),
	routes
})

export default router