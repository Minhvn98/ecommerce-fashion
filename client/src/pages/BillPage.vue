<template>
  <div class="container1">
    <div class="bill-wrapper">
      <h1 class="title">Các đơn mua</h1>
      <div class="table">
        <div class="nav">
          <router-link
            :to="{ name: 'bill', query: { status: 'all', tabSelected: 0 } }"
          >
            <button
              @click="handleSelectTab(0)"
              :class="{ 'tab-active': tabSelected == 0 }"
            >
              Tất cả
            </button>
          </router-link>
          <router-link
            :to="{
              name: 'bill',
              query: { status: 'Chờ thanh toán', tabSelected: 1 },
            }"
          >
            <button
              @click="handleSelectTab(1)"
              :class="{ 'tab-active': tabSelected == 1 }"
            >
              Chờ thanh toán
            </button>
          </router-link>
          <router-link
            :to="{
              name: 'bill',
              query: { status: 'Đã thanh toán', tabSelected: 2 },
            }"
          >
            <button
              @click="handleSelectTab(2)"
              :class="{ 'tab-active': tabSelected == 2 }"
            >
              Đã thanh toán
            </button>
          </router-link>
          <router-link
            :to="{
              name: 'bill',
              query: { status: 'Thanh toán thất bại', tabSelected: 3 },
            }"
          >
            <button
              @click="handleSelectTab(3)"
              :class="{ 'tab-active': tabSelected == 3 }"
            >
              Bị hủy
            </button>
          </router-link>
        </div>
        <div class="row-wrap" v-for="order in orders" :key="order.orderId">
          <router-link
            :to="{
              name: 'bill-detail',
              params: { id: order.orderId },
            }"
          >
            <div class="row-item">
              <div
                class="item"
                v-for="product in order.products"
                :key="product.id"
              >
                <div class="image">
                  <img
                    :src="`${BASE_URL_IMAGE}${product.product_images[0].uri}`"
                    alt="image"
                  />
                </div>
                <div class="product-name">{{ product.name }}</div>
              </div>
            </div>
            <div class="info">
              <p>{{ order.created_at }}</p>
              <p class="text-blue">{{ order.status }}</p>
            </div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getOrdersByStatus } from "../services/order.service.js";

const BASE_URL_IMAGE = process.env.VUE_APP_BASE_URL_IMAGE;

export default {
  name: "BillPage",
  data() {
    return {
      BASE_URL_IMAGE,
      tabSelected: 0,
      orders: null,
    };
  },
  methods: {
    handleSelectTab(index) {
      this.tabSelected = index;
    },
  },
  beforeRouteUpdate(to) {
    getOrdersByStatus(to.query.status).then((res) => {
      this.orders = res.data;
      this.tabSelected = to.query.tabSelected;
    });
  },
  created() {
    document.title = "Danh sách hóa đơn";
    this.tabSelected = this.$route.query.tabSelected;
    getOrdersByStatus(this.$route.query.status).then((res) => {
      this.orders = res.data;
    });
  },
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  background: #e9e9e9;
}
a {
  text-decoration: none;
  color: black;
}
.container1 {
  width: 80%;
  margin: auto;
}

.bill-wrapper {
  margin-top: 115px;
  width: 100%;
}

.title {
  font-size: 23px;
  font-weight: 400;
  color: #ee4d2d;
}

.table {
  /* display: flex; */
  margin-top: 30px;
  background: #fff;
  border-radius: 3px;
}

button {
  padding: 20px 0;
  width: 200px;
  background-color: white;
  border: none;
  cursor: pointer;
  font-size: 16px;
}
button:hover {
  color: tomato;
}
.tab-active {
  color: tomato;
  border-bottom: 1px solid tomato;
}
.row-wrap {
  padding: 20px 10px;
  border-bottom: 1px solid rgb(182, 182, 182);
  transition: all 0.2s ease-in;
}
.row-wrap:hover {
  background-color: rgb(241, 241, 241);
}
.row-wrap > a {
  display: flex;
}
.image {
  width: 100px;
  height: 100px;
}
.image > img {
  width: 100px;
  height: 100px;
}
.row-item {
  display: flex;
  flex: 70%;
  flex-wrap: wrap;
}
.item {
  display: flex;
}
.item:not(.item:last-child) {
  padding-bottom: 10px;
}
.info {
  flex: 30%;
  display: flex;
  justify-content: space-between;
  flex-direction: column;
  align-items: flex-end;
}
.info > p {
  display: inline-block;
}
.product-name {
  padding-left: 20px;
}
.text-blue {
  color: blue;
}
</style>
