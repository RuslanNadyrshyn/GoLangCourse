import axios from "axios";

const state = {
  url: "http://localhost:8080/get_suppliers",
  suppliers: [],
  suppliersTypes: [],
  selectedSupplier: { id: Number, name: String, image: String },
  errors: [],
  loaded: false,
};

const mutations = {
  setSuppliers(state, suppliers) {
    state.suppliers = suppliers;
  },
  setSuppliersTypes(state, suppliersTypes) {
    state.suppliersTypes = suppliersTypes;
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
        let types = [];
        for (let i = 0; i < res.data.length; i++) {
          if (types.includes(res.data[i].type) === false) {
            types.push(res.data[i].type);
          }
        }
        context.commit("setSuppliersTypes", types);
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
