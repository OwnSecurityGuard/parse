<template>
  <el-card class="playback-recorder">
    <div class="table-container">
      <!-- 客户端 → 服务端 -->
      <div class="table-panel">
        <div class="panel-title">客户端 → 服务端</div>
        <el-table-v2
          class="v2-table"
          :columns="s2cColumns"
          :data="playback.c2s"
          :row-height="40"
          :height="tableHeight"
        />
      </div>

      <!-- 服务端 → 客户端 -->
      <div class="table-panel">
        <div class="panel-title">服务端 → 客户端</div>
        <el-table-v2
          class="v2-table"
          :columns="s2cColumns"
          :data="playback.s2c"
          :row-height="40"
          :height="tableHeight"
        />
      </div>
    </div>

    <!-- JSON 弹窗查看 -->
    <el-dialog v-model="jsonDialogVisible" title="JSON 内容" width="600px">
      <pre>{{ dialogJsonContent }}</pre>
    </el-dialog>
  </el-card>
</template>

<script lang="ts" setup>
import { ref, h } from 'vue'
import type { Column } from 'element-plus'
import {  Playback, } from '../types/playback'
import  playbackMockData from '../types/playback'

// // 播放数据源（由外部注入或实时推送）
const playback = ref<Playback>(playbackMockData)

// JSON 弹窗查看
const jsonDialogVisible = ref(false)
const dialogJsonContent = ref('')
const handleViewJson = (data: string) => {
  try {
    dialogJsonContent.value = JSON.stringify(JSON.parse(data), null, 2)
  } catch {
    dialogJsonContent.value = '解析失败'
  }
  jsonDialogVisible.value = true
}

// 公共列定义
const baseColumns: Column[] = [
  {
    key: 'time',
    dataKey: 'time',
    title: '时间(s)',
    width: 80,
  },
  {
    key: 'name',
    dataKey: 'name',
    title: '协议名',
    width: 200,
  },
  {
    key: 'data',
    dataKey: 'data',
    title: '数据',
    width: 120,
    cellRenderer({ rowData }) {
      return h(
        'el-button',
        {
          type: 'primary',
          link: true,
          size: 'small',
          onClick: () => handleViewJson(rowData.data),
        },
        '查看 JSON'
      )
    },
  },
]

// 表头列
// const c2sColumns: Column[] = [
//   ...baseColumns,
//   {
//     key: 'params',
//     dataKey: 'params',
//     title: '参数',
//     width: 100,
//     cellRenderer({ rowData }) {
//       const len = rowData.params?.length || 0
//       return h('span', len ? `${len} 个` : '-')
//     },
//   },
// ]

const s2cColumns: Column[] = [...baseColumns]

// 表格高度（可以设置为 window 自适应）
const tableHeight = 400
</script>

<style scoped>
.playback-recorder {
  margin: 20px;
}

.table-container {
  display: flex;
  gap: 20px;
  justify-content: space-between;
}

.table-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.panel-title {
  font-weight: bold;
  margin-bottom: 10px;
}

.v2-table {
  border: 1px solid var(--el-border-color);
  border-radius: 4px;
  overflow: hidden;
}

pre {
  font-size: 12px;
  background: #f9f9f9;
  padding: 10px;
  white-space: pre-wrap;
  word-break: break-word;
  max-height: 500px;
  overflow: auto;
}
</style>
