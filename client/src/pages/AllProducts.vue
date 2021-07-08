<template>
  <section class="container category-wrap">
    <div class="category">
      <h3>Danh mục sản phẩm</h3>
      <select
        name="category"
        id="category"
        v-model="idCategory"
        @change="onFilterByCategory"
      >
        <option value="0">Tất cả sản phẩm</option>
        <option v-for="cate in categories" :value="cate.id" :key="cate.id">
          {{ cate.name }}
        </option>
      </select>
    </div>

    <div class="sort-products">
      <label for="sort-product">Sắp xếp</label>
      <select name="sort-product" id="sort-product" @change="onSortHandler">
        <option value="">Sắp xếp</option>
        <option value="asc">Giá tăng dần</option>
        <option value="desc">Giá giảm dần</option>
      </select>
    </div>
  </section>

  <div class="container">
    <div class="row">
      <product
        v-for="product in products"
        :product="product"
        :key="product.id"
      ></product>
    </div>
  </div>
</template>

<script>
import Product from "../components/Product.vue";

import {
  fetchProducts,
  fetchProductByCategory,
} from "../services/products.service";

import { fetchCategories } from "../services/categories.service";

export default {
  name: "AllProducts",
  components: {
    Product,
  },

  data() {
    return {
      idCategory: 0,
      categories: [],
      products: [],
    };
  },

  methods: {
    async onFilterByCategory() {
      const responseProduct = await fetchProductByCategory(this.idCategory);
      this.products = responseProduct.data;
    },

    onSortHandler(event) {
      const sortType = event.target.value;

      if (sortType === "asc") {
        this.products.sort((a, b) => a.price - b.price);
      }

      if (sortType === "desc") this.products.sort((a, b) => b.price - a.price);
    },
  },

  async created() {
    const productsResponse = await fetchProducts();
    const categoriesResponse = await fetchCategories();

    this.categories = categoriesResponse.data;
    this.products = productsResponse.data;
  },
};
</script>


<style scoped>
:root {
  --primary-color: #ee4d2d;
}

/*

----------Product -----------

*/
body {
  background-color: #eee;
}

.container {
  width: 80%;
  margin: auto;
  margin-top: 15px;
}

.row {
  display: flex;
  flex-wrap: wrap;
  /* justify-content: center; */
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

.price-sort {
  display: flex;
  align-items: center;
}

#sort-product {
  width: 150px;
  height: 39px;
  padding: 10px;
  border: transparent;
  background: #fff;
  border-radius: 3px;
  font-size: 16px;
  margin-left: 15px;
}

.sort-products {
  font-size: 16px;
}

#category {
  margin-top: 15px;
  width: 200px;
  height: 40px;
  padding: 10px;
  font-size: 16px;
  background: #fff;
  border-radius: 3px;
  border: transparent;
}

.category-wrap {
  margin-top: 115px;
  width: calc(80% - 15px);
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
}
/*

  ------Responsive -------

 */

@media screen and (max-width: 992px) {
  .card {
    width: calc(33.333% - 15px);
  }
}

@media screen and (max-width: 576px) {
  .container {
    width: 100%;
  }

  .card {
    width: 100%;
  }

  .category-wrap {
    width: calc(100% - 15px);
  }

  .sort-products label {
    display: none;
  }
}
</style>
