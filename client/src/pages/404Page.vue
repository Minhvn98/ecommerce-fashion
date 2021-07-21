<template>
  <div class="page-404">
    <h2>404</h2>
  </div>
</template>

<script>
export default {
  name: "404Page",
  created() {
    document.title = "Không tìm được trang yêu cầu";
  }
};
</script>

<style scoped>
.page-404 {
  width: 100%;
  height: 600px;
  display: flex;
  align-items: center;
  justify-content: center;
}
h2 {
  font-size: 40px;
}
</style>
