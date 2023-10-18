<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { signIn } from '../services/auth.js'
import { required, helpers } from '@vuelidate/validators'
import { useVuelidate } from '@vuelidate/core'
import { useToast } from 'vue-toastification'
import { useStore, mapState } from 'vuex'
import { useRouter } from 'vue-router'
import Talk from 'talkjs';

const router = useRouter()

const store = useStore()

const toast = useToast()

const input = reactive({
  username: '',
  password: ''
})

const rules = computed(() => {
  return {
    username: {
      required: helpers.withMessage('Username is required!', required)
    },
    password: {
      required: helpers.withMessage('Password is required!', required)
    }
  }
})

const v$ = useVuelidate(rules, input)

const processSubmit = ref(false)

const handleSignIn = async () => {
  processSubmit.value = true
  const isFormCorrect = await v$.value.$validate()
  if (!isFormCorrect) {
    toast.error(v$.value.$errors[0].$message, {
      timeout: 3000
    })

    input.username = ''
    input.password = ''

    processSubmit.value = false

    return
  }

  try {
    const user = await signIn(input.username, input.password)
    input.username = ''
    input.password = ''

    v$.value.$reset()

    store.dispatch('auth/login', { user: user });
    localStorage.setItem('user', JSON.stringify(user))

    toast.success('Sign In Success!', {
      timeout: 3000
    })

    router.go({ name: 'home' })
  } catch (error) {
    toast.error(error, {
      timeout: 5000
    })

    processSubmit.value = false
    return
  }

  processSubmit.value = false
}
</script>

<template>
  <main class="w-full">
    <div class="container mx-auto">
      <div class="flex items-center justify-center w-full pt-20 pb-10">
        <div class="w-full max-w-xl px-8 py-6 bg-white rounded-md shadow-md">
          <form @submit.prevent="handleSignIn">
            <h2 class="w-full mb-8 text-3xl font-bold text-center">Sign In</h2>

            <div class="w-full mb-4">
              <div class="mb-2 font-semibold">Username</div>
              <input
                type="text"
                class="w-full px-4 py-2 rounded"
                placeholder="Username"
                v-model="input.username"
                @blur="v$.username.$touch"
                :class="{ 'border-red-500 border-2': v$.username.$error }"
                :disabled="processSubmit"
              />
              <span v-if="v$.username.$error" class="px-2 text-sm text-red-500">{{
                v$.username.$errors[0].$message
              }}</span>
            </div>

            <div class="w-full mb-8">
              <div class="mb-2 font-semibold">Password</div>
              <input
                type="password"
                class="w-full px-4 py-2 rounded"
                placeholder="Password"
                v-model="input.password"
                @blur="v$.password.$touch"
                :class="{ 'border-red-500 border-2': v$.password.$error }"
                :disabled="processSubmit"
              />
              <span v-if="v$.password.$error" class="px-2 text-sm text-red-500">{{
                v$.password.$errors[0].$message
              }}</span>
            </div>

            <div>
              <button
                type="submit"
                class="w-full px-4 py-2 font-bold text-white bg-blue-500 rounded-md hover:bg-blue-700 disabled:bg-gray-500"
                :disabled="processSubmit"
              >
                Sign In
                <font-awesome-icon class="ml-2" icon="fa-solid fa-right-to-bracket" />
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </main>
</template>
