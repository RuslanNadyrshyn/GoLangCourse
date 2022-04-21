import axios from "axios";

const state = {
  getOrderURL: "http://localhost:8080/get_order",
  postOrderURL: "http://localhost:8080/post_order",
  order: Object,
  errors: [],
  loaded: false,
};

const mutations = {
  setOrder(state, order) {
    state.order = order;
  },
  setErrors(state, errors) {
    state.errors = errors;
    console.log("errors: ", errors);
  },
  setLoaded(state, loaded) {
    state.loaded = loaded;
  },
};
// fetchOrder(context) {
//     context.commit("setLoaded", false);
//     axios
//         .get(context.getters.getOrderURL)
//         .then((res) => {
//             let types = [];
//             for (let i = 0; i < res.data.length; i++) {
//                 if (types.includes(res.data[i].type) === false)
//                     types.push(res.data[i].type);
//             }
//             context.commit("setOrder", res.data);
//             context.commit("setSortedSuppliers", res.data);
//             context.commit("setSuppliersTypes", types);
//         })
//         .catch((err) => [context.commit("setErrors", [err])])
//         .finally(() => {
//             context.commit("setLoaded", true);
//         });
// },
const actions = {
  fetchOrderPOST(context, order) {
    console.log("start_fetch");
    axios
      .post(context.getters.getPostURL, order)
      .then((res) => console.log(res.data))
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => {
        context.commit("setLoaded", true);
      });
  },
};

const getters = {
  getOrder: () => {
    return state.order;
  },
  getPostURL: () => {
    return state.postOrderURL;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
