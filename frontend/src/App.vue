<script setup>
import { onUpdated, onMounted, reactive, ref } from 'vue';
import { RouterLink, RouterView } from 'vue-router'
import { NDivider } from 'naive-ui';
import Global from './var.js';
import 'katex/dist/katex.min.css';

const data = reactive({
    text: ref("未登录，点击登录"),
    to: ref("/login"),
    showRegister: ref(true),
})

onMounted(() => {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/user/session", true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var resp = JSON.parse(xhr.responseText);
                Global.user = resp.user;
            } else {
                Global.user = ref(null);
                data.text = "未登录，点击登录";
            }
        }
        textChange();
    }
})

onUpdated(() => {
    textChange();
})

function textChange() {
    if (Global.user == null) {
        data.text = "未登录，点击登录";
        data.to = "/login";
        data.showRegister = true;
    }
    else {
        data.text = Global.user + " 的个人空间";
        data.to = "/profile";
        data.showRegister = false;
    }
}

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
                        <li class="mr-3" v-show="data.showRegister">
                            <RouterLink class="inline-block py-2 px-4 text-white no-underline" to="/register">注册
                            </RouterLink>
                        </li>
                        <li class="mr-3">
                            <p class="inline-block py-2 px-4 text-white no-underline">
                                <RouterLink :to="data.to">{{data.text}}</RouterLink>
                            </p>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </header>
    <main>
        <div class="container mx-auto mt-28 md:mt-16 h-screen ">
            <!-- <n-notification-provider> -->
            <RouterView />
            <!-- </n-notification-provider> -->
        </div>
    </main>
    <footer>
        <n-divider />
        <RouterLink to="/opensource" class="float-right m-3">开放源代码声明</RouterLink>
    </footer>
</template>

<style scoped>
@import "../node_modules/katex/dist/katex.min.css";
</style>
