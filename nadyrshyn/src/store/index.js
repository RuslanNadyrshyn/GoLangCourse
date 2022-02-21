import Vue from "vue";
import Vuex from "vuex";
import productState from "./product";
import supplierState from "./supplier";
import basketState from "./basket";

Vue.use(Vuex);

const modules = {
  products: productState,
  suppliers: supplierState,
  basket: basketState,
};

export default new Vuex.Store({
  modules,
});
