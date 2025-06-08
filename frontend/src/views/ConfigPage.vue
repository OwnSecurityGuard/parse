<template>
  <el-container class="config-page">
    <!-- 左侧区域 -->
    <el-aside width="220px" class="sidebar">
      <div class="sidebar-header">
        <div class="tech-circle"></div>
        <h2 class="sidebar-title">请求类型</h2>
        <div class="tech-circle"></div>
      </div>
      
      <!-- 请求名称选择下拉框 -->
      <div class="request-selector">
        <el-select
          v-model="selectedReqName"
          placeholder="选择请求名称"
          @change="handleReqNameChange"
          class="req-select"
        >
          <el-option
            v-for="name in uniqueRequestNames"
            :key="name"
            :label="name"
            :value="name"
          />
        </el-select>
      </div>

      <!-- 请求实例列表 -->
      <div v-if="selectedReqName" class="request-instance-list">
        <div class="list-header">{{ selectedReqName }} 实例</div>
        <div 
          v-for="req in filteredRequestsByName"
          :key="`${req.name}|${req.time}`"
          :class="['request-instance-item', { 'is-active': selectedReqKey === `${req.name}|${req.time}` }]"
          @click="handleReqInstanceClick(req)"
        >
          <span class="instance-time">@ {{ req.time }}</span>
        </div>
        <div v-if="filteredRequestsByName.length === 0" class="empty-instances">
          <el-empty description="无请求实例" :image-size="50" />
        </div>
      </div>
      
      <div class="tech-decoration">
        <div class="tech-line"></div>
        <div class="tech-dot"></div>
        <div class="tech-line"></div>
      </div>
    </el-aside>

    <!-- 右侧主内容 -->
    <el-main class="main-panel">
      <div v-if="selectedReq" class="config-content">
        <div class="section-header">
          <h2 class="section-title">
            <span class="tech-badge">配置</span>
            {{ selectedReq.name }} @ {{ selectedReq.time }}
          </h2>
          <div class="tech-indicator"></div>
        </div>

        <!-- 响应绑定选择 -->
        <el-card class="config-card tech-card">
          <el-form label-width="120px">
            <el-form-item label="绑定响应包">
              <el-select
                v-model="selectedS2CKey"
                placeholder="选择一个响应包"
                @change="handleS2CChange"
                class="response-select"
              >
                <el-option
                  v-for="(msg, idx) in filteredS2C"
                  :key="idx"
                  :label="`${msg.name} @ ${msg.time}`"
                  :value="`${msg.name}|${msg.time}`"
                />
              </el-select>
            </el-form-item>
          </el-form>
        </el-card>

        <!-- 参数配置列表 -->
        <el-card class="config-card param-config-card tech-card">
          <template #header>
            <div class="card-header">
              <div class="header-title">
                <div class="tech-icon"></div>
                <span>参数配置</span>
              </div>
              <el-button type="primary" @click="addParam" size="small" class="tech-button">
                <el-icon><Plus /></el-icon>
                <span>添加</span>
              </el-button>
            </div>
          </template>

          <div class="param-list">
            <div
              v-for="(param, index) in configMap[selectedReqKey]"
              :key="index"
              class="param-row"
            >
              <el-row :gutter="16" align="middle">
                <el-col :span="9">
                  <el-select 
                    v-model="param.srcPath" 
                    placeholder="选择来源字段" 
                    filterable
                    class="param-select"
                  >
                    <el-option
                      v-for="key in availableS2CKeys"
                      :key="key"
                      :label="key"
                      :value="key"
                    />
                  </el-select>
                </el-col>
                <el-col :span="9">
                  <el-input 
                    v-model="param.tarPath" 
                    placeholder="目标路径" 
                    class="param-input"
                  />
                </el-col>
                <el-col :span="2">
                  <el-button
                    type="danger"
                    circle
                    @click="removeParam(index)"
                    class="delete-btn"
                  >
                    <el-icon><Delete /></el-icon>
                  </el-button>
                </el-col>
              </el-row>
            </div>
            <div v-if="configMap[selectedReqKey]?.length === 0" class="empty-params">
              <el-empty description="暂无参数配置" />
            </div>
          </div>
        </el-card>

        <div class="save-section">
          <el-button type="success" size="large" @click="saveConfig" class="save-btn tech-button">
            <el-icon><Check /></el-icon>
            <span>保存配置</span>
          </el-button>
        </div>
      </div>
      
      <div v-else class="empty-selection">
        <div class="tech-empty">
          <div class="tech-circle-large"></div>
          <el-empty description="请从左侧选择一个请求类型" />
          <div class="tech-line-horizontal"></div>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { reactive, ref, computed, onMounted } from 'vue'
import { usePlaybackStore } from '../store/playback'
import { Plus, Delete, Check } from '@element-plus/icons-vue'
import type { Req, Param} from '../types/playback'
  
 
 

function GetKey(s2cName: string): string[] {
  if (s2cName === 'LoginResp') return ['data.user.id', 'data.token']
  if (s2cName === 'MoveResp') return ['pos.x', 'pos.y']
  return ['data']
}

const playback = usePlaybackStore()
onMounted(() => playback.loadMockData())

// 唯一的请求名称列表
const uniqueRequestNames = computed<string[]>(() => {
  const names = new Set<string>()
  playback.c2s.forEach(req => names.add(req.name))
  return Array.from(names)
})

// 选中的请求名称
const selectedReqName = ref('')

// 根据选中的名称过滤出的请求实例
const filteredRequestsByName = computed<Req[]>(() => {
  if (!selectedReqName.value) return []
  return playback.c2s.filter(req => req.name === selectedReqName.value)
})

// 选中的请求键值（格式：name|time）
const selectedReqKey = ref('')

// 当前选中的请求对象
const selectedReq = computed(() => {
  if (!selectedReqKey.value) return null
  
  const [name, timeStr] = selectedReqKey.value.split('|')
  const time = parseFloat(timeStr)
  
  return playback.c2s.find(req => req.name === name && req.time === time) || null
})

const configMap = reactive<Record<string, Param[]>>({})
const selectedS2CKey = ref('')
const currentS2C = ref<{ name: string; data: string } | null>(null)
const currentS2CName = ref('')

// 处理请求名称选择变化
function handleReqNameChange(val: string) {
  selectedReqName.value = val
  selectedReqKey.value = '' // 清空当前选中的具体请求
  currentS2C.value = null // 清空响应预览
}

// 处理请求实例点击
function handleReqInstanceClick(req: Req) {
  const key = `${req.name}|${req.time}`
  selectedReqKey.value = key
  if (!configMap[key]) {
    configMap[key] = []
  }
}

const filteredS2C = computed(() => playback.s2c)

function handleS2CChange(val: string) {
  const [name, timeStr] = val.split('|')
  const time = parseFloat(timeStr)
  const msg = playback.s2c.find(m => m.name === name && m.time === time)
  if (msg) {
    currentS2C.value = msg
    currentS2CName.value = msg.name
  }
}

const availableS2CKeys = computed(() => GetKey(currentS2CName.value))

function addParam() {
  if (!selectedReqKey.value) return
  configMap[selectedReqKey.value].push({
    index: configMap[selectedReqKey.value].length,
    srcPath: '',
    tarPath: '',
  })
}

function removeParam(index: number) {
  if (!selectedReqKey.value) return
  configMap[selectedReqKey.value].splice(index, 1)
}

function saveConfig() {
  console.log('保存配置：', selectedReqKey.value, configMap[selectedReqKey.value])
}

// function formatJson(str: string): string {
//   try {
//     return JSON.stringify(JSON.parse(str), null, 2)
//   } catch {
//     return str
//   }
// }
</script>

<style scoped>
.config-page {
  height: 100vh;
  background-color: var(--bg-dark);
}

/* 侧边栏样式 */
.sidebar {
  background-color: var(--bg-primary);
  border-right: 1px solid var(--border-color);
  padding: 20px;
  position: relative;
  overflow: hidden;
}

.sidebar::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background: linear-gradient(90deg, var(--primary-color), var(--primary-light));
}

.sidebar-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.tech-circle {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: var(--primary-color);
  box-shadow: 0 0 10px var(--primary-color);
}

.sidebar-title {
  flex: 1;
  font-size: 1.2rem;
  color: var(--text-primary);
  text-align: center;
  margin: 0;
}

/* 主面板样式 */
.main-panel {
  background-color: var(--bg-secondary);
  padding: 20px;
}

/* 卡片样式 */
.tech-card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius);
  margin-bottom: 20px;
  position: relative;
  overflow: hidden;
}

.tech-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background: linear-gradient(90deg, var(--primary-color), var(--primary-light));
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 15px;
  border-bottom: 1px solid var(--border-color);
}

.header-title {
  display: flex;
  align-items: center;
  gap: 10px;
  color: var(--text-primary);
}

.tech-icon {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background-color: var(--primary-color);
  box-shadow: 0 0 8px var(--primary-color);
}

/* 按钮样式 */
.tech-button {
  background: linear-gradient(135deg, var(--primary-color), var(--primary-dark));
  border: none;
  padding: 8px 16px;
  border-radius: var(--border-radius);
  color: var(--text-primary);
  transition: all 0.3s ease;
}

.tech-button:hover {
  transform: translateY(-1px);
  box-shadow: 0 0 15px rgba(79, 70, 229, 0.4);
}

/* JSON预览区域样式 */
.json-container {
  background-color: var(--bg-dark);
  border-radius: var(--border-radius);
  padding: 20px;
  position: relative;
}

.json-preview {
  color: var(--text-primary);
  font-family: 'Fira Code', monospace;
  margin: 0;
  white-space: pre-wrap;
}

/* 装饰元素 */
.tech-decoration-corner {
  position: absolute;
  width: 10px;
  height: 10px;
  border: 1px solid var(--primary-color);
}

.top-left {
  top: 10px;
  left: 10px;
  border-right: none;
  border-bottom: none;
}

.top-right {
  top: 10px;
  right: 10px;
  border-left: none;
  border-bottom: none;
}

.bottom-left {
  bottom: 10px;
  left: 10px;
  border-right: none;
  border-top: none;
}

.bottom-right {
  bottom: 10px;
  right: 10px;
  border-left: none;
  border-top: none;
}

/* 空状态样式 */
.empty-selection {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  color: var(--text-secondary);
}

.tech-empty {
  text-align: center;
  position: relative;
}

.tech-circle-large {
  width: 40px;
  height: 40px;
  border: 2px solid var(--primary-color);
  border-radius: 50%;
  margin: 0 auto 20px;
  position: relative;
}

.tech-circle-large::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 8px;
  height: 8px;
  background-color: var(--primary-color);
  border-radius: 50%;
  box-shadow: 0 0 15px var(--primary-color);
}

.tech-line-horizontal {
  width: 100px;
  height: 1px;
  background: linear-gradient(90deg, transparent, var(--primary-color), transparent);
  margin: 20px auto 0;
}

.request-instance-list {
  margin-top: 20px;
  max-height: calc(100vh - 200px); /* 调整高度以适应侧边栏 */
  overflow-y: auto;
  padding-right: 5px;
}

.list-header {
  font-size: 1rem;
  color: var(--text-secondary);
  margin-bottom: 10px;
  padding-left: 5px;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 5px;
}

.request-instance-item {
  padding: 8px 10px;
  cursor: pointer;
  color: var(--text-primary);
  border-radius: var(--border-radius);
  transition: background-color 0.2s ease, color 0.2s ease;
  margin-bottom: 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.request-instance-item:hover {
  background-color: rgba(var(--primary-color-rgb), 0.2);
  color: var(--primary-light);
}

.request-instance-item.is-active {
  background-color: var(--primary-color);
  color: var(--bg-primary);
  font-weight: bold;
}

.instance-time {
  font-size: 0.85rem;
  color: var(--text-light);
}

.request-instance-item.is-active .instance-time {
  color: var(--bg-primary);
}

.empty-instances {
  text-align: center;
  padding: 20px 0;
}
.param-row {
  margin-bottom: 15px; /* 增加行之间的间隔 */
}
</style>
