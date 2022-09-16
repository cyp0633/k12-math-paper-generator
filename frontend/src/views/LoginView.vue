<script setup>
import { NInput, NButton, useMessage } from 'naive-ui';
import sha256 from 'js-sha256';

var data = {
    loginForm: {},
}

// 登录
function postLogin() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", "/api/user/session",true); 
    xmlHttp.setRequestHeader("Content-Type", "application/json");
    xmlHttp.withCredentials=true; // 允许携带cookie
    var hash=sha256.create();
    hash.update(data.rawPassword);
    data.loginForm.password=hash.hex(); // 计算密码的 SHA256 哈希值
    console.log(JSON.stringify(data.loginForm));
    xmlHttp.send(JSON.stringify(data.loginForm));
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState == XMLHttpRequest.DONE) { // 等待后期换一个更漂亮的提示框
            const statuscode = xmlHttp.status;
            if (statuscode == 200) {
                alert("登录成功");
            } else {
                alert("登录失败");
            }
        }
    }
}

// 获取登陆状态
function getLoginStatus() {
    var xhr=new XMLHttpRequest();
    xhr.open("GET","/api/user/session",true);
    xhr.send();
    xhr.onreadystatechange=function(){
        if(xhr.readyState==4){
            if(xhr.status==200){
                alert("已登录")
            }else{
                alert("未登录")
            }
        }
    }
}
</script>

<template>
    <main>
        <div class="container mx-auto text-center w-2/6 space-y-6 pt-20">
            <h1 class="font-sans font-medium text-5xl">登录</h1>
            <div class="space-y-2">
                <n-input v-model:value="data.loginForm.username" type="text" placeholder="手机号/邮箱" />
                <n-input v-model:value="data.rawPassword" type="password" placeholder="密码" />
            </div>
            <n-button @click="postLogin">登录</n-button>
            <n-button @click="getLoginStatus">查询登录状态</n-button>
        </div>
    </main>
</template> 