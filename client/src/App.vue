<template>
  <the-header
    v-if="['register', 'login'].indexOf($route.name) === -1"
  ></the-header>
  <router-view></router-view>
  <the-footer></the-footer>
</template>

<script>
import TheHeader from "./components/TheHeader.vue";
import TheFooter from "./components/TheFooter.vue";
import { getUserByCookie } from "./services/users.service.js";

export default {
  name: "App",
  components: {
    TheHeader,
    TheFooter
  },
  created() {
    getUserByCookie().then(res => {
      this.$store.dispatch("setUser", res.data);
      this.$store.dispatch("getCartProduct");
    });
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
:root {
  --primary-color: #ee4d2d;
}
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  background-color: #eee;
}
</style>
