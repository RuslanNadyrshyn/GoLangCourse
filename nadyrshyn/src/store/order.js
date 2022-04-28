import axios from "axios";

const state = {
  orderURL: "http://localhost:8080/get_order",
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

const actions = {
  fetchOrderPOST(context, order) {
    console.log("start_fetch");
    axios
      .post(context.getters.getPostOrderURL, order)
      .then((res) => {
        alert("Спасибо за заказ. Номер вашего заказа: " + res.data);
      })
      .catch((err) => [context.commit("setErrors", [err])])
      .finally(() => {
        context.commit("setLoaded", true);
      });
  },
  fetchOrder(context, orderId) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getOrderURL, {
        params: {
          id: orderId,
        },
      })
      .then((res) => {
        context.commit("setOrder", res.data);
      })
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
  getOrderProducts: () => {
    return state.order.products;
  },
  getPostOrderURL: () => {
    return state.postOrderURL;
  },
  getOrderURL: () => {
    return state.orderURL;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
