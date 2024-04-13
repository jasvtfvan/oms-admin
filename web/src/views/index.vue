<template>
  <article class="login-box">
    <section class="content-box">
      <div class="login-logo">
        <div class="svg-wrap">
          <svg-icon name="logo"></svg-icon>
        </div>
        <div class="title">全域运营管理系统</div>
      </div>
      <a-form layout="horizontal" :model="loginFormModel" @submit.prevent="handleSubmit">
        <a-form-item>
          <a-input v-model="loginFormModel.username" size="large" placeholder="admin">
            <template #prefix> <user-outlined /> </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-input
            v-model="loginFormModel.password"
            size="large"
            type="password"
            placeholder="a123456"
            autocomplete="new-password"
          >
            <template #prefix> <lock-outlined /></template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-input
            v-model="loginFormModel.verifyCode"
            placeholder="验证码"
            :maxlength="4"
            size="large"
          >
            <template #prefix> <safety-outlined /> </template>
            <template #suffix>
              <img
                :src="captcha"
                class="absolute right-0 h-full cursor-pointer"
                @click="updateCaptcha"
              />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" size="large" :loading="loading" block>
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </section>
  </article>
</template>

<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'
import { UserOutlined, LockOutlined, SafetyOutlined } from '@ant-design/icons-vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)
const captcha = ref('')
const loginFormModel = ref({
  username: 'admin',
  password: 'a123456',
  verifyCode: '',
  captchaId: ''
})

const updateCaptcha = async () => {
  // const data = await Api.captcha.captchaCaptchaByImg({ width: 100, height: 50 })
  // captcha.value = data.img
  // loginFormModel.value.captchaId = data.id
}
updateCaptcha()

const handleSubmit = async () => {
  const { username, password, verifyCode } = loginFormModel.value
  if (username.trim() == '' || password.trim() == '') {
    return message.warning('用户名或密码不能为空！')
  }
  if (!verifyCode) {
    return message.warning('请输入验证码！')
  }
  message.loading('登录中...', 0)
  loading.value = true

  userStore
    .Login(loginFormModel.value)
    .then((res) => {
      message.success('登录成功！')
      setTimeout(() => router.replace(route.query.redirect || '/'))
    })
    .catch((_) => {
      updateCaptcha()
    })
    .finally(() => {
      loading.value = false
      message.destroy()
    })
}
</script>

<style lang="less" scoped>
.login-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100vw;
  height: 100vh;
  background: url('@/assets/images/login.svg');
  background-size: 100%;
  .content-box {
    position: relative;
    top: -40px;
    .login-logo {
      display: flex;
      justify-content: center;
      align-items: center;
      margin-bottom: 30px;
      .svg-wrap {
        width: 45px;
        height: 45px;
      }
      .title {
        margin-left: 8px;
        font-size: 30px;
      }
    }
    :deep(.ant-form) {
      width: 400px;
      .ant-col {
        width: 100%;
      }
      .ant-form-item-label {
        padding-right: 6px;
      }
      .ant-input-prefix {
        margin-right: 10px;
      }
    }
  }
}
</style>
