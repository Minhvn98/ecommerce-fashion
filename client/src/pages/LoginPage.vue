<template>
  <div class="login-wrapper">
    <div class="background-login">
      <div class="form-wrapper">
        <h1 class="login-title">Đăng Nhập</h1>
        <form @submit.prevent="onSubmitLogin" class="login-form">
          <input
            type="text"
            class="form-input"
            placeholder="Vui lòng nhập username"
            v-model="inputUserName"
          />
          <input
            type="password"
            class="form-input"
            placeholder="Vui lòng nhập password"
            v-model="inputPassword"
          />
          <div v-if="errMessage" class="error-message">{{ errMessage }}</div>
          <button class="btn-submit" type="submit">ĐĂNG NHẬP</button>
          <div class="text-forgot">
            <a href="#">Quên mật khẩu</a>
            <router-link :to="{ name: 'home-page' }"
              >Quay lại trang chủ</router-link
            >
          </div>
          <div class="text-register">
            Bạn mới biết đến Shopee?
            <router-link :to="{ name: 'register' }">Đăng ký</router-link>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>


<script>
import { loginHandler } from "../services/users.service";

export default {
  name: "LoginPage",

  data() {
    return {
      inputUserName: "",
      inputPassword: "",
      errMessage: "",
    };
  },

  methods: {
    async onSubmitLogin() {
      this.errMessage = "";
      if (!this.inputUserName || !this.inputPassword) {
        return (this.errMessage = "Vui lòng nhập đầy đủ thông tin!");
      }

      try {
        const { data: user } = await loginHandler(
          this.inputUserName,
          this.inputPassword
        );

        this.$store.dispatch("setUser", user);
        this.$router.push({ name: "home-page" });
      } catch (error) {
        const { status } = error.response;

        if (status === 400)
          return (this.errMessage = "Tài khoản, mật khẩu không chính xác!");
      }
    },
  },

  created() {
    document.title = "Đăng nhập";
  },
};
</script>

<style scoped>
.login-wrapper {
  background: #ee4d2d;
}

.error-message {
  color: #dc3413;
}

.text-forgot {
  display: flex;
  justify-content: space-between;
}

.text-register {
  color: #444;
  font-size: 15px;
  line-height: 25px;
}

.text-register a {
  text-decoration: none;
  color: #1328c5;
}

.btn-submit {
  width: 100%;
  padding: 10px;
  color: #fff;
  font-size: 16px;
  border-radius: 3px;
  border: transparent;
  background: #ed4c2d;
  margin-bottom: 20px;
  margin-top: 30px;
  cursor: pointer;
}

.form-input {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border-radius: 3px;
  border: 1px solid #bbb;
  margin-bottom: 35px;
}

.login-title {
  font-weight: 500;
  font-size: 23px;
  margin-bottom: 35px;
  margin-top: 10px;
}

.form-wrapper {
  background: #fefefe;
  height: 450px;
  width: 400px;
  border-radius: 3px;
  padding: 15px;
  margin-bottom: 30px;
}

.text-forgot a {
  margin-top: 10px;
  margin-bottom: 10px;
  display: block;
  text-decoration: none;
  font-size: 15px;
  color: #1328c5;
}
.btn-submit:hover {
  background: #c3452d;
}
.background-login {
  background-image: url("../assets/login-background.png");
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center center;
  display: flex;
  justify-content: flex-end;
  padding-right: 50px;
  padding-top: 35px;
  width: 80%;
  margin: auto;
}

@media screen and (max-width: 992px) {
}

@media screen and (max-width: 768px) {
  .background-login {
    padding: 35px 0;
    background-image: none !important;
    justify-content: center;
  }
}

@media screen and (max-width: 576px) {
  .background-login {
    padding: 35px 0;
    background-image: none !important;
    justify-content: center;
  }

  .form-wrapper {
    width: 95%;
  }
}
</style>
