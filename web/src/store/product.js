import axios from "axios";

//const ip = "http://45.148.29.14:8080";
const ip = "http://localhost:8080";

const state = {
  url: ip + "/get_products",
  urlGetTypes: ip + "/get_prod_types",
  urlGetById: ip + "/get_prod_by_id",
  urlGetByType: ip + "/get_prod_by_type",
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
  getByType(context, type) {
    if (type === "все") {
      actions.fetchProducts(context);
    } else {
      axios
        .get(context.getters.getByTypeURL, {
          params: { type: type },
        })
        .then((res) => context.commit("setProducts", res.data))
        .catch((err) => context.commit("setErrors", err))
        .finally(() => context.commit("setLoaded", true));
    }
  },
  getByParams(context, params) {
    console.log("getting by params ", params);
    context.commit("setSelectedType", params.prodType);

    if (params.supType === "все") {
      params.supType = "";
      // actions.getByType(context, params.prodType);
    }
    if (params.prodType === "все") params.prodType = "";
    if (params.supType === "Открыто") params.supType = "workingHours";
    axios
      .get("http://localhost:8080/get_prod_by_params", {
        params: {
          sup_id: params.supId,
          sup_type: params.supType,
          prod_type: params.prodType,
        },
      })
      .then((res) => {
        console.log(res.data);
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
  getByTypeURL: (state) => {
    return state.urlGetByType;
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
