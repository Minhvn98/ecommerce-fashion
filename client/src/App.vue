<template>
  <the-header
    v-if="['register', 'login'].indexOf($route.name) === -1"
  ></the-header>
  <!-- <shopping-cart></shopping-cart> -->
  <router-view></router-view>
  <the-footer></the-footer>
</template>

<script>
// import HomePage from "./pages/HomePage.vue";
import TheHeader from "./components/TheHeader.vue";
import TheFooter from "./components/TheFooter.vue";

// import DetailProduct from "./pages/DetailProduct.vue";

export default {
  name: "App",
  components: {
    TheHeader,
    TheFooter,
    // HomePage,
    // DetailProduct,
    // AllProducts,
  },
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
