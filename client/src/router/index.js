import { createRouter, createWebHistory } from "vue-router";

import HomePage from "../pages/HomePage.vue";
import AllProducts from "../pages/AllProducts.vue";
import ProductDetail from "../pages/ProductDetail.vue";
import ShoppingCart from "../pages/ShoppingCart.vue";
import LoginPage from "../pages/LoginPage.vue";

const routes = [
  {
    path: "/",
    name: "home-router",
    component: HomePage
  },
  {
    path: "/products",
    name: "products-router",
    component: AllProducts
  },
  {
    path: "/products/:id",
    name: "products-detail-router",
    component: ProductDetail
  },
  {
    path: "/shopping-cart",
    name: "shopping-cart-router",
    component: ShoppingCart
  },
  {
    path: "/categories/:category",
    name: "products-category",
    component: AllProducts
  },
  {
    path: "/login",
    name: "login",
    component: LoginPage
  }
];

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes
});

export default router;
