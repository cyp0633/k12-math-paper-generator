import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Vue3Katex from '@hsorby/vue3-katex';
import 'katex/dist/katex.min.css';

const app = createApp(App)

app.use(Vue3Katex, {
    globalOptions: {
        //... Define globally applied KaTeX options here
    }
})

app.use(router)

app.mount('#app')
