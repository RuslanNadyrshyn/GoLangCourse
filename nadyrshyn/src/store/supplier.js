import axios from "axios";

const state = {
  url: "http://localhost:8080/get_suppliers",
  suppliers: [],
  sortedSuppliers: [],
  suppliersTypes: [],
  supplierSortType: "",
  errors: [],
  loaded: false,
};

const mutations = {
  setSuppliers(state, suppliers) {
    state.suppliers = suppliers;
  },
  setSortedSuppliers(state, sortedSuppliers) {
    state.sortedSuppliers = sortedSuppliers;
  },
  setSuppliersTypes(state, suppliersTypes) {
    state.suppliersTypes = suppliersTypes;
  },
  setSupplierSortType(state, type) {
    state.supplierSortType = type;
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
  sortByType(context, type) {
    let suppliers = context.getters.getSuppliers;
    context.commit("setSupplierSortType", type);
    if (type === "all") {
      context.commit("setSortedSuppliers", suppliers);
    } else {
      let SortedArray = suppliers.filter((supplier) => supplier.type === type);
      context.commit("setSortedSuppliers", SortedArray);
    }
  },
};

const getters = {
  getSupplierURL: (state) => {
    return state.url;
  },
  getSuppliers: (state) => {
    return state.suppliers;
  },
  getSortedSuppliers: (state) => {
    if (state.supplierSortType === "") {
      return state.suppliers;
    } else {
      return state.sortedSuppliers;
    }
  },
  getSuppliersTypes: (state) => {
    return state.suppliersTypes;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
