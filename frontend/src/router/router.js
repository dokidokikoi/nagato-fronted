// import Vue from 'vue'
import {createRouter, createWebHashHistory} from 'vue-router'

const floderTree = () => import(/* webpackChunkName: 'wel' */'@/views/foldertree/index.vue')
const blank = () => import(/* webpackChunkName: 'wel' */'@/views/blank/index.vue')
const blankedit = () => import(/* webpackChunkName: 'wel' */'@/views/blank/edit/index.vue')
const blankview = () => import(/* webpackChunkName: 'wel' */'@/views/blank/view/index.vue')
const downloads = () => import(/* webpackChunkName: 'wel' */'@/views/download/index.vue')
const uploads = () => import(/* webpackChunkName: 'wel' */'@/views/upload/index.vue')
const history = () => import(/* webpackChunkName: 'wel' */'@/views/history/index.vue')
const share = () => import(/* webpackChunkName: 'wel' */'@/views/share/index.vue')
const me = () => import(/* webpackChunkName: 'wel' */'@/views/me/index.vue')

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
    path: '/blank/view/:id',
    name: 'blankview',
    component: blankview,
  },
  {
    path: '/downloads',
    name: 'downloads',
    component: downloads,
  },
  {
    path: '/uploads',
    name: 'uploads',
    component: uploads,
  },
  {
    path: '/history',
    name: 'history',
    component: history,
  },
  {
    path: '/share',
    name: 'share',
    component: share,
  },
  {
    path: '/me',
    name: 'me',
    component: me,
  },
]

const router = createRouter({
  // base: process.env.NODE_ENV === 'production' ? '/newProject/index/': 'index', 
	history: createWebHashHistory(),
	routes
})

export default router