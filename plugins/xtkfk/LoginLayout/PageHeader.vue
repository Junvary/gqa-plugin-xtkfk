<template>
    <header>
        <div class="logo">
            <GqaAvatar size="xl" :src="gqaFrontend.logo" />
            <span style="margin-left: 10px">
                {{ gqaFrontend.mainTitle }}
            </span>
        </div>
        <ul>
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-banner')" label="首页" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-news')" label="最新要闻" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-project')" label="项目进度" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-weapon')" label="武器库" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-honour')" label="荣誉认证" />
            <q-btn flat rounded push size="lg" text-color="white" @click="handleScroll('gqa-resource')" label="资源查阅" />
            <q-btn flat rounded push size="lg" text-color="white" @click="openBlankAbbreviate" label="简称大全" />
            <q-btn flat rounded push size="lg" text-color="white" label="版本信息">
                <GqaVersion />
            </q-btn>
        </ul>
    </header>
</template>

<script setup>
import GqaVersion from 'src/components/GqaVersion'
import GqaAvatar from 'src/components/GqaAvatar'
import { onBeforeUnmount, onMounted } from 'vue';
import useCommon from 'src/composables/useCommon';
import { useRouter } from 'vue-router';

const router = useRouter()
const { gqaFrontend } = useCommon()

const flow = () => {
    // 监视下滑，改变导航栏
    window.addEventListener('scroll', () => {
        var header = document.querySelector('header')
        header.classList.toggle('sticky', window.scrollY > 0)
    })
}
const handleScroll = (ele) => {
    document.getElementById(ele).scrollIntoView({
        behavior: 'smooth',
    })
}
onMounted(() => {
    flow()
})
onBeforeUnmount(() => {
    window.removeEventListener('scroll', () => {
        var header = document.querySelector('header')
        header.classList.toggle('sticky', window.scrollY > 0)
    })
})
const openBlankAbbreviate = () => {
    /*
        1.在 src/router/routes文件的 /new-tab 路由的children列表加入
        {
            path: "/new-tab/abbreviate",
            component: () => import('src/plugins/xtkfk/NewTabLayout/Abbreviate.vue')
        }
        2.在 src/composables/useCommon.js 中的 const AllowList = ['/login', '/init-db'] 添加 '/new-tab/abbreviate'
    */
    // new-tab前缀 为固定写法，不喜欢可以替换src/router/routes中的路由path
    //新窗口打开
    const r = router.resolve({ path: '/new-tab/abbreviate' })
    window.open(r.href)
    // 老窗口跳转
    // router.push({ path: '/new-tab/abbreviate' })
}

</script>

<style lang="scss" scoped>
header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    transition: 0.6s;
    padding: 40px 100px;
    z-index: 100;
    box-sizing: border-box;
}

header.sticky {
    padding: 10px 100px;
    background: rgba(255, 255, 255, 0.8);
    box-shadow: 0px 1px 5px #888888;
}

header .logo {
    position: relative;
    font-weight: 700;
    color: #ffffff;
    text-decoration: none;
    font-size: 25px;
    text-transform: uppercase;
    letter-spacing: 2px;
    transition: 0.6s;
    display: flex;
    align-items: center;
    user-select: none;
    cursor: pointer;
}

header ul {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
}

header.sticky .logo,
header.sticky {
    color: black;

    .q-btn--rounded {
        color: black !important;
    }
}
</style>