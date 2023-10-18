<script setup>
import { ref, reactive, computed } from 'vue'
import { signUp } from '../services/auth.js'
import { required, minLength, helpers, email, sameAs } from '@vuelidate/validators'
import { useVuelidate } from '@vuelidate/core'
import { useToast } from 'vue-toastification'

const toast = useToast()

const input = reactive({
  username: '',
  country_code: '62',
  phone: '',
  email: '',
  password: '',
  confirm_password: ''
})

const rules = computed(() => {
  return {
    username: {
      required: helpers.withMessage('Username is required!', required),
      minLength: helpers.withMessage('Username must minimal 6 characters!', minLength(6))
    },
    country_code: {
      required: helpers.withMessage('Country Code is required!', required),
    },
    phone: {
      required: helpers.withMessage('Phone is required!', required),
    },
    email: {
      required: helpers.withMessage('Email is required!', required),
      email: helpers.withMessage('Not valid email format!', email)
    },
    password: {
      required: helpers.withMessage('Password is required!', required),
      minLength: helpers.withMessage('Password must minimal 8 characters!', minLength(8))
    },
    confirm_password: {
      required: helpers.withMessage('Confirm Password is required!', required),
      minLength: helpers.withMessage('Confirm Password must minimal 8 characters!', minLength(8)),
      sameAs: helpers.withMessage('Confirm Password must same as password', sameAs(input.password))
    }
  }
})

const v$ = useVuelidate(rules, input)

const processSubmit = ref(false)

const handleSignUp = async () => {
  processSubmit.value = true
  const isFormCorrect = await v$.value.$validate()
  if (!isFormCorrect) {
    toast.error(v$.value.$errors[0].$message, {
      timeout: 3000
    })

    input.username = ''
    input.country_code = '62'
    input.phone = ''
    input.email = ''
    input.password = ''
    input.confirm_password = ''

    processSubmit.value = false

    return
  }

  try {
    await signUp(input.username, input.email, input.country_code, input.phone, input.password, input.confirm_password)
    input.username = ''
    input.country_code = '62'
    input.phone = ''
    input.email = ''
    input.password = ''
    input.confirm_password = ''

    v$.value.$reset()

    toast.success('Sign Up Success!', {
      timeout: 3000
    })
  } catch (error) {
    Object.values(error.response.data.errors).forEach((a) => {
      toast.error(a, {
        timeout: 5000
      })
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
          <form @submit.prevent="handleSignUp">
            <h2 class="w-full mb-8 text-3xl font-bold text-center">Sign Up</h2>

            <div class="w-full mb-4">
              <div class="mb-2 font-semibold">Username</div>
              <input
                type="text"
                class="w-full px-4 py-2 rounded"
                placeholder="Username"
                v-model="input.username"
                :class="{ 'border-red-500 border-2': v$.username.$error }"
                @blur="v$.username.$touch"
                :disabled="processSubmit"
              />
              <span v-if="v$.username.$error" class="px-2 text-sm text-red-500">{{
                v$.username.$errors[0].$message
              }}</span>
            </div>

            <div class="w-full mb-4">
              <div class="mb-2 font-semibold">Phone</div>

              <div class="flex items-center w-full space-x-4">
                <input
                  type="text"
                  class="w-20 px-4 py-2 rounded"
                  placeholder="Country Code"
                  v-model="input.country_code"
                  @blur="v$.country_code.$touch"
                  :disabled="processSubmit"
                  readonly
                />

                <input
                  type="text"
                  class="w-full px-4 py-2 rounded"
                  placeholder="Phone"
                  v-model="input.phone"
                  @blur="v$.phone.$touch"
                  :class="{ 'border-red-500 border-2': v$.phone.$error }"
                  :disabled="processSubmit"
                />
              </div>
              <span v-if="v$.phone.$error" class="px-2 text-sm text-red-500">{{
                v$.phone.$errors[0].$message
              }}</span>
            </div>

            <div class="w-full mb-4">
              <div class="mb-2 font-semibold">Email</div>
              <input
                type="text"
                class="w-full px-4 py-2 rounded"
                placeholder="Email"
                v-model="input.email"
                @blur="v$.email.$touch"
                :class="{ 'border-red-500 border-2': v$.email.$error }"
                :disabled="processSubmit"
              />
              <span v-if="v$.email.$error" class="px-2 text-sm text-red-500">{{
                v$.email.$errors[0].$message
              }}</span>
            </div>

            <div class="w-full mb-4">
              <div class="mb-2 font-semibold">Password</div>
              <input
                type="password"
                class="w-full px-4 py-2 rounded"
                placeholder="Password"
                @blur="v$.password.$touch"
                v-model="input.password"
                :class="{ 'border-red-500 border-2': v$.password.$error }"
                :disabled="processSubmit"
              />
              <span v-if="v$.password.$error" class="px-2 text-sm text-red-500">{{
                v$.password.$errors[0].$message
              }}</span>
            </div>

            <div class="w-full mb-8">
              <div class="mb-2 font-semibold">Confirm Password</div>
              <input
                type="password"
                class="w-full px-4 py-2 rounded"
                placeholder="Confirm Password"
                v-model="input.confirm_password"
                @blur="v$.confirm_password.$touch"
                :disabled="processSubmit"
                :class="{ 'border-red-500 border-2': v$.confirm_password.$error }"
              />
              <span v-if="v$.confirm_password.$error" class="px-2 text-sm text-red-500">{{
                v$.confirm_password.$errors[0].$message
              }}</span>
            </div>

            <div>
              <button
                type="submit"
                class="w-full px-4 py-2 font-bold text-white bg-blue-500 rounded-md hover:bg-blue-700 disabled:bg-gray-500"
                :disabled="processSubmit"
              >
                Sign Up <font-awesome-icon class="ml-2" icon="fa-solid fa-user-plus" />
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </main>
</template>
