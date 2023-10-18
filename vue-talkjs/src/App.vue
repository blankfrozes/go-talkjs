<script setup>
import {computed} from 'vue'
import { RouterLink, RouterView } from 'vue-router'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

const router = useRouter()
const store = useStore()
const loggedIn = computed(() => store.state.auth.status.loggedIn)

const user = computed(() => store.state.auth.user)

const handleSignOut = async () => {
  store.dispatch('auth/logout')
  localStorage.removeItem('user')
  return router.push({ name: 'signin' })
}

</script>

<template>
  <header class="sticky top-0 w-full text-white bg-blue-500 shadow-lg">
    <div class="py-2 mx-auto containenr">
      <div class="flex justify-end w-full">

        <div class="flex items-center justify-end space-x-4" v-if="loggedIn">
          <div>
            Hi, {{ user.user.username }}
          </div>

          <button class="px-4 py-2 bg-red-500 rounded-md hover:bg-red-700" @click="handleSignOut">
            Sign Out<font-awesome-icon icon="fa-solid fa-right-from-bracket" class="ml-2" />
          </button>
        </div>

        <div class="flex items-center justify-end space-x-4" v-else>
          <RouterLink to="/signin"
            >Sign In <font-awesome-icon class="ml-2" icon="fa-solid fa-right-to-bracket"
          /></RouterLink>

          <RouterLink to="/signup" class="px-4 py-2 bg-yellow-500 rounded-md hover:bg-yellow-400"
            >Sign Up<font-awesome-icon class="ml-2" icon="fa-solid fa-user-plus"
          /></RouterLink>
        </div>

      </div>
    </div>
  </header>
  <RouterView />
</template>
