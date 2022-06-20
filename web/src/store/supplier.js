import axios from "axios";
import config from "./config.js";

const state = {
  url: config.hostname + "/get_suppliers",
  urlGetByType: config.hostname + "/get_sup_by_type",
  urlGetByWorkingHours: config.hostname + "/get_sup_by_working_hours",
  suppliers: [],
  suppliersTypes: [],
  selectedType: "все",
  selectedSupplierId: 0,
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
  setSelectedType(state, selected) {
    state.selectedType = selected;
  },
  setSelectedSupplier(state, selectedId) {
    state.selectedSupplierId = selectedId;
  },
  setErrors(state, errors) {
    state.errors = errors;
  },
  setLoaded(state, loaded) {
    state.loaded = loaded;
  },
};

const actions = {
  fetchSuppliers(context) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getSupplierURL)
      .then((res) => {
        context.commit("setSuppliers", res.data.Suppliers);
        context.commit("setSuppliersTypes", res.data.Types);
      })
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => context.commit("setLoaded", true));
  },
  getByType(context, type) {
    context.commit("setSelectedType", type);
    context.commit("setSelectedSupplier", 0);
    if (type === "все") type = "";
    if (type === "Открыто") type = "workingHours";
    axios
      .get(context.getters.getByTypeURL, {
        params: { type: type },
      })
      .then((res) => context.commit("setSuppliers", res.data))
      .catch((err) => context.commit("setErrors", err))
      .finally(() => context.commit("setLoaded", true));
  },
};

const getters = {
  getSupplierURL: (state) => {
    return state.url;
  },
  getByTypeURL: (state) => {
    return state.urlGetByType;
  },
  getSuppliers: (state) => {
    return state.suppliers;
  },
  getSuppliersTypes: (state) => {
    return state.suppliersTypes;
  },
  getSelectedType: (state) => {
    return state.selectedType;
  },
  getSelectedSupplier: (state) => {
    return state.selectedSupplierId;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
