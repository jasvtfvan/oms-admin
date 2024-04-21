<template>
  <a-layout-header class="layout-header">
    <section class="header-left">
      <a-space :size="20">
        <a-icon
          :name="collapsed ? 'MenuUnfoldOutlined' : 'MenuFoldOutlined'"
          @click="() => emitCollapsed('update:collapsed', !collapsed)"
        />
      </a-space>
    </section>
    <section class="header-menu"></section>
    <section class="header-right">
      <a-space :size="20">
        <a-tooltip title="全屏" placement="bottom" :mouse-enter-delay="0.5">
          <a-icon
            :name="isFullscreen ? 'FullscreenExitOutlined' : 'FullscreenOutlined'"
            @click="toggle"
          />
        </a-tooltip>
        <a-dropdown placement="bottomRight" :arrow="{ pointAtCenter: true }">
          <a-avatar :src="avatar" :alt="userProfile.username">{{ userProfile.username }}</a-avatar>
          <template #overlay>
            <a-menu>
              <a-menu-item @click="openChangePwd">修改密码 </a-menu-item>
              <a-menu-item @click="confirmLogout">退出登录 </a-menu-item>
              <a-menu-divider />
              <a-menu-item @click="confirmRefreshAuth">刷新权限 </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </a-space>
    </section>
  </a-layout-header>
</template>

<script setup>
import { nextTick } from 'vue'
import { Modal, message } from 'ant-design-vue'
import { useUserStore } from '@/stores/user'
import { useFullscreen } from '@vueuse/core'
import avatar from '@/assets/images/avatar.png'
import $bus from '@/utils/bus'

const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false
  },
  theme: {
    type: String
  }
})

const userStore = useUserStore()

const emitCollapsed = defineEmits(['update:collapsed'])
const { toggle, isFullscreen } = useFullscreen()
const userProfile = userStore.userProfile

// 修改密码
const openChangePwd = async () => {
  $bus.emit('changePassword', { open: true })
}

// 是否退出
const confirmLogout = async () => {
  Modal.confirm({
    title: '确定退出系统？',
    okText: '确定',
    cancelText: '取消',
    onOk: () => {
      doLogout()
    }
  })
}

// 退出
const doLogout = async () => {
  await userStore.Logout()
  window.location.reload()
}

// 是否刷新权限
const confirmRefreshAuth = async () => {
  Modal.confirm({
    title: '确定刷新权限？',
    okText: '确定',
    cancelText: '取消',
    onOk: () => {
      doRefreshAuth()
    }
  })
}

// 刷新权限
const doRefreshAuth = async () => {
  try {
    message.loading('权限刷新中...', 0)
    await userStore.RefreshAuth()
  } catch (error) {
    console.warn(error)
    nextTick(() => {
      Modal.confirm({
        title: '刷新失败',
        content: '可选择【继续使用】或【退出系统】',
        okText: '退出系统',
        cancelText: '继续使用',
        onOk: () => {
          doLogout()
        }
      })
    })
  } finally {
    message.destroy()
  }
}
</script>

<style lang="scss" scoped>
.layout-header {
  background: rgba(255, 255, 255, 0.85);
  display: flex;
  position: sticky;
  z-index: 10;
  top: 0;
  align-items: center;
  justify-content: space-between;
  height: var(--app-header-height);
  padding: 0 20px;
  .header-right {
    cursor: pointer;
  }
  .header-menu {
    flex: 1;
    align-items: center;
    min-width: 0;
  }
}
</style>
