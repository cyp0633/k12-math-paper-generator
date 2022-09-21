<script setup>
import { NInput, NButton, useMessage, useDialog } from 'naive-ui';
import { reactive, ref } from 'vue';
import sha256 from 'js-sha256';

const data = reactive({
    sent: false,
    password: ref(null),
    password2: ref(null),
})

const message = useMessage();
const dialog = useDialog();

function sendVerification() {
    if (data.sent) {
        message.info("阿里云钱不多了，请不要重复发送");
        return;
    }
    if (data.username == null) {
        message.error("请输入邮箱或手机号");
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/user", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({ "username": data.username }))
    xhr.onreadystatechange = function () {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            switch (xhr.status) {
                case 200:
                    message.success("验证码已发送");
                    data.sent = true;
                    break;
                case 400:
                    message.error("手机号/邮箱格式错误");
                    break;
                case 409:
                    message.error("手机号/邮箱已被注册");
                    break;
                case 500:
                    dialog.error({
                        title: "验证码发送失败",
                        content: "请检查邮箱/手机号是否正确\n特别注意：未在阿里云登记过的手机号无法接收验证码",
                        negativeText: "好",
                    })
                default:
                    message.error("未知错误");
            }
        }
    }
}

function tryRegister() {
    if (!data.sent) {
        dialog.warning({
            title: "未发送验证码",
            content: "请先发送验证码",
            negativeText: "好",
        })
        return;
    }
    if (data.username == null || data.verification == null || data.password == null || data.password2 == null) {
        dialog.error({
            title: "信息不完整",
            content: "请填写用户名、验证码和两次确认密码",
            negativeText: "好",
        })
        return;
    }
    if (data.password != data.password2) {
        dialog.error({
            title: "两次密码不一致",
            content: "请重新输入",
            negativeText: "好",
        })
        return;
    }
    var format = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{6,10}$/.test(data.password);
    if (!format) {
        dialog.error({
            title: "密码格式错误",
            content: "密码必须包含大小写字母和数字，长度为6-10位",
            negativeText: "好",
        })
        return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/user", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    var hash = sha256.create(); // 加密密码
    hash.update(data.password);
    xhr.send(JSON.stringify({ "username": data.username, "verify": data.verification, "password": hash.hex() }))
    xhr.onreadystatechange = function () {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            switch (xhr.status) {
                case 200:
                    message.success("注册成功，请登录");
                    window.location.href = "/login";
                    break;
                case 400:
                    dialog.error({
                        title: "验证码错误",
                        content: "请重新输入",
                        negativeText: "好",
                    })
                    break;
                default:
                    message.error("未知错误");
            }
        }
    }
}
</script>
    
<template>
    <main>
        <main>
            <div class="container mx-auto text-center w-2/6 space-y-6">
                <h1 class="font-sans font-medium text-5xl">注册</h1>
                <p class="font-size-xl">请输入您的手机号或邮箱，我们会将验证码发送到此处。<br>密码需同时包含大、小写字母以及数字，长度 6-10
                    位。<br>短信服务仅限对应账户的测试手机号使用，其他情况请输入邮箱。</p>
                <div class="space-y-2">
                    <n-input v-model:value="data.username" type="text" placeholder="手机号/邮箱" />
                    <div class="flex space-x-2">
                        <n-input v-model:value="data.verification" type="text" placeholder="验证码" />
                        <n-button class="flex w-2/6" @click="sendVerification">发送验证码</n-button>
                    </div>
                </div>
                <div class="space-y-2">
                    <n-input v-model:value="data.password" type="password" show-password-on="mousedown"
                        placeholder="密码" />
                    <n-input v-model:value="data.password2" type="password" show-password-on="mousedown"
                        placeholder="确认密码" />
                </div>
                <n-button @click="tryRegister" size="large">注册</n-button>
            </div>
        </main>
    </main>
</template>