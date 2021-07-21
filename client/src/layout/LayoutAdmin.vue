<template>
  <div id="wrapper">
    <Navbar />
    <div id="content-wrapper" class="d-flex flex-column">
      <div id="content">
        <HeaderAdmin />
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from "../admin/components/Navbar.vue";
import HeaderAdmin from "../admin/components/HeaderAdmin.vue";
export default {
  name: "LayoutAdmin",
  components: {
    Navbar,
    HeaderAdmin
  }
};
</script>
