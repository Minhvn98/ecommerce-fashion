<template>
  <the-header
    v-if="['register', 'login'].indexOf($route.name) === -1"
  ></the-header>
  <router-view></router-view>
  <the-footer></the-footer>
</template>

<script>
import TheHeader from "../components/TheHeader.vue";
import TheFooter from "../components/TheFooter.vue";
export default {
  name: "LayoutShop",
  components: {
    TheHeader,
    TheFooter,
  },
};
</script>
