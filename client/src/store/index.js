import { createStore } from 'vuex';

import { searchProducts as apiSearch } from '../services/products.service';

const store = createStore({
  state() {
    return {
      user: {},
      count: 0,
      textSearch: '',
      productsSearch: []
    };
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
    }
  }
});

export default store;
