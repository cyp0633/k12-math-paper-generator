<script setup>
import { NInput, NButton, useMessage, useDialog, NStatistic, NGrid, NGi } from 'naive-ui';
import { reactive, ref, onMounted } from 'vue';
import { sha256 } from 'js-sha256';
import Global from '../var';
import router from '../router';

const data = reactive({
    oldPassword: ref(null),
    newPassword: ref(null),
    confirmPassword: ref(null),
});

const stats = reactive({
    total: 0,
    correct: 0,
})

onMounted(() => {
    // 检测登录
    if (Global.title.loginUser == null) {
        dialog.warning({
            title: "请先登录",
            content: "你不登录，怎么知道你要修改什么呢？",
            negativeText: "好",
            onNegativeClick: () => {
                router.push("/login");
            },
            onMaskClick: () => {
                router.push("/login");
            }
        })
    }
    // 获取统计数据
    var xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/user/stats", true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                var resp = JSON.parse(xhr.responseText);
                stats.total = resp.total;
                stats.correct = resp.correct;
            } else {
                message.error("获取统计数据失败");
            }
        }
    }
});

const message = useMessage();
const dialog = useDialog();

function changePassword() {
    if (Global.title.loginUser == null) {
        dialog.warning({
            title: "未登录",
            content: "请先登录",
            negativeText: "好",
        })
        return;
    }
    if (data.newPassword == null || data.confirmPassword == null || data.oldPassword == null) {
        dialog.warning({
            title: "信息不完整",
            content: "请填写完整信息",
            negativeText: "好",
        })
        return;
    }
    if (data.newPassword != data.confirmPassword) {
        dialog.warning({
            title: "信息不完整",
            content: "两次输入的密码不一致",
            negativeText: "好",
        })
        return;
    }
    var format = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{6,10}$/.test(data.newPassword)
    if (!format) {
        dialog.warning({
            title: "密码格式错误",
            content: "密码必须包含大小写字母和数字，长度为6-10位",
            negativeText: "好",
        })
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("PATCH", "/api/user", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    var hash1 = sha256.create(), hash2 = sha256.create();
    hash1.update(data.oldPassword);
    hash2.update(data.newPassword);
    xhr.send(JSON.stringify({
        "old_password": hash1.hex(),
        "new_password": hash2.hex()
    }))
    xhr.onreadystatechange = function () {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            switch (xhr.status) {
                case 200:
                    message.success("修改成功");
                    break;
                case 400:
                    message.error("没有登录");
                    break;
                case 401:
                    message.error("旧密码错误");
                    break;
                default:
                    message.error("未知错误");
            }
        }
    }
}

function logout() {
    if (Global.title.loginUser == null) {
        dialog.warning({
            title: "未登录",
            content: "请先登录",
            negativeText: "好",
        })
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/api/user/session", true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            switch (xhr.status) {
                case 200:
                    message.success("注销成功");
                    Global.title.loginUser = null;
                    Global.textChange();
                    router.replace('/');
                    vm.$forceUpdate();
                    break;
                case 401:
                    message.error("未登录");
                    break;
                default:
                    message.error("未知错误");
            }
        }
    }
}
</script>

<template>
    <div class="container mx-auto text-center w-4/6 space-y-10">
        <h1 class="text-5xl">
            个人中心
        </h1>
        <div class="space-y-5 mx-auto text-center">
            <h2 class="text-3xl">历史统计</h2>
            <div class="space-y-5 mx-auto text-center w-1/3 min-w-min">
                <n-grid x-gap="12" :cols="3">
                    <n-gi>
                        <NStatistic label="已做题数" :value="stats.total" />
                    </n-gi>
                    <n-gi>
                        <NStatistic label="正确题数" :value="stats.correct" />
                    </n-gi>
                    <n-gi>
                        <n-statistic label="正确率"
                            :value="stats.total==0?0:Math.round(stats.correct/stats.total*10000)/100+'%'" />
                    </n-gi>
                </n-grid>
            </div>
        </div>
        <div class="container space-y-5">
            <h2 class="text-3xl">修改密码</h2>
            <p>密码需同时包含大、小写字母以及数字，长度 6-10 位。</p>
            <div class="container space-y-2 w-1/4 mx-auto">
                <n-input show-password-on="mousedown" type="password" placeholder="原密码"
                    v-model:value="data.oldPassword" />
                <n-input show-password-on="mousedown" type="password" placeholder="新密码" :maxlength="10"
                    v-model:value="data.newPassword" />
                <n-input show-password-on="mousedown" type="password" placeholder="确认新密码" :maxlength="10"
                    v-model:value="data.confirmPassword" />
                <n-button class="w-1/2 flex mx-auto" @click="changePassword">确认修改</n-button>
            </div>
        </div>
        <div class="container space-y-5">
            <h2 class="text-3xl">退出账户</h2>
            <div class="container space-y-2 w-1/4 mx-auto">
                <n-button class="w-1/2 flex mx-auto text" type="error" text-color="#d03050" @click="logout">确认退出
                </n-button>
            </div>
        </div>
    </div>
</template>