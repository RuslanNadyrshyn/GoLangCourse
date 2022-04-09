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
    let prod = [];
    for (let i = 0; i < suppliers.length; i++) {
      for (let j = 0; j < suppliers[i].menu.length; j++) {
        prod.push(suppliers[i].menu[j]);
      }
    }
    // console.log(prod);
  },
  setSortedSuppliers(state, sortedSuppliers) {
    state.sortedSuppliers = sortedSuppliers;
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
};

const actions = {
  fetchSuppliers(context) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getSupplierURL)
      .then((res) => {
        context.commit("setSuppliers", res.data);
        context.commit("setSortedSuppliers", res.data);
        let prod = [];
        let types = [];
        for (let i = 0; i < res.data.length; i++) {
          if (types.includes(res.data[i].type) === false) {
            types.push(res.data[i].type);
          }
          for (let j = 0; j < res.data[i].menu.length; j++) {
            prod.push(res.data[i].menu[j]);
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
    if (type === "all") context.commit("setSortedSuppliers", suppliers);
    else {
      let SortedArray = suppliers.filter((supplier) => supplier.type === type);
      context.commit("setSortedSuppliers", SortedArray);
    }
  },
  sortByWorkingHours(context) {
    let suppliers = context.getters.getSuppliers;
    let time = new Date();
    let hours = time.getHours();
    let minutes = time.getMinutes();
    let sorted = [];

    if (hours < 10) hours = "0" + hours;
    if (minutes < 10) minutes = "0" + minutes;

    let timestamp = hours + ":" + minutes;
    for (let i = 0; i < suppliers.length; i++) {
      if (
        timestamp > suppliers[i].workingHours.opening &&
        timestamp < suppliers[i].workingHours.closing
      )
        sorted.push(suppliers[i]);
    }
    context.commit("setSortedSuppliers", sorted);
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
    return state.sortedSuppliers;
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
