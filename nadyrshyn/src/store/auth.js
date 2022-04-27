const state = {
  urlLogin: "http://localhost:8080/login",
  urlSignIn: "http://localhost:8080/sign_in",
  Access: false,
  userId: 0,
  accessToken: "",
  refreshToken: "",
  errors: [],
  loaded: false,
};

const mutations = {
  addAccessToken(state, token) {
    state.accessToken = token;
    // state.userId = token.uid;
  },
  addRefreshToken(state, token) {
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
    console.log(tokens.access_token.id);
    console.log(tokens.refresh_token);
    console.log("id:", tokens.access_token.uid);
    context.commit("setAccess", true);
  },
};

const getters = {
  getAccessToken: () => {
    return state.accessToken;
  },
  getRefreshToken: () => {
    return state.refreshToken;
  },
  getAccess: (state) => {
    return state.Access;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
