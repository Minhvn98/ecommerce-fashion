<template>
  <div class="payment-wrapper">
    <h1 class="title">Thông tin thanh toán</h1>
    <form class="form-payment">
      <input
        class="info"
        type="text"
        v-model="firstName"
        placeholder="Vui lòng nhập họ và đệm"
      />
      <input
        class="info"
        type="text"
        v-model="lastName"
        placeholder="Vui lòng nhập tên"
      />
      <input
        class="info"
        v-model="phone"
        type="text"
        placeholder="Vui lòng nhập số điện thoại"
      />
      <input
        class="info"
        type="email"
        v-model="email"
        placeholder="Vui lòng nhập email"
      />
      <input
        class="info"
        type="text"
        v-model="location"
        placeholder="Vui lòng nhập địa chỉ"
      />

      <div class="payment-method">
        <div class="method">Phương thức thanh toán</div>
        <div class="method-group">
          <div
            @click="selectPaymentMethod('Thanh toán khi nhận hàng')"
            class="method-item"
            :class="{
              'payment-method-active': payment === 'Thanh toán khi nhận hàng'
            }"
          >
            Thanh toán khi nhận hàng
          </div>
          <div
            @click="selectPaymentMethod('Thanh toán momo')"
            class="method-item"
            :class="{ 'payment-method-active': payment === 'Thanh toán momo' }"
          >
            Thanh toán momo
          </div>
        </div>
      </div>

      <button type="button" @click="checkout" class="btn-payment">
        Thanh Toán
      </button>
    </form>
  </div>
</template>

<script>
import { checkoutHandler } from "../services/checkout.service.js";
export default {
  name: "ThePayment",
  data() {
    return {
      firstName: "",
      lastName: "",
      email: "",
      phone: "",
      location: "",
      payment: ""
    };
  },
  methods: {
    selectPaymentMethod(payment) {
      this.payment = payment;
    },
    checkout() {
      checkoutHandler(
        this.firstName,
        this.lastName,
        this.email,
        this.phone,
        this.location,
        this.payment
      )
        .then(res => {
          if (res.data.paymentType === 0) {
            this.$router.push({
              name: "bill-detail",
              params: { id: res.data.orderId }
            });
          }
          if (res.data.paymentType === 1) {
            window.location.href = res.data.url;
          }
        })
        .catch(err => {
          console.log(err.body);
        });
    }
  }
};
</script>

<style scoped>
.method {
  margin-top: 10px;
}
.payment-method-active {
  color: tomato;
  border-color: tomato !important;
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
  cursor: pointer;
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
  top: 100px;
  left: calc(50% - 225px);
  z-index: 99;
}
</style>
