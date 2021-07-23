import { createStore } from 'vuex';
import router from '../router';
import { searchProducts as apiSearch } from '../services/products.service';
import {
  addToCart,
  getCart,
  changeQuantityProductCart,
  deleteProductCart as apiDeleteProductCart
} from '../services/cart.service';

const store = createStore({
  state() {
    return {
      user: null,
      count: 0,
      textSearch: '',
      productsSearch: [],
      cart: [],
      layout: 'LayoutShop'
    };
  },

  getters: {
    cartNumber(state) {
      console.log(state.cart.length);
      return state.cart.reduce((total, product) => {
        return (total += product.quantity_in_cart);
      }, 0);
    },
    isAuthenticated: state => !!state.user
  },

  mutations: {
    setProducts(state, products) {
      state.productsSearch = products;
    },

    setText(state, text) {
      state.textSearch = text;
    },

    setUser(state, user) {
      state.user = user;
    },

    setCart(state, products) {
      state.cart = [...products];
    },

    setCartNumber(state, cartNumber) {
      state.cartNumber = cartNumber;
    },

    updateLayout(state, layout) {
      state.layout = layout;
    }
  },

  actions: {
    async searchProducts({ commit }, { textSearch }) {
      const { data: products } = await apiSearch(textSearch);

      commit('setProducts', products);
      commit('setText', textSearch);
    },

    setUser({ commit }, user) {
      commit('setUser', user);
    },

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
  }
});

export default store;
