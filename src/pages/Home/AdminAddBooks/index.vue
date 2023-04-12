<template>
	<el-form ref="form" label-width="80px">
  <el-form-item label="书籍名称" label-width="80px">
	<el-col :span="12">
    	<el-input v-model="bookName"></el-input>
	  </el-col>
  </el-form-item>
  <el-form-item label="书籍作者" label-width="80px">
	  <el-col :span="6">
    	<el-input v-model="author"></el-input>
	  </el-col>
  </el-form-item>
  <el-form-item label="总库存" label-width="80px">
	    <el-col :span="6">
    		<el-input v-model="amount" min="0" type="number"></el-input>
	  </el-col>
  </el-form-item>
  <el-form-item label="图书位置" label-width="80px">
	  <el-col :span="6">
	<el-autocomplete
	popper-class="my-autocomplete"
	v-model="position"
	:fetch-suggestions="querySearch"
	placeholder="请输入内容"
	@select="handleSelect">
	<i
		class="el-icon-edit el-input__icon"
		slot="suffix"
	>
	</i>
	<template slot-scope="{ item }">
		<div class="name">{{ item.position }}</div>
		<!-- <span class="addr">{{ item.address }}</span> -->
	</template>
	</el-autocomplete>
	  </el-col>	  

	 </el-form-item>  
	 <el-form-item>
	  <el-col :span="6">
    <el-button type="primary" @click="addBook">立即添加</el-button>
		  
	  </el-col>		 
  </el-form-item>
</el-form>
</template>

<script>
import { mapState } from "vuex";
import {addBooks} from '@/api'
import qs from 'qs'
export default {
	name: 'AdminAddBooks',

  data() {
      return {
		bookName: '',
		author:'',
		amount:'',
        position: ''
      }
    },
 computed: {
    ...mapState({
      booksList(state) {
        return state.Books.booksList;
      }
     
    })
  },
	 methods: {
      querySearch(queryString, cb) {
        cb(this.booksList);
      },
      handleSelect(item) {
		// console.log(item);
		this.position = item.position
	  },
	  addBook(){
		  let {bookName,author,amount,position} = this
		  let infoObj = {bookName,author,amount,position}
		  addBooks(qs.stringify(infoObj)).then(res=>{
			  console.log(res);
			  if(res.status == 200){
				    this.$message({
                        showClose: true,
                        message: res.msg,
                        type: 'success',
					});
					 this.bookName =  '',
					this.author = '',
					this.amount = '',
					this.position =  ''
                    this.$store.dispatch('initBooksList')

			  }else{
				   this.$message({
                        showClose: true,
                        message: res.msg,
                        type: 'error',
                    });
			  }
			   
			  
		  },err=>{
			  console.log(err.message);
			  
		  })
	  }
    }
	
};
</script>

<style lang="less" scoped>
.my-autocomplete {
  li {
    line-height: normal;
    padding: 7px;

    .name {
      text-overflow: ellipsis;
      overflow: hidden;
    }
    .addr {
      font-size: 12px;
      color: #b4b4b4;
    }

    .highlighted .addr {
      color: #ddd;
    }
  }
}
</style>