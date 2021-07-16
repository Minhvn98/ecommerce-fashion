<template>
  <div class="card">
    <div class="thumbnail">
      <img
        :src="`${BASE_URL_IMAGE}${product.product_images[0].uri}`"
        :alt="product.name"
      />
    </div>
    <div class="detail">
      <div class="detail-left">
        <div class="name">
          {{ product.name }}
        </div>
        <div class="price">{{ formatMoney(product.price) }}</div>
      </div>
      <div class="detail-right">
        <div class="quantity">
          <button
            @click="
              decreaseProduct();
              onChangeQuantity(product.id);
            "
            class="btn-quantity"
          >
            -
          </button>
          <input
            type="number"
            min="1"
            disabled
            @change="onChangeQuantity(product.id)"
            v-model="quantitySelected"
          />
          <button
            @click="
              increaseProduct();
              onChangeQuantity(product.id);
            "
            class="btn-quantity"
          >
            +
          </button>
        </div>
        <div v-if="errMessage" class="errMessage">{{ errMessage }}</div>

        <div class="delete" @click="onDeleteHandler(product.id, product.name)">
          Xóa
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { formatMoney } from "../utils/money.util";

const BASE_URL_IMAGE = process.env.VUE_APP_BASE_URL_IMAGE;

export default {
  name: "ProductCart",
  props: {
    product: Object,
    quantity: Number,
  },
  data() {
    return {
      BASE_URL_IMAGE,
      quantitySelected: this.quantity,
      errMessage: "",
    };
  },

  methods: {
    formatMoney,

    onDeleteHandler(idProduct, nameProduct) {
      this.$emit("delete-product", idProduct, nameProduct);
      // console.log("Delete", idProduct);
    },

    onChangeQuantity(idProduct) {
      let quantity = this.quantitySelected;
      if (!quantity) {
        this.quantitySelected = 1;
        alert("Nhập tử tế vào bạn ơi.");
      }
      this.$emit("change-quantity", idProduct, quantity);
      console.log("Change", idProduct, quantity);
    },

    decreaseProduct() {
      if (this.quantitySelected <= 1) {
        this.quantitySelected = 1;
        return this.onDeleteHandler(this.product.id, this.product.name);
      }
      this.errMessage = "";
      this.quantitySelected -= 1;
    },

    increaseProduct() {
      if (this.quantitySelected >= this.product.quantity) {
        this.errMessage = "Bạn đã chọn hết số lượng hàng trong kho";
        console.log(this.product.quantity);

        return (this.quantitySelected = this.product.quantity);
      }
      this.quantitySelected += 1;
    },
  },

  created() {},
};
</script>

<style scoped>
.errMessage {
  color: #e43917;
}
.detail-right {
  display: flex;
  align-items: center;
}

.quantity {
  display: flex;
  width: 205px;
  margin-right: 50px;
}

.btn-quantity {
  width: 65px;
  height: 35px;
  font-size: 32px;
  color: #555;
  background: #fff;
  border: 0.5px solid #ccc;
  cursor: pointer;
  line-height: 16px;
}

.delete:hover {
  text-decoration: underline;
}

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
  background: #eee;
  padding: 6px 15px;
  border-radius: 3px;
}

.price {
  padding-top: 15px;
  color: #ee4d2d;
}

.detail {
  width: 80%;
  padding-left: 30px;
  display: flex;
  justify-content: space-between;
}

input[type="number"] {
  padding-left: 10px;
  width: 55px;
  font-size: 1.1rem;
  border: 1px solid #ddd;
  height: 35px;
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
