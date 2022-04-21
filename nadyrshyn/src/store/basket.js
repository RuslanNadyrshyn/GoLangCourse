const state = {
  url: "http://localhost:8080/get_basket",
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
  clearBasket(state) {
    state.products = [];
  },
  deleteProduct(state, productId) {
    state.products.splice(productId, 1);
  },
  updateProduct(state, prod) {
    for (let i = 0; i < state.products.length; i++)
      if (state.products[i].id === prod.id) {
        state.products[i] = prod;
        break;
      }
  },
  setTotalPrice(state, total) {
    state.totalPrice = total;
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
  addProduct(context, product) {
    let inBasket = false;
    for (let i = 0; i < state.products.length; i++)
      if (state.products[i].id === product.id) {
        state.products[i].counter++;
        inBasket = true;
        break;
      }
    if (inBasket === false) {
      product.counter++;
      context.commit("addProduct", product);
    }
  },
  deleteProduct(context, productId) {
    for (let i = 0; i < context.state.products.length; i++)
      if (context.state.products[i].id === productId) {
        context.commit("deleteProduct", i);
        break;
      }
  },
  calcTotalPrice(context) {
    let total = 0;
    for (let i = 0; i < state.products.length; i++)
      total += state.products[i].price * state.products[i].counter;

    total = total.toFixed(2);
    context.commit("setTotalPrice", total);
  },
  updateProduct(context, prod) {
    context.commit("updateProduct", prod);
  },
};

const getters = {
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
