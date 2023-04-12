<template>
  <el-table
    :data="reportList"
    style="width: 100%"
    v-loading.fullscreen.lock="loading"
    element-loading-text="正在处理..."
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" class="demo-table-expand">
          <el-form-item label="评论书籍:"
            >&nbsp;
            <span style="color: #409eff">{{ props.row.bookName }}</span>
          </el-form-item>
          <el-form-item label="评论内容:"
            >&nbsp;
            <span style="color: #409eff">{{ props.row.content }}</span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column label="举报日期">
      <template slot-scope="scope">
        <i class="el-icon-time"></i>
        <span style="margin-left: 10px">{{ scope.row.reportdate }}</span>
      </template>
    </el-table-column>
    <el-table-column label="举报用户">
      <template slot-scope="scope">
        <el-popover trigger="hover" placement="top">
          <p>
            姓名:
            {{
              !scope.row.readerName ? scope.row.readerId : scope.row.readerName
            }}
          </p>
          <div slot="reference" class="name-wrapper">
            <el-tag size="medium">{{
              !scope.row.readerName ? scope.row.readerId : scope.row.readerName
            }}</el-tag>
          </div>
        </el-popover>
      </template>
    </el-table-column>

    <el-table-column label="审核状态">
      <template slot-scope="scope">
        <el-button size="mini" disabled slot="reference">{{
          scope.row.status
        }}</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "ReaderReport",
  data() {
    return {
      loading: false,
    };
  },
  computed: {
    ...mapState({
      reportList(state) {
        return state.Report.stuReport;
      },
      readerId(state) {
        return state.User.readerInfo.readerId;
      },
    }),
  },
  mounted() {
    this.$store.dispatch("initStuReport", { readerId: this.readerId });
  },
};
</script>

<style scoped lang="less">
</style>
