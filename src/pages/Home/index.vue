<template>
    <div class="clearfix">
        <header>
            <img src="./images/library.jpeg" alt="">
        </header>
        
        <el-row :gutter="20">
            <el-col :span="6">
                    <AdminBanner v-if="isAdmin"/>
                    <ReaderBanner v-else/>
            </el-col>
            <el-col :span="14">  
                <div class="tablemain">
                    <router-view />
                </div>
            </el-col>
        </el-row>
    </div>
</template>
<script>
import axios from 'axios'
import qs from "qs" 
import {mapState} from 'vuex'
export default {
    name: 'Home',
    data() {
        return {
            
        };
    },
    computed:{
        ...mapState({
           isAdmin(state){
               return state.User.isAdmin
           },
           userName(state){
               if(this.isAdmin){
                   return state.User.adminName
               }else{
                   return state.User.readerName
               }
           },
           readerId(state){
               return state.User.readerInfo.readerId
           }
        })
    },
   
};
</script>

<style lang="less" scoped>
header{
    text-align: center;
    color: rgb(16, 148, 93);
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
    img{
        height: 120px;
    }
}
.tablemain{
     box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
     border: 1px solid #eee;
     border-radius: 1px 1px 1px ;
     padding: 20px;
     min-height:450px
}
</style>