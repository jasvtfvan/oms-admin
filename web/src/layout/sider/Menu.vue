<template>
  <section class="menu-container">
    <div class="select-wrap">
      <a-select
        v-model:value="selectGroupCode"
        :options="groups"
        :default-active-first-option="true"
        style="width: 100%"
        show-search
      ></a-select>
    </div>
    <a-menu :theme="props.theme" :collapsed="props.collapsed" collapsible mode="inline">
      <template v-for="item in menus" :key="item.name">
        <MySubMenu :item="item"></MySubMenu>
      </template>
    </a-menu>
  </section>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useUserStore } from '@/stores/user'
import MySubMenu from './MySubMenu.vue'

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

const selectGroupCode = ref('')
const groups = computed(() =>
  userStore.groups.map((item) => ({
    label: item.shortName,
    value: item.orgCode
  }))
)
if (groups.value[0]) {
  selectGroupCode.value = groups.value[0].value
}

const menus = computed(() => userStore.menus)
</script>

<style lang="scss" scoped>
.menu-container {
  height: calc(100vh - var(--app-header-height));
  width: 100%;
  overflow: auto;
  .select-wrap {
    padding: 16px 16px 8px 16px;
  }
  > .ant-menu {
    width: 100%;
  }
  &::-webkit-scrollbar {
    width: 0;
    height: 0;
  }
}
</style>
