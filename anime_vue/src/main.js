import { createApp } from 'vue';
import App from './App.vue';
import router from './router/'
import axios from 'axios';

const app = createApp(App);

// 全局配置 Axios
app.config.globalProperties.$axios = axios;

// 使用 Vue Router
app.use(router);

app.mount('#app');
