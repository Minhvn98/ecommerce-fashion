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
      <div class="title-page">
        Kết quả tìm kiếm cho từ khóa
        <span>'{{ textSearch }}'</span>
      </div>
      <div class="category-link-wrap"></div>
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
      <div class="message" v-if="!productsSearch">
        <div>Không tìm thấy kết quả nào</div>
        <span>Hãy thử sử dụng từ khóa chung chung hơn</span>
      </div>
      <product
        v-for="product in productsSearch"
        :product="product"
        :key="product.id"
      ></product>
    </div>
    <button
      class="btn-view-more"
      v-if="productsSearch"
      @click="onClickViewMore"
    >
      {{ viewMoreText }}
    </button>
  </div>
</template>

<script>
import Product from "../components/Product.vue";

export default {
  name: "SearchPage",
  components: {
    Product,
  },

  data() {
    return {
      idCategory: 0,
      categories: [],
      products: [],
      offSet: 12,
      viewMoreText: "Xem Thêm",
    };
  },

  computed: {
    productsSearch() {
      console.log(this.$store.state.productsSearch);
      return this.$store.state.productsSearch;
    },

    textSearch() {
      return this.$store.state.textSearch;
    },
  },

  methods: {
    async onFilterByCategory() {},

    onSortHandler(event) {
      const sortType = event.target.value;

      if (sortType === "asc") {
        this.products.sort((a, b) => a.price - b.price);
      }

      if (sortType === "desc") this.products.sort((a, b) => b.price - a.price);
    },

    async onClickViewMore() {},
  },

  async created() {},
};
</script>


<style scoped>
:root {
  --primary-color: #ee4d2d;
}
.message {
  margin: auto;
  text-align: center;
  font-size: 20px;
  padding: 50px;
  line-height: 35px;
}

.message span {
  color: #6d6666;
}
.title-page {
  color: #3c3c3c;
  font-size: 20px;
}
.title-page span {
  color: #ee4d2d;
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
