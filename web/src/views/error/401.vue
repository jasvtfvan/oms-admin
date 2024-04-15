<template>
  <div>
    <section class="error-page">
      <div class="info">
        <img src="@/assets/images/401.svg" />
        <h2 class="error-title">401 暂无权限</h2>
        <p>你没有权限去该页面。</p>
        <p>请访问其他页面</p>
        <nav>
          <RouterLink :to="{ path: '/', replace: true }">Back To Home</RouterLink>
        </nav>
      </div>
    </section>
  </div>
</template>
<script setup>
</script>


<style lang="scss" scoped>
.error-page {
  background: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  box-sizing: border-box;
  text-align: center;

  .error-title {
    font-size: 35px;
    padding-bottom: 3px;
    font-weight: bold;
  }

  p {
    font-size: 13px;
    padding: 5px 0;
  }

  img {
    width: 100%;
  }
}
</style>
