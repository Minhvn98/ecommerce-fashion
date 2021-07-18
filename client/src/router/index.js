import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'home-page',
    component: () => import('../pages/HomePage')
  },
  {
    path: '/products',
    name: 'products-page',
    component: () => import('../pages/ProductsCategory')
  },
  {
    path: '/products/:slug',
    name: 'products-detail-page',
    component: () => import('../pages/ProductDetail')
  },
  {
    path: '/shopping-cart',
    name: 'shopping-cart-page',
    component: () => import('../pages/ShoppingCart')
  },
  {
    path: '/categories/:slug',
    name: 'products-category',
    component: () => import('../pages/ProductsCategory')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../pages/LoginPage')
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../pages/RegisterPage')
  },
  {
    path: '/products/search',
    name: 'search',
    component: () => import('../pages/SearchPage')
  },
  {
    path: '/bill',
    name: 'bill',
    component: () => import('../pages/BillPage')
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
