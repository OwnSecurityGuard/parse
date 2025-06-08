<template>
  <el-container class="app-container">
    <!-- 左侧侧边栏导航 -->
    <el-aside width="var(--sidebar-width)" class="sidebar">
      <div class="logo-container">
        <div class="logo-icon">
          <el-icon><VideoPlay /></el-icon>
        </div>
        <div class="logo-text">流量录制工具</div>
      </div>
      
      <el-menu
        :default-active="activeMenu"
        class="sidebar-menu"
        router
        :collapse="isCollapse"
      >
        <el-menu-item index="/recorder">
          <el-icon><Monitor /></el-icon>
          <template #title>录制</template>
        </el-menu-item>
        <el-menu-item index="/config">
          <el-icon><Setting /></el-icon>
          <template #title>配置</template>
        </el-menu-item>
      </el-menu>
      
      <div class="sidebar-footer">
        <el-button 
          type="text" 
          class="collapse-btn" 
          @click="toggleSidebar"
        >
          <el-icon>
            <component :is="isCollapse ? 'Expand' : 'Fold'" />
          </el-icon>
        </el-button>
      </div>
    </el-aside>

    <!-- 主内容区域 -->
    <el-container class="main-container">
      <el-header height="var(--header-height)" class="main-header">
        <div class="header-title">
          {{ getPageTitle() }}
        </div>
        <div class="header-actions">
          <el-button type="primary" plain size="small">
            <el-icon><Refresh /></el-icon>
            刷新
          </el-button>
          <el-button type="success" plain size="small">
            <el-icon><Download /></el-icon>
            导出
          </el-button>
        </div>
      </el-header>
      
      <el-main class="main-content">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { Monitor, Setting, VideoPlay, Refresh, Download,  } from '@element-plus/icons-vue'

const route = useRoute()
const activeMenu = computed(() => route.path)
const isCollapse = ref(false)

const toggleSidebar = () => {
  isCollapse.value = !isCollapse.value
  document.documentElement.style.setProperty(
    '--sidebar-width', 
    isCollapse.value ? '64px' : '220px'
  )
}

const getPageTitle = () => {
  switch (route.path) {
    case '/recorder':
      return '流量录制'
    case '/config':
      return '配置管理'
    default:
      return '流量录制工具'
  }
}
</script>

<style scoped>
.app-container {
  height: 100vh;
  overflow: hidden;
}

.sidebar {
  background-color: var(--bg-dark);
  color: #fff;
  height: 100%;
  display: flex;
  flex-direction: column;
  transition: width 0.3s;
  overflow: hidden;
}

.logo-container {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px 0;
  gap: 10px;
}

.logo-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  background-color: var(--primary-color);
  border-radius: 8px;
  font-size: 20px;
}

.logo-text {
  font-size: 18px;
  font-weight: bold;
  white-space: nowrap;
}

.sidebar-menu {
  border-right: none;
  background-color: transparent;
  flex: 1;
}

.sidebar-menu :deep(.el-menu-item) {
  color: #e2e8f0;
}

.sidebar-menu :deep(.el-menu-item.is-active) {
  background-color: rgba(59, 130, 246, 0.2);
  color: var(--primary-light);
  border-right: 3px solid var(--primary-color);
}

.sidebar-menu :deep(.el-menu-item:hover) {
  background-color: rgba(255, 255, 255, 0.1);
}

.sidebar-footer {
  padding: 10px;
  display: flex;
  justify-content: center;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.collapse-btn {
  color: #e2e8f0;
}

.main-container {
  background-color: var(--bg-secondary);
}

.main-header {
  background-color: var(--bg-primary);
  box-shadow: var(--shadow-sm);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
}

.header-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
}

.header-actions {
  display: flex;
  gap: 10px;
}

.main-content {
  padding: 20px;
  overflow-y: auto;
}

/* 页面过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
