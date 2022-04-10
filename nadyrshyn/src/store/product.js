import axios from "axios";

const state = {
  url: "http://localhost:8080/get_products",
  products: [],
  sortedProducts: [],
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
  setSortedProducts(state, sortedProducts) {
    state.sortedProducts = sortedProducts;
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
        context.commit("setSortedProducts", res.data);
        let types = [];
        for (let i = 0; i < res.data.length; i++) {
          if (types.includes(res.data[i].type) === false)
            types.push(res.data[i].type);
        }
        context.commit("setProductsTypes", types);
      })
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => {
        context.commit("setLoaded", true);
      });
  },
  fetchP(context, products) {
    context.commit("setProducts", products);
    context.commit("setSortedProducts", products);
    let types = [];
    for (let i = 0; i < products.length; i++) {
      if (types.includes(products[i].type) === false)
        types.push(products[i].type);
    }
    context.commit("setProductsTypes", types);
    context.commit("setLoaded", true);
  },
  sortByType(context, type) {
    let products = context.getters.getProducts;
    if (type === "all") context.commit("setSortedProducts", products);
    else {
      let SortedArray = products.filter((product) => product.type === type);
      context.commit("setSortedProducts", SortedArray);
    }
  },
  sortBySupplier(context, products) {
    context.commit("setSortedProducts", products);
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
    return state.sortedProducts;
  },
  getProductsTypes: (state) => {
    return state.productsTypes;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
