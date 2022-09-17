import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import GetProblemView from '../views/GetProblemView.vue'
import ProblemView from '../views/ProblemView.vue'



const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
      // component: AboutView

    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue')
      // component: RegisterView
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
      // component: LoginView
    },
    {
      path: '/problem/:id',
      name: 'problem',
      component: () => import('../views/ProblemView.vue')
      // component: ProblemView
    },
    {
      path: '/getproblem',
      name: 'getproblem',
      component: () => import('../views/GetProblemView.vue')
      // component: GetProblemView
    },
    {
      path: '/result',
      name: 'result',
      component: () => import('../views/ResultView.vue')
    }
  ]
})

export default router
