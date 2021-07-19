import router from '../../router';

import {
  addToCart,
  getCart,
  changeQuantityProductCart,
  deleteProductCart as apiDeleteProductCart
} from '../../services/cart.service';

const state = () => ({
  cartProducts: []
});

const getters = {
  cartNumber(state) {
    console.log(state.cartProducts.length);
    return state.cartProducts.reduce((total, product) => {
      return (total += product.quantity_in_cart);
    }, 0);
  }
};

const mutations = {
  setCart(state, products) {
    state.cartProducts = products;
  }
};

const actions = {
  async addProductToCart({ commit }, { idProduct, quantity }) {
    const { data: products, status } = await addToCart(idProduct, quantity);

    if (status === 401) {
      return router.push({ name: 'login' });
    }

    commit('setCart', products);
  },

  async getCartProduct({ commit }) {
    const { data: products, status } = await getCart();
    console.log('z', products);
    if (status === 401) {
      return router.push({ name: 'login' });
    }

    commit('setCart', products);
  },

  async changeCartQuantity({ commit }, { idProduct, quantity }) {
    const { data: products } = await changeQuantityProductCart(
      idProduct,
      quantity
    );

    commit('setCart', products);
  },

  async deleteProductCart({ commit }, { idProduct }) {
    const { data: products } = await apiDeleteProductCart(idProduct);

    commit('setCart', products);
  }
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
};
