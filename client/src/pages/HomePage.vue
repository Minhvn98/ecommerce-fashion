<template>
  <the-banner></the-banner>

  <section class="container category-wrap">
    <div class="category">
      <h3>Danh mục sản phẩm</h3>
      <div class="category-link-wrap">
        <router-link :to="{ name: 'products-router' }" class="category-link">
          Tất cả sản phẩm
        </router-link>
        <router-link
          v-for="category in categories"
          :to="{
            name: 'products-category',
            params: { category: category.name },
          }"
          class="category-link"
          :key="category.id"
        >
          {{ category.name }}
        </router-link>
      </div>
    </div>

    <div class="sort-products">
      <label for="sort-product">Sắp xếp</label>
      <select name="sort-product" id="sort-product">
        <option>Sắp xếp</option>
        <option value="">Giá tăng dần</option>
        <option value="">Giá giảm dần</option>
      </select>
    </div>
  </section>

  <section class="container">
    <h3 class="title">Sản phẩm mới</h3>
    <div class="row">
      <product
        v-for="product in products"
        :product="product"
        :key="product.id"
      ></product>
    </div>
  </section>
</template>

<script>
import TheBanner from "../components/TheBanner.vue";
import Product from "../components/Product.vue";

import { fetchProducts } from "../services/products.service";

import { fetchCategories } from "../services/categories.service";

export default {
  name: "HomePage",

  components: {
    TheBanner,
    Product,
  },

  data() {
    return {
      products: [],
      categories: [],
    };
  },

  async created() {
    document.title =
      "Shopee Việt Nam | Mua và Bán Trên Ứng Dụng Di Động Hoặc Website";

    const productsResponse = await fetchProducts();
    const categoriesResponse = await fetchCategories();

    this.products = productsResponse.data.slice(0, 4);
    this.categories = categoriesResponse.data;
  },
};
</script>

<style src="./HomePage.css" scoped></style>
