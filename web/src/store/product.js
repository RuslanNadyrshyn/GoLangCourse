import axios from "axios";
import config from "./config.js";

const state = {
  url: config.hostname + "/get_products",
  urlGetById: config.hostname + "/get_prod_by_id",
  urlGetByParams: config.hostname + "/get_prod_by_params",
  products: [],
  productsTypes: [],
  selectedType: "все",
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
  fetchProducts(context) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getProductURL)
      .then((res) => {
        context.commit("setProducts", res.data.Products);
        context.commit("setProductsTypes", res.data.Types);
      })
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => context.commit("setLoaded", true));
  },
  getByParams(context, params) {
    context.commit("setSelectedType", params.prodType);
    if (params.supType === "все") params.supType = "";
    if (params.prodType === "все") params.prodType = "";
    if (params.supType === "Открыто") params.supType = "workingHours";
    axios
      .get(context.getters.getByParamsURL, {
        params: {
          sup_id: params.supId,
          sup_type: params.supType,
          prod_type: params.prodType,
        },
      })
      .then((res) => {
        context.commit("setProducts", res.data.Products);
        context.commit("setProductsTypes", res.data.Types);
      })
      .catch((err) => context.commit("setErrors", err))
      .finally(() => context.commit("setLoaded", true));
  },
};

const getters = {
  getProductURL: (state) => {
    return state.url;
  },
  getByIdURL: (state) => {
    return state.urlGetById;
  },
  getByParamsURL: (state) => {
    return state.urlGetByParams;
  },
  getProducts: (state) => {
    return state.products;
  },
  getProductsTypes: (state) => {
    return state.productsTypes;
  },
  getSelectedType: (state) => {
    return state.selectedType;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
