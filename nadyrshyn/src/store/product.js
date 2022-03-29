import axios from "axios";

const state = {
  url: "http://localhost:8080/get_products",
  products: [],
  productsTypes: [],
  errors: [],
  loaded: false,
};

const mutations = {
  setProducts(state, products) {
    state.products = products;
  },
  setProductsTypes(state, productsTypes) {
    state.productsTypes = productsTypes;
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
        let types = [];
        for (let i = 0; i < res.data.length; i++) {
          if (types.includes(res.data[i].type) === false) {
            types.push(res.data[i].type);
          }
        }
        context.commit("setProductsTypes", types);
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
  getProductsByType: (state) => {
    return state.products.filter((type) => type === state.products.type);
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
