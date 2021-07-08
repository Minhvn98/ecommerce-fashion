<template>
  <div class="container">
    <div class="shopping-cart">
      <h1>Giỏ hàng</h1>
      <div class="not-products" v-if="products.length <= 0">
        <div>Không có sản phẩm trong giỏ hàng</div>
        <button class="back-to-shop" @click="onBackToShop">
          Quay lại mua hàng
        </button>
      </div>
      <div class="products">
        <product-cart
          v-for="product in products"
          :key="product.id"
          :product="product"
          @change-quantity="onChangeQuantityHandler"
          @delete-product="onDeleteHandler"
        ></product-cart>
      </div>
      <!-- end products -->

      <div class="shopping-cart-footer" v-if="products.length > 0">
        <div class="discount">
          <div>
            <label for="discount">Mã giảm giá</label>
            <input
              type="text"
              name="discount"
              id="discount"
              placeholder="Vui lòng nhập mã giảm giá"
              @change="onEnterDiscount"
              v-model.trim="discountCodeUser"
            />
          </div>
          <span>{{ textMessage }}</span>
        </div>
        <div class="checkout">
          <div class="summary">
            <ul>
              <li>
                Tổng sản phẩm:
                <strong> {{ totalNumberProducts }} sản phẩm</strong>
              </li>
              <li>
                Tổng tiền hàng:
                <strong> {{ formatMoney(totalAmountCart) }}</strong>
              </li>
              <li v-if="totalDiscount !== 0">
                Giảm giá ({{ discount }}%):
                <strong>- {{ formatMoney(totalDiscount) }}</strong>
              </li>
              <li>
                VAT(10%): <strong> {{ formatMoney(totalVat) }}</strong>
              </li>
              <li>
                <b>Tổng thanh toán: </b>
                <strong> {{ formatMoney(totalPayment) }}</strong>
              </li>
            </ul>
          </div>
          <button type="button" class="btn-checkout" @click="onCheckOutHandle">
            Mua hàng
          </button>
        </div>
      </div>
    </div>
  </div>

  <the-footer></the-footer>

  <transition name="fade">
    <div v-if="isShowModalDelete" class="background-black"></div>
  </transition>

  <transition name="fade">
    <modal-delete
      v-if="isShowModalDelete"
      :nameProduct="productDelete.name"
      @confirm="onConfirmProduct"
    ></modal-delete>
  </transition>
</template>

<script>
import ProductCart from "../components/ProductCart";
import ModalDelete from "../components/ModalDelete.vue";

export default {
  name: "ShoppingCart",
  components: {
    ProductCart,
    ModalDelete,
  },
  data() {
    return {
      // products: [],
      products: [
        {
          id: 1,
          image:
            "https://cf.shopee.vn/file/bab72e15497c082dba30b6764f46ccd0_tn",
          name: "Đầm Ngủ Mặc Nhà Lụa Satin Hisexy Không Mút Ngực M06",
          description: "Description for product item number 1",
          price: 210000,
          quantity: 2,
        },
        {
          id: 2,
          image:
            "https://cf.shopee.vn/file/f935953ec0595d44b2cc2ece5c74d49a_tn",
          name: "Váy ngủ lụa hai dây chất liệu lụa hàn cao cấp chất mát,mềm mịn",
          description: "Description for product item number 1",
          price: 500000,
          quantity: 1,
        },
        {
          id: 3,
          image:
            "https://cf.shopee.vn/file/c8d9f913bc978fff2117b22a911d887b_tn",
          name: "Mùa hè địu nữ váy ngủ mùa hè mỏng dệt bông dễ thương ngọt ngào nơ phong cách cổ tích váy ngủ không tay",
          description: "Description for product item number 1",
          price: 300000,
          quantity: 3,
        },
      ],
      productsFake: [],
      discountCodes: [
        {
          code: "300475",
          discount: 60,
        },
        {
          code: "020945",
          discount: 50,
        },
        {
          code: "COVID19",
          discount: 30,
        },
      ],
      discountCodeUser: "",
      discount: 0,
      textMessage: "",
      productDelete: {},
      isConfirm: false,
      isShowModalDelete: false,
    };
  },

  computed: {
    totalAmountCart() {
      return this.products.reduce(
        (total, product) => (total += product.price * product.quantity),
        0
      );
    },

    totalNumberProducts() {
      return this.products.reduce(
        (total, product) => (total += product.quantity),
        0
      );
    },

    totalVat() {
      return this.totalAmountCart * 0.1;
    },

    totalDiscount() {
      return this.totalAmountCart * (this.discount / 100);
    },

    totalPayment() {
      return this.totalAmountCart + this.totalVat - this.totalDiscount;
    },
  },

  methods: {
    formatMoney(price) {
      return price.toLocaleString("it-IT", {
        style: "currency",
        currency: "VND",
      });
    },

    onChangeQuantityHandler(idProduct, quantity) {
      quantity = parseInt(quantity);
      console.log("Changeeee", idProduct, quantity);

      return this.products.forEach((product) => {
        if (product.id === idProduct) {
          product.quantity = quantity;
        }
      });
    },

    onDeleteHandler(idProduct, nameProduct) {
      this.productDelete.name = nameProduct;
      this.productDelete.id = idProduct;
      this.isShowModalDelete = true;
      console.log(this.productDelete);
      // if (!this.isConfirm) return;

      // return (this.products = this.products.filter(
      //   (product) => product.id !== idProduct
      // ));
    },

    onEnterDiscount() {
      const discountObj = this.discountCodes.filter(
        (discount) =>
          discount.code.toLowerCase() === this.discountCodeUser.toLowerCase()
      );

      if (discountObj.length > 0) {
        this.discount = discountObj[0].discount;
        return (this.textMessage = `Bạn đã được giảm giá ${this.discount}%`);
      }
      this.textMessage = `Mã bạn nhập không chính xác`;
      console.log(this.totalPayment);
    },

    onConfirmProduct(isConfirm) {
      this.isConfirm = isConfirm;
      console.log(isConfirm);

      if (!this.isConfirm) return (this.isShowModalDelete = false);

      this.isShowModalDelete = false;
      return (this.products = this.products.filter(
        (product) => product.id !== this.productDelete.id
      ));
    },
    onBackToShop() {
      this.products = [...this.productsFake];
    },

    onCheckOutHandle() {
      alert("Vội thế ! Vội thế! Đợi thời gian nữa nhé <3");
    },
  },
};
</script>

<style scoped>
.not-products {
  width: 100%;
  background: #fff;
  margin-bottom: 15px;
  padding: 20px 15px;
  border-radius: 4px;
}

.summary li {
  list-style: none;
  padding-bottom: 10px;
}
strong {
  color: #ee4d2d;
}
input#discount {
  padding: 10px 20px;
  font-size: 1.1rem;
  border-radius: 5px;
  border: 1px solid #ccc;
}

.discount {
  width: 70%;
}

.discount label {
  color: #ee4d2d;
  padding-right: 15px;
}

.shopping-cart-footer {
  font-size: 1.1rem;
  background: #fff;
  border-radius: 5px;
  padding: 15px;
  display: flex;
  width: 100%;
  z-index: 11;
}

.products {
  width: 100%;
}

.shopping-cart h1 {
  color: #ee4d2d;
  padding: 20px 0;
}

.shopping-cart {
  margin-top: 115px;
  width: 100%;
}

.container {
  width: 80%;
  margin: auto;
}
.btn-checkout {
  border: 1px solid transparent;
  background: #ee4d2d;
  font-size: 1.1rem;
  color: #fff;
  border-radius: 4px;
  width: 100%;
  height: 40px;
}
.btn-checkout:hover {
  background: #d6563d;
}
.delete:hover {
  text-decoration: underline;
}

.checkout {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
}

.background-black {
  position: fixed;
  top: 0;
  z-index: 19;
  width: 100%;
  height: 900px;
  background: rgba(0, 0, 0, 0.507);
}

.discount span {
  color: #111;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.back-to-shop {
  background: #ee4d2d;
  padding: 10px 20px;
  font-size: 16px;
  border: none;
  color: #fff;
  margin-top: 20px;
  cursor: pointer;
  border-radius: 4px;
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
  }

  .card {
    align-items: center;
  }

  #discount {
    width: calc(100% - 111px);
  }
  .products {
    margin-bottom: 15px;
  }
}
</style>
