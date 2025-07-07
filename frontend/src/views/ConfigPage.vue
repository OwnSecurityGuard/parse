<template>
  <div class="config-manager-tool">
    <!-- 主内容区 -->
    <div class="main-content">
      <!-- 顶部工具栏 -->
      <div class="toolbar">
        <div class="page-title">
          <h1>配置管理</h1>
          <p class="subtitle">配置文件</p>
        </div>
        <div class="toolbar-actions">
          <el-button type="primary" @click="refreshConfig" class="refresh-btn">
            <i class="el-icon-refresh"></i>
          </el-button>
          <el-button type="success" @click="exportConfig" class="export-btn">
            <i class="el-icon-download"></i>
          </el-button>
        </div>
      </div>

      <!-- 配置选择区 -->
      <div class="selection-panel">
        <div class="selection-row">
          <div class="selection-item">
            <label>配置文件</label>
            <el-select
              v-model="ctxConfigStore.selectedPath"
              placeholder="选择配置文件"
              size="small"
              class="custom-select"
            >
              <el-option
                v-for="item in ctxConfigStore.filaPaths"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </div>

          <div class="selection-item">
            <label>请求类型</label>
            <el-select
              v-model="ctxConfigStore.selectedReq"
              placeholder="选择请求类型"
              size="small"
              class="custom-select"
            >
              <el-option
                v-for="item in ctxConfigStore.c2s"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </div>

          <div class="action-buttons">
            <el-button type="primary" @click="openRuleDialog" size="small">添加规则</el-button>
            <el-button type="success" @click="submitConfig" size="small">提交配置</el-button>
          </div>
        </div>

        <div class="key-path-container">
          <label>唯一性字段路径</label>
          <el-input
            v-if="ctxConfigStore.configData[ctxConfigStore.selectedPath]"
            v-model="ctxConfigStore.configData[ctxConfigStore.selectedPath].keyPath"
            placeholder="如 data.user.id"
            size="small"
            class="custom-input"
          />
        </div>
      </div>

      <!-- 规则列表区 -->
      <div class="rules-container">
        <div
          v-for="(rule, index) in ctxConfigStore.configData[ctxConfigStore.selectedPath]?.keyRules[ctxConfigStore.selectedReq] || []"
          :key="index"
          class="rule-card"
        >
          <div class="rule-header">
            <h3>规则 #{{ index + 1 }}</h3>
            <el-button
              circle
              size="small"
              type="danger"
              @click="removeRule(index)"
              class="delete-btn"
            >
              <i class="el-icon-delete"></i>
            </el-button>
          </div>

          <div class="rule-body">
            <el-table :data="[rule]" border size="small" class="rule-table">
              <el-table-column label="源路径" width="200">
                <template #default="scope">
                  <el-input
                    v-model="scope.row.SrcPath"
                    placeholder="如 data.items.id"
                    size="small"
                    clearable
                  />
                </template>
              </el-table-column>
              <el-table-column label="目标键名" width="200">
                <template #default="scope">
                  <el-input
                    v-model="scope.row.DstKey"
                    placeholder="如 itemId"
                    size="small"
                    clearable
                  />
                </template>
              </el-table-column>
              <el-table-column label="目标路径" width="200">
                <template #default="scope">
                  <el-input
                    v-model="scope.row.DstPath"
                    placeholder="如 payload.item_id"
                    size="small"
                    clearable
                  />
                </template>
              </el-table-column>
            </el-table>

            <div class="filters-section">
              <div class="filters-header">
                <span>筛选条件</span>
                <el-button
                  type="primary"
                  size="small"
                  @click="ctxConfigStore.AddFilter(index)"
                  class="add-filter-btn"
                >
                  + 添加筛选
                </el-button>
              </div>
              <el-table
                :data="rule.Fs || []"
                border
                size="small"
                class="filters-table"
              >
                <el-table-column label="#" type="index" width="40" />
                <el-table-column label="路径" width="160">
                  <template #default="scope">
                    <el-input
                      v-model="scope.row.Path"
                      placeholder="data.items.0"
                      size="small"
                      clearable
                    />
                  </template>
                </el-table-column>
                <el-table-column label="字段名" width="140">
                  <template #default="scope">
                    <el-input
                      v-model="scope.row.ExpectKey"
                      placeholder="字段名"
                      size="small"
                      clearable
                    />
                  </template>
                </el-table-column>
                <el-table-column label="操作符" width="120">
                  <template #default="scope">
                    <el-select
                      v-model="scope.row.Op"
                      size="small"
                      placeholder="操作符"
                      class="w-full"
                    >
                      <el-option label="=" value="=" />
                      <el-option label="!=" value="!=" />
                      <el-option label="包含" value="contains" />
                    </el-select>
                  </template>
                </el-table-column>
                <el-table-column label="期望值">
                  <template #default="scope">
                    <el-input
                      v-model="scope.row.ExpectVal"
                      placeholder="期望值"
                      size="small"
                      clearable
                    />
                  </template>
                </el-table-column>
                <el-table-column label="操作" width="60">
                  <template #default="scope">
                    <el-button
                      circle
                      size="small"
                      type="danger"
                      @click="removeFilter(index, scope.$index)"
                    >
                      <i class="el-icon-delete"></i>
                    </el-button>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </div>
        </div>

        <!-- 空状态提示 -->
        <div v-if="!ctxConfigStore.configData[ctxConfigStore.selectedPath]?.keyRules[ctxConfigStore.selectedReq]?.length" class="empty-state">
          <p class="empty-text">暂无规则，点击上方按钮添加规则</p>
        </div>
      </div>

      <!-- 添加规则弹窗 -->
      <el-dialog
        v-model="ruleDialogVisible"
        title="添加新规则"
        width="520px"
        class="custom-dialog"
        destroy-on-close
      >
        <div class="dialog-body">
          <el-table :data="[newRule]" border size="small" class="dialog-table">
            <el-table-column label="源路径" width="200">
              <template #default="scope">
                <el-input
                  v-model="scope.row.SrcPath"
                  placeholder="如 data.items.id"
                  clearable
                />
              </template>
            </el-table-column>
            <el-table-column label="目标键名" width="200">
              <template #default="scope">
                <el-input
                  v-model="scope.row.DstKey"
                  placeholder="如 itemId"
                  clearable
                />
              </template>
            </el-table-column>
            <el-table-column label="目标路径" width="200">
              <template #default="scope">
                <el-input
                  v-model="scope.row.DstPath"
                  placeholder="如 payload.item_id"
                  clearable
                />
              </template>
            </el-table-column>
          </el-table>
        </div>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="ruleDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="addNewRule">添加</el-button>
          </div>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useCtxConfigStore } from '../store/ctxConfig';
 
const ctxConfigStore = useCtxConfigStore();
const ruleDialogVisible = ref(false);
const newRule = ref ({ SrcPath: '', DstKey: '', DstPath: '', Fs: [] });

function openRuleDialog() {
  ruleDialogVisible.value = true;
  newRule.value = { SrcPath: '', DstKey: '', DstPath: '', Fs: [] };
}

function addNewRule() {
  if (!ctxConfigStore.configData[ctxConfigStore.selectedPath]) {
    ctxConfigStore.configData[ctxConfigStore.selectedPath] = {
      keyPath: '',
      keyRules: {}
    };
  }
  const rules = ctxConfigStore.configData[ctxConfigStore.selectedPath].keyRules;
  if (!rules[ctxConfigStore.selectedReq]) {
    rules[ctxConfigStore.selectedReq] = [];
  }
  rules[ctxConfigStore.selectedReq].push(JSON.parse(JSON.stringify(newRule.value)));
  ruleDialogVisible.value = false;
}

function submitConfig() {
  ctxConfigStore.Upload();
}

function refreshConfig() {
  ctxConfigStore.load();
}

function exportConfig() {
  console.log('Exporting configuration...');
}

function removeRule(index: number) {
  if (
    ctxConfigStore.configData[ctxConfigStore.selectedPath] &&
    ctxConfigStore.configData[ctxConfigStore.selectedPath].keyRules &&
    ctxConfigStore.configData[ctxConfigStore.selectedPath].keyRules[ctxConfigStore.selectedReq]
  ) {
    ctxConfigStore.configData[ctxConfigStore.selectedPath].keyRules[ctxConfigStore.selectedReq].splice(index, 1);
  }
}

function removeFilter(ruleIndex: number, filterIndex: number) {
  if (
    ctxConfigStore.configData[ctxConfigStore.selectedPath] &&
    ctxConfigStore.configData[ctxConfigStore.selectedPath].keyRules &&
    ctxConfigStore.configData[ctxConfigStore.selectedPath].keyRules[ctxConfigStore.selectedReq]
  ) {
    ctxConfigStore.configData[ctxConfigStore.selectedPath].keyRules[ctxConfigStore.selectedReq][ruleIndex].Fs.splice(filterIndex, 1);
  }
}

onMounted(() => {
  ctxConfigStore.load();
});
</script>

<style scoped>
:root {
  --primary-color: #409eff;
  --success-color: #67c23a;
  --danger-color: #f56c6c;
  --light-bg: #f5f7fa;
  --card-bg: #ffffff;
  --text-color: #333333;
  --text-light: #606266;
  --border-color: #e4e7ed;
  --shadow-sm: 0 2px 4px rgba(0, 0, 0, 0.05);
  --spacing-sm: 8px;
  --spacing-md: 16px;
  --spacing-lg: 24px;
  --border-radius: 4px;
}

.config-manager-tool {
  min-height: 100vh;
  background-color: var(--light-bg);
}

.main-content {
  padding: var(--spacing-lg);
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}

/* 顶部工具栏样式 */
.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
}

.page-title h1 {
  margin: 0;
  font-size: 20px;
  font-weight: 600;
  color: var(--text-color);
}

.page-title .subtitle {
  font-size: 14px;
  color: var(--text-light);
  margin-top: 4px;
}

.toolbar-actions {
  display: flex;
  gap: var(--spacing-sm);
}

/* 配置选择区样式 */
.selection-panel {
  background-color: var(--card-bg);
  border-radius: var(--border-radius);
  padding: var(--spacing-md);
  box-shadow: var(--shadow-sm);
  margin-bottom: var(--spacing-lg);
}

.selection-row {
  display: flex;
  gap: var(--spacing-md);
  align-items: center;
  margin-bottom: var(--spacing-md);
}

.selection-item {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
  flex: 1;
}

.selection-item label {
  font-size: 14px;
  color: var(--text-light);
}

.custom-select {
  width: 100%;
}

.action-buttons {
  display: flex;
  gap: var(--spacing-md);
  justify-content: flex-end;
}

.key-path-container {
  margin-bottom: var(--spacing-md);
}

.key-path-container label {
  font-size: 14px;
  color: var(--text-light);
  margin-bottom: 8px;
  display: block;
}

.custom-input {
  width: 100%;
}

/* 规则列表区样式 */
.rules-container {
  background-color: var(--card-bg);
  border-radius: var(--border-radius);
  padding: var(--spacing-md);
  box-shadow: var(--shadow-sm);
  margin-top: var(--spacing-lg);
}

.rule-card {
  background-color: var(--card-bg);
  border-radius: var(--border-radius);
  margin-bottom: var(--spacing-md);
  border: 1px solid var(--border-color);
}

.rule-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md);
  background-color: rgba(0, 0, 0, 0.02);
  border-bottom: 1px solid var(--border-color);
}

.rule-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 500;
}

.delete-btn {
  padding: 6px;
}

.rule-body {
  padding: var(--spacing-md);
}

.rule-table {
  width: 100%;
  margin-bottom: var(--spacing-md);
}

/* 筛选条件样式 */
.filters-section {
  margin-top: var(--spacing-md);
}

.filters-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-sm);
}

.add-filter-btn {
  padding: 6px 12px;
}

.filters-table {
  width: 100%;
}

/* 空状态样式 */
.empty-state {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-lg);
  background-color: var(--card-bg);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow-sm);
  min-height: 200px;
}

.empty-text {
  color: var(--text-light);
  text-align: center;
}

/* 对话框样式 */
.dialog-body {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.dialog-table {
  width: 100%;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
}
</style>