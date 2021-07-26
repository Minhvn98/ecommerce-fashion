<template>
  <div class="wrapper-404">
    <h1 class="title">404</h1>
    <div class="description">Trang Không Tồn Tại!</div>
    <router-link :to="{ name: 'home-page' }" class="back-to-home"
      >Trở lại trang chủ</router-link
    >
  </div>
</template>

<script>
export default {
  name: "404Page",

  created() {
    document.title = "Không tìm thấy trang";
  },
};
</script>

<style scoped>
.description {
  font-size: 25px;
}

.wrapper-404 {
  width: 400px;
  margin: auto;
  margin-top: 115px !important;
  text-align: center;
  padding: 30px;
  border-radius: 3px;
}

.title {
  color: #ee4d2d;
  font-size: 115px;
}

.back-to-home {
  color: #ee4d2d;
  text-decoration: none;
  font-size: 18px;
}
</style>
