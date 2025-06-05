<script setup lang="ts">
import { ref } from 'vue'
import { main } from '../../wailsjs/go/models'

// 定义组件的props
const props = defineProps<{
  playback: main.Playback | null
}>()

// 弹窗相关状态
const dialogVisible = ref(false)
const currentRow = ref<main.Tmp | null>(null)
const selectedResp = ref<main.Tmp | null>(null)
const filteredResps = ref<main.Tmp[]>([])

// 格式化时间戳为可读时间
const formatTimestamp = (timestamp: number): string => {
  const date = new Date(timestamp)
  return date.toLocaleString()
}

// 格式化JSON字符串，如果是有效的JSON则美化显示
const formatJson = (jsonStr: string): string => {
  try {
    const parsed = JSON.parse(jsonStr)
    return JSON.stringify(parsed, null, 2)
  } catch (e) {
    return jsonStr // 如果不是有效的JSON，则原样返回
  }
}

// 打开设置弹窗
const openSettingsDialog = (row: main.Tmp) => {
  currentRow.value = row
  // 筛选时间戳小于当前行的Resp数据
  if (props.playback && props.playback.Resp) {
    filteredResps.value = props.playback.Resp.filter(resp => resp.Timestamp < row.Timestamp)
  } else {
    filteredResps.value = []
  }
  dialogVisible.value = true
}

// 处理选择的Resp变更
const handleRespChange = (value: main.Tmp) => {
  selectedResp.value = value
}

// 确认设置
const confirmSettings = () => {
  // 这里可以添加保存设置的逻辑
  console.log('已选择的Resp:', selectedResp.value)
  dialogVisible.value = false
}
</script>

<template>
  <div class="playback-container">
    <h3>请求记录</h3>
    <el-table v-if="playback && playback.Req && playback.Req.length > 0" :data="playback.Req" stripe border>
      <el-table-column label="时间" width="180">
        <template #default="scope">
          {{ formatTimestamp(scope.row.Timestamp) }}
        </template>
      </el-table-column>
      <el-table-column prop="Uri" label="Uri" width="250" />
      <el-table-column label="Msg">
        <template #default="scope">
          <el-popover
            placement="right"
            trigger="click"
            :width="400"
            :content="formatJson(scope.row.Msg)"
          >
            <template #reference>
              <el-button size="small" type="primary">查看内容</el-button>
            </template>
          </el-popover>
          <div class="msg-preview">{{ scope.row.Msg.substring(0, 50) }}{{ scope.row.Msg.length > 50 ? '...' : '' }}</div>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="120">
        <template #default="scope">
          <el-button size="small" type="success" @click="openSettingsDialog(scope.row)">设置</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div v-else class="no-data">暂无请求数据</div>
    
    <!-- 设置弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      title="数据设置"
      width="50%"
    >
      <div v-if="currentRow" class="dialog-content">
        <p><strong>当前选中的请求：</strong> {{ currentRow.Uri }}</p>
        <p><strong>时间戳：</strong> {{ formatTimestamp(currentRow.Timestamp) }}</p>
        
        <div class="select-container">
          <p><strong>选择时间戳小于当前行的响应数据：</strong></p>
          <el-select
            v-model="selectedResp"
            placeholder="请选择响应数据"
            style="width: 100%"
            @change="handleRespChange"
          >
            <el-option
              v-for="(resp, index) in filteredResps"
              :key="index"
              :label="`${formatTimestamp(resp.Timestamp)} - ${resp.Uri}`"
              :value="resp"
            >
              <div class="option-content">
                <div>{{ formatTimestamp(resp.Timestamp) }}</div>
                <div>{{ resp.Uri }}</div>
              </div>
            </el-option>
          </el-select>
        </div>
        
        <div v-if="selectedResp" class="selected-resp-preview">
          <p><strong>选中的响应内容：</strong></p>
          <pre>{{ formatJson(selectedResp.Msg) }}</pre>
        </div>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmSettings">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.playback-container {
  margin: 20px 0;
}

.msg-preview {
  margin-top: 5px;
  color: #666;
  font-size: 12px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.no-data {
  text-align: center;
  padding: 20px;
  color: #999;
}

.dialog-content {
  margin-bottom: 20px;
}

.select-container {
  margin: 20px 0;
}

.option-content {
  display: flex;
  justify-content: space-between;
}

.selected-resp-preview {
  margin-top: 20px;
  padding: 10px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.selected-resp-preview pre {
  white-space: pre-wrap;
  word-break: break-all;
  max-height: 200px;
  overflow-y: auto;
}
</style>