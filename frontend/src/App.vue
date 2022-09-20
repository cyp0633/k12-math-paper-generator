<script setup>
import { watch, onUpdated, onMounted, reactive, ref } from 'vue';
import { RouterLink, RouterView } from 'vue-router'
import { NDivider, NMessageProvider } from 'naive-ui';
import Global from './var.js';
import 'katex/dist/katex.min.css';

// const data = reactive({
//     text: ref("未登录，点击登录"),
//     to: ref("/login"),
//     showRegister: ref(true),
// })

onMounted(() => {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/user/session", true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var resp = JSON.parse(xhr.responseText);
                Global.title.loginUser = resp.user;
            } else {
                Global.title.loginUser = null;
            }
            Global.textChange();
        }

    }
})

onUpdated(() => {
    Global.textChange();
})

watch(() => Global.title.loginUser, async () => {
    Global.textChange();
}, { deep: true })
</script>

<template>
    <header>
        <nav class="bg-gray-800 p-2 mt-0 fixed w-full z-10 top-0">
            <div class="container mx-auto flex flex-wrap items-center">
                <div class="flex w-full md:w-1/2 justify-center md:justify-start text-white font-extrabold">
                    <a class="text-white no-underline hover:text-white hover:no-underline" href="#">
                        <span class="text-2xl pl-2"><i class="em em-grinning"></i>数学学习平台</span>
                    </a>
                </div>
                <div class="flex w-full pt-2 content-center justify-between md:w-1/2 md:justify-end">
                    <ul class="list-reset flex justify-between flex-1 md:flex-none items-center">
                        <li class="mr-3">
                            <RouterLink class="inline-block py-2 px-4 text-white no-underline" to="/">主页
                            </RouterLink>
                        </li>
                        <li class="mr-3">
                            <RouterLink class="inline-block py-2 px-4 text-white no-underline" to="/getproblem">学习
                            </RouterLink>
                        </li>
                        <li class="mr-3" v-show="Global.title.showRegister">
                            <RouterLink class="inline-block py-2 px-4 text-white no-underline" to="/register">注册
                            </RouterLink>
                        </li>
                        <li class="mr-3">
                            <p class="inline-block py-2 px-4 text-white no-underline">
                                <RouterLink :to="Global.title.to">{{Global.title.text}}</RouterLink>
                            </p>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </header>
    <main>
        <div id="mainframe" class="min-h-screen py-16 ">
            <div class="container mx-auto mt-28 md:mt-16 pt-16 pb-8 acrylic">
                <n-message-provider>
                    <RouterView />
                </n-message-provider>
                <n-divider />
                <RouterLink to="/opensource" class="float-right m-3">开放源代码声明</RouterLink>
            </div>
        </div>
    </main>
</template>

<style scoped>
@import "../node_modules/katex/dist/katex.min.css";

#mainframe {
    background-image: url(./assets/catalina.jpg);
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
    background-attachment: fixed;
}

.acrylic {
    border-radius: 8px;
    border-width: thin;
    border-style: solid;
    border-color: #e5e5e5;
    -webkit-backdrop-filter: blur(40px);
    backdrop-filter: blur(40px);
    background-color: #F7F7F7CC;
    overflow-y: auto;
    background-image: url('./assets/noise.png');
    background-repeat: repeat;
    box-shadow: 0px 25.6px 57.6px rgb(0 0 0 / 14%), 0px 0px 16.4px rgb(0 0 0 / 12%);
}

/* set the background image of the whole website to ./assets/catalina.jpg */
/* * {
    background-image: url(./assets/catalina.jpg);
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
    background-attachment: fixed;
} */
</style>