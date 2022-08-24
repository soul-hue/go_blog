<script setup>
import {Fold} from '@element-plus/icons-vue'
import router from '../../router/index'
function loginOut(){
    window.sessionStorage.setItem('token','')
    router.push('/login')
}
</script>


<template>
<div class="header">
    <!-- 折叠图标 -->
    <div class="collapse left">
        <div>
            <el-icon><Fold /></el-icon>
        </div>
    </div>
    <div class="right">
        <el-button type="danger" @click="loginOut">退出</el-button>
    </div>
</div>
</template>

<style scoped lang="less">
    .header{
        width: 100%;
        height: 100%;
        background-color:#81ecec ;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        .collapse{
            .el-icon{
                font-size: 2rem;
            }
        }
    }

</style>