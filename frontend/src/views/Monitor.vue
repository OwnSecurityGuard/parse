<template>
    <el-button @click="monitorStore.Start">assad</el-button>
    <el-button @click="monitorStore.Pull">assad</el-button>
    <el-card class="playback-recorder">



        <div class="table-container">
            <!-- 客户端 → 服务端 -->
            <div class="table-panel">
                <div class="panel-title">客户端 → 服务端</div>
                <el-table :data="monitorStore.c2s">
                    <el-table-column v-for="k in monitorStore.columns" :label="k.label" width="180">
                        <template #default="scope">
                            {{ scope.row.key_values[k.prop][0] }}

                        </template>

                    </el-table-column>



                </el-table>


            </div>

            <!-- 服务端 → 客户端 -->
            <!-- <div class="table-panel">
        <div class="panel-title">服务端 → 客户端</div>
        <el-table :data="monitorStore.c2s">
          <el-table-column prop="time" label="Date" width="180" />
          <el-table-column prop="name" label="Name" width="180" />
          <el-table-column prop="data" label="data" />

        </el-table>
      </div> -->
        </div>

        <!-- JSON 弹窗查看 -->

    </el-card>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import { useMonitorStore } from '../store/monitor'



// // 播放数据源（由外部注入或实时推送）
const monitorStore = useMonitorStore()
onMounted(() => {
    monitorStore.GetColumns()
})


// 公共列定义



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
