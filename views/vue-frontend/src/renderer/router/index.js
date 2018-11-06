import Vue from 'vue'
import Router from 'vue-router'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: require('@/components/Login').default
    },
    {
      path: '/register',
      name: 'Register',
      component: require('@/components/Register').default
    },
    {
      path: '/home',
      name: 'Home',
      component: require('@/components/Home').default
    }
  ]
})
