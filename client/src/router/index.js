import { createRouter, createWebHistory } from "vue-router";

// import HomePage from "../pages/HomePage.vue";
// import AllProducts from "../pages/AllProducts.vue";
// import ProductDetail from "../pages/ProductDetail.vue";
// import ShoppingCart from "../pages/ShoppingCart.vue";
// import LoginPage from "../pages/LoginPage.vue";

const routes = [
  {
    path: "/",
    name: "home-page",
    component: () => import("../pages/HomePage")
  },
  {
    path: "/products",
    name: "products-page",
    component: () => import("../pages/AllProducts")
  },
  {
    path: "/products/:slug",
    name: "products-detail-page",
    component: () => import("../pages/ProductDetail")
  },
  {
    path: "/shopping-cart",
    name: "shopping-cart-page",
    component: () => import("../pages/ShoppingCart")
  },
  {
    path: "/categories/:category",
    name: "products-category",
    component: () => import("../pages/AllProducts")
  },
  {
    path: "/login",
    name: "login",
    component: () => import("../pages/LoginPage")
  },
  {
    path: "/register",
    name: "register",
    component: () => import("../pages/RegisterPage")
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
