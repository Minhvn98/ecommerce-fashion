import { searchProducts as apiSearch } from '../../services/products.service';

const state = () => ({
  textSearch: '',
  productsSearch: []
});

const mutations = {
  setProducts(state, products) {
    state.productsSearch = products;
  },

  setText(state, text) {
    state.textSearch = text;
  }
};

const actions = {
  async searchProducts({ commit }, { textSearch }) {
    const { data: products } = await apiSearch(textSearch);

    commit('setProducts', products);
    commit('setText', textSearch);
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
