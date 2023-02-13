<template>
  <main>
    <header>
      <a href="/">Pusheen <span>Chat</span></a>
      <div class="button-group">
        <button class="b-red" title="Leave">x</button>
        <button class="b-yellow">x</button>
        <button class="b-green" title="Visit Github">x</button>
      </div>
    </header>
    <NuxtPage></NuxtPage>
  </main>
</template>



<script>
import { useMainStore } from '~~/stores/mainStore'

export default {
  setup() {
    const mainStore = useMainStore()
    return { mainStore }
  },
  created() {
    let char
    if (process.client) {
       char = localStorage.getItem("char")
    }

    if (char) {
      this.mainStore.setChar(char)
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Nerko+One&family=Work+Sans&display=swap');



main {
  margin: 1vmin;
  background: rgb(245, 210, 210);
  height: 98vh;
  border-radius: 10px;
  box-shadow: 5px 5px 5px 5px #000000;
}

body {
  background: rgb(145, 145, 145);
}


header {
  display: flex;
  align-items: center;
  background: rgb(255, 149, 181);
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  padding: 5px;
  min-height: 60px;
  border-bottom: 2px solid black;
}

header a {
  font-size: 2.5rem;
  color: rgb(51, 51, 51);
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  margin-left: 5%
}

header a span {
  color: rgb(250, 250, 250)
}

.button-group {
  margin-left: auto;
  margin-right: 2%
}

.button-group button {
  width: 10px;
  height: 10px;
  margin: 20px 3px;
  border-radius: 60%;
  padding: 10% 0;
}

.b-green {
  background: rgb(47, 175, 47);
  color: rgba(0, 0, 0, 0)
}

.b-yellow {
  background: gold;
  color: rgba(0, 0, 0, 0)
}

.b-red {
  background: crimson;
  color: rgba(0, 0, 0, 0)
}
</style>
