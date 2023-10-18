import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SignIn from '../views/SiginView.vue'
import SignUp from '../views/SignupView.vue'
import { useStore } from 'vuex'
import { computed } from 'vue'

function authMiddleware(to) {
  const store = useStore()
  const loggedIn = computed(() => store.state.auth.status.loggedIn)

  if (!loggedIn.value) {
    return router.push({ name: 'signin' });
  }
}

function signInMiddleware(to) {
  const store = useStore()
  const loggedIn = computed(() => store.state.auth.status.loggedIn)

  if (loggedIn.value) {
    return router.push({ name: 'home' });
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      beforeEnter: authMiddleware
    },
    {
      path: '/signin',
      name: 'signin',
      component: SignIn,
      beforeEnter: signInMiddleware
    },
    {
      path: '/signup',
      name: 'signup',
      component: SignUp,
      beforeEnter: signInMiddleware
    },
  ]
})

export default router
