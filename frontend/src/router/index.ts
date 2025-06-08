import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import PlaybackRecorder from '../views/PlaybackRecorder.vue';
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
    component: PlaybackRecorder,
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