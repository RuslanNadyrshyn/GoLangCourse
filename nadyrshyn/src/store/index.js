import Vue from "vue";
import Vuex from "vuex";
import productState from "./product";
import supplierState from "./supplier";
import basketState from "./basket";
import orderState from "./order";

Vue.use(Vuex);

const modules = {
  products: productState,
  suppliers: supplierState,
  basket: basketState,
  orders: orderState,
};

export default new Vuex.Store({
  modules,
});
