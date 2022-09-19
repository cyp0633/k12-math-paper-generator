<script setup>
import { NH1, NH2, NH3, NA } from 'naive-ui';
</script>

<template>
    <div class="container mx-auto space-y-6 pt-10">
        <n-h1 prefix="bar">开放源代码声明</n-h1>
        <div class="container mx-auto space-y-2">
            <n-h2 prefix="bar">本项目</n-h2>
            <p>本软件在 MIT 协议下开源，源代码地址：<n-a href="https://github.com/cyp0633/k12-math-paper-generator">GitHub</n-a>
            </p>
        </div>
        <n-h2 prefix="bar">后端依赖</n-h2>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">Gin</n-h3>
            <p>本项目后端使用 Gin Web 框架，Gin 遵循 MIT 协议。源代码地址：<n-a href="https://github.com/gin-gonic/gin">GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">alibaba-cloud-sdk-go</n-h3>
            <p>阿里云 SDK 用于发送短信，遵循 Apache 协议。源代码地址：<n-a href="https://github.com/aliyun/alibaba-cloud-sdk-go">GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">go-sqlite</n-h3>
            <p>go-sqlite 用于读写数据库，遵循 BSD 协议。源代码地址：<n-a href="https://github.com/glebarez/go-sqlite">GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">ini</n-h3>
            <p>ini 用于读取配置文件，遵循 Apache 协议。源代码地址：<n-a href="https://github.com/go-ini/ini">GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">gorm</n-h3>
            <p>gorm 用于进行对象关系转换，遵循 MIT 协议。源代码地址：<n-a href="https://github.com/go-gorm/gorm">GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">gomail</n-h3>
            <p>gomail 用于收发邮件，遵循 MIT 协议。源代码地址：<n-a href="https://github.com/go-gomail/gomail">GitHub</n-a>
            </p>
        </div>
        <n-h2 prefix="bar">前端依赖</n-h2>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">Vue 3</n-h3>
            <p>Vue 3 是本项目前端的主要框架。Vue 3 遵循 MIT 协议，源代码地址：<n-a href="https://github.com/vuejs/vue">GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">Naive UI</n-h3>
            <p>本项目使用 High Level 的 Naive UI 作为组件库，能提高姿势水平。Naive UI 遵循 MIT 协议，源代码地址：<n-a
                    href="https://github.com/tusen-ai/naive-ui">
                    GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">vue3-katex</n-h3>
            <p>做题页面使用 vue3-katex 渲染数学公式。vue3-katex 遵循 MIT 协议，源代码地址：<n-a href="https://github.com/hsorby/vue3-katex">
                    GitHub</n-a>
            </p>
        </div>
        <div class="container mx-auto space-y-2">
            <n-h3 prefix="bar">Tailwind CSS</n-h3>
            <p>网站使用 Tailwind CSS 实现简化的样式编辑。Tailwind CSS 遵循 MIT 协议，源代码地址：<n-a
                    href="https://github.com/tailwindlabs/tailwindcss">GitHub</n-a>
            </p>
            <p>另有许多使用的开源项目因间接引用等原因无法详述，在此感谢。</p>
        </div>
    </div>
</template>