<script setup>
import { NButton } from 'naive-ui'
import Global from '../var'
</script>

<template>
    <main>
        <div class="container mx-auto text-center space-y-7 pt-10">
            <h1
                class="font-sans font-extrabold text-7xl text-transparent bg-clip-text bg-gradient-to-r from-purple-400 to-pink-600">
                欢迎使用数学学习平台</h1>
            <div class="container mx-auto text-center space-y-7" v-show="Global.title.showRegister">
                <p class="text-2xl">请先登录/注册，支持使用手机号与邮箱。</p>
                <div class="content-center place-content-center grid-cols-2 gap-5 flex">
                    <n-button @click="$router.push('/register')">注册</n-button>
                    <n-button @click="$router.push('/login')">登录</n-button>
                </div>
            </div>
            <div class="container mx-auto text-center space-y-7" v-show="!Global.title.showRegister">
                <p class="text-2xl">您已登录，点击下方按钮开始学习。</p>
                <div class="content-center place-content-center grid-cols-2 gap-5 flex">
                    <n-button @click="$router.push('/getproblem')">开始学习</n-button>
                </div>
            </div>
        </div>
    </main>
</template>