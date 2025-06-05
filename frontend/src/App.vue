<script setup lang="ts">
import { onMounted, ref } from 'vue'
import PlaybackTable from './components/PlaybackTable.vue'
import { main } from '../wailsjs/go/models'
import { A } from '../wailsjs/go/main/App'
// 示例数据或从后端获取的数据
const playbackData = ref<main.Playback | null>(null)

// 假设您从后端获取数据的函数
const fetchPlaybackData = async () => {
  // 这里替换为实际的数据获取逻辑
  // 例如：playbackData.value = await YourBackendAPI.GetPlaybackData()
  playbackData.value = await A()
  console.log(playbackData)
  // 示例数据
  // playbackData.value = main.Playback.createFrom({
  //   Req: [
  //     { Timestamp: Date.now(), Uri: "/api/example", Msg: '{"key":"value","data":123}' },
  //     { Timestamp: Date.now() - 60000, Uri: "/api/test", Msg: '{"status":"success"}' }
  //   ],
  //   Resp: []
  // })
}

// 组件挂载时获取数据
onMounted(() => {
  fetchPlaybackData()
})
</script>

<template>
    <div>
    <!-- 其他内容 -->
    <PlaybackTable :playback="playbackData" />
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
.logo.typescript:hover {
  filter: drop-shadow(0 0 2em #1296DB);
}
.logo.pinia:hover {
  filter: drop-shadow(0 0 2em #FFCF48);
}
</style>
