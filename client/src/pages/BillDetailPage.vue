<template>
  <div class="container">
    <div class="bill-wrapper">
      <h1 class="title">Chi Tiết Hóa Đơn</h1>
      <div class="content">
        <div class="info">
          <div class="info-customer">
            <ul>
              <li class="item">
                <strong>Khách Hàng:</strong> {{ order?.first_name }}
                {{ order?.last_name }}
              </li>
              <li class="item">
                <strong>Địa Chỉ:</strong> {{ order?.location }}
              </li>
              <li class="item"><strong>SĐT:</strong> {{ order?.phone }}</li>
              <li class="item"><strong>Email:</strong> {{ order?.email }}</li>
            </ul>
          </div>

          <div class="info-bill">
            <ul>
              <li class="item"><strong>Hóa Đơn:</strong> {{ order?.id }}</li>
              <li class="item">
                <strong>Ngày Thanh Toán:</strong> {{ order?.created_at }}
              </li>
              <li class="item">
                <strong>Phương thức Thanh Toán:</strong> {{ order?.payment }}
              </li>
              <li class="item">
                <strong>Trạng Thái: </strong>
                <b class="blue"> {{ order?.status }}</b>
              </li>
              <h1 class="total">
                Tổng tiền: {{ formatPrice(order?.total_price) }}
              </h1>
            </ul>
          </div>
        </div>

        <div class="bill-detail">
          <!--                <h3 class="title">Chi tiết</h3>-->
          <div
            class="product"
            v-for="product in products"
            :key="product.product_info.id"
          >
            <div class="product-image">
              <img
                :src="getLinkImage(product.product_info.product_images[0].uri)"
                alt="image"
              />
            </div>
            <div class="product-name">
              {{ product.product_info.name }}
            </div>
            <div class="product-price">
              {{
                formatPrice(
                  product.product_info.price * product.quantity_in_cart
                )
              }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getOrderById } from "../services/order.service.js";
import { formatMoney } from "../utils/money.util.js";
const BASE_URL_IMAGE = process.env.VUE_APP_BASE_URL_IMAGE;
export default {
  name: "BillDetailPage",
  data() {
    return {
      order: null,
      products: null
    };
  },
  methods: {
    formatPrice(price) {
      return formatMoney(price);
    },
    getLinkImage(image) {
      return `${BASE_URL_IMAGE}/${image}`;
    }
  },
  created() {
    getOrderById(this.$route.params.id).then(res => {
      this.order = res.data;
      this.products = res.data.product;
      console.log(res.data);
    });
  }
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.blue {
  color: blue;
}
.total {
  font-size: 22px;
  font-weight: 400;
  margin-top: 20px;
  border-top: 1px solid #ddd;
  padding: 12px 0;
  color: #ee4d2d;
}

strong {
  color: #ee4d2d;
}

.bill-detail {
  width: 70%;
  padding: 15px;
}

.info {
  width: 30%;
  padding: 15px;
  background: #f7f7f7;
}

.content {
  display: flex;
  margin-top: 30px;
  background: #fff;
  border-radius: 3px;
}

.item {
  list-style: none;
  padding: 5px 0;
}

body {
  background: #e9e9e9;
}

.product-price {
  width: 17%;
  color: #ee4d2d;
  display: flex;
  align-items: center;
}

.product-name {
  width: 70%;
}

.product-image img {
  width: 100%;
  border-radius: 3px;
}

.product-image {
  width: 12%;
  margin-right: 15px;
}

.product {
  display: flex;
  background: #fff;
  padding: 10px;
  border-radius: 3px;
  margin: 10px 0;
  border: 1px solid #ddd;
}

.title {
  font-size: 23px;
  font-weight: 400;
  color: #ee4d2d;
}

.container {
  width: 80%;
  margin: auto;
}

.bill-wrapper {
  margin-top: 115px;
  width: 100%;
}
</style>
