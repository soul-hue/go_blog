<script setup>
import { ref, reactive, getCurrentInstance } from "vue";
import router from '../router'
let ruleForm = reactive({
  username: "admin2",
  password: "admin123",
});
// let rules = []
const ruleFormRef = ref(null);
const rules = reactive({
  username: [
    { required: true, message: "请输出用户名", trigger: "blur" },
    { min: 3, max: 10, message: "用户名长度在3或10个字符之间", trigger: "blur" },
  ],
  password: [
    {
      required: true,
      message: "请输入密码",
      trigger: "blur",
    },
    { min: 3, max: 10, message: "密码长度在3或10个字符之间", trigger: "blur" },
  ],
});
const ctx = getCurrentInstance();
const { appContext } = getCurrentInstance()

// 验证
function loginValidate() {
  ruleFormRef.value.validate(async (valid) => {
    if (!valid){
        ElMessage.error('请输入正确的用户名')
        return
    }
       
      const {data:res} = await ctx.proxy.$Api({
        url : "login",
        method : 'post',
        data : ruleForm
      })
      console.log(res);
    if(res.status != 200){
        return ElMessage.error('用户名或密码不正确',appContext)
    }
    window.sessionStorage.setItem('token',res.token)
    ElMessage.success('登录成功')
    // ctx.$route.push('/admin')
    router.push('/admin')
  });
}
</script>

<template>
  <el-container>
    <el-card>
      <el-form
        ref="ruleFormRef"
        :model="ruleForm"
        :rules="rules"
        label-width="120px"
        class="demo-ruleForm"
        label-position="left"
      >
        <el-form-item label="用户名" prop="username">
          <el-input v-model="ruleForm.username" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="ruleForm.password" />
        </el-form-item>
      </el-form>
      <div>
        <el-button type="primary" @click="loginValidate">登录</el-button>
      </div>
      <div>
        <el-button :link="true">没有账号？立即注册</el-button>
      </div>
    </el-card>
  </el-container>
</template>

<style scoped lang="less">
.el-container {
  height: 100vh;
  // position: relative;
  background: #2c3e50;
}

.el-card {
  position: absolute;
  width: 25rem;
  height: 15rem;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  :deep(.el-form-item__content) {
    margin-left: 0px;
  }

  .el-button {
    width: 100%;
  }
}
</style>
