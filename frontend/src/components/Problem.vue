<script setup>
import { onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { NButton } from 'naive-ui';
import global from '../var';

const route = useRoute();

var id = route.params.id;


</script>

<template>
    <main>
        <div class="container mx-auto text-center w-4/6 space-y-6 pt-20">
            <h1 class="text-4xl">
                第 {{id}} 题
            </h1>
            <div>
                <katex-element :expression="global.problems[id-1]" class="text-4xl" />
            </div>
            <n-button @click="test">测试</n-button>
        </div>
    </main>
</template>
