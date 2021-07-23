<template>
  <div v-if="product.id" class="container1">
    <div class="breadcrumb">
      <router-link :to="{ name: 'home-page' }">Shopee</router-link>
      > <a href="#">{{ product.category.name }}</a>
    </div>

    <div class="product-wrapper">
      <div class="product-image-wrapper">
        <div class="primary-image">
          <img
            :src="`${SRC_IMAGE}${
              primaryImage || product.product_images[0].uri
            }`"
            :alt="product.name"
            class="product-image"
          />
        </div>
        <div class="all-image">
          <div
            v-for="image in product.product_images"
            class="item-image"
            :key="image.id"
          >
            <img
              :src="`${SRC_IMAGE}${image.uri}`"
              :alt="product.name"
              class="product-image"
              @mouseover="onChangeImage"
            />
          </div>
        </div>
      </div>

      <div class="product-info">
        <div class="info-small">
          <h3 class="product-name">{{ product.name }}</h3>
          <div class="price">
            <div v-if="product.sale_percent > 0">
              <span class="price-origin"
                ><del>{{ priceOrigin }}</del></span
              >
              <span class="price-sale"
                >{{ priceSale }}
                <span class="sale-detail"
                  >GIẢM {{ product.sale_percent }}%</span
                >
              </span>
            </div>

            <div v-else>
              <span class="price-sale">{{ priceOrigin }} </span>
            </div>
          </div>

          <div class="quantity">
            <button class="btn1 btn-down" @click="decreaseProduct">-</button>
            <input
              type="number"
              name="num-product"
              id="num-product"
              v-model="quantitySelected"
              disabled
            />
            <button class="btn1 btn-up" @click="increaseProduct">+</button>
            <span class="quantity-text"
              >{{ product.quantity }} sản phẩm có sẵn</span
            >
          </div>
          <div v-if="errText" class="err-text">{{ errText }}</div>
          <button class="btn-add-to-cart" @click="addToCart">
            <i class="fas fa-cart-plus"></i> Thêm vào giỏ hàng
          </button>
        </div>

        <div class="product-description">
          <strong>Mô tả sản phẩm: </strong> <br />
          {{ product.description }}
        </div>
      </div>
    </div>
  </div>

  <section
    v-if="products.length > 0"
    class="container1 products-recommend-wrapper"
  >
    <h3 class="title1">Có thể bạn thích</h3>
    <div class="product-recommnend">
      <product
        v-for="product in products"
        :product="product"
        :key="product.id"
      ></product>
    </div>
  </section>

  <teleport to="body">
    <div v-if="isMessage" class="message">
      <div class="meesage-check"><i class="fas fa-check"></i></div>
      <div>Sản phẩm đã được thêm vào Giỏ hàng</div>
    </div>
  </teleport>
</template>

<script>
const SRC_IMAGE = process.env.VUE_APP_BASE_URL_IMAGE;

import Product from "../components/Product.vue";

import {
  getProductsByCategory,
  getProductById,
} from "../services/products.service";

import { addToCart as apiAddToCart } from "../services/cart.service";

import { getCategoryIdAndProductId } from "../utils/slug.util";
import { formatMoney } from "../utils/money.util";

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
      quantitySelected: 1,
      SRC_IMAGE,
      errText: "",
      primaryImage: "",
      isMessage: false,
    };
  },

  computed: {
    priceOrigin() {
      return formatMoney(this.product.price);
    },

    priceSale() {
      const price =
        this.product.price -
        this.product.price * (this.product.sale_percent / 100);
      return formatMoney(price);
    },
  },

  methods: {
    async addToCart() {
      await apiAddToCart(this.product.id, this.quantitySelected);
      const idProduct = this.product.id;
      const quantity = this.quantitySelected;

      this.$store.dispatch("addProductToCart", { idProduct, quantity });

      this.isMessage = true;
      setTimeout(() => {
        this.isMessage = false;
      }, 1000);
    },

    async fetchData() {
      if (!this.$route.params.slug) return; // tranh error khi thoat khoi trang product detail

      const [idCategory, idProduct] = getCategoryIdAndProductId(
        this.$route.params.slug
      );

      const { data: product } = await getProductById(idProduct);
      const { data: products } = await getProductsByCategory(idCategory, 4, 4);

      this.product = product;
      this.products = products;

      this.quantitySelected = 1;
      this.primaryImage = "";
      document.title = this.product.name;
    },

    decreaseProduct() {
      if (this.quantitySelected <= 1) {
        return (this.quantitySelected = 1);
      }
      this.errText = "";
      this.quantitySelected -= 1;
    },

    increaseProduct() {
      if (this.quantitySelected >= this.product.quantity) {
        this.errText = "Bạn đã chọn hết số lượng hàng trong kho";
        return (this.quantitySelected = this.product.quantity);
      }
      this.quantitySelected += 1;
    },

    onChangeImage(e) {
      const src = e.target.src.split("/").slice(-1)[0];
      this.primaryImage = src;
      console.log(src);
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
.message {
  width: 400px;
  height: 200px;
  background: rgba(0, 0, 0, 0.65);
  position: fixed;
  top: 35%;
  left: 50%;
  color: #fff;
  font-size: 20px;
  padding: 20px;
  border-radius: 3px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  transform: translateX(-50%);
}
.meesage-check {
  margin-bottom: 30px;
  width: 60px;
  height: 60px;
  background: #26aa99;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 50%;
}
.product-image {
  max-width: 100%;
  border-radius: 3px;
}

.item-image {
  width: calc(25%);
  margin-right: 10px;
  cursor: pointer;
}

.item-image:last-child {
  margin-right: 0;
}

.all-image {
  margin-top: 10px;
  display: flex;
  width: 100%;
}
.err-text {
  margin: 15px 0;
  color: #e61c1c;
}

.product-description {
  margin-top: 30px;
  font-size: 18px;
  color: #333;
}

.sale-detail {
  margin-left: 15px;
  background: #ee4d2d;
  color: #fff;
  font-size: 15px;
  padding: 2px;
  border-radius: 3px;
}
.price {
  padding: 20px 10px;
  margin: 20px 0;
  font-size: 20px;
  color: #ee4d2d;
  background: #f5f5f5;
  border-radius: 3px;
}

#num-product {
  font-size: 18px;
}
.router-link {
  text-decoration: none;
  color: #fff;
}

.price-origin {
  color: #8c8787;
  margin-right: 15px;
}

.btn-add-to-cart {
  margin-top: 25px;
  padding: 12px 30px;
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
  height: 36px;
  text-align: center;
}

.btn1 {
  width: 70px;
  font-size: 30px;
  border: 1px solid #e6e6e6;
  background: #fff;
  cursor: pointer;
  height: 36px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.quantity {
  display: flex;
  align-items: center;
}

.quantity-text {
  margin-left: 15px;
  color: #5a5959;
}

.product-name {
  font-size: 25px;
  font-weight: 300;
  color: #333;
}

.product-wrapper {
  display: flex;
  background: #fdfdfd;
  padding: 15px;
  border-radius: 3px;
}

.product-info {
  width: 60%;
  padding-left: 15px;
}

.product-image {
  width: 100%;
}

.product-image-wrapper {
  width: calc(40% - 15px);
  margin-right: 15px;
}

.breadcrumb {
  margin-top: 115px;
  padding: 20px 10px;
  background: #eae7e7;
  border-radius: 3px;
}

.breadcrumb a {
  text-decoration: none;
  color: rgb(4 67 214);
  margin: 0 10px;
}

.container1 {
  width: 80%;
  margin: auto;
}

.products-recommend-wrapper {
  width: calc(80% + 18px);
  margin-bottom: 10px;
}

.title1 {
  font-size: 22px;
  font-weight: 400;
  margin: 25px 0 15px 10px;
}

:root {
  --primary-color: #ee4d2d;
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
  .product-wrapper {
    flex-direction: column;
  }

  .product-image-wrapper {
    width: 100%;
  }

  .product-info {
    width: 100%;
  }

  .container1 {
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

