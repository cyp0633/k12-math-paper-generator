<script setup>
import { NButton, NSpin, NProgress, useMessage, useDialog } from 'naive-ui';
import { reactive, onMounted } from 'vue';
import Global from '../var';

const data = reactive({
    seen: 0,
    score: 0,
})

const message = useMessage();
const dialog = useDialog();

onMounted(() => { // 提交答案，获取成绩
    if (Global.answers == []) {
        message.info("请先完成试卷");
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/paper", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({ "answers": Global.answers }));
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var res;
                res = JSON.parse(xhr.responseText);
                data.seen = 0;
                data.score = res.score;
                Global.answers = [];
            }
            else {
                dialog.error({
                    title: "未知错误",
                    content: "请联系管理员",
                    negativeText: "好",
                })
            }
        }
    }
})
</script>

<template>
    <main>
        <div class="container mx-auto text-center space-y-6">
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