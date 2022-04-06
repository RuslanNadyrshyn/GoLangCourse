import axios from "axios";

const state = {
  url: "http://localhost:8080/get_products",
  products: [],
  sortedProducts: [],
  productsTypes: [],
  sortType: "",
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
  setSortedProducts(state, sortedProducts) {
    state.sortedProducts = sortedProducts;
  },
  setSortType(state, type) {
    state.sortType = type;
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
  sortByType(context, type) {
    let products = context.getters.getProducts;
    context.commit("setSortType", type);
    if (type === "all") {
      context.commit("setSortedProducts", products);
    } else {
      let SortedArray = products.filter((product) => product.type === type);
      context.commit("setSortedProducts", SortedArray);
    }
  },
};

const getters = {
  getProductURL: (state) => {
    return state.url;
  },
  getProducts: (state) => {
    return state.products;
  },
  getSortedProducts: (state) => {
    if (state.sortType === "") {
      return state.products;
    } else {
      return state.sortedProducts;
    }
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
