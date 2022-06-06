import Vue from "vue";
import VueRouter from "vue-router";
import ProductList from "@/views/ProductList";
import ProductView from "@/views/ProductView";
import BasketView from "@/views/BasketView";
import OrderView from "@/views/OrderView";
import LoginView from "@/views/LoginView";
import LoadedOrderView from "@/views/LoadedOrderView";
import UserView from "@/views/UserView";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
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
  {
    path: "/order",
    name: "Order",
    component: OrderView,
  },
  {
    path: "/order/:id",
    name: "LoadedOrder",
    component: LoadedOrderView,
  },
  {
    path: "/login",
    name: "Login",
    component: LoginView,
  },
  {
    path: "/user/:id",
    name: "User",
    component: UserView,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
