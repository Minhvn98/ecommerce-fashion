<template>
  <div class="card">
    <router-link
      class="name-product"
      :to="{
        name: 'products-detail-page',
        params: { slug: generateSlugFromName },
      }"
    >
      <div class="thumbnail">
        <div v-if="product.sale_percent" class="sale-wrapper">
          <div class="logo-sale">
            <span>Giảm</span>
            <span class="logo-sale-number">{{ product.sale_percent }}%</span>
          </div>
          <img src="../assets/sale-bg.png" class="sale-bg" alt="" />
        </div>
        <img
          class="card-image"
          :src="`${BASE_URL_IMAGE}${product.product_images[0].uri}`"
          :alt="product.name"
        />
      </div>

      <div class="card-title">
        {{ product.name }}
      </div>
      <div class="card-body">
        <span>{{ priceSale }}</span>
      </div>
    </router-link>
  </div>
</template>

<script>
import { generateSlugs } from "../utils/slug.util";
import { formatMoney } from "../utils/money.util";

const BASE_URL_IMAGE = process.env.VUE_APP_BASE_URL_IMAGE;

export default {
  name: "Product",

  props: {
    product: Object,
  },

  data() {
    return {
      BASE_URL_IMAGE,
    };
  },
  methods: {},

  computed: {
    generateSlugFromName() {
      return generateSlugs(
        this.product.name,
        this.product.category.id,
        this.product.id
      );
    },

    priceSale() {
      const price =
        this.product.price -
        this.product.price * (this.product.sale_percent / 100);
      return formatMoney(price);
    },
  },
};
</script>

<style scoped>
.logo-sale-number {
  color: red;
}

.logo-sale {
  position: absolute;
  background: #fdd637e6;
  top: 1px;
  right: 1px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 3px 5px;
  color: #fff;
  border-radius: 2px;
  font-size: 13px;
  text-transform: uppercase;
}

.sale-bg {
  position: absolute;
  width: 90%;
  bottom: 0;
  left: 0;
}

.thumbnail {
  position: relative;
}

.name-product {
  text-decoration: none;
  color: #111;
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
  border-radius: 3px 3px 0 0;
}

.card {
  background: #fff;
  width: calc(25% - 15px);
  border-radius: 3px;
  margin: 7.5px;
  border: 1px solid #eee;
  transition: 0.3s;
  cursor: pointer;
}

.card:hover {
  transform: translateY(-2px);
  border: 1px solid #ee4d2d;
}

/*

  ------Responsive -------

 */

@media screen and (max-width: 992px) {
  .card {
    width: calc(33.333% - 15px);
  }
}

@media screen and (max-width: 768px) {
  .card {
    width: calc(50% - 15px);
  }
}

@media screen and (max-width: 576px) {
  .card {
    width: 100%;
  }
}
</style>
