<template>
  <component :is="componentSelected"></component>
</template>

<script>
import LayoutShop from "./layout/LayoutShop.vue";
import LayoutAdmin from "./layout/LayoutAdmin.vue";
// import { getUserByCookie } from "./services/users.service.js";

export default {
  name: "App",
  components: {
    LayoutShop,
    LayoutAdmin
  },
  computed: {
    componentSelected() {
      return this.$store.state.layout;
    }
  }
  // async created() {
  //   let res = await getUserByCookie();
  //   this.$store.dispatch("setUser", res.data);
  //   if (res.data.role === "customer") {
  //     this.$store.dispatch("getCartProduct");
  //     this.$store.commit("updateLayout", "shop");
  //   }
  //   if (res.data.role === "admin") {
  //     this.$store.commit("updateLayout", "admin");
  //   }
  // }
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
a {
  text-decoration: none;
}
</style>
