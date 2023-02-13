<template>
  <section class="start-page">
    <h2> Select your Character and Start Chatting</h2>
    <div class="character-selection">
      <div class="char" @click="setRoute('Pusheen')">
        <img src="../assets/pusheen.png">
        <h3>Pusheen</h3>
      </div>
      <div class="char">
        <img src="../assets/tomcat.png" @click="setRoute('Tomcat')">
        <h3>Tomcat</h3>
      </div>

    </div>
  </section>
</template>

<script>
import { useMainStore } from '~~/stores/mainStore';
export default {
  setup() {
    const mainStore = useMainStore()
    return { mainStore }
  },
  methods: {
    setRoute(name) {
      if(process.client){
          localStorage.setItem("char", name);
      }
      this.mainStore.setChar(name)
      this.$router.push("chat")
    }
  }
}
</script>

<style scoped>
.start-page {
  display: flex;
  flex-direction: column;
  align-items: center;
}

h2 {
  font-family: 'Nerko One', cursive;
  margin: 1%;
  font-size: 2rem;
}

.character-selection {
  display: flex;

}

.char {
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 10px;
  box-shadow: 5px 4px 3px 5px #000000;
  margin: 2%
}

.char:hover {
  background: pink;
  cursor: pointer;
}

.char img {
  width: 20vmin;
}

.char h3 {
  font-family: 'Nerko One', cursive;
  font-size: 1.7rem;
}
</style>