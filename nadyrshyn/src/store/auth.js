import axios from "axios";

const state = {
  urlLogin: "http://45.148.29.14:8080/login",
  urlSignIn: "http://45.148.29.14:8080/sign_in",
  urlUser: "http://45.148.29.14:8080/profile",
  urlOrders: "http://45.148.29.14:8080/orders",
  urlRefresh: "http://45.148.29.14:8080/refresh",
  Access: false,
  user: [],
  orders: [],
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
  SignIn(context, user) {
    console.log(user);
    axios
      .post(context.getters.getSignInURL, user)
      .then(() => {
        let login = {
          email: user.email,
          password: user.password,
        };
        actions.Login(context, login);
      })
      .catch((err) => {
        console.log(err);
      });
  },
  Login(context, login) {
    axios
      .post(context.getters.getLoginURL, login)
      .then((res) => {
        localStorage.setItem("delivery_tokens", JSON.stringify(res.data));
        actions.fetchProfile(context);
        setTimeout(async () => {
          if (context.getters.getAccess) await actions.RefreshTokens(context);
        }, 3500000);
      })
      .catch((err) => console.log(err))
      .finally(() => {
        console.log("finally logged in)");
      });
  },
  fetchProfile(context) {
    context.commit("setLoaded", false);
    context.commit("setAccess", false);

    if (localStorage.getItem("delivery_tokens").length) {
      let tokens = JSON.parse(localStorage.getItem("delivery_tokens"));
      axios
        .get(context.getters.getUserURL, {
          headers: { Authorization: "Bearer " + tokens.access_token },
        })
        .then((res) => {
          console.log("user fetched");
          context.commit("setUser", res.data);
          context.commit("setAccess", true);
        })
        .catch((err) => context.commit("setErrors", err))
        .finally(() => {
          context.commit("setLoaded", true);
        });
    }
  },
  fetchOrders(context) {
    if (localStorage.getItem("delivery_tokens").length) {
      let tokens = JSON.parse(localStorage.getItem("delivery_tokens"));
      axios
        .get(context.getters.getOrdersURL, {
          headers: { Authorization: "Bearer " + tokens.access_token },
        })
        .then((res) => {
          if (res.data != null) context.commit("setOrders", res.data);
          else context.commit("setOrders", []);
        })
        .catch((err) => console.log(err));
    }
  },
  Logout(context) {
    context.commit("setUser", []);
    context.commit("setAccess", false);
    localStorage.setItem("delivery_tokens", JSON.stringify(""));
  },
  RefreshTokens(context) {
    let tokens = JSON.parse(localStorage.getItem("delivery_tokens"));
    let refreshRequest = {
      user_id: context.getters.getUserID,
      access_token: tokens.access_token,
      refresh_token: tokens.refresh_token,
    };
    axios
      .post(context.getters.getRefreshURL, refreshRequest)
      .then((res) => {
        console.log(res.data);
        localStorage.setItem("delivery_tokens", JSON.stringify(res.data));
        setTimeout(async () => {
          console.log("refresh");
          if (context.getters.getAccess) await actions.RefreshTokens(context);
        }, 3500000);
      })
      .catch((err) => console.log(err))
      .finally(() => {
        console.log("finally logged in)");
      });
  },
};

const getters = {
  getSignInURL: () => {
    return state.urlSignIn;
  },
  getLoginURL: () => {
    return state.urlLogin;
  },
  getUserURL: () => {
    return state.urlUser;
  },
  getOrdersURL: () => {
    return state.urlOrders;
  },
  getRefreshURL: () => {
    return state.urlRefresh;
  },
  getAccess: (state) => {
    return state.Access;
  },
  getUser: (state) => {
    return state.user;
  },
  getUserID: (state) => {
    return state.user.id;
  },
  getOrders: (state) => {
    return state.orders;
  },
  getLoaded: () => {
    return state.loaded;
  },
  getErrors: () => {
    return state.errors;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
