<template>
  <div class="login-wrapper">
    <div class="form-wrapper">
      <h1 class="login-title">Đăng Nhập</h1>
      <form action="#" class="login-form">
        <input
          type="email"
          class="form-input"
          placeholder="Vui lòng nhập email"
        />
        <input
          type="password"
          class="form-input"
          placeholder="Vui lòng nhập password"
        />
        <button class="btn-submit" type="submit">ĐĂNG NHẬP</button>
        <div class="text-forgot">
          <a href="#">Quên mật khẩu</a>
        </div>
        <div class="text-register">
          Bạn mới biết đến Shopee? <a href="#">Đăng ký</a>
        </div>
      </form>
    </div>
  </div>
</template>


<script>
export default {
  name: "LoginPage",
};
</script>

<style scoped>
.text-register {
  color: #444;
  font-size: 15px;
}

.text-register a {
  text-decoration: none;
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
}

.text-forgot a {
  margin-top: 10px;
  margin-bottom: 10px;
  display: block;
  text-decoration: none;
  font-size: 15px;
}
.btn-submit:hover {
  background: #c3452d;
}
.login-wrapper {
  background-image: url("../assets/login-background.png");
  background-size: cover;
  height: 100%;
  min-height: 35.5rem;
  background-repeat: no-repeat;
  background-position: center center;
  display: flex;
  justify-content: flex-end;
  padding-right: 50px;
  padding-top: 35px;
}
.menu-top {
  display: none;
}
@media screen and (max-width: 992px) {
}

@media screen and (max-width: 768px) {
}

@media screen and (max-width: 576px) {
  .login-wrapper {
    padding: 35px 0;
    justify-content: center;
  }

  .form-wrapper {
    width: 95%;
  }
}
</style>
