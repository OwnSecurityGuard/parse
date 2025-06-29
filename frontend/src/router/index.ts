import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Monitor from '../views/Monitor.vue';
import ConfigPage from '../views/ConfigPage.vue';

const routes: RouteRecordRaw[] = [
  // {
  //   path: '/',
  //   name: 'root',
  //   redirect: '/recorder',
  //   component: undefined, // 明确指定 component 为 undefined
  // },
  {
    path: '/recorder',
    name: 'Recorder',
    component: Monitor,
  },
  {
    path: '/config',
    name: 'Config',
    component: ConfigPage,
  }
];

export const router = createRouter({
  history: createWebHistory(),
  routes
});