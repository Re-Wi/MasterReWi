<template>
  <div class="app-container">
    <el-card shadow="always">
      <!-- 查询 -->
      <el-form
          :model="queryParams"
          ref="queryForm"
          :inline="true"
          label-width="68px"
      >
        <el-form-item label="表名称" prop="tableName">
          <el-input
              placeholder="请输入表名称模糊查询"
              clearable
              @keyup.enter="handleQuery"
              style="width: 240px"
              v-model="queryParams.tableName"
          />
        </el-form-item>
        <el-form-item label="表描述" prop="tableComment">
          <el-input
              placeholder="请输入表描述模糊查询"
              clearable
              @keyup.enter="handleQuery"
              style="width: 240px"
              v-model="queryParams.tableComment"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" plain @click="handleQuery">
            <SvgIcon name="elementSearch"/>
            搜索
          </el-button>
          <el-button @click="resetQuery">
            <SvgIcon name="elementRefresh"/>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span class="card-header-text">生成表列表</span>
          <div>
            <el-button
                type="primary"
                plain
                @click="onOpenAddModule"
                v-auth="'develop:code:add'"
            >
              <SvgIcon name="elementDownload"/>
              导入
            </el-button>
            <el-button
                type="danger"
                plain
                :disabled="multiple"
                @click="onTabelRowDel"
                v-auth="'develop:code:delete'"
            >
              <SvgIcon name="elementDelete"/>
              删除
            </el-button>
          </div>
        </div>
      </template>
      <!--数据表格-->
      <el-table
          v-loading="loading"
          :data="tableList"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" align="center"/>
        <el-table-column label="序号" width="55" align="center" prop="tableId"/>
        <el-table-column label="表名称" align="center" prop="tableName" :show-overflow-tooltip="true"/>
        <el-table-column label="表描述" align="center" prop="tableComment" :show-overflow-tooltip="true"/>
        <el-table-column label="模型名称" align="center" prop="className" :show-overflow-tooltip="true"/>
        <el-table-column label="创建时间" align="center" prop="createTime">
          <template #default="scope">
            <span>{{ dateStrFormat(scope.row.createTime) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="更新时间" align="center" prop="updateTime">
          <template #default="scope">
            <span>{{ dateStrFormat(scope.row.updateTime) }}</span>
          </template>
        </el-table-column>
        <el-table-column
            label="操作"
            align="center"
        >
          <template #default="scope">
            <el-popover  placement="left">
              <template #reference>
                <el-button type="primary" circle ><SvgIcon name="elementStar"/></el-button>
              </template>
              <div>
                <el-button text type="primary"
                           v-auth="'develop:code:view'"
                           @click="handlePreview(scope.row)"
                >
                  <SvgIcon name="elementView"/>
                  预览
                </el-button>
              </div>
              <div>
                <el-button text type="primary"
                           v-auth="'develop:code:edit'"
                           @click="onOpenEditModule(scope.row)"
                >
                  <SvgIcon name="elementEdit"/>
                  参数编辑
                </el-button>
              </div>
                <div>
                  <el-button
                          slot="reference"
                          text type="primary"
                          v-auth="'develop:code:code'"
                          @click="handleToProject(scope.row)"
                  >
                    <SvgIcon name="elementDownload"/>
                    代码生成
                  </el-button>
                </div>
                <div>
                  <el-button
                          slot="reference"
                          text type="primary"
                          @click="openConfigureDialog(scope.row)"
                  ><SvgIcon name="elementView" />生成配置</el-button>
                </div>
                <div>
                  <el-button
                          text type="primary"
                          v-auth="'develop:code:delete'"
                          @click="onTabelRowDel(scope.row)"
                  >
                    <SvgIcon name="elementDelete"/>
                    删除
                  </el-button>
                </div>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页设置-->
      <div v-show="total > 0">
        <el-divider></el-divider>
        <el-pagination
            background
            :total="total"
            :current-page="queryParams.pageNum"
            :page-size="queryParams.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
    <!-- 添加或修改岗位对话框 -->
    <ImportTable ref="importTableRef" :title="title"/>
    <EditTable ref="editModelRef"/>
    <el-dialog v-model="isShowConfigureDialog" width="769px" center>
      <el-form :model="configureForm" label-width="80px">
        <el-form-item label="上级菜单" prop="parentId">
          <el-cascader
              v-model="configureForm.parentId"
              :options="menuOptions"
              class="w100"
              :props="{
                  label: 'menuName',
                  value: 'menuId',
                  checkStrictly: true,
                  emitPath: false,
                }"
              clearable
              filterable
              placeholder="选择上级菜单"
              :show-all-levels="false"
          ></el-cascader>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="onCancel">取 消</el-button>
          <el-button type="primary" @click="handleToDB" :loading="loading">编 辑</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog class="preview" :header="preview.title" v-model="preview.open" width="80%" center>
      <div class="el-dialog-container">
        <div class="tag-group">
          <el-tag v-for="(value, key) in preview.data" @click="codeChange(key)">
            <template #>
              {{ key.substring(key.lastIndexOf('/') + 1, key.indexOf('.template')) }}
            </template>
          </el-tag>
        </div>
        <div id="codemirror">
          <Codemirror ref="cmEditor" :value="codestr" border :options="cmOptions" @change="change"/>
        </div>
      </div>
      <template #footer>
                <span class="dialog-footer">
                  <el-button type="primary" @click="preview.open = false">确 定</el-button>
                </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">

</script>
<style lang="scss" scoped>
.el-dialog-container {
  overflow: hidden;

  .el-scrollbar__view {
    height: 100%;
  }

  .pre {
    height: 546px;
    overflow: hidden;

    .el-scrollbar {
      height: 100%;
    }
  }

  .el-scrollbar__wrap::-webkit-scrollbar {
    display: none;
  }
}

.el-dialog__body {
  padding: 0 20px;
  margin: 0;
}

.tag-group {
  margin: 0 0 10px -10px;
}

.tag-group .el-tag {
  margin-left: 10px;
}

.el-tag {
  cursor: pointer;
}
</style>

<style lang="scss">
#codemirror {
  height: auto;
  margin: 0;
  overflow: auto;
}

.CodeMirror {
  border: 1px solid #eee;
  height: 600px;
}
</style>