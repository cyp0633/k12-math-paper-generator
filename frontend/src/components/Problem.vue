<script setup>
import { watch, ref, onUpdated, reactive, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { NButton, NRadioGroup, NRadioButton, NPagination, NSpace, useMessage, useDialog } from 'naive-ui';
import global from '../var';
import router from '../router';

const route = useRoute();

var id = reactive(route.params.id);

const message = useMessage();
const dialog = useDialog();

const data = reactive({
    id: ref(id),
    insist: ref(0),
})

onUpdated(() => {
    id = route.params.id;
});

watch(() => data.page,
    async (newPage) => {
        router.replace('/problem/' + newPage)
        data.id = newPage
    })

onMounted(() => {
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
})

function handin() {
    var i = 0;
    for (i = 0; i < global.problems.length; i++) {
        if (global.answers[i] == NaN && data.insist == 0) {
            message.info("未完成所有题目，如坚持交卷，请再次点击交卷");
            data.insist = 1;
            return;
        }
    }
    router.replace('/result')

}

</script>

<template>
    <main>
        <div class="container mx-auto text-center w-4/6 space-y-6">
            <h1 class="text-4xl">
                第 {{data.id}} 题
            </h1>
            <div>
                <katex-element :expression="global.problems[data.id-1].problem_str" class="text-4xl" />
            </div>
            <div class="space-y-2">
                请选择结果<br>
                <n-radio-group v-model:value="global.answers[data.id-1]" name="result" size="large">
                    <n-radio-button v-for="problem in global.problems[data.id-1].options" :key="problem"
                        :value="problem">
                        {{problem}}</n-radio-button>
                </n-radio-group>
            </div>
            <n-space justify="center">
                <n-pagination v-model:page="data.page" :page-count="global.problems.length" class="flex" />
            </n-space>
            <n-button @click="handin" size="large">交卷</n-button>
        </div>
    </main>
</template>
