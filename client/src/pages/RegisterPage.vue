<template>
  <div class="login-wrapper">
    <div class="background-login">
      <div class="form-wrapper">
        <h1 class="login-title">Đăng Ký</h1>
        <form @submit.prevent="onSubmitRegister" class="login-form">
          <input
            type="email"
            class="form-input"
            placeholder="Vui lòng nhập email"
            v-model="inputEmail"
          />
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
          <input
            type="password"
            class="form-input"
            placeholder="Vui lòng xác nhận lại password"
            v-model="inputConfirmPassword"
          />
          <div class="message-error">
            {{ errMessage }}
          </div>
          <button class="btn-submit" type="submit">ĐĂNG KÝ</button>

          <div class="text-register">
            Bạn đã có tài khoản shopee?
            <router-link :to="{ name: 'login' }"> Đăng nhập</router-link>
            <div>
              <router-link :to="{ name: 'home-page' }"
                >Quay lại trang chủ</router-link
              >
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { registerHandler, loginHandler } from "../services/users.service";

export default {
  name: "RegisterPage",

  data() {
    return {
      inputEmail: "",
      inputUserName: "",
      inputPassword: "",
      inputConfirmPassword: "",

      errMessage: "",
    };
  },

  methods: {
    async onSubmitRegister() {
      if (this.inputPassword !== this.inputConfirmPassword) {
        return (this.errMessage = "Xác nhận mật khẩu không chính xác!");
      }
      this.errMessage = "";

      try {
        const { status } = await registerHandler(
          this.inputEmail,
          this.inputUserName,
          this.inputPassword,
          this.inputConfirmPassword
        );

        if (status === 200) {
          const { data: user } = await loginHandler(
            this.inputUserName,
            this.inputPassword
          );

          this.$store.dispatch("setUser", user);
          this.$router.push({ name: "home-page" });
        }
      } catch (error) {
        const { status } = error.response;

        if (status >= 400)
          return (this.errMessage = "Đăng ký thất bại, vui lòng thử lại!");
      }
    },
  },

  created() {
    document.title = "Đăng ký";
  },
};
</script>

<style scoped>
.message-error {
  margin-bottom: 15px;
  color: red;
}
.login-wrapper {
  background: #ee4d2d;
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
  /* margin-top: 30px; */
  cursor: pointer;
}

.form-input {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border-radius: 3px;
  border: 1px solid #bbb;
  margin-bottom: 25px;
}

.login-title {
  font-weight: 500;
  font-size: 23px;
  margin-bottom: 35px;
  margin-top: 10px;
}

.form-wrapper {
  background: #fefefe;
  height: 495px;
  width: 400px;
  border-radius: 3px;
  padding: 15px;
  margin-bottom: 30px;
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
