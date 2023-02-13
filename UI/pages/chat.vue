<template>
  <section class="chat">
    <div class="chatbox">
      <h2>Welcome <span>{{ character }}</span></h2>
      <div class="status-bar">
        <p>{{ status }}</p>
      </div>
      <div class="messages">
        <div class="message" v-for="message in messages" :key="message.length" :class="message.ownMessage ? 'message-own': 'message-other'">
          <img :src="`../assets/${message.user}.png`">
          <div class="message-content">
            <p class="message-username">{{ message.user }}</p>
            <p class="message-payload">{{ message.payload }}</p>
            <p class="message-timestamp"> {{ message.Timestamp }}</p>
          </div>
        </div>
      </div>
      <div class="input-bar">
        <input v-model="message" @keypress.enter="sendMessageToServer()">
        <button @click="sendMessageToServer()">SEND</button>
      </div>
    </div>
    <div class="sidebar"></div>
  </section>
</template>

<script>
var connection
import { useMainStore } from '~~/stores/mainStore'

export default {

  setup() {
    const mainStore = useMainStore()
    return { mainStore }
  },
  data() {
    return {
      message: null
    }
  },
  computed: {
    character() {
      return this.mainStore.getCharacter
    },
    messages() {
      return this.mainStore.getMessages
    },
    status() {
      return this.mainStore.getStatus
    },

  },
  methods: {
    sendMessageToServer() {
      if (this.message.length > 1) {
        this.mainStore.sendMessage("send_message", this.message)
      }
      this.message = ""
    }
  },
  mounted() {
    this.mainStore.connect()
  }
}
</script>

<style scoped>
.chat {
  display: grid;
  grid-template-areas:
    "sidebar chatbox";
  grid-template-columns: 29% 69%;
  margin: 2vh 0;
  min-height: 90%;
}

.chatbox {
  grid-area: chatbox;
  display: flex;
  flex-direction: column;
  border-radius: 10px;
  box-shadow: 5px 5px 5px 5px #000000;
  min-height: 100%;
}




.chatbox h2 {
  font-family: 'Nerko One', cursive;
  font-size: 2.5em !important;
  background: rgb(255, 149, 181);
  padding-left: 3%;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  height: 5%
}

.chatbox h2 span {
  color: white;
}

.status-bar {
  background: rgb(250, 188, 188);
  padding: 0 20px;
  height: 2%
}

.messages {
  background: rgb(255, 223, 221);
  height: 86%;
  display: flex;
  flex-direction: column;
}

.input-bar {
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px;
  background: rgb(250, 188, 188);
  height: 8%;
  display: flex;
  align-items: center;
}

.input-bar input {
  width: 80%;
  margin: 0 5%;
  height: 70%;
  border-radius: 10px;
  font-family: 'Nerko One', cursive;
  padding: 10px;
  font-size: 1.4em;
  box-shadow: 3px 3px 3px 3px #000000;
}

.input-bar button {
  background: rgb(255, 149, 181);
  padding: 10px;
  border-radius: 5px;
  box-shadow: 3px 3px 3px 3px #000000;
}

.input-bar button:active {
  transform: scale(0.9);
}

.message {
  margin: 1%;
  border: solid black 1px;
  padding: 5px;
  box-shadow: 3px 3px 3px 3px #000000;
  font-family: 'Nerko One', cursive;
  display: flex;
  width: 93%;
}

.message-other{
  background: rgb(255, 149, 181);
  border-bottom-right-radius: 10px;
  border-top-right-radius: 10px;
  border-top-left-radius: 10px;
}

.message-own{
  background: rgb(144, 190, 241);
  border-bottom-left-radius: 10px;
  border-top-right-radius: 10px;
  border-top-left-radius: 10px;
  align-self: flex-end;
}

.message img {
  width: 70px;
  background: white;
  border-radius: 70%;
  height: 70px;
  box-shadow: 0px 0px 10px 1px #f1862e;
}

.message-content{
  display:flex;
  flex-direction: column;
  margin-left: 1%;
  
}


.message-payload{
  color: black;
  font-size: 20px;
  border-top:1px solid pink;
  border-bottom:1px solid pink;
}

.message-timestamp{
  color:grey;
}
</style>