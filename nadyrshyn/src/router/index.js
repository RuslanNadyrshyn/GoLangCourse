import Vue from "vue";
import VueRouter from "vue-router";
import ProductList from "@/views/ProductList";
import ProductView from "@/views/ProductView";
import BasketView from "@/views/BasketView";

Vue.use(VueRouter);

const routes = [
  {
    path: "/products",
    name: "Products",
    component: ProductList,
  },
  {
    path: "/product/:id",
    name: "Product",
    component: ProductView,
  },
  {
    path: "/basket",
    name: "Basket",
    component: BasketView,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
