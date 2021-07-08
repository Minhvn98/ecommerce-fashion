<template>
  <div class="container">
    <div class="breadcrumb">
      <a href="/">Shopee</a> > <a href="#">Đồ ngủ </a>
    </div>
    <div class="product-wrapper">
      <div class="product-image-wrapper">
        <img :src="product.image" :alt="product.name" class="product-image" />
      </div>
      <div class="product-info">
        <div class="info-small">
          <h3 class="product-name">{{ product.name }}</h3>
          <div class="price">{{ formatMoney }}</div>
          <div class="quantity">
            <button class="btn btn-down" @click="quantityProduct--">-</button>
            <input
              type="number"
              name="num-product"
              id="num-product"
              v-model="quantityProduct"
            />
            <button class="btn btn-up" @click="quantityProduct++">+</button>
          </div>
          <div class="product-description">
            {{ product.description }}
          </div>
        </div>
        <button class="btn-add-to-cart">
          <router-link
            :to="{ name: 'shopping-cart-router' }"
            class="router-link"
            ><i class="fas fa-cart-plus"></i> Thêm vào giỏ hàng</router-link
          >
        </button>
      </div>
    </div>
  </div>

  <section
    v-if="products.length > 0"
    class="container products-recommend-wrapper"
  >
    <h3>Có thể bạn thích</h3>
    <div class="product-recommnend">
      <product
        v-for="product in products"
        :product="product"
        :key="product.id"
      ></product>
    </div>
  </section>
</template>

<script>
import Product from "../components/Product.vue";

import {
  fetchProductById,
  fetchProductByCategory,
} from "../services/products.service";

export default {
  name: "ProductDetail",
  components: {
    Product,
  },

  data() {
    return {
      product: {
        price: 0,
      },
      products: [],
      quantityProduct: 1,
    };
  },

  computed: {
    formatMoney() {
      return this.product.price.toLocaleString("it-IT", {
        style: "currency",
        currency: "VND",
      });
    },
  },

  methods: {
    async fetchData() {
      const idProduct = this.$route.params.id;
      const productsResponse = await fetchProductById(idProduct);

      this.product = productsResponse.data;

      const responseProductCategory = await fetchProductByCategory(
        this.product.category_id
      );

      this.products = responseProductCategory.data;
    },
  },

  created() {
    this.$watch(
      () => this.$route.params,
      () => {
        this.fetchData();
      },
      // fetch the data when the view is created and the data is
      // already being observed
      { immediate: true }
    );
  },
};
</script>

<style scoped>
.product-description {
  margin: 25px 0;
}
#num-product {
  font-size: 18px;
}
.router-link {
  text-decoration: none;
  color: #fff;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
:root {
  --shopee: #ee4d2d;
}

.btn-add-to-cart {
  margin-top: 25px;
  padding: 15px 30px;
  font-size: 17px;
  background: #ee4d2d;
  cursor: pointer;
  color: #fff;
  border: none;
  border-radius: 3px;
}

.btn-add-to-cart:hover {
  background: #e05e45;
}

#num-product {
  padding-left: 10px;
  width: 50px;
  border: 1px solid #eee;
}

.btn {
  width: 70px;
  font-size: 30px;
  border: 1px solid #eee;
  background: #fff;
}

.quantity {
  display: flex;
}

.price {
  padding: 30px 0;
  font-size: 20px;
  color: #ee4d2d;
}

.product-name {
  font-size: 25px;
  font-weight: 300;
}

.product-wrapper {
  display: flex;
  background: #fdfdfd;
  padding: 15px;
  border-radius: 3px;
}

.product-info {
  width: 60%;
  padding-left: 25px;
}

.product-image {
  width: 100%;
}

.product-image-wrapper {
  width: calc(30% - 15px);
  margin-right: 15px;
}

.breadcrumb {
  margin-top: 115px;
  padding: 20px 10px;
}

.breadcrumb a {
  text-decoration: none;
  color: #111;
}

body {
  background: #eee;
}

.container {
  width: 80%;
  margin: auto;
}

.products-recommend-wrapper {
  width: calc(80% + 18px);
  margin-bottom: 10px;
}

.products-recommend-wrapper h3 {
  font-size: 20px;
  font-weight: 400;
  margin: 15px 0 15px 10px;
}

:root {
  --primary-color: #ee4d2d;
}

/* menu-top  */
.logo {
  width: 150px;
}

.logo img {
  width: 150px;
}

.menu-top {
  background: linear-gradient(-180deg, #f53d2ded, #f05c2af0);
  width: 100%;
  height: 100px;
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

/* footer  */
.footer-wrap {
  width: 100%;
  background: #f5f5f5;
  color: #8c8787;
  padding-bottom: 30px;
}

.footer {
  display: flex;
  width: calc(80% - 30px);
  margin: auto;
}

.footer-item {
  width: 25%;
}

.footer-item li {
  list-style: none;
  padding: 10px 0;
  cursor: pointer;
}

.footer-item li:hover {
  color: var(--primary-color);
}

.footer-title {
  padding: 20px 0;
  font-size: 16px;
}

.footer-item i {
  color: #333;
  margin-right: 15px;
}

.footer-app {
  display: flex;
}

.qrcode {
  max-width: 60%;
}

.card-body {
  padding: 0px 15px 25px 15px;
  color: #ee4d2d;
  font-size: 1.1rem;
}

.card-title {
  padding: 15px;
  font-size: 1.1rem;
}

.card-image {
  width: 100%;
  border-radius: 5px 5px 0 0;
}

.card {
  background: #fff;
  width: calc(25% - 15px);
  border-radius: 5px;
  margin: 7.5px;
  border: 1px solid #eee;
  transition: 0.3s;
}

.card:hover {
  transform: translateY(-2px);
}

.product-recommnend {
  display: flex;
  flex-wrap: wrap;
  /* justify-content: center; */
}

/*
  --------Responsive------
  */

@media screen and (max-width: 992px) {
  .card {
    width: calc(33.333% - 15px);
  }
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

  .footer {
    width: calc(100% - 30px);
    flex-wrap: wrap;
  }

  .footer-item {
    width: 50%;
  }

  .product-wrapper {
    flex-direction: column;
  }

  .product-image-wrapper {
    width: 100%;
  }

  .product-info {
    width: 100%;
  }

  .container {
    width: calc(100% - 20px);
  }

  .product-info {
    padding-left: 0;
  }

  .quantity {
    width: 100%;
  }

  .btn-add-to-cart {
    width: 100%;
  }

  .product-name {
    padding-top: 20px;
    font-size: 22px;
  }

  .card {
    width: 100%;
  }
}
</style>
