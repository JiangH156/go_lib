<template>
    <div class="main">
        <el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
        <el-radio-button :label="false">展开</el-radio-button>
        <el-radio-button :label="true">收起</el-radio-button>
        </el-radio-group>
         <el-menu
      default-active="1"
      class="el-menu-vertical-demo"
      :collapse="isCollapse"
      @open="handleOpen"
      @close="handleClose"
      active-text-color="#ffd04b">
      <el-menu-item index="1" @click="$router.push('/home/introduce')"> 
        <i class="el-icon-s-home"></i>
        <span>首页</span>
      </el-menu-item>
      <el-menu-item index="2" @click="$router.push('/home/search')">
        <i class="el-icon-search"></i>
        <span slot="title">查询图书</span>
      </el-menu-item>
      <el-menu-item index="3" @click="$router.push('/home/readerreserve')">
        <i class="el-icon-s-promotion"></i>
        <span slot="title">预约记录</span>
      </el-menu-item>
      <el-menu-item index="4" @click="$router.push('/home/readerborrows')">
        <i class="el-icon-document-copy"></i>
        <span slot="title">借阅记录</span>
      </el-menu-item>
      <el-menu-item index="5" @click="$router.push('/home/comment')">
        <i class="el-icon-s-comment"></i>
        <span slot="title">交流社区</span>
      </el-menu-item>
        <el-menu-item index="6" @click="$router.push('/home/readerreport')">
        <i class="el-icon-question"></i>
        <span slot="title">举报反馈</span>
      </el-menu-item>
      <el-submenu index="7">
        <template slot="title">
          <i class="el-icon-setting"></i>
          <span>设置</span>
        </template>
        <el-submenu index="1-1">
          <template slot="title">
          <i class="el-icon-user"></i>
                <span>我的</span>
          </template>
          <el-menu-item index="1-1-1" @click="$router.push('/home/readerreserve')">
              <i class="el-icon-date"></i>
              <span>预约记录</span>
              </el-menu-item>
          <el-menu-item index="1-1-2" @click="$router.push('/home/readerborrows')">
              <i class="el-icon-document-copy"></i>
              <span>借阅记录</span>
              </el-menu-item>
          <el-menu-item index="1-1-3" @click="toggleUser">
              <i class="el-icon-s-fold"></i>
              <span>切换账号</span>
        </el-menu-item>
        </el-submenu>
         <el-menu-item index="1-2" @click="toggleUser">
        <i class="el-icon-switch-button"></i>
        <span slot="title">退出账号</span>
      </el-menu-item>
      </el-submenu>
    </el-menu>
    </div>
</template>

<script>
import {mapState} from 'vuex'
export default {
  name: "ReaderBanner",
  data() {
    return {
      isCollapse: true
    };
  },

  mounted() {
    // this.$store.dispatch('initCommentsList')
    // this.$store.dispatch('initBorrows',{readerId:this.readerId})
    // this.$store.dispatch('initReserve',{readerId:this.readerId})
    // this.$store.dispatch('initStuReport',{readerId:this.readerId})
  },
 computed:{
    ...mapState({
      readerId(state){
        return state.User.readerInfo.readerId
      }
    })
  },
  methods: {
    handleOpen() {
      console.log(1);
    },
    handleClose() {
      console.log(2);
    },
    toggleUser() {
      this.$router.push("/LoginRegister");
    }
  }
};
</script>

<style lang="less" scoped>
.main{
    float: left;
    margin-right: 200px;
}
</style>