import { createRouter, createWebHistory } from 'vue-router';
import store from '../store/index.js';
import { getUserByCookie } from '../services/users.service.js';

const routes = [
  // Route for shop
  {
    path: '/:pathMatch(.*)*',
    name: '404-page',
    component: () => import('../pages/404Page'),
    meta: { roles: ['customer', 'admin', 'guest'] }
  },
  {
    path: '/',
    name: 'home-page',
    component: () => import('../pages/HomePage'),
    meta: { roles: ['customer', 'guest'] }
  },
  {
    path: '/products',
    name: 'products-page',
    component: () => import('../pages/ProductsCategory'),
    meta: { roles: ['customer', 'guest'] }
  },
  {
    path: '/products/:slug',
    name: 'products-detail-page',
    component: () => import('../pages/ProductDetail'),
    meta: { roles: ['customer', 'guest'] }
  },
  {
    path: '/shopping-cart',
    name: 'shopping-cart-page',
    component: () => import('../pages/ShoppingCart'),
    meta: { requiresAuth: true, roles: ['customer'] }
  },
  {
    path: '/categories/:slug',
    name: 'products-category',
    component: () => import('../pages/ProductsCategory'),
    meta: { roles: ['customer', 'guest'] }
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../pages/LoginPage'),
    meta: { roles: ['customer', 'guest'] }
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../pages/RegisterPage'),
    meta: { roles: ['customer', 'guest'] }
  },
  {
    path: '/products/search',
    name: 'search',
    component: () => import('../pages/SearchPage'),
    meta: { roles: ['customer', 'guest'] }
  },
  {
    path: '/bill',
    name: 'bill',
    component: () => import('../pages/BillPage'),
    meta: { requiresAuth: true, roles: ['customer'] }
  },
  {
    path: '/bill-detail/:id',
    name: 'bill-detail',
    component: () => import('../pages/BillDetailPage'),
    meta: { requiresAuth: true, roles: ['customer'] }
  },
  // Route for admin
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('../admin/pages/Dashboard'),
    meta: { requiresAuth: true, roles: ['admin'] }
  }
];

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes
});

router.beforeEach(async (to, from, next) => {
  try {
    if (!store.getters.isAuthenticated) {
      const res = await getUserByCookie();
      store.dispatch('setUser', res.data);
      if (res.data.role === 'customer') {
        store.dispatch('getCartProduct');
        store.commit('updateLayout', 'shop');
      }
      if (res.data.role === 'admin') {
        store.commit('updateLayout', 'admin');
      }
    }
  } catch (error) {
    next();
  }
  let role = store.state.user ? store.state.user.role : 'guest';
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (role === 'guest') {
      next('/login');
    }
    if (store.getters.isAuthenticated) {
      if (to.meta.roles.includes(role)) {
        next();
      } else {
        next('/404');
      }
    } else {
      next('/login');
    }
  } else {
    if (to.meta.roles.includes(role)) {
      next();
    } else {
      next('/404');
    }
  }
});

export default router;
