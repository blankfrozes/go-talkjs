import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueTransitions from '@morev/vue-transitions';
import '@morev/vue-transitions/styles';
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faRightToBracket, faUserPlus, faCirclePlus, faTrashCan, faPenToSquare, faAngleLeft, faAnglesLeft, faAngleRight, faAnglesRight, faXmark, faFloppyDisk, faRightFromBracket, faCircleCheck, faFilter } from '@fortawesome/free-solid-svg-icons'
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons'
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";
import { createStore } from 'vuex';

import { auth } from './store/auth.module';

const store = createStore({
  modules: {
    auth
  }
});

library.add(faRightToBracket, faUserPlus, faCirclePlus, faTrashCan, faPenToSquare, faAngleLeft, faAnglesLeft, faAngleRight, faAnglesRight, faXmark, faCircleXmark, faFloppyDisk, faRightFromBracket, faCircleCheck, faFilter)

const app = createApp(App).component('font-awesome-icon', FontAwesomeIcon)

const options = {
  transition: "Vue-Toastification__bounce",
  maxToasts: 20,
  newestOnTop: true
};

app.use(router)
app.use(VueTransitions)
app.use(Toast, options)
app.provide('store', store);

app.mount('#app')
