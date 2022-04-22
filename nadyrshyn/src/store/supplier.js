import axios from "axios";

const state = {
  url: "http://localhost:8080/get_suppliers",
  suppliers: [],
  sortedSuppliers: [],
  suppliersTypes: [],
  selectedType: "все",
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
  fetchSuppliers(context) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getSupplierURL)
      .then((res) => {
        let types = [];
        for (let i = 0; i < res.data.length; i++) {
          if (types.includes(res.data[i].type) === false)
            types.push(res.data[i].type);
          for (let j = 0; j < res.data[i].menu.length; j++) {
            res.data[i].menu[j].counter = 0;
            res.data[i].menu[j].supplier_id = res.data[i].id;
            res.data[i].menu[j].supplier_name = res.data[i].name;
            res.data[i].menu[j].supplier_image = res.data[i].image;
          }
        }
        context.commit("setSuppliers", res.data);
        context.commit("setSortedSuppliers", res.data);
        context.commit("setSuppliersTypes", types);
      })
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => context.commit("setLoaded", true));
  },
  sortByType(context, type) {
    context.commit("setSelectedType", type);
    let suppliers = context.getters.getSuppliers;
    if (type !== "все")
      suppliers = suppliers.filter((supplier) => supplier.type === type);
    context.commit("setSortedSuppliers", suppliers);
    return suppliers;
  },
  sortByWorkingHours(context) {
    context.commit("setSelectedType", "Открыто");
    let suppliers = context.getters.getSuppliers;
    let time = new Date();
    let hours = time.getHours();
    let minutes = time.getMinutes();
    let sorted = [];

    if (hours < 10) hours = "0" + hours;
    if (minutes < 10) minutes = "0" + minutes;

    let timestamp = hours + ":" + minutes;
    for (let i = 0; i < suppliers.length; i++)
      if (
        timestamp >= suppliers[i].workingHours.opening &&
        timestamp < suppliers[i].workingHours.closing
      )
        sorted.push(suppliers[i]);
    context.commit("setSortedSuppliers", sorted);
    return sorted;
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
