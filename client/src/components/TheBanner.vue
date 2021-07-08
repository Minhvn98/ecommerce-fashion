<template>
  <section class="container banner-wrap">
    <div class="banner">
      <div class="banner-left">
        <div class="banner-img-left"></div>
        <!-- <img class="banner-img-left" src="./banner.png" alt="" /> -->
      </div>
      <div class="banner-right">
        <div class="banner-right-top">
          <div class="banner-img-top"></div>
        </div>
        <div class="banner-right-bottom">
          <div class="banner-img-bottom"></div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "TheBanner",
};
</script>

<style scoped>
.container {
  width: 80%;
  margin: auto;
  margin-top: 15px;
}

.banner {
  display: flex;
  justify-content: center;
}

.banner-wrap {
  margin-top: 115px !important;
}

.banner-bottom {
  width: 100%;
}

.banner-top {
  width: 100%;
}

.banner-img-bottom {
  background-image: url("../assets/banner-bottom.png");
  background-size: cover;
  min-height: 120px;
  border-radius: 3px;
}

.banner-img-top {
  min-height: 125px;
  background-image: url("../assets/banner-top.png");
  background-size: cover;
  margin-bottom: 15px;
  border-radius: 3px;
}

.banner-left {
  width: calc(66.666% - 10px);
}

.banner-img-left {
  background-image: url("../assets/banner.png");
  background-size: cover;
  min-height: 260px;
  margin-right: 15px;
  border-radius: 3px;
}

.banner-right {
  width: calc(33.333% - 10px);
}

@media screen and (max-width: 768px) {
  .banner-right {
    display: none;
  }

  .banner-left {
    width: calc(100%);
    margin-left: 10px;
    margin-right: -5px;
  }

  .banner-img-left {
    min-height: 165px;
  }
}

@media screen and (max-width: 576px) {
  .container {
    width: 100%;
  }
}
</style>