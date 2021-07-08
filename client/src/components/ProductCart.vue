<template>
  <div class="card">
    <div class="thumbnail">
      <img :src="product.image" :alt="product.name" />
    </div>
    <div class="detail">
      <div class="name">
        {{ product.name }}
      </div>
      <div class="price">{{ formatMoney(product.price) }}</div>
    </div>
    <div class="right">
      <div class="quantity">
        <input
          type="number"
          :value="product.quantity"
          min="1"
          @change="onChangeQuantity($event, product.id)"
        />
      </div>
      <div
        class="delete"
        @click="onDeleteHandler($event, product.id, product.name)"
      >
        Xóa
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductCart",
  props: {
    product: Object,
  },
  data() {
    return {};
  },
  methods: {
    formatMoney(price) {
      return price.toLocaleString("it-IT", {
        style: "currency",
        currency: "VND",
      });
    },

    onDeleteHandler($event, idProduct, nameProduct) {
      this.$emit("delete-product", idProduct, nameProduct);
      // console.log("Delete", idProduct);
    },

    onChangeQuantity($event, idProduct) {
      let quantity = $event.target.value;
      if (!quantity) {
        alert("Nhập tử tế vào bạn ơi.");
        quantity = 0;
      }
      this.$emit("change-quantity", idProduct, quantity);
      console.log("Change", quantity, idProduct);
    },
  },
};
</script>

<style scoped>
.thumbnail img {
  width: 100%;
  border-radius: 4px;
}

.thumbnail {
  width: 15%;
  margin-left: 15px;
}

.card {
  width: 100%;
  display: flex;
  padding: 15px 0;
  background: #fff;
  margin-bottom: 15px;
  border-radius: 5px;
  cursor: pointer;
  transition: 0.2s;
}

.card:hover {
  transform: scale(1.005);
}

.name {
  font-size: 1.2rem;
}

.delete {
  color: #ee4d2d;
  cursor: pointer;
}

.price {
  padding-top: 15px;
  color: #ee4d2d;
}
.detail {
  width: 55%;
  padding: 0 30px;
}

input[type="number"] {
  padding: 10px;
  width: 100%;
  font-size: 1.1rem;
  border: 1px solid #ddd;
}

.quantity {
  width: 60px;
}

.right {
  display: flex;
  width: 30%;
  justify-content: space-around;
  align-items: center;
}

@media screen and (max-width: 768px) {
  .container {
    width: calc(100% - 30px);
    margin: 0 15px;
  }

  .quantity {
    width: 60px;
  }

  .discount {
    width: 100%;
    padding-bottom: 15px;
  }

  .shopping-cart-footer {
    flex-direction: column;
    width: calc(100% - 30px);
  }

  .card {
    align-items: center;
  }

  #discount {
    width: calc(100% - 111px);
  }
  .products {
    margin-bottom: 330px;
  }
}
</style>
