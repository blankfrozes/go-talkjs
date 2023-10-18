<script setup>
import { computed, onMounted, ref } from 'vue'
import { useStore } from 'vuex'
import Talk from 'talkjs';

const store = useStore()
const user = computed(() => store.state.auth.user)

const openConversationRef = ref();

onMounted( async () => {
  await Talk.ready;

  const me = new Talk.User({
    id: user.value.user.username,
    name: user.value.user.username,
    email: user.value.user.email,
    role: "user",
  })

  const session = new Talk.Session({
    appId: 'tzlO2unG',
    me: me,
  });

  const openConversation = async (orderCode) => { 
    const conversation = session.getOrCreateConversation("order-buy-" + orderCode);

    const chatbox = session.createChatbox();

    chatbox.select(conversation);
    chatbox.mount(document.getElementById('talkjs'));
  }
  
  openConversationRef.value = openConversation
})

const callOpenConversation = (orderCode) => {
  openConversationRef.value(orderCode);
};


</script>

<template>
  <main class="min-h-[100vh]">
    <div class="container flex py-20 mx-auto space-x-6">
      <div class="w-40">
        <ul class="w-full text-center">
          <li class="px-4 py-2 mb-4 text-black bg-gray-200 cursor-pointer last:mb-0 hover:bg-green-500 hover:text-white" @click="callOpenConversation('BJA24KS8MJ1')">BJA24KS8MJ1</li>
          <li class="px-4 py-2 mb-4 text-black bg-gray-200 cursor-pointer last:mb-0 hover:bg-green-500 hover:text-white" @click="callOpenConversation('BJA24KS8MJ2')">BJA24KS8MJ2</li>
          <li class="px-4 py-2 mb-4 text-black bg-gray-200 cursor-pointer last:mb-0 hover:bg-green-500 hover:text-white" @click="callOpenConversation('BJA24KS8MJ3')">BJA24KS8MJ3</li>
          <li class="px-4 py-2 mb-4 text-black bg-gray-200 cursor-pointer last:mb-0 hover:bg-green-500 hover:text-white" @click="callOpenConversation('BJA24KS8MJ4')">BJA24KS8MJ4</li>
          <li class="px-4 py-2 mb-4 text-black bg-gray-200 cursor-pointer last:mb-0 hover:bg-green-500 hover:text-white" @click="callOpenConversation('BJA24KS8MJ5')">BJA24KS8MJ5</li>
        </ul>
      </div>

      <div id="talkjs" class="flex-1 h-[60vh]">
        <i>Loading chat...</i>
      </div>

    </div>
  </main>
</template>
