<template>
  <a-modal
    v-model:open="pwdModalOpen"
    :closable="false"
    :maskClosable="false"
    :cancelText="pwdModalOpenForce ? '退出登录' : '取消修改'"
    title="修改密码"
    okText="确定修改"
    @ok="handlePwdOk"
    @cancel="handlePwdCancel"
  >
    <section class="pwd-modal">
      <p class="item-wrap">若忘记密码，请联系组织管理员->重置密码</p>
      <p class="item-wrap">
        <a-input-password
          :maxlength="20"
          placeholder="输入旧密码"
          autocomplete="new-password"
          v-model:value="pwdObj.oldPassword"
        >
        </a-input-password>
      </p>
      <p class="item-wrap">
        <a-input-password
          placeholder="输入新密码"
          autocomplete="new-password"
          :maxlength="20"
          v-model:value="pwdObj.newPassword"
        >
        </a-input-password>
      </p>
      <p class="item-wrap">
        <a-input-password
          placeholder="再次输入新密码"
          autocomplete="new-password"
          :maxlength="20"
          v-model:value="pwdObj.reNewPassword"
        >
        </a-input-password>
      </p>
    </section>
  </a-modal>
</template>

<script setup>
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
import { decryptPwd, decryptSecret } from '@/utils/cryptoLoginSecret'
import { rsaEncryptOAEP } from '@/utils/rsaEncrypt'
import { useUserStore } from '@/stores/user'
import $bus from '@/utils/bus'
import { Modal, message } from 'ant-design-vue'
import { isValidPassword, passwordErrorMessage } from '@/utils/util'
import { postChangePwd } from '@/api/common/user'
import { postCompareSecret } from '@/api/common/base'

const userStore = useUserStore()
const username = userStore.userProfile.username

const pwdModalOpen = ref(false) // 是否显示修改密码弹出框
const pwdModalOpenForce = ref(false) // 是否是强制弹出
const pwdObj = ref({
  cachePwd: '', // 缓存的密码
  oldPassword: '', // 旧密码
  newPassword: '', // 新密码
  reNewPassword: '' // 再次输入新密码
})
let timer

const doOpenLogoutModal = (msg) => {
  nextTick(() => {
    if (pwdModalOpen.value) {
      if (timer) clearTimeout(timer)
      timer = setTimeout(() => {
        pwdModalOpen.value = false
        openLogoutModal(msg)
      }, 500)
    } else {
      openLogoutModal(msg)
    }
  })
}

const openLogoutModal = (msg) => {
  Modal.error({
    title: msg || '缓存密码被篡改',
    content: `${msg || '缓存密码被篡改'}，需退出登录`,
    okText: '退出登录',
    onOk: async () => {
      await userStore.Logout()
      window.location.reload()
    }
  })
}

const judgePwd = async () => {
  const password = decryptPwd()
  const enSecret = userStore.encryptedSecret
  if (!enSecret) {
    console.warn('userStore.encryptedSecret未取到')
    doOpenLogoutModal()
  }
  try {
    const deSecret = decryptSecret(enSecret)
    const secret = await rsaEncryptOAEP(JSON.stringify({ username, password }))
    const { code } = await postCompareSecret({ s1: secret, s2: deSecret })
    // 缓存的密钥相同，没有被篡改
    if (/^[2-3]0\d$/.test(code)) {
      pwdObj.value.cachePwd = password // 把解析出来的旧密码赋值给cachePwd
    } else {
      // 如果两次secret不一致，则弹出退出提示框
      doOpenLogoutModal()
    }
  } catch (error) {
    doOpenLogoutModal(error.msg)
  }
}
judgePwd()

// changePassword弹出框ok
const handlePwdOk = async () => {
  const { cachePwd, oldPassword, newPassword, reNewPassword } = pwdObj.value
  if (!cachePwd) {
    doOpenLogoutModal()
    return
  }
  if (!oldPassword) {
    message.error('旧密码不能为空')
    return
  }
  if (!newPassword || !reNewPassword) {
    message.error('新密码不能为空')
    return
  }
  if (newPassword == oldPassword) {
    message.error('新密码和旧密码不能相同')
    return
  }
  if (!isValidPassword(newPassword) || !isValidPassword(reNewPassword)) {
    message.error(passwordErrorMessage)
    return
  }
  if (newPassword != reNewPassword) {
    message.error('新密码必须相同')
    return
  }
  if (cachePwd != oldPassword) {
    message.error('旧密码输入错误')
    return
  }
  // 打包加密密码，调用后台接口
  try {
    const secret = await rsaEncryptOAEP(JSON.stringify({ oldPassword, newPassword, reNewPassword }))
    await postChangePwd({ secret })
    message.success('修改成功，即将重新登录', 3, async () => {
      await userStore.Logout()
      window.location.reload()
    })
  } catch (_) {
    pwdModalOpen.value = true
  }
}
// changePassword弹出框cancel
const handlePwdCancel = async () => {
  if (pwdModalOpenForce.value) {
    await userStore.Logout()
    window.location.reload()
  }
}

// 强制修改密码
const getChangePasswordForce = ({ open = false }) => {
  pwdModalOpen.value = open
  pwdModalOpenForce.value = true
}
// 修改密码
const getChangePassword = ({ open = false } = {}) => {
  pwdModalOpen.value = open
  pwdModalOpenForce.value = false
}

onMounted(() => {
  console.log('layout-changePwd-onMounted')
  $bus.on('changePassword', getChangePassword)
  $bus.on('changePasswordForce', getChangePasswordForce)
})
onUnmounted(() => {
  console.log('layout-changePwd-onUnmounted')
  $bus.off('changePassword', getChangePassword)
  $bus.off('changePasswordForce', getChangePasswordForce)
})
</script>

<style lang="scss" scoped>
.pwd-modal {
  margin: 16px 0px 24px 0px;
  .item-wrap {
    margin-bottom: 12px;
  }
}
</style>
