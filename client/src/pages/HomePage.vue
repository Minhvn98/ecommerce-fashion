<template>
  <the-banner></the-banner>

  <section class="container category-wrap">
    <div class="category mobile">
      <h3>Danh mục sản phẩm</h3>
      <select
        name="category"
        id="category"
        v-model="idCategory"
        @change="onFilterByCategory"
      >
        <option value="0">
          <router-link
            :to="{
              name: 'products-category',
              params: { slug: `tat-ca-san-pham-cat.0` },
            }"
            class="category-link"
          >
            Tất cả sản phẩm
          </router-link>
        </option>

        <router-link
          v-for="category in categories"
          :to="{
            name: 'products-category',
            params: { slug: `${category.name}-cat.${category.id}` },
          }"
          class="category-link"
          :key="category.id"
          >
          <option>
            {{ category.name }}
          </option>
        </router-link>
      </select>
    </div>

    <div class="category pc">
      <h3>Danh mục sản phẩm</h3>
      <div class="category-link-wrap">
        <router-link
          :to="{
            name: 'products-category',
            params: { slug: `tat-ca-san-pham-cat.0` },
          }"
          class="category-link"
        >
          Tất cả sản phẩm
        </router-link>
        <router-link
          v-for="category in categories"
          :to="{
            name: 'products-category',
            params: { slug: `${category.name}-cat.${category.id}` },
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

import { getProducts } from "../services/products.service";
import { getCategories } from "../services/categories.service";

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
      idCategory: 0,
    };
  },

  async created() {
    document.title =
      "Shopee Việt Nam | Mua và Bán Trên Ứng Dụng Di Động Hoặc Website";

    const { data: products } = await getProducts();
    const { data: categories } = await getCategories();

    this.products = products;
    this.categories = categories;
  },
};
</script>

<style scoped>
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

/*

----------Product -----------

*/

.container {
  width: 80%;
  margin: auto;
  margin-top: 15px;
}

.category-wrap {
  width: calc(80% - 15px);
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
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
  cursor: pointer;
}

.card:hover {
  transform: translateY(-2px);
}

.category-link-wrap {
  margin-top: 15px;
  display: flex;
}

.category-link {
  padding: 10px 20px;
  background: #fff;
  text-decoration: none;
  color: #111;
  display: block;
  margin-right: 10px;
  border-radius: 3px;
}
.category-link:hover {
  color: #ee513c;
}

.title {
  font-weight: 400;
  padding: 10px;
  font-size: 20px;
}

.mobile {
  display: none;
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

  .mobile {
    display: block;
  }

  .pc {
    display: none;
  }
}
</style>
