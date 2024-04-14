<template>
  <article class="login-box">
    <section class="content-box">
      <div class="login-logo">
        <div class="svg-wrap">
          <svg-icon name="logo"></svg-icon>
        </div>
        <div class="title">{{ fullName }}</div>
      </div>
      <div class="brand">{{ brand }}</div>
      <a-form layout="horizontal" :model="loginFormModel" @submit.prevent="handleSubmit">
        <a-form-item>
          <a-input v-model="loginFormModel.username" size="large" placeholder="用户名">
            <template #prefix> <user-outlined /> </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-input
            v-model="loginFormModel.password"
            size="large"
            type="password"
            placeholder="密码"
            autocomplete="new-password"
          >
            <template #prefix> <lock-outlined /></template>
          </a-input>
        </a-form-item>
        <a-form-item v-if="loginFormModel.openCaptcha">
          <a-input
            v-model="loginFormModel.captcha"
            placeholder="验证码"
            :maxlength="4"
            size="large"
          >
            <template #prefix> <safety-outlined /> </template>
            <template #suffix>
              <img
                :src="picPath"
                class="absolute right-0 h-full cursor-pointer"
                @click="updateCaptcha"
              />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" size="large" :loading="loginLoading" block>
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </section>
    <footer class="login-footer">
      <div>{{ footerInfo }}</div>
    </footer>
  </article>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { message, Modal } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'
import { UserOutlined, LockOutlined, SafetyOutlined } from '@ant-design/icons-vue'
import { postInitCheck, postInitDb } from '@/api/common/db'
import setting from '@/setting.js'
import { useLoading } from '@/hooks/useLoading'

// use
const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const { showLoading, hideLoading } = useLoading()

/** 网站信息 */
const fullName = ref(setting.websiteInfo.fullName)
const brand = ref(setting.websiteInfo.brand)
const copyright = ref(setting.websiteInfo.copyright)
const forTheRecord = ref(setting.websiteInfo.forTheRecord)
const footerInfo = computed(() => `${copyright.value} | ${forTheRecord.value}`)

/** 表单数据 */
const loginLoading = ref(false)
const loginFormModel = ref({
  username: '',
  password: '',
  picPath: '',
  openCaptcha: false,
  captchaLength: 0,
  captchaId: '',
  captcha: ''
})

// 已经初始化
const initReady = ref(false)
// 全局loadingTimer
let loadingTimer // 不需要响应式

// 弹出提示框
const openModal = () => {
  Modal.confirm({
    title: '尚未初始化，请选择操作',
    content: '点击【开始】开始初始化，点击【检查】重新检查',
    okText: '开始',
    cancelText: '检查',
    onOk: () => {
      initDb()
    },
    onCancel: () => {
      initCheck()
    }
  })
}

// 检查初始化
const initCheck = async () => {
  try {
    if (loadingTimer) clearTimeout(loadingTimer)
    showLoading()
    await postInitCheck()
    initReady.value = true
    updateCaptcha()
  } catch (error) {
    initReady.value = false
    openModal()
  } finally {
    loadingTimer = setTimeout(() => {
      hideLoading()
    }, 200)
  }
}
initCheck()
// 执行初始化
const initDb = async () => {
  try {
    if (loadingTimer) clearTimeout(loadingTimer)
    showLoading()
    const { msg } = await postInitDb({initPwd: '123'})
    message.success(msg || '初始化成功')
    initReady.value = true
    updateCaptcha()
  } catch (error) {
    initReady.value = false
    const { needRefresh } = error.data || {}
    message.error({
      content: error.msg || '初始化失败',
      onClose: () => {
        if (needRefresh) {
          initCheck()
        } else {
          openModal()
        }
      }
    })
  } finally {
    loadingTimer = setTimeout(() => {
      hideLoading()
    }, 200)
  }
}

// 获取验证码
const getCaptcha = async () => {
  try {
    const data = await userStore.Captcha({
      width: 100,
      height: 50
    })
    const { captchaId, picPath, captchaLength, openCaptcha } = data
    loginFormModel.value.captchaId = captchaId
    loginFormModel.value.picPath = picPath
    loginFormModel.value.captchaLength = captchaLength
    loginFormModel.value.openCaptcha = openCaptcha
  } catch (error) {
    message.error(error.msg || '获取验证码失败')
  }
}
const updateCaptcha = async () => {
  if (!initReady.value) {
    return message.warning('系统尚未初始化')
  }
  getCaptcha()
}

// 登录方法
const handleSubmit = async () => {
  if (!initReady.value) {
    return message.warning('系统尚未初始化')
  }
  const { username, password, captcha, captchaId, openCaptcha } = loginFormModel.value
  if (username.trim() == '' || password.trim() == '') {
    return message.warning('用户名或密码不能为空')
  }
  if (openCaptcha) {
    if (!captcha) {
      return message.warning('请输入验证码')
    }
    if (!captchaId) {
      return message.warning('验证码ID丢失，请重新获取', () => {
        loginFormModel.value.captcha = ''
      })
    }
  }
  message.loading('登录中...', 0)
  loading.value = true

  try {
    const { captcha, captchaId, username, password } = loginFormModel.value
    await userStore.Login({ captcha, captchaId, username, password })
    message.success('登录成功')
    setTimeout(() => router.replace(route.query.redirect || '/home'))
  } catch (_) {
    updateCaptcha()
  } finally {
    loading.value = false
    message.destroy()
  }
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
      margin-bottom: 9px;
      .svg-wrap {
        width: 45px;
        height: 45px;
      }
      .title {
        margin-left: 8px;
        font-size: 30px;
      }
    }
    .brand {
      text-align: center;
      font-size: 14px;
      color: rgba(0, 0, 0, 0.45);
      margin-bottom: 30px;
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
  .login-footer {
    position: absolute;
    bottom: 0;
    padding: 25px;
    font-size: 12px;
    color: rgba(0, 0, 0, 0.45);
  }
}
</style>
