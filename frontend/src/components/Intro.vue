<script setup>
import { NH1, NButton } from 'naive-ui'
</script>

<template>
    <main>
        <div class="container mx-auto text-center pt-40 space-y-7">
            <h1
                class="font-sans font-extrabold text-7xl text-transparent bg-clip-text bg-gradient-to-r from-purple-400 to-pink-600">
                欢迎使用数学学习平台</h1>
            <p class="text-2xl">请先登录/注册，支持使用手机号与邮箱。</p>
            <div class="content-center place-content-center grid-cols-2 gap-5 flex">
                <n-button @click="$router.push('/register')">注册</n-button>
                <n-button @click="$router.push('/login')">登录</n-button>
            </div>
        </div>
    </main>
</template>