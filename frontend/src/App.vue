<script setup>
import { VueElement } from 'vue';
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import 'katex/dist/katex.min.css';

</script>

<template>
    <div>
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
                            <!-- <a class="inline-block py-2 px-4 text-white no-underline" @click="router.push('/home')">主页</a> -->
                            <RouterLink class="inline-block py-2 px-4 text-white no-underline" to="/home">主页
                            </RouterLink>
                        </li>
                        <li class="mr-3">
                            <RouterLink
                                class="inline-block text-gray-600 no-underline hover:text-gray-200 hover:text-underline py-2 px-4"
                                to="/about">关于</RouterLink>
                        </li>
                        <li class="mr-3">
                            <RouterLink
                                class="inline-block text-gray-600 no-underline hover:text-gray-200 hover:text-underline py-2 px-4"
                                to="/register">注册</RouterLink>
                        </li>
                        <li class="mr-3">
                            <RouterLink class="inline-block text-gray-600 no-underline hover:text-gray-200 hover:text-underline py-2 px-4"
                                to="/login">登录</RouterLink>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
        <div class="container shadow-lg mx-auto mt-24 md:mt-16 h-screen">
            <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

            <div class="wrapper">
                <HelloWorld msg="You did it!" />
                <div v-katex:auto class="h-8">
                    \(\frac{a_i}{1+x}\)
                </div>
                <nav>
                    <RouterLink to="/">Home</RouterLink>
                    <RouterLink to="/about">About</RouterLink>
                </nav>
            </div>
            <RouterView />
        </div>

    </div>
</template>

<style scoped>
@import "../node_modules/katex/dist/katex.min.css";
/* header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
} */
</style>
