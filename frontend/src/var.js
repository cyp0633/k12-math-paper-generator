import { reactive, ref } from "vue"

var problems = [];
var answers = [];

var title = reactive({
    text: "未登录，点击登录",
    to: "/login",
    loginUser: null,
    showRegister: true,
})

function textChange() {
    if (title.loginUser == null) {
        title.text = "未登录，点击登录";
        title.to = "/login";
        title.showRegister = true;
    }
    else {
        title.text = title.loginUser + " 的个人空间";
        title.to = "/profile";
        title.showRegister = false;
    }
}

export default { problems, answers, textChange, title }
