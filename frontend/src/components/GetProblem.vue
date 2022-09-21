<script setup>
import { NRadioGroup, NRadioButton, NSpace, NInput, useMessage, useDialog } from 'naive-ui';
import { useRouter } from 'vue-router';
import { onMounted, ref } from 'vue';
import global from '../var'

const data = {
    timestr: "上午"
};

const router = useRouter();
const message = useMessage();
const dialog = useDialog();

const difficulty = [
    {
        value: 1,
        label: "小学",
    },
    {
        value: 2,
        label: "初中",
    },
    {
        value: 3,
        label: "高中",
    },
];

onMounted(() => { // 自动切换页面上的时间文字
    if (global.title.loginUser == null) {
        dialog.warning({
            title: "请先登录",
            content: "抱歉，需要登录才能做题",
            negativeText: "好",
            onNegativeClick: () => {
                router.push("/login");
            },
            onMaskClick: () => {
                router.push("/login");
            }
        })
    }
    var date = new Date();
    var hour = date.getHours();
    if (hour >= 0 && hour < 12) {
        data.timestr = "上午";
    } else if (hour >= 12 && hour < 18) {
        data.timestr = "下午";
    } else {
        data.timestr = "晚上";
    }
})

function getProblems() {
    if (data.difficulty == null || data.problemNum == null) {
        dialog.warning({
            title: "信息不完整",
            content: "请选择难度和题目数量",
            negativeText: "好",
        })
        return;
    }
    if (isNaN(data.problemNum)) { // 判断是否为数字
        dialog.warning({
            title: "信息不完整",
            content: "题目数量必须为数字",
            negativeText: "好",
        })
        return;
    }
    if (data.problemNum < 10 || data.problemNum > 30) {
        dialog.warning({
            title: "信息不完整",
            content: "题目数量应在 10 与 30 之间",
            negativeText: "好",
        })
        return;
    }
    var params = "level=" + data.difficulty + "&num=" + data.problemNum;
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/paper?" + params, true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var res = JSON.parse(xhr.responseText);
                global.problems = res.data;
                if (res.code == 0) {
                    message.success("生成试卷成功");
                    var i = ref(0);
                    global.answers = []; // 清理答案数组
                    for (i = 0; i < global.problems.length; i++) { // 初始化答案均为 NaN
                        global.answers[i] = NaN;
                    }
                    router.push('/problem/1');
                } else {
                    message.error("生成试卷失败");
                }
            } else {
                message.error("生成试卷失败");
            }
        }
    }
}
</script>

<template>
    <main>
        <div class="container mx-auto text-center w-4/6 space-y-6">
            <h1 class="font-sans font-bold text-5xl">{{data.timestr}}好，{{global.title.loginUser}}!</h1>
            <h1 class="font-sans font-medium text-3xl">请选择您的题目等级与数量</h1>
            <div class="space-y-2">
                <p>难度选择</p>
                <n-space vertical>
                    <n-radio-group v-model:value="data.difficulty" name="difficultyButtonGroup" size="large">
                        <n-radio-button v-for="item in difficulty" :key="item.value" :value="item.value">{{item.label}}
                        </n-radio-button>
                    </n-radio-group>
                </n-space>
            </div>
            <div class="space-y-2 w-1/3 mx-auto">
                <p>题目数量</p>
                <n-input v-model:value="data.problemNum" type="text" placeholder="建议在 1-30 之间" />
            </div>
            <button
                class="rounded-full w-40 h-20 text-3xl text-white bg-gradient-to-r from-purple-400 via-pink-500 to-red-500"
                @click="getProblems">开始生成</button>
        </div>
    </main>
</template>