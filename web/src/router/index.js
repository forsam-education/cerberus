import Vue from 'vue';
import Router from 'vue-router';

import { SiteName, TitleSeparator, RouterMode } from '@Config';

import Login from '@View/Login.vue';

Vue.use(Router);

const router = new Router({
  mode: RouterMode,
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
      meta: {
        title: 'Login',
      },
    },
  ],
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title + TitleSeparator + SiteName;
  next();
});

export default router;
