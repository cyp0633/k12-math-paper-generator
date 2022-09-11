<script setup>
    import {NH1} from 'naive-ui'

    const data=[
        {
            content:"abc",
            number:"1",
        },
        {
            content:"def",
            number:"2",
        },
        {
            content:"ghi",
            number:"3",
        },
    ]
</script>

<template>
    <div class="test w-40">
        <n-h1>测试</n-h1>
        <div v-for="item in data" :key="item.number">
            
        </div>
    </div>
</template>