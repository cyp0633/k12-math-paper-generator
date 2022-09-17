<script setup>
import { NButton, NSpin, NProgress } from 'naive-ui';
import { ref, reactive, onMounted } from 'vue';
import Global from '../var';

const data = reactive({
    seen: 0,
    score: 0,
})

onMounted(() => { // 提交答案，获取成绩
    if(Global.answers==[]){
        alert("请先完成试卷");
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/paper", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({ "answers": Global.answers }));
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                res = JSON.parse(xhr.responseText);
                data.seen = 0;
                data.score = res.score;
                Global.answers = [];
            }
            else {
                alert("提交失败");
            }
        }
    }
})
</script>

<template>
    <main>
        <div class="container mx-auto text-center space-y-6 pt-20">
            <div v-show="data.seen" class="space-y-2">
                <n-spin size="large" />
                <p class="text-lg">正在提交答案并获取结果</p>
            </div>
            <div v-show="!data.seen" class="space-y-6">
                <p class="text-xl">您的得分为</p>
                <div class="grid-cols-2 space-x-3 space-y-3 justify-center">
                    <n-progress type="circle" :percentage="data.score" status="success" class="flex-shrink" />
                    <p class="text-6xl flex-shrink">{{data.score}}</p>
                    <n-button @click="$router.push('/getproblem')">再来一套</n-button>
                    <n-button @click="$router.push('/')">回到主页</n-button>
                </div>

            </div>
        </div>
    </main>
</template>