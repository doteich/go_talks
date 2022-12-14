<template>
  <section class="chat">
    <div class="chatbox">
      <h2>Welcome <span>{{ character }}</span></h2>
      <div class="status-bar">
        <p>Connected</p>
      </div>
      <div class="messages">

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
  computed: {
    character() {
      return this.mainStore.getCharacter
    }
  },


  methods: {
    connect() {
      connection = new WebSocket("ws://localhost:3001/ws")
      setTimeout(() => {
        connection.send("MOIN")
      }, 2000)
    },
  },
  mounted() {
    this.connect()
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
}

.chatbox {
  grid-area: chatbox;
  display: flex;
  flex-direction: column;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  box-shadow: 5px 5px 5px 5px #000000;

}

.chatbox h2 {
  font-family: 'Nerko One', cursive;
  font-size: 2.5em;
  background: rgb(255, 149, 181);
  padding-left: 3%;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  
}
.chatbox h2 span {
  color: white;
}

.status-bar{
  background: rgb(250, 188, 188);
  padding: 0 20px;
}

.messages{
  background: rgb(255, 223, 221); 
  height: 80vh
}



</style>