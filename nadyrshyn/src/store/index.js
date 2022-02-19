import Vue from "vue";
import Vuex from "vuex";
import productState from "./product";
import supplierState from "./supplier";

Vue.use(Vuex);

const modules = {
  products: productState,
  suppliers: supplierState,
};

export default new Vuex.Store({
  modules,
});
