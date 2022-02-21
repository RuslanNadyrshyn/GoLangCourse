import axios from "axios";

const state = {
  url: "http://localhost:8080/get_suppliers",
  suppliers: [],
  selectedSupplier: { Number, String },
  errors: [],
  loaded: false,
};

const mutations = {
  setSuppliers(state, suppliers) {
    state.suppliers = suppliers;
  },
  setErrors(state, errors) {
    state.errors = errors;
  },
  setLoaded(state, loaded) {
    state.loaded = loaded;
  },
  setSelectedSupplier(state) {
    state.selectedSupplier = {
      id: 5,
      name: "Pizza Club",
      image: "https://eda.ua/images/506509-195-195-burger_club_harkov.jpg",
    };
  },
};

const actions = {
  fetchSuppliers(context) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getSupplierURL)
      .then((res) => {
        context.commit("setSuppliers", res.data);
      })
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => {
        context.commit("setLoaded", true);
      });
  },
};

const getters = {
  getSupplierURL: (state) => {
    return state.url;
  },
  getSuppliers: (state) => {
    return state.suppliers;
  },
  getSelectedSupplier: (state) => {
    return state.selectedSupplier;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
