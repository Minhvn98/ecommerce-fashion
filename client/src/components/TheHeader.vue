<template>
  <header class="menu-top">
    <div class="navbar">
      <div class="logo">
        <router-link :to="{ name: 'home-router' }">
          <img src="../assets/logo.jpg" alt="logo" />
        </router-link>
      </div>
      <div class="box-search">
        <input
          type="text"
          name="search"
          id="search"
          placeholder="Nhập sản phẩm bạn cần tìm"
          v-model.trim="textSearch"
          @input="onEmitTextSearch"
        />
        <button class="btn-search" @click="onEmitTextSearch">
          <i class="fas fa-search"></i>
        </button>
      </div>
      <div class="cart-user">
        <div class="navbar-cart">
          <router-link :to="{ name: 'shopping-cart-router' }"
            ><i class="fas fa-cart-plus"></i
          ></router-link>
        </div>
        <div class="cart-login">
          <router-link :to="{ name: 'login' }">
            <i class="far fa-user"></i>
          </router-link>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: "TheHeader",

  data() {
    return {
      textSearch: "",
    };
  },

  methods: {
    onEmitTextSearch() {
      console.log(this.textSearch);
      this.$emit("on-search", this.textSearch);
    },
  },
};
</script>


<style  scoped>
/* menu-top  */
.logo {
  width: 150px;
}

.logo img {
  width: 150px;
}

.menu-top {
  background: linear-gradient(-180deg, #ec4c3e, #fa6735);
  width: 100%;
  height: 100px;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 9;
}

.navbar {
  width: calc(80% - 30px);
  margin: auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100px;
}

.btn-search {
  color: #fff;
  font-size: 23px;
  background: #ee4d2d;
  border: none;
  border-radius: 4px;
  padding: 6px 18px;
  margin-left: -63px;
}

.box-search {
  width: 70%;
  display: flex;
  height: 65px;
  align-items: center;
}

#search {
  padding-left: 15px;
  width: 100%;
  font-size: 1.1rem;
  outline: none;
  border-radius: 4px;
  border: none;
  height: 45px;
}

.cart-user i {
  font-size: 25px;
  cursor: pointer;
  color: #fff;
}

.navbar-cart {
  width: 40px;
}

.cart-user {
  margin-left: 15px;
  width: 80px;
  display: flex;
  justify-content: space-between;
}
@media screen and (max-width: 768px) {
  .navbar {
    width: calc(100% - 30px);
  }
}

@media screen and (max-width: 576px) {
  .logo {
    width: 50px;
    overflow: hidden;
    margin-right: 15px;
  }

  #search {
    font-size: 14px;
    padding: 10px;
  }
}
</style>
