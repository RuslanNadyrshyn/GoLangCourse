import axios from "axios";

const state = {
  url: "http://localhost:8080/get_products",
  products: [],
  errors: [],
  loaded: false,
};

const mutations = {
  setProducts(state, products) {
    state.products = products;
  },
  setErrors(state, errors) {
    state.errors = errors;
  },
  setLoaded(state, loaded) {
    state.loaded = loaded;
  },
};

const actions = {
  fetchProducts(context) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getProductURL)
      .then((res) => {
        context.commit("setProducts", res.data);
      })
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => {
        context.commit("setLoaded", true);
      });
  },
};

const getters = {
  getProductURL: (state) => {
    return state.url;
  },
  getProducts: (state) => {
    return state.products;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
