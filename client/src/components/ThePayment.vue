<template>
  <div class="payment-wrapper">
    <h1 class="title">Thông tin thanh toán</h1>
    <form action="#" class="form-payment">
      <input class="info" type="text" placeholder="Vui lòng nhập họ tên" />
      <input
        class="info"
        type="text"
        placeholder="Vui lòng nhập số điện thoại"
      />
      <input class="info" type="email" placeholder="Vui lòng nhập email" />
      <input class="info" type="text" placeholder="Vui lòng nhập địa chỉ" />

      <div class="payment-method">
        <div class="method">Phương thức thanh toán</div>
        <div class="method-group">
          <div class="method-item">Thanh toán khi nhận hàng</div>
          <div class="method-item">Thanh toán paypal</div>
        </div>
      </div>

      <button type="submit" class="btn-payment">Thanh Toán</button>
    </form>
  </div>
</template>

<script>
export default {
  name: "ThePayment",
};
</script>

<style scoped>
.method {
  margin-top: 10px;
}

.btn-payment {
  width: 100%;
  padding: 13px;
  font-size: 18px;
  background: #ee4d2d;
  border-radius: 3px;
  border: transparent;
  color: #fff;
  cursor: pointer;
  margin: 25px 0 15px 0;
}

.btn-payment:hover {
  background: #dc310f;
}

.method-item:hover {
  border-color: #ee4d2d;
  color: #ee4d2d;
}

.method-item {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 3px;
}

.method-group {
  display: flex;
  justify-content: space-between;
  margin: 20px 0 10px 0;
}

.info {
  width: 100%;
  padding: 15px 20px;
  font-size: 16px;
  margin: 10px 0;
  border-radius: 3px;
  border: 1px solid #ddd;
}

.title {
  font-size: 24px;
  font-weight: 400;
}

.container {
  width: 80%;
  margin: auto;
}

.payment-wrapper {
  width: 450px;
  background: #fff;
  padding: 20px;
  border-radius: 3px;
  position: fixed;
  top: 130px;
  left: calc(50% - 225px);
  z-index: 99;
}
</style>
