import * as Vue from "vue";
import './public/css/base.less'
// import './style.css'
import App from "./App.vue";
import Router from "./router";
import './plugins/elementui'
import axios from "axios";
import VueAxios from 'vue-axios'
const app = Vue.createApp(App)
const API = axios.create({
	baseURL:'http://localhost:3000/api/v1', //请求后端数据的基本地址，自定义
	timeout: 2000,                   //请求超时设置，单位ms
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8',
      }, // 跨域请求时是否需要使用凭证
})
// axios.defaults.baseURL = ""

app.config.globalProperties.$Api = API

app
    .use(Router)


app.mount("#app");