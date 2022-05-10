import axios from "axios";

const state = {
  url: "http://localhost:8080/get_products",
  urlGetById: "http://localhost:8080/prod",
  product: null,
  products: [],
  sortedProducts: [],
  productsTypes: [],
  selectedType: "все",
  errors: [],
  loaded: false,
};

const mutations = {
  setProduct(state, product) {
    state.product = product;
  },
  setProducts(state, products) {
    state.products = products;
  },
  setProductsTypes(state, productsTypes) {
    state.productsTypes = productsTypes;
  },
  setSortedProducts(state, sortedProducts) {
    state.sortedProducts = sortedProducts;
  },
  setSelectedType(state, selected) {
    state.selectedType = selected;
  },
  setErrors(state, errors) {
    state.errors = errors;
  },
  setLoaded(state, loaded) {
    state.loaded = loaded;
  },
};

const actions = {
  fetchProducts(context, products) {
    context.commit("setLoaded", false);
    let types = [];
    for (let i = 0; i < products.length; i++)
      if (types.includes(products[i].type) === false)
        types.push(products[i].type);
    context.commit("setProducts", products);
    context.commit("setSortedProducts", products);
    context.commit("setProductsTypes", types);
    context.commit("setLoaded", true);
  },
  fetchById(context, id) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getByIdURL, {
        params: { id: id },
      })
      .then((res) => {
        context.commit("setProduct", res.data);
      })
      .catch((err) => console.log(err))
      .finally(() => {
        context.commit("setLoaded", true);
      });
  },
  sortByType(context, type) {
    context.commit("setSelectedType", type);
    let products = context.getters.getSortedProducts;
    if (type === "все") context.commit("setSortedProducts", products);
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
  getByIdURL: (state) => {
    return state.urlGetById;
  },
  getProducts: (state) => {
    return state.products;
  },
  getProduct: (state) => {
    return state.product;
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
