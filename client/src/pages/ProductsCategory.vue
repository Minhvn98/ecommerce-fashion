<template>
  <section class="container category-wrap">
    <div class="category mobile">
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
    <button class="btn-view-more" @click="onClickViewMore">
      {{ viewMoreText }}
    </button>
  </div>
</template>

<script>
import Product from "../components/Product.vue";

import { getProductsByCategory } from "../services/products.service";
import { getCategories } from "../services/categories.service";

export default {
  name: "ProductsCategory",
  components: {
    Product,
  },

  data() {
    return {
      idCategory: 0,
      categories: [],
      products: [],
      offSet: 0,
      viewMoreText: "Xem Thêm",
    };
  },

  methods: {
    async onFilterByCategory() {
      const { data: products } = await getProductsByCategory(this.idCategory);
      this.products = products;
    },

    onSortHandler(event) {
      const sortType = event.target.value;

      if (sortType === "asc") {
        this.products.sort((a, b) => a.price - b.price);
      }

      if (sortType === "desc") this.products.sort((a, b) => b.price - a.price);
    },

    async onClickViewMore() {
      this.offSet += 8;

      const { data: products } = await getProductsByCategory(
        this.idCategory,
        this.offSet
      );
      console.log(products);

      if (!products) {
        return (this.viewMoreText = "Đã Hết Sản Phẩm");
      }

      this.products = this.products.concat(products);
    },

    async fetchData() {
      this.offSet = 0;
      this.viewMoreText = "Xem Thêm";

      if (!this.$route.params.slug) return; // tranh error khi thoat khoi trang product detail

      this.idCategory = this.$route.params.slug.split("-cat.").slice(-1)[0];
      console.log(this.idCategory);

      // document.title = this.product.name;

      const { data: products } = await getProductsByCategory(this.idCategory);
      const { data: categories } = await getCategories();
      this.categories = categories;
      this.products = products;
    },
  },

  async created() {
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
:root {
  --primary-color: #ee4d2d;
}

.btn-view-more {
  display: block;
  margin: auto;
  padding: 12px 50px;
  font-size: 17px;
  border: transparent;
  background: #ee4d2d;
  border-radius: 3px;
  color: #fff;
  margin-top: 10px;
  margin-bottom: 15px;
  cursor: pointer;
}

.btn-view-more:hover {
  background: #f05d40;
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
