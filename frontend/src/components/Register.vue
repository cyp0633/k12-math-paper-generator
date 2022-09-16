<script setup>
    import { NInput, NButton } from 'naive-ui';
    import sha256 from 'js-sha256';
    
    const data = {
        sent: false,
    }
    
    function sendVerification() {
        if (data.sent) {
            alert("阿里云钱不多了，请不要重复发送");
            return;
        }
        if (data.username == null) {
            alert("请输入邮箱或手机号");
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
                        alert("验证码已发送");
                        data.sent = true;
                        break;
                    case 400:
                        alert("手机号/邮箱格式错误");
                        break;
                    case 409:
                        alert("手机号/邮箱已被注册");
                        break;
                    default:
                        alert("未知错误");
                }
            }
        }
    }
    
    function tryRegister() {
        if (!data.sent) {
            alert("请先发送验证码");
            return;
        }
        if (data.username == null || data.verification == null || data.password == null) {
            alert("请填写完整信息");
            return;
        }
        var xhr = new XMLHttpRequest();
        xhr.open("PUT", "/api/user", true);
        xhr.setRequestHeader("Content-Type", "application/json");
        var hash = sha256.create(); // 加密密码
        hash.update(data.password);
        xhr.send(JSON.stringify({ "username": data.username, "verify": data.verification, "password": hash.hex() }))
    }
    </script>
    
    <template>
        <main>
            <main>
                <div class="container mx-auto text-center w-2/6 space-y-6 pt-20">
                    <h1 class="font-sans font-medium text-5xl">注册</h1>
                    <p class="font-size-xl">请输入您的手机号或邮箱，我们会将验证码发送到此处。</p>
                    <div class="space-y-2">
                        <n-input v-model:value="data.username" type="text" placeholder="手机号/邮箱" />
                        <div class="flex space-x-2">
                            <n-input v-model:value="data.verification" type="text" placeholder="验证码" />
                            <n-button class="flex w-2/6" @click="sendVerification">发送验证码</n-button>
                        </div>
                        <n-input v-model:value="data.password" type="password" placeholder="密码" />
                    </div>
                    <n-button @click="tryRegister" size="large" type="success">注册</n-button>
                </div>
            </main>
        </main>
    </template> 