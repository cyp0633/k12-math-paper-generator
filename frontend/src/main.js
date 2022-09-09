import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import './assets/main.css'

import Vue3Katex from '@hsorby/vue3-katex';

const app = createApp(App).use(Vue3Katex, {
    globalOptions: {
        //... Define globally applied KaTeX options here
    }
});

app.use(router)

app.mount('#app')
