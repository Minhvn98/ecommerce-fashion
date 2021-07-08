<template>
  <div class="card">
    <img class="card-image" :src="product.image" :alt="product.name" />
    <div class="card-title">
      <router-link
        class="name-product"
        :to="{
          name: 'products-detail-router',
          params: { id: product.id },
        }"
        >{{ product.name }}</router-link
      >
    </div>
    <div class="card-body">
      <span>{{ formatMoney(product.price) }}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "Product",

  props: {
    product: Object,
  },

  methods: {
    formatMoney(price) {
      return price.toLocaleString("it-IT", {
        style: "currency",
        currency: "VND",
      });
    },
  },
};
</script>

<style scoped>
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