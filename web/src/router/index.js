import Vue from 'vue';
import Router from 'vue-router';
import Dashboard from '@/views/Dashboard';
import Profile from '@/views/Profile';
import Tables from '@/views/Tables';
import Maps from '@/views/Maps';
import NotFound from '@/views/NotFound';

import { SiteName, TitleSeparator, RouterMode } from '../config';

Vue.use(Router);

const router = new Router({
  mode: RouterMode,
  routes: [
    {
      path: '/',
      name: 'Dashboard',
      component: Dashboard,
      props: { page: 1 },
      meta: {
        title: 'Dashboard'
      }
    },
    {
      path: '/profile',
      name: 'Profile',
      props: { page: 2 },
      component: Profile,
      meta: {
        title: 'Dashboard'
      }
    },
    {
      path: '/tables',
      name: 'Tables',
      props: { page: 3 },
      component: Tables,
      meta: {
        title: 'Dashboard'
      }
    },
    {
      path: '/maps',
      name: 'Maps',
      props: { page: 4 },
      component: Maps,
      meta: {
        title: 'Dashboard'
      }
    },
    {
      path: '*',
      name: 'NotFound',
      props: { page: 5 },
      component: NotFound,
      meta: {
        title: 'Dashboard'
      }
    }
  ]
});

router.beforeEach((to, from, next) => {
  document.title = to.meta.title + TitleSeparator + SiteName;
  next();
});

export default router;
