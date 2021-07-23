<template>
  <TheHeader v-if="['register', 'login'].indexOf($route.name) === -1" />
  <router-view></router-view>
  <TheFooter />
</template>

<script>
import TheHeader from "../components/TheHeader.vue";
import TheFooter from "../components/TheFooter.vue";
export default {
  name: "LayoutShop",
  components: {
    TheHeader,
    TheFooter
  }
};
</script>
