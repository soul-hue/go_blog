import {
    createApp
} from "vue";
// import './style.css'
import App from "../App.vue";
// const app = createApp(App)
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

import * as VueRouter from 'vue-router'

const routes = [
    {
        path :  '/',
        redirect : '/login'

    },
    {
        path : '/login',
        component : Login,
        // direct : '/'
    },
    {
        path : '/admin',
        component : Admin
    }
]

const router = VueRouter.createRouter({
    history: VueRouter.createWebHashHistory(),
    routes,
})

export default router