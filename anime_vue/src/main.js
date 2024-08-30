import { createApp } from 'vue';
import App from './App.vue';
import axios from 'axios';

// 创建 Vue 应用实例
const app = createApp(App);

// 如果你想全局使用 axios，可以将其添加到全局属性中
app.config.globalProperties.$axios = axios;

// 挂载应用到指定的 DOM 元素
app.mount('#app');
