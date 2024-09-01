import { createRouter, createWebHistory } from 'vue-router';
import Home from '../components/HomePage.vue';
import LoginForm from '../components/LoginForm.vue';
import RegisterForm from '../components/RegisterForm.vue';
import VideoPlay from '../components/VideoPlay.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/login',
        name: 'Login',
        component: LoginForm,
    },
    {
        path: '/regist',
        name: 'Regist',
        component: RegisterForm,
    },
    {
        path: '/videoplay/:id',
        name: 'VidePlay',
        component: VideoPlay,
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
