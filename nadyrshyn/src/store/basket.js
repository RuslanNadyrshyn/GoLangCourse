const state = {
  url: "http://localhost:8080/get_baskets",
  products: JSON.parse(localStorage.getItem("delivery_basket")),
  totalPrice: 0,
  errors: [],
  loaded: false,
};

const mutations = {
  addProduct(state, product) {
    state.products.push(product);
    localStorage.setItem("delivery_basket", JSON.stringify(state.products));
  },
  incCount(state, productId) {
    for (let i = 0; i < state.products.length; i++) {
      if (state.products[i].id === productId) {
        state.products[i].counter++;
        break;
      }
    }
  },
  decCount(state, productId) {
    for (let i = 0; i < state.products.length; i++) {
      if (
        state.products[i].id === productId &&
        state.products[i].counter !== 1
      ) {
        state.products[i].counter--;
        break;
      }
    }
  },
  clearBasket(state) {
    state.products = [];
  },
  deleteProduct(state, productId) {
    for (let i = 0; i < state.products.length; i++) {
      if (state.products[i].id === productId) {
        state.products.splice(i, 1);
        break;
      }
    }
  },
  setTotalPrice(state) {
    let total = 0;
    for (let i = 0; i < state.products.length; i++) {
      total += state.products[i].price * state.products[i].counter;
    }
    state.totalPrice = total.toFixed(2);
    localStorage.setItem("delivery_basket", JSON.stringify(state.products));
  },
  setErrors(state, errors) {
    state.errors = errors;
  },
  setLoaded(state, loaded) {
    state.loaded = loaded;
  },
};

const actions = {
  deleteProduct(context, productId) {
    context.commit("deleteProduct", productId);
  },
  calcTotalPrice(context) {
    context.commit("setTotalPrice");
  },
};

const getters = {
  getBasketURL: (state) => {
    return state.url;
  },
  getBasket: () => {
    return state.products;
  },
  getTotalPrice: (state) => {
    return state.totalPrice;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
