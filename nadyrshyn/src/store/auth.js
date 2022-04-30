import axios from "axios";

const state = {
  urlLogin: "http://localhost:8080/login",
  urlSignIn: "http://localhost:8080/sign_in",
  urlUser: "http://localhost:8080/user",
  urlOrders: "http://localhost:8080/orders",
  Access: false,
  user: null,
  orders: [],
  accessToken: "",
  refreshToken: "",
  errors: [],
  loaded: false,
};

const mutations = {
  setUser(state, user) {
    state.user = user;
  },
  setOrders(state, orders) {
    state.orders = orders;
  },
  setAccessToken(state, token) {
    state.accessToken = token;
  },
  setRefreshToken(state, token) {
    state.refreshToken = token;
  },
  clearTokens(state) {
    state.accessToken = "";
    state.refreshToken = "";
  },
  setAccess(state, value) {
    state.Access = value;
  },
  setErrors(state, errors) {
    state.errors = errors;
  },
  setLoaded(state, loaded) {
    state.loaded = loaded;
  },
};

const actions = {
  AddTokens(context, tokens) {
    context.commit("addAccessToken", tokens.access_token);
    context.commit("addRefreshToken", tokens.refresh_token);

    context.commit("setAccess", true);
  },
  fetchUser(context, id) {
    context.commit("setLoaded", false);
    axios
      .get(context.getters.getUserURL, {
        params: { id: id },
      })
      .then((res) => {
        context.commit("setUser", res.data);
      })
      .catch((err) => context.commit("setErrors", err))
      .finally(() => {
        context.commit("setLoaded", true);
      });
  },
  fetchOrders(context, userId) {
    axios
      .get(context.getters.getOrdersURL, {
        params: { userId: userId },
      })
      .then((res) => {
        context.commit("setOrders", res.data);
      })
      .catch((err) => console.log(err));
  },
};

const getters = {
  getUserURL: () => {
    return state.urlUser;
  },
  getOrdersURL: () => {
    return state.urlOrders;
  },
  getAccessToken: () => {
    return state.accessToken;
  },
  getRefreshToken: () => {
    return state.refreshToken;
  },
  getAccess: (state) => {
    return state.Access;
  },
  getUser: (state) => {
    return state.user;
  },
  getOrders: (state) => {
    return state.orders;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
