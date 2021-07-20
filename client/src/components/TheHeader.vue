<template>
  <header class="menu-top">
    <div class="navbar">
      <div class="logo">
        <router-link :to="{ name: 'home-page' }">
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
          @change="onSearchHandler"
        />
        <button class="btn-search" @click="onSearchHandler">
          <i class="fas fa-search"></i>
        </button>
      </div>
      <div class="cart-user">
        <div class="navbar-cart">
          <router-link :to="{ name: 'shopping-cart-page' }"
            ><span v-if="cartNumber" class="cart-number">{{ cartNumber }}</span
            ><i class="fas fa-cart-plus"></i
          ></router-link>
        </div>
        <div v-if="isLogin" class="cart-login">
          <i class="far fa-user"></i>
          <div class="wrap-user">
            <p>Xin chào, {{ $store.state.user.username }}</p>
            <p>
              <router-link
                :to="{ name: 'bill', query: { status: 'all', tabSelected: 0 } }"
              >
                Đơn mua
              </router-link>
            </p>
            <p @click="handleLogout">Đăng xuất</p>
          </div>
        </div>
        <div v-if="!isLogin" class="cart-login">
          <router-link :to="{ name: 'login' }">
            <span class="login-text">Đăng nhập</span>
          </router-link>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import { logout } from "../services/users.service.js";
export default {
  name: "TheHeader",
  data() {
    return {
      textSearch: ""
    };
  },
  computed: {
    cartNumber() {
      return this.$store.getters.cartNumber;
    },
    isLogin() {
      return Boolean(this.$store.state.user);
    }
  },
  methods: {
    onSearchHandler() {
      document.title = this.textSearch + " | Tìm kiếm";

      this.$router.push({
        name: "search",
        query: { text: this.textSearch, limit: 100, offset: 0 }
      });

      this.$store.dispatch("searchProducts", { textSearch: this.textSearch });
    },
    handleLogout() {
      logout().then(() => {
        this.$store.dispatch("setUser", null);
        this.$store.commit("setCart", []);
      });
    }
  },

  async created() {
    // this.$store.dispatch("getCartProduct");
  }
};
</script>

<style scoped>
/* menu-top  */
.logo {
  width: 150px;
}
.logo img {
  width: 150px;
}
.navbar-cart {
  margin-right: 10px;
}
.cart-login {
  position: relative;
}
.cart-login:hover .wrap-user {
  opacity: 1;
  visibility: visible;
}
.wrap-user {
  position: absolute;
  width: 200px;
  background: white;
  left: -90px;
  top: 45px;
  transition: all 0.3s ease-in;
  opacity: 0;
  visibility: hidden;
}
.wrap-user::before {
  content: "";
  width: 20px;
  height: 20px;
  background: white;
  position: absolute;
  left: 50%;
  transform: translate(-50%, -50%) rotateZ(45deg);
}
.wrap-user > p {
  padding: 10px;
  cursor: pointer;
}
.wrap-user > p:hover {
  color: tomato;
}
.wrap-user a {
  color: black;
}
.wrap-user a:hover {
  color: tomato;
}
a {
  text-decoration: none;
}
.login-text {
  color: white;
  position: relative;
  top: 3px;
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
  width: calc(80% - 10px);
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
  cursor: pointer;
}
.box-search {
  width: 60%;
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
  width: auto;
  display: flex;
  justify-content: space-between;
  position: relative;
}

.cart-number[data-v-9a9f6144] {
  background: #fff;
  padding: 0px 10px;
  border-radius: 12px;
  color: #ff5722;
  position: absolute;
  top: -15px;
  left: 12px;
  border: 2px solid #f53e04;
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
