<template>
  <el-table
    height="450"
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
          <el-form-item label="举报用户ID">
            <span style="color: #409eff">{{ props.row.reporterId }}</span>
          </el-form-item>
          <el-form-item label="被举报者ID">
            <span style="color: #409eff">{{ props.row.readerId }}</span>
          </el-form-item>
          <el-form-item label="评论日期">
            <span style="color: #409eff">
              &nbsp; &nbsp;{{ props.row.date }}</span
            >
          </el-form-item>
          <el-form-item label="书籍 ID">
            <span style="color: #409eff"
              >&nbsp;&nbsp;&nbsp;&nbsp; {{ props.row.bookId }}</span
            >
          </el-form-item>
          <el-form-item label="书籍名称">
            <span style="color: #409eff"
              >&nbsp;&nbsp;&nbsp;{{ props.row.bookName }}</span
            >
          </el-form-item>
          <el-form-item label="评论内容">
            <span style="color: #409eff"
              >&nbsp;&nbsp;{{ props.row.content }}</span
            >
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
          <p>姓名: {{ scope.row.reporterName }}</p>
          <div slot="reference" class="name-wrapper">
            <el-tag size="medium">{{ scope.row.reporterName }}</el-tag>
          </div>
        </el-popover>
      </template>
    </el-table-column>
    <el-table-column label="被举报用户">
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

    <el-table-column label="操作">
      <template slot-scope="scope">
        <el-popconfirm
          v-show="scope.row.status == '审核中'"
          title="确认删除吗？"
          @confirm="auditDel(scope.$index, scope.row)"
        >
          <el-button
            size="mini"
            type="danger"
            slot="reference"
            >删除评论</el-button
          >
        </el-popconfirm>&nbsp;
        <el-popconfirm
          v-show="scope.row.status == '审核中'"
          title="确认驳回吗？"
          @confirm="auditBack(scope.$index, scope.row)"
        >
          <el-button
          size="mini"
          type="warning"
          slot="reference"
          >驳回</el-button
        >
        </el-popconfirm>
       
        <el-button
          size="mini"
          type="warning"
          disabled
          v-show="scope.row.status != '审核中'"
          >{{ scope.row.status }}</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { mapState } from "vuex";
import { auditComment } from "@/api";
import qs from "qs";
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
        return state.Report.reportList;
      },
    }),
  },
  methods: {
    auditDel(index, row) {
      console.log(row);
      const {
        commentId,
        reporterId,
        reportdate,
        email,
        readerId,
        bookId,
        date,
      } = row;
      var infoObj = {
        commentId,
        reporterId,
        reportdate,
        email,
        readerId,
        bookId,
        date,
        status: 0,
      };
      auditComment(qs.stringify(infoObj)).then(
        (res) => {
          if (res.status == 200)
            this.$message({
              showClose: true,
              message: "删除评论成功！",
              type: "success",
            });
          this.$store.dispatch("initReportList");
          this.$store.dispatch("initCommentsList");

          console.log(res);
        },
        (err) => {
          console.log(err.message);
        }
      );
    },
    auditBack(index, row) {
      const { commentId, reporterId, reportdate, email } = row;
      var infoObj = { commentId, reporterId, reportdate, email, status: 1 };
      auditComment(qs.stringify(infoObj)).then(
        (res) => {
          if (res.status == 200)
            this.$message({
              showClose: true,
              message: "驳回成功！",
              type: "success",
            });
          this.$store.dispatch("initReportList");
          this.$store.dispatch("initCommentsList");
          console.log(res.data);
        },
        (err) => {
          console.log(err.message);
        }
      );
    },
  },
  mounted(){
    this.$store.dispatch('initReportList')
  }
};
</script>

<style scoped lang="less">
</style>
