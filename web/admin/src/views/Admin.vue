<script setup>
import { ref } from "vue";
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from "@element-plus/icons-vue";
import Footer from '../components/admin/footer.vue'
import Header from '../components/admin/header.vue'
const isCollapse = ref(false);
const handleOpen = (key, keyPath) => {
  console.log(key, keyPath);
};
const handleClose = (key, keyPath) => {
  console.log(key, keyPath);
};
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px"
        ><!-- 左边折叠栏 -->
        <div>
          <div class="logo">
            <!-- logo -->
          </div>
          <el-menu
            default-active="2"
            class="el-menu-vertical-demo"
            :collapse="isCollapse"
            @open="handleOpen"
            @close="handleClose"
          >
            <el-sub-menu>
              <template #title>
                <el-icon><location /></el-icon>
                <span>Navigator One</span>
              </template>
              <el-menu-item-group>
                <el-menu-item index="1-1">用户列表</el-menu-item>
                <el-menu-item index="1-2">item two</el-menu-item>
              </el-menu-item-group>
            </el-sub-menu>
          </el-menu>
        </div></el-aside
      >
      <el-container>
        <el-header><Header></Header></el-header>
        <el-main>Main</el-main>
        <el-footer><Footer></Footer></el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped lang="less">
.logo {
  width: 200px;
  height: 10vh;
  background-color: black;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 90vh;
}
</style>
